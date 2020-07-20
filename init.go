package acfundanmu

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/valyala/fastjson"
)

// 检查错误
func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// 登陆acfun账号
func login(username string, password string) (cookieContainer []*http.Cookie, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("login() error: %w", err)
		}
	}()

	client := &http.Client{}

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
	if v.GetInt("result") != 0 {
		log.Panicf("以注册用户的身份登陆AcFun失败，响应为 %s", string(body))
	}

	userID := v.GetInt("userId")
	content := fmt.Sprintf(safetyIDContent, userID)
	req, err := http.NewRequest("POST", acfunSafetyIDURL, strings.NewReader(content))
	checkErr(err)

	for _, cookie := range resp.Cookies() {
		cookieContainer = append(cookieContainer, cookie)
	}
	resp, err = client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("code") != 0 {
		log.Panicf("获取safetyid失败，响应为 %s", string(body))
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

	client := &http.Client{}

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
		req, err = http.NewRequest("POST", getTokenURL, strings.NewReader(form.Encode()))
		checkErr(err)
		for _, cookie := range cookieContainer {
			req.AddCookie(cookie)
		}
	} else {
		form.Set(sid, visitor)
		req, err = http.NewRequest("POST", loginURL, strings.NewReader(form.Encode()))
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
	if v.GetInt("result") != 0 {
		log.Panicf("获取AcFun token失败，响应为 %s", string(body))
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
	resp, err = http.PostForm(play, form)
	checkErr(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		log.Panicf("获取直播详细信息失败，响应为 %s", string(body))
	}

	liveID := string(v.GetStringBytes("data", "liveId"))
	enterRoomAttach := string(v.GetStringBytes("data", "enterRoomAttach"))
	availableTickets := v.GetArray("data", "availableTickets")
	var tickets []string
	for _, ticket := range availableTickets {
		tickets = append(tickets, string(ticket.GetStringBytes()))
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
	}

	err = t.updateGiftList(cookieContainer, deviceID)
	checkErr(err)

	return t, nil
}

// 更新礼物列表
func (t *token) updateGiftList(cookieContainer []*http.Cookie, deviceID string) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("updateGiftList() error: %w", err)
		}
	}()

	if t == nil {
		log.Panicln("获取token失败，可能主播不在直播")
	}

	var giftList string
	if cookieContainer != nil {
		giftList = fmt.Sprintf(giftURL, t.userID, deviceID, midgroundSt, t.serviceToken)
	} else {
		giftList = fmt.Sprintf(giftURL, t.userID, deviceID, visitorSt, t.serviceToken)
	}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	resp, err := http.PostForm(giftList, form)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		log.Panicf("获取礼物列表失败，响应为 %s", string(body))
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

// 获取在线观众列表
func (t *token) watchingList(cookieContainer []*http.Cookie, deviceID string) (watchList map[int]string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	if t == nil {
		log.Panicln("获取token失败，可能主播不在直播")
	}

	var watchURL string
	if cookieContainer != nil {
		watchURL = fmt.Sprintf(watchingURL, t.userID, deviceID, midgroundSt, t.serviceToken)
	} else {
		watchURL = fmt.Sprintf(watchingURL, t.userID, deviceID, visitorSt, t.serviceToken)
	}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	resp, err := http.PostForm(watchURL, form)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		log.Panicf("获取在线观众列表失败，响应为 %s", string(body))
	}

	watchList = make(map[int]string)
	for _, watch := range v.GetArray("data", "list") {
		watchList[watch.GetInt("userId")] = string(watch.GetStringBytes("nickname"))
	}

	return watchList, nil
}
