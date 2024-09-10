package acfundanmu

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"go.uber.org/atomic"
)

// 登陆 acfun 账号
func login(account, password string) (cookies Cookies, e error) {
	defer func() {
		if err := recover(); err != nil {
			cookies = nil
			e = fmt.Errorf("login() error: %v", err)
		}
	}()

	if account == "" || password == "" {
		panic(fmt.Errorf("AcFun 帐号邮箱或密码为空，无法登陆"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("username", account)
	form.Set("password", password)
	form.Set("key", "")
	form.Set("captcha", "")

	client := &httpClient{
		url:         acfunSignInURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
	}
	body, cookies, err := client.getCookies()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("以注册用户的身份登陆AcFun失败，响应为 %s", string(body)))
	}

	return cookies, nil
}

func loginWithQRCode(qrCodeCallback func(QRCode), scannedCallback func()) (cookies Cookies, e error) {
	defer func() {
		if err := recover(); err != nil {
			cookies = nil
			e = fmt.Errorf("loginWithQRCode() error: %v", err)
		}
	}()

	t := time.Now().UnixMilli()
	client := &httpClient{url: fmt.Sprintf(startScanQRURL, t), method: "GET", noReqID: true}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取登陆二维码失败，响应为 %s", string(body)))
	}

	expireTime := v.GetInt64("expireTime")
	if expireTime <= 0 {
		panic(fmt.Errorf("获取登陆二维码失效时间失败，响应为 %s", string(body)))
	}
	qrLoginSignature := string(v.GetStringBytes("qrLoginSignature"))
	if len(qrLoginSignature) == 0 {
		panic(fmt.Errorf("获取qrLoginSignature失败，响应为 %s", string(body)))
	}
	imageData := string(v.GetStringBytes("imageData"))
	if len(imageData) == 0 {
		panic(fmt.Errorf("获取imageData失败，响应为 %s", string(body)))
	}
	qrLoginToken := string(v.GetStringBytes("qrLoginToken"))
	if len(qrLoginToken) == 0 {
		panic(fmt.Errorf("获取qrLoginToken失败，响应为 %s", string(body)))
	}
	qrCodeCallback(QRCode{ExpireTime: t + expireTime, ImageData: imageData})

	// 留 10 秒多余的时间
	expired := time.Duration((expireTime/1000 + 10) * int64(time.Second))
	fasthttpClient := &fasthttp.Client{MaxIdleConnDuration: expired, ReadTimeout: expired, WriteTimeout: expired}
	t = time.Now().UnixMilli()
	client = &httpClient{
		client:  fasthttpClient,
		url:     fmt.Sprintf(scanQRResultURL, qrLoginToken, qrLoginSignature, t),
		method:  "GET",
		noReqID: true}
	body, err = client.request()
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		if v.GetInt("result") == 100400002 && string(v.GetStringBytes("error_msg")) == "token expired" {
			return nil, nil
		}

		panic(fmt.Errorf("获取登陆二维码扫描状态失败，响应为 %s", string(body)))
	}
	if string(v.GetStringBytes("status")) != "SCANNED" {
		panic(fmt.Errorf("获取登陆二维码扫描状态失败，响应为 %s", string(body)))
	}

	qrLoginSignature = string(v.GetStringBytes("qrLoginSignature"))
	if len(qrLoginSignature) == 0 {
		panic(fmt.Errorf("获取qrLoginSignature失败，响应为 %s", string(body)))
	}
	scannedCallback()

	t = time.Now().UnixMilli()
	client = &httpClient{
		client:  fasthttpClient,
		url:     fmt.Sprintf(acceptQRResultURL, qrLoginToken, qrLoginSignature, t),
		method:  "GET",
		noReqID: true}
	body, cookies, err = client.getCookies()
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		result := v.GetInt("result")
		errorMsg := string(v.GetStringBytes("error_msg"))
		if result == 100400004 && errorMsg == "token canceled" {
			return nil, nil
		}
		if result == 100400002 && errorMsg == "token expired" {
			return nil, nil
		}

		panic(fmt.Errorf("扫描二维码登陆失败，响应为 %s", string(body)))
	}
	if string(v.GetStringBytes("status")) != "ACCEPTED" {
		panic(fmt.Errorf("扫描二维码登陆失败，响应为 %s", string(body)))
	}

	return cookies, nil
}

// 获取 AcFun 帐号的 token
func (t *token) getAcFunToken() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getAcFunToken() error: %v", err)
		}
	}()

	deviceID, err := getDeviceID()
	checkErr(err)
	t.DeviceID = deviceID

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	var client *httpClient
	if len(t.Cookies) != 0 {
		form.Set(sid, midground)
		client = &httpClient{
			url:     getTokenURL,
			body:    form.QueryString(),
			cookies: t.Cookies,
			referer: t.livePage,
		}
	} else {
		form.Set(sid, visitor)
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)
		cookie.SetKey("_did")
		cookie.SetValue(t.DeviceID)
		client = &httpClient{
			url:     loginURL,
			body:    form.QueryString(),
			cookies: []*fasthttp.Cookie{cookie},
			referer: t.livePage,
		}
	}
	client.method = "POST"
	client.contentType = formContentType
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取AcFun token失败，响应为 %s", string(body)))
	}

	// 获取 userId 和对应的令牌
	userID := v.GetInt64("userId")
	var serviceToken, securityKey string
	if len(t.Cookies) != 0 {
		securityKey = string(v.GetStringBytes("ssecurity"))
		serviceToken = string(v.GetStringBytes(midgroundSt))
	} else {
		securityKey = string(v.GetStringBytes("acSecurity"))
		serviceToken = string(v.GetStringBytes(visitorSt))
	}

	t.UserID = userID
	t.SecurityKey = securityKey
	t.ServiceToken = serviceToken

	return nil
}

// 获取直播间的 token
func (t *token) getLiveToken() (stream StreamInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveToken() error: %v", err)
		}
	}()

	if t.liverUID == 0 {
		return stream, nil
	}

	var play string
	if len(t.Cookies) != 0 {
		// 需要 userId、deviceID 和 serviceToken
		play = fmt.Sprintf(playURL, t.UserID, t.DeviceID, midgroundSt, t.ServiceToken)
	} else {
		play = fmt.Sprintf(playURL, t.UserID, t.DeviceID, visitorSt, t.ServiceToken)
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	// authorId 就是主播的 uid
	form.Set("authorId", strconv.FormatInt(t.liverUID, 10))
	form.Set("pullStreamType", "FLV")
	client := &httpClient{
		url:         play,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
		referer:     t.livePage, // 会验证 Referer
		noReqID:     true,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
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
	t.appID = 0
	t.instanceID = 0
	t.sessionKey = nil
	t.seqID = atomic.NewInt64(1)
	t.headerSeqID = atomic.NewInt64(1)
	t.heartbeatSeqID = 0
	t.ticketIndex = atomic.NewUint32(0)
	t.err = atomic.NewError(nil)

	giftList, err := t.getGiftList(t.liveID)
	checkErr(err)
	t.gifts = giftList

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

// 获取全部 token
func (t *token) getToken() (stream StreamInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getToken() error: %v", err)
		}
	}()

	err := t.getAcFunToken()
	checkErr(err)
	stream, err = t.getLiveToken()
	checkErr(err)

	return stream, nil
}

// 获取设备 ID
func getDeviceID() (devideID string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getDeviceID() error: %v", err)
		}
	}()

	client := &httpClient{
		url:    liveHost,
		method: "GET",
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	didCookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(didCookie)
	didCookie.SetKey("_did")
	if !resp.Header.Cookie(didCookie) {
		panic("无法获取 didCookie")
	}

	return string(didCookie.Value()), nil
}

// 获取礼物列表
func (t *token) getGiftList(liveID string) (giftList map[int64]GiftDetail, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getGiftList() error: %v", err)
		}
	}()

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(giftURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播间的礼物列表失败，响应为 %s", string(body)))
	}

	return updateGiftList(v), nil
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
			case "vupArId":
				g.VupArID = v.GetInt()
			case "description":
				g.Description = string(v.GetStringBytes())
			case "redpackPrice":
				g.RedpackPrice = v.GetInt()
			case "cornerMarkerText":
				g.CornerMarkerText = string(v.GetStringBytes())
			default:
				log.Printf("礼物列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
		gifts[g.GiftID] = g
	}

	return gifts
}
