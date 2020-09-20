package acfundanmu

import (
	"fmt"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

type httpClient struct {
	client      *fasthttp.Client
	url         string
	body        []byte
	method      string
	cookies     []*fasthttp.Cookie
	contentType string
	referer     string
}

// http请求，调用后需要 defer fasthttp.ReleaseResponse(resp)
func (c *httpClient) doRequest() (resp *fasthttp.Response, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("doRequest() error: %w", err)
			fasthttp.ReleaseResponse(resp)
		}
	}()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp = fasthttp.AcquireResponse()

	if c.client == nil {
		panic(fmt.Errorf("client不能为nil"))
	}

	if c.url != "" {
		req.SetRequestURI(c.url)
	} else {
		panic(fmt.Errorf("请求的url不能为空"))
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

	err := c.client.Do(req, resp)
	checkErr(err)

	return resp, nil
}

// 登陆acfun账号
func login(username, password string) (cookies []string, e error) {
	defer func() {
		if err := recover(); err != nil {
			cookies = nil
			e = fmt.Errorf("login() error: %w", err)
		}
	}()

	if username == "" || password == "" {
		panic(fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("username", username)
	form.Set("password", password)
	form.Set("key", "")
	form.Set("captcha", "")

	c := &fasthttp.Client{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	client := &httpClient{
		client:      c,
		url:         acfunSignInURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("以注册用户的身份登陆AcFun失败，响应为 %s", string(body)))
	}

	resp.Header.VisitAllCookie(func(key, value []byte) {
		cookies = append(cookies, string(value))
	})

	userID := v.GetInt("userId")
	content := fmt.Sprintf(safetyIDContent, userID)
	client = &httpClient{
		client: c,
		url:    acfunSafetyIDURL,
		body:   []byte(content),
		method: "POST",
	}
	resp, err = client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body = resp.Body()

	v, err = p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("code") || v.GetInt("code") != 0 {
		panic(fmt.Errorf("获取safetyid失败，响应为 %s", string(body)))
	}

	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)
	cookie.SetKey("safety_id")
	cookie.SetValueBytes(v.GetStringBytes("safety_id"))
	cookie.SetDomain(".acfun.cn")
	cookies = append(cookies, cookie.String())

	return cookies, nil
}

// 获取相应的token
func (t *token) getToken() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getToken() error: %w", err)
		}
	}()

	client := &httpClient{
		client: t.client,
		url:    t.livePage,
		method: "GET",
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	// 获取did（device ID）
	didCookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(didCookie)
	resp.Header.VisitAllCookie(func(key, value []byte) {
		if string(key) == "_did" {
			err = didCookie.ParseBytes(value)
			checkErr(err)
		}
	})
	deviceID := string(didCookie.Value())

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	if len(t.cookies) != 0 {
		form.Set(sid, midground)
		var cookies []*fasthttp.Cookie
		for _, c := range t.cookies {
			cookie := fasthttp.AcquireCookie()
			defer fasthttp.ReleaseCookie(cookie)
			err = cookie.Parse(c)
			checkErr(err)
			cookies = append(cookies, cookie)
		}
		client = &httpClient{
			url:     getTokenURL,
			body:    form.QueryString(),
			cookies: cookies,
		}
	} else {
		form.Set(sid, visitor)
		client = &httpClient{
			url:     loginURL,
			body:    form.QueryString(),
			cookies: []*fasthttp.Cookie{didCookie},
		}
	}
	client.client = t.client
	client.method = "POST"
	client.contentType = formContentType
	resp, err = client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取AcFun token失败，响应为 %s", string(body)))
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

	form = fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	// authorId就是主播的uid
	form.Set("authorId", strconv.FormatInt(t.uid, 10))
	form.Set("pullStreamType", "FLV")
	client = &httpClient{
		client:      t.client,
		url:         play,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
		referer:     t.livePage, // 会验证 Referer
	}
	resp, err = client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body = resp.Body()

	v, err = p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播详细信息失败，响应为 %s", string(body)))
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

	resp, err := t.fetchKuaiShouAPI(giftURL)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取礼物列表失败，响应为 %s", string(body)))
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
		list := gift.GetArray("allowBatchSendSizeList")
		g.AllowBatchSendSizeList = make([]int, len(list))
		for i, num := range list {
			g.AllowBatchSendSizeList[i] = num.GetInt()
		}
		t.gifts[g.GiftID] = g
	}

	return nil
}

// 通过快手API获取数据，调用后需要 defer fasthttp.ReleaseResponse(resp)
func (t *token) fetchKuaiShouAPI(url string) (*fasthttp.Response, error) {
	if t == nil {
		panic(fmt.Errorf("获取token失败，可能主播不在直播"))
	}

	var apiURL string
	if len(t.cookies) != 0 {
		apiURL = fmt.Sprintf(url, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		apiURL = fmt.Sprintf(url, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	client := &httpClient{
		client:      t.client,
		url:         apiURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
	}
	return client.doRequest()
}
