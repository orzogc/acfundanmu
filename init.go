package acfundanmu

import (
	"fmt"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

type httpClient struct {
	url         string
	body        []byte
	method      string
	cookies     []*fasthttp.Cookie
	contentType string
	referer     string
}

func (c *httpClient) httpRequest() (response *fasthttp.Response, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("httpRequest() error: %w", err)
		}
	}()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if c.url != "" {
		req.SetRequestURI(c.url)
	} else {
		panicln(fmt.Errorf("请求的url不能为空"))
	}

	if len(c.body) != 0 {
		req.SetBody(c.body)
	}

	if c.method != "" {
		req.Header.SetMethod(c.method)
	} else {
		// 默认为GET
		req.Header.SetMethod("GET")
	}

	if len(c.cookies) != 0 {
		for _, cookie := range c.cookies {
			req.Header.SetCookieBytesKV(cookie.Key(), cookie.Value())
		}
	}

	if c.contentType != "" {
		req.Header.SetContentType(c.contentType)
	}

	if c.referer != "" {
		req.Header.SetReferer(c.referer)
	}

	client := &fasthttp.Client{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := client.Do(req, resp)
	checkErr(err)

	response = fasthttp.AcquireResponse()
	resp.CopyTo(response)

	return response, nil
}

// 登陆acfun账号
func login(username, password string) (cookies []*fasthttp.Cookie, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("login() error: %w", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("username", username)
	form.Set("password", password)
	form.Set("key", "")
	form.Set("captcha", "")

	client := &httpClient{
		url:         acfunSignInURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: "application/x-www-form-urlencoded",
	}
	resp, err := client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panicln(fmt.Errorf("以注册用户的身份登陆AcFun失败，响应为 %s", string(body)))
	}

	resp.Header.VisitAllCookie(func(key, value []byte) {
		cookie := fasthttp.AcquireCookie()
		err = cookie.ParseBytes(value)
		cookies = append(cookies, cookie)
	})

	userID := v.GetInt("userId")
	content := fmt.Sprintf(safetyIDContent, userID)
	client = &httpClient{
		url:    acfunSafetyIDURL,
		body:   []byte(content),
		method: "POST",
	}
	resp, err = client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body = resp.Body()

	v, err = p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("code") || v.GetInt("code") != 0 {
		panicln(fmt.Errorf("获取safetyid失败，响应为 %s", string(body)))
	}

	cookie := fasthttp.AcquireCookie()
	cookie.SetKey("safety_id")
	cookie.SetValueBytes(v.GetStringBytes("safety_id"))
	cookie.SetDomain(".acfun.cn")
	cookies = append(cookies, cookie)

	return cookies, nil
}

// 初始化，获取相应的token，cookies为nil时为游客模式
func initialize(uid int, cookies []*fasthttp.Cookie) (t *token, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("initialize() error: %w", err)
		}
	}()

	livePage := liveURL + strconv.Itoa(uid)

	client := &httpClient{
		url:    livePage,
		method: "GET",
	}
	resp, err := client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	// 获取did（device ID）
	didCookie := fasthttp.AcquireCookie()
	resp.Header.VisitAllCookie(func(key, value []byte) {
		if string(key) == "_did" {
			err = didCookie.ParseBytes(value)
			checkErr(err)
		}
	})
	deviceID := string(didCookie.Value())

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	if cookies != nil {
		form.Set(sid, midground)
		client = &httpClient{
			url:     getTokenURL,
			body:    form.QueryString(),
			method:  "POST",
			cookies: cookies,
		}
	} else {
		form.Set(sid, visitor)
		client = &httpClient{
			url:     loginURL,
			body:    form.QueryString(),
			method:  "POST",
			cookies: []*fasthttp.Cookie{didCookie},
		}
	}
	client.contentType = "application/x-www-form-urlencoded"
	resp, err = client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panicln(fmt.Errorf("获取AcFun token失败，响应为 %s", string(body)))
	}

	// 获取userId和对应的令牌
	userID := v.GetInt64("userId")
	var play, serviceToken, securityKey string
	if cookies != nil {
		securityKey = string(v.GetStringBytes("ssecurity"))
		serviceToken = string(v.GetStringBytes(midgroundSt))
		// 需要userId、deviceID和serviceToken
		play = fmt.Sprintf(playURL, userID, deviceID, midgroundSt, serviceToken)
	} else {
		securityKey = string(v.GetStringBytes("acSecurity"))
		serviceToken = string(v.GetStringBytes(visitorSt))
		play = fmt.Sprintf(playURL, userID, deviceID, visitorSt, serviceToken)
	}

	form = fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	// authorId就是主播的uid
	form.Set("authorId", strconv.Itoa(uid))
	form.Set("pullStreamType", "FLV")
	client = &httpClient{
		url:         play,
		body:        form.QueryString(),
		method:      "POST",
		contentType: "application/x-www-form-urlencoded",
		referer:     livePage, // 会验证Referer
	}
	resp, err = client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body = resp.Body()

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
		medalParser:     fastjson.ParserPool{},
		watchParser:     fastjson.ParserPool{},
	}

	err = t.updateGiftList(cookies)
	checkErr(err)

	return t, nil
}

// 更新礼物列表
func (t *token) updateGiftList(cookies []*fasthttp.Cookie) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("updateGiftList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var giftList string
	if cookies != nil {
		giftList = fmt.Sprintf(giftURL, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		giftList = fmt.Sprintf(giftURL, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	client := &httpClient{
		url:         giftList,
		body:        form.QueryString(),
		method:      "POST",
		contentType: "application/x-www-form-urlencoded",
		referer:     liveMainPage,
	}
	resp, err := client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

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
func (t *token) watchingList(cookies []*fasthttp.Cookie) (watchList *[]WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	if t == nil {
		panicln(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var watchURL string
	if cookies != nil {
		watchURL = fmt.Sprintf(watchingURL, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		watchURL = fmt.Sprintf(watchingURL, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	client := &httpClient{
		url:         watchURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: "application/x-www-form-urlencoded",
		referer:     liveMainPage,
	}
	resp, err := client.httpRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

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
