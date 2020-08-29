package acfundanmu

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fastjson"
)

// 检查错误
func checkErr(err error) {
	if err != nil {
		panicln(err)
	}
}

// 打印错误然后panic
func panicln(err error) {
	log.Println(err)
	panic(err)
}

// 登陆acfun账号
func login(username string, password string) (cookieContainer []*http.Cookie, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("login() error: %w", err)
		}
	}()

	form := url.Values{}
	form.Set("username", username)
	form.Set("password", password)
	form.Set("key", "")
	form.Set("captcha", "")
	resp, err := http.PostForm(acfunSignInURL, form)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panicln(fmt.Errorf("以注册用户的身份登陆AcFun失败，响应为 %s", string(body)))
	}

	for _, cookie := range resp.Cookies() {
		cookieContainer = append(cookieContainer, cookie)
	}

	userID := v.GetInt("userId")
	content := fmt.Sprintf(safetyIDContent, userID)
	resp, err = http.Post(acfunSafetyIDURL, "", strings.NewReader(content))
	checkErr(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("code") != 0 {
		panicln(fmt.Errorf("获取safetyid失败，响应为 %s", string(body)))
	}

	cookie := &http.Cookie{Name: "safety_id", Value: string(v.GetStringBytes("safety_id")), Domain: ".acfun.cn"}
	cookieContainer = append(cookieContainer, cookie)

	return cookieContainer, nil
}

// 初始化，获取相应的token，cookieContainer为nil时为游客模式
func initialize(uid int, cookieContainer []*http.Cookie) (t *token, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("initialize() error: %w", err)
		}
	}()

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := http.Get(liveURL + strconv.Itoa(uid))
	checkErr(err)
	defer resp.Body.Close()

	// 获取did（device ID）
	var didCookie *http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_did" {
			didCookie = cookie
		}
	}
	deviceID := didCookie.Value

	var req *http.Request
	form := url.Values{}
	if cookieContainer != nil {
		form.Set(sid, midground)
		req, err = http.NewRequest(http.MethodPost, getTokenURL, strings.NewReader(form.Encode()))
		checkErr(err)
		for _, cookie := range cookieContainer {
			req.AddCookie(cookie)
		}
	} else {
		form.Set(sid, visitor)
		req, err = http.NewRequest(http.MethodPost, loginURL, strings.NewReader(form.Encode()))
		checkErr(err)
		req.AddCookie(didCookie)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panicln(fmt.Errorf("获取AcFun token失败，响应为 %s", string(body)))
	}

	// 获取userId和对应的令牌
	userID := v.GetInt64("userId")
	var play, serviceToken, securityKey string
	if cookieContainer != nil {
		securityKey = string(v.GetStringBytes("ssecurity"))
		serviceToken = string(v.GetStringBytes(midgroundSt))
		// 需要userId、deviceID和serviceToken
		play = fmt.Sprintf(playURL, userID, deviceID, midgroundSt, serviceToken)
	} else {
		securityKey = string(v.GetStringBytes("acSecurity"))
		serviceToken = string(v.GetStringBytes(visitorSt))
		play = fmt.Sprintf(playURL, userID, deviceID, visitorSt, serviceToken)
	}

	form = url.Values{}
	// authorId就是主播的uid
	form.Set("authorId", strconv.Itoa(uid))
	form.Set("pullStreamType", "FLV")
	req, err = http.NewRequest(http.MethodPost, play, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 会验证Referer
	req.Header.Set("Referer", liveURL+strconv.Itoa(uid))
	resp, err = client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panicln(fmt.Errorf("获取直播详细信息失败，响应为 %s", string(body)))
	}

	liveID := string(v.GetStringBytes("data", "liveId"))
	enterRoomAttach := string(v.GetStringBytes("data", "enterRoomAttach"))
	availableTickets := v.GetArray("data", "availableTickets")
	tickets := make([]string, len(availableTickets))
	for i, ticket := range availableTickets {
		tickets[i] = string(ticket.GetStringBytes())
	}

	t = &token{
		userID:          userID,
		securityKey:     securityKey,
		serviceToken:    serviceToken,
		liveID:          liveID,
		enterRoomAttach: enterRoomAttach,
		tickets:         tickets,
		instanceID:      0,
		sessionKey:      "",
		seqID:           1,
		headerSeqID:     1,
		heartbeatSeqID:  1,
		ticketIndex:     0,
		deviceID:        deviceID,
	}

	err = t.updateGiftList(cookieContainer)
	checkErr(err)

	return t, nil
}

// 更新礼物列表
func (t *token) updateGiftList(cookieContainer []*http.Cookie) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("updateGiftList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var giftList string
	if cookieContainer != nil {
		giftList = fmt.Sprintf(giftURL, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		giftList = fmt.Sprintf(giftURL, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	req, err := http.NewRequest(http.MethodPost, giftList, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", liveMainPage)
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panicln(fmt.Errorf("获取礼物列表失败，响应为 %s", string(body)))
	}

	t.gifts = make(map[int]Giftdetail)
	for _, gift := range v.GetArray("data", "giftList") {
		g := Giftdetail{
			ID:          gift.GetInt("giftId"),
			Name:        string(gift.GetStringBytes("giftName")),
			Price:       gift.GetInt("giftPrice"),
			WebpPic:     string(gift.GetArray("webpPicList")[0].GetStringBytes("url")),
			PngPic:      string(gift.GetArray("pngPicList")[0].GetStringBytes("url")),
			SmallPngPic: string(gift.GetArray("smallPngPicList")[0].GetStringBytes("url")),
			Description: string(gift.GetStringBytes("description")),
		}
		t.gifts[g.ID] = g
	}

	return nil
}

// 获取直播间排名前50的在线观众信息列表
func (t *token) watchingList(cookieContainer []*http.Cookie) (watchList *[]WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var watchURL string
	if cookieContainer != nil {
		watchURL = fmt.Sprintf(watchingURL, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		watchURL = fmt.Sprintf(watchingURL, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	req, err := http.NewRequest(http.MethodPost, watchURL, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", liveMainPage)
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panicln(fmt.Errorf("获取在线观众列表失败，响应为 %s", string(body)))
	}

	watchArray := v.GetArray("data", "list")
	watchingUserList := make([]WatchingUser, len(watchArray))
	for i, watch := range watchArray {
		watchingUserList[i] = WatchingUser{
			UserInfo: UserInfo{
				UserID:   watch.GetInt64("userId"),
				Nickname: string(watch.GetStringBytes("nickname")),
				Avatar:   string(watch.GetStringBytes("avatar", "0", "url")),
			},
			AnonymousUser:          watch.GetBool("anonymousUser"),
			DisplaySendAmount:      string(watch.GetStringBytes("displaySendAmount")),
			CustomWatchingListData: string(watch.GetStringBytes("customWatchingListData")),
		}
	}

	return &watchingUserList, nil
}
