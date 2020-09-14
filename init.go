package acfundanmu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fastjson"
)

// 登陆acfun账号
func (t *token) login(username, password string) (e error) {
	defer func() {
		if err := recover(); err != nil {
			t.cookies = nil
			e = fmt.Errorf("login() error: %w", err)
		}
	}()

	if username == "" || password == "" {
		panicln(fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆"))
	}

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

	t.cookies = resp.Cookies()

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
	t.cookies = append(t.cookies, cookie)

	return nil
}

// 获取相应的token
func (t *token) getToken() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getToken() error: %w", err)
		}
	}()

	t.livePage = liveURL + strconv.FormatInt(t.uid, 10)

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := http.Get(t.livePage)
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
	if len(t.cookies) != 0 {
		form.Set(sid, midground)
		req, err = http.NewRequest(http.MethodPost, getTokenURL, strings.NewReader(form.Encode()))
		checkErr(err)
		for _, cookie := range t.cookies {
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
	if len(t.cookies) != 0 {
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
	form.Set("authorId", strconv.FormatInt(t.uid, 10))
	form.Set("pullStreamType", "FLV")
	req, err = http.NewRequest(http.MethodPost, play, strings.NewReader(form.Encode()))
	checkErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 会验证Referer
	req.Header.Set("Referer", t.livePage)
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

	v = v.Get("data")
	liveID := string(v.GetStringBytes("liveId"))
	enterRoomAttach := string(v.GetStringBytes("enterRoomAttach"))
	availableTickets := v.GetArray("availableTickets")
	tickets := make([]string, len(availableTickets))
	for i, ticket := range availableTickets {
		tickets[i] = string(ticket.GetStringBytes())
	}

	t.userID = userID
	t.securityKey = securityKey
	t.serviceToken = serviceToken
	t.liveID = liveID
	t.enterRoomAttach = enterRoomAttach
	t.tickets = tickets
	t.instanceID = 0
	t.sessionKey = ""
	t.seqID = 1
	t.headerSeqID = 1
	t.heartbeatSeqID = 1
	t.ticketIndex = 0
	t.deviceID = deviceID

	err = t.updateGiftList()
	checkErr(err)

	return nil
}

// 更新礼物列表
func (t *token) updateGiftList() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("updateGiftList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var giftList string
	if len(t.cookies) != 0 {
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
	req.Header.Set("Referer", t.livePage)
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

	t.gifts = make(map[int64]Giftdetail)
	for _, gift := range v.GetArray("data", "giftList") {
		g := Giftdetail{
			GiftID:        gift.GetInt64("giftId"),
			GiftName:      string(gift.GetStringBytes("giftName")),
			ARLiveName:    string(gift.GetStringBytes("arLiveName")),
			PayWalletType: gift.GetInt("payWalletType"),
			Price:         gift.GetInt("giftPrice"),
			WebpPic:       string(gift.GetStringBytes("webpPicList", "0", "url")),
			PngPic:        string(gift.GetStringBytes("pngPicList", "0", "url")),
			SmallPngPic:   string(gift.GetStringBytes("smallPngPicList", "0", "url")),
			CanCombo:      gift.GetBool("canCombo"),
			MagicFaceID:   gift.GetInt("magicFaceId"),
			Description:   string(gift.GetStringBytes("description")),
			RedpackPrice:  gift.GetInt("redpackPrice"),
		}
		t.gifts[g.GiftID] = g
	}

	return nil
}

// 获取直播间排名前50的在线观众信息列表
func (t *token) watchingList() (watchList *[]WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var watchURL string
	if len(t.cookies) != 0 {
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
	req.Header.Set("Referer", t.livePage)
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	p := t.watchParser.Get()
	defer t.watchParser.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panicln(fmt.Errorf("获取在线观众列表失败，响应为 %s", string(body)))
	}

	watchArray := v.GetArray("data", "list")
	watchingUserList := make([]WatchingUser, len(watchArray))
	for i, watch := range watchArray {
		o := watch.GetObject()
		w := WatchingUser{}
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				w.UserID = v.GetInt64()
			case "nickname":
				w.Nickname = string(v.GetStringBytes())
			case "avatar":
				w.Avatar = string(v.GetStringBytes("0", "url"))
			case "anonymousUser":
				w.AnonymousUser = v.GetBool()
			case "displaySendAmount":
				w.DisplaySendAmount = string(v.GetStringBytes())
			case "customWatchingListData":
				w.CustomWatchingListData = string(v.GetStringBytes())
			case "managerType":
				w.ManagerType = ManagerType(v.GetInt())
			}
		})
		watchingUserList[i] = w
	}

	return &watchingUserList, nil
}

// 初始化
func (t *token) initialize(usernameAndPassword ...string) error {
	if len(usernameAndPassword) == 2 && usernameAndPassword[0] != "" && usernameAndPassword[1] != "" {
		err := t.login(usernameAndPassword[0], usernameAndPassword[1])
		if err != nil {
			return err
		}
	}
	return t.getToken()
}
