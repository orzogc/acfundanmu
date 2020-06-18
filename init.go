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
func login(username string, password string) (cookieContainer []*http.Cookie) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from panic in login(), the error is:", err)
			cookieContainer = nil
		}
	}()

	client := &http.Client{}

	form := url.Values{}
	form.Set("username", username)
	form.Set("password", password)
	form.Set("key", "")
	form.Set("captcha", "")
	req, err := http.NewRequest("POST", acfunSignInURL, strings.NewReader(form.Encode()))
	checkErr(err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 0 {
		return nil
	}

	userID := v.GetInt("userId")
	content := fmt.Sprintf(safetyIDContent, userID)
	req, err = http.NewRequest("POST", acfunSafetyIDURL, strings.NewReader(content))
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
		return nil
	}

	cookie := &http.Cookie{Name: "safety_id", Value: string(v.GetStringBytes("safety_id")), Domain: ".acfun.cn"}
	cookieContainer = append(cookieContainer, cookie)

	return cookieContainer
}

// 初始化，获取相应的token，cookieContainer为nil时为游客模式
func initialize(uid int, cookieContainer []*http.Cookie) (deviceID string, t *token) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from panic in initialize(), the error is:", err)
			deviceID, t = "", nil
		}
	}()

	client := &http.Client{}

	req, err := http.NewRequest("GET", liveURL+strconv.Itoa(uid), nil)
	checkErr(err)
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	// 获取did（device ID）
	var didCookie *http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_did" {
			didCookie = cookie
		}
	}
	deviceID = didCookie.Value

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
		return "", nil
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
	req, err = http.NewRequest("POST", play, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		return "", nil
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

	return deviceID, t
}

// 更新礼物列表
func updateGiftList(cookieContainer []*http.Cookie, deviceID string, t *token) map[int]string {
	if t == nil {
		log.Println("获取token失败，可能主播不在直播")
		return nil
	}

	var giftList string
	if cookieContainer != nil {
		giftList = fmt.Sprintf(giftURL, t.userID, deviceID, midgroundSt, t.serviceToken)
	} else {
		giftList = fmt.Sprintf(giftURL, t.userID, deviceID, visitorSt, t.serviceToken)
	}

	client := &http.Client{}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	req, err := http.NewRequest("POST", giftList, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		return nil
	}

	gifts := make(map[int]string)
	for _, gift := range v.GetArray("data", "giftList") {
		gifts[gift.GetInt("giftId")] = string(gift.GetStringBytes("giftName"))
	}

	return gifts
}

// 获取在线观众列表
func watchingList(cookieContainer []*http.Cookie, deviceID string, t *token) map[int]string {
	if t == nil {
		log.Println("获取token失败，可能主播不在直播")
		return nil
	}

	var watchURL string
	if cookieContainer != nil {
		watchURL = fmt.Sprintf(watchingURL, t.userID, deviceID, midgroundSt, t.serviceToken)
	} else {
		watchURL = fmt.Sprintf(watchingURL, t.userID, deviceID, visitorSt, t.serviceToken)
	}

	client := &http.Client{}

	form := url.Values{}
	form.Set("visitorId", strconv.Itoa(int(t.userID)))
	form.Set("liveId", t.liveID)
	req, err := http.NewRequest("POST", watchURL, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		return nil
	}

	watchList := make(map[int]string)
	for _, watch := range v.GetArray("data", "list") {
		watchList[watch.GetInt("userId")] = string(watch.GetStringBytes("nickname"))
	}

	return watchList
}
