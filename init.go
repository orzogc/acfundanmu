package acfundanmu

import (
	"fmt"
	"log"
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

var defaultClient = &fasthttp.Client{
	MaxIdleConnDuration: 90 * time.Second,
	ReadTimeout:         10 * time.Second,
	WriteTimeout:        10 * time.Second,
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
		c.client = defaultClient
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

	client := &httpClient{
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

	return cookies, nil
}

// 获取AcFun帐号的token
func (t *token) getAcFunToken() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getAcFunToken() error: %w", err)
		}
	}()

	client := &httpClient{
		url:    t.livePage,
		method: "GET",
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	// 获取did（device ID）
	didCookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(didCookie)
	didCookie.SetKey("_did")
	if !resp.Header.Cookie(didCookie) {
		panic("无法获取didCookie")
	}
	deviceID := string(didCookie.Value())

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	if len(t.cookies) != 0 {
		form.Set(sid, midground)
		cookies := make([]*fasthttp.Cookie, len(t.cookies))
		for i, c := range t.cookies {
			cookie := fasthttp.AcquireCookie()
			defer fasthttp.ReleaseCookie(cookie)
			err = cookie.Parse(c)
			checkErr(err)
			cookies[i] = cookie
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
	var serviceToken, securityKey string
	if len(t.cookies) != 0 {
		securityKey = string(v.GetStringBytes("ssecurity"))
		serviceToken = string(v.GetStringBytes(midgroundSt))
	} else {
		securityKey = string(v.GetStringBytes("acSecurity"))
		serviceToken = string(v.GetStringBytes(visitorSt))
	}

	t.userID = userID
	t.securityKey = securityKey
	t.serviceToken = serviceToken
	t.deviceID = deviceID

	return nil
}

// 获取直播间的token
func (t *token) getLiveToken() (stream StreamInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveToken() error: %w", err)
		}
	}()

	if t.uid == 0 {
		return stream, nil
	}

	var play string
	if len(t.cookies) != 0 {
		// 需要userId、deviceID和serviceToken
		play = fmt.Sprintf(playURL, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		play = fmt.Sprintf(playURL, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	// authorId就是主播的uid
	form.Set("authorId", strconv.FormatInt(t.uid, 10))
	form.Set("pullStreamType", "FLV")
	client := &httpClient{
		url:         play,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
		referer:     t.livePage, // 会验证 Referer
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
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

	t.liveID = liveID
	t.enterRoomAttach = enterRoomAttach
	t.tickets = tickets
	t.instanceID = 0
	t.sessionKey = ""
	t.seqID = 1
	t.headerSeqID = 1
	t.heartbeatSeqID = 1
	t.ticketIndex = 0

	err = t.getGiftList()
	checkErr(err)

	stream = StreamInfo{
		LiveID:        liveID,
		Title:         string(v.GetStringBytes("caption")),
		LiveStartTime: v.GetInt64("liveStartTime"),
		Panoramic:     v.GetBool("panoramic"),
	}
	videoPlayRes := v.GetStringBytes("videoPlayRes")
	v, err = p.ParseBytes(videoPlayRes)
	checkErr(err)
	stream.StreamName = string(v.GetStringBytes("streamName"))
	representation := v.GetArray("liveAdaptiveManifest", "0", "adaptationSet", "representation")
	stream.StreamList = make([]StreamURL, len(representation))
	for i, r := range representation {
		stream.StreamList[i] = StreamURL{
			URL:         string(r.GetStringBytes("url")),
			Bitrate:     r.GetInt("bitrate"),
			QualityType: string(r.GetStringBytes("qualityType")),
			QualityName: string(r.GetStringBytes("name")),
		}
	}

	return stream, nil
}

// 获取全部token
func (t *token) getToken() (stream StreamInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getToken() error: %w", err)
		}
	}()

	err := t.getAcFunToken()
	checkErr(err)
	stream, err = t.getLiveToken()
	checkErr(err)

	return stream, nil
}

// 获取礼物列表
func (t *token) getGiftList() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("updateGiftList() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(giftURL, nil)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播间的礼物列表失败，响应为 %s", string(body)))
	}

	t.gifts = updateGiftList(v)

	return nil
}

// 返回礼物数据
func updateGiftList(v *fastjson.Value) map[int64]GiftDetail {
	list := v.GetArray("data", "giftList")
	gifts := make(map[int64]GiftDetail, len(list))
	for _, gift := range list {
		o := gift.GetObject()
		g := GiftDetail{}
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "giftId":
				g.GiftID = v.GetInt64()
			case "giftName":
				g.GiftName = string(v.GetStringBytes())
			case "arLiveName":
				g.ARLiveName = string(v.GetStringBytes())
			case "payWalletType":
				g.PayWalletType = v.GetInt()
			case "giftPrice":
				g.Price = v.GetInt()
			case "webpPicList":
				g.WebpPic = string(v.GetStringBytes("0", "url"))
			case "pngPicList":
				g.PngPic = string(v.GetStringBytes("0", "url"))
			case "smallPngPicList":
				g.SmallPngPic = string(v.GetStringBytes("0", "url"))
			case "allowBatchSendSizeList":
				list := v.GetArray()
				g.AllowBatchSendSizeList = make([]int, len(list))
				for i, num := range list {
					g.AllowBatchSendSizeList[i] = num.GetInt()
				}
			case "canCombo":
				g.CanCombo = v.GetBool()
			case "canDraw":
				g.CanDraw = v.GetBool()
			case "magicFaceId":
				g.MagicFaceID = v.GetInt()
			case "description":
				g.Description = string(v.GetStringBytes())
			case "redpackPrice":
				g.RedpackPrice = v.GetInt()
			default:
				log.Printf("礼物列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
		gifts[g.GiftID] = g
	}

	return gifts
}

// 通过快手API获取数据，form为nil时采用默认form，调用后需要 defer fasthttp.ReleaseResponse(resp)
func (t *token) fetchKuaiShouAPI(url string, form *fasthttp.Args) (*fasthttp.Response, error) {
	var apiURL string
	if len(t.cookies) != 0 {
		apiURL = fmt.Sprintf(url, t.userID, t.deviceID, midgroundSt, t.serviceToken)
	} else {
		apiURL = fmt.Sprintf(url, t.userID, t.deviceID, visitorSt, t.serviceToken)
	}

	if form == nil {
		form = fasthttp.AcquireArgs()
		defer fasthttp.ReleaseArgs(form)
		form.Set("visitorId", strconv.FormatInt(t.userID, 10))
		form.Set("liveId", t.liveID)
	}
	client := &httpClient{
		url:         apiURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
		referer:     t.livePage,
	}
	return client.doRequest()
}
