package acfundanmu

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"facette.io/natsort"
	"github.com/valyala/fasthttp"
)

const maxIdleConnDuration = 90 * time.Second
const timeout = 10 * time.Second
const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"

var numRune = []rune("0123456789ABCDEF")

// 默认 HTTP 客户端
var defaultClient = &fasthttp.Client{
	MaxIdleConnDuration: maxIdleConnDuration,
	ReadTimeout:         timeout,
	WriteTimeout:        timeout,
}

// 生成随机数字的字符串
func genRandomNum() string {
	num := rand.Intn(1e9)
	chars := make([]rune, 7)
	for i := range chars {
		chars[i] = numRune[rand.Intn(len(numRune))]
	}

	return fmt.Sprintf("%d%s", num, string(chars))
}

// HTTP 客户端
type httpClient struct {
	client      *fasthttp.Client
	url         string
	body        []byte
	method      string
	cookies     Cookies
	contentType string
	referer     string
	userAgent   string
	deviceID    string
	noReqID     bool
}

// 完成 http 请求，调用后需要 defer fasthttp.ReleaseResponse(resp)
func (c *httpClient) doRequest() (resp *fasthttp.Response, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("doRequest() error: %v", err)
			fasthttp.ReleaseResponse(resp)
		}
	}()

	if c.client == nil {
		c.client = defaultClient
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp = fasthttp.AcquireResponse()

	if c.url != "" {
		req.SetRequestURI(c.url)
	} else {
		panic(fmt.Errorf("请求的 url 不能为空"))
	}

	if len(c.body) != 0 {
		req.SetBody(c.body)
	}

	if c.method != "" {
		req.Header.SetMethod(c.method)
	} else {
		// 默认为 GET
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

	var referer string
	if c.referer != "" {
		referer = c.referer
	} else {
		// 默认 referer
		referer = liveHost
	}
	req.Header.SetReferer(referer)

	if c.userAgent != "" {
		req.Header.SetUserAgent(c.userAgent)
	} else {
		req.Header.SetUserAgent(userAgent)
	}

	if c.deviceID != "" {
		// 设置 did 的 cookie，否则可能会被反爬
		req.Header.SetCookie("_did", c.deviceID)
	}

	if !c.noReqID {
		reqID := fmt.Sprintf("%s_self_%x", genRandomNum(), md5.Sum([]byte(referer)))
		req.Header.SetCookie("cur_req_id", reqID)
		req.Header.SetCookie("cur_group_id", reqID+"_0")
	}

	req.Header.Set("Accept-Encoding", "gzip")

	err := c.client.Do(req, resp)
	checkErr(err)

	return resp, nil
}

// http 请求，返回响应 body
func (c *httpClient) request() (body []byte, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("request() error: %v", err)
		}
	}()

	resp, err := c.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	return getBody(resp), nil
}

// http 请求，返回响应 body 和 cookies
func (c *httpClient) getCookies() (body []byte, cookies Cookies, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getCookies() error: %v", err)
		}
	}()

	resp, err := c.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)

	resp.Header.VisitAllCookie(func(key, value []byte) {
		cookie := fasthttp.AcquireCookie()
		err = cookie.ParseBytes(value)
		checkErr(err)
		cookies = append(cookies, cookie)
	})

	return getBody(resp), cookies, nil
}

// 通过快手 API 获取数据，form 为 nil 时采用默认 form，sign 为 true 时会对请求签名
func (t *token) fetchKuaiShouAPI(url string, form *fasthttp.Args, sign bool) (body []byte, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("fetchKuaiShouAPI() error: %v", err)
		}
	}()

	var apiURL string
	if len(t.Cookies) != 0 {
		apiURL = fmt.Sprintf(url, t.UserID, t.DeviceID, midgroundSt, t.ServiceToken)
	} else {
		apiURL = fmt.Sprintf(url, t.UserID, t.DeviceID, visitorSt, t.ServiceToken)
	}

	if form == nil {
		form = t.defaultForm(t.liveID)
		defer fasthttp.ReleaseArgs(form)
	}
	if sign {
		clientSign, err := t.genClientSign(apiURL, form)
		checkErr(err)
		form.Set("__clientSign", clientSign)
	}
	client := &httpClient{
		url:         apiURL,
		body:        form.QueryString(),
		method:      "POST",
		contentType: formContentType,
		referer:     t.livePage,
		noReqID:     true,
	}

	return client.request()
}

// 默认 form，调用后需要 defer fasthttp.ReleaseArgs(form)
func (t *token) defaultForm(liveID string) *fasthttp.Args {
	form := fasthttp.AcquireArgs()
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	form.Set("liveId", liveID)
	return form
}

// 获取响应 body
func getBody(resp *fasthttp.Response) []byte {
	if string(resp.Header.Peek("content-encoding")) == "gzip" || string(resp.Header.Peek("Content-Encoding")) == "gzip" {
		body, err := resp.BodyGunzip()
		if err == nil {
			return body
		}
	}

	body := append([]byte{}, resp.Body()...)

	return body
}

// 生成 client sign
func (t *token) genClientSign(url string, form *fasthttp.Args) (clientSign string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("genClientSign() error: %v", err)
		}
	}()

	uri := fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(uri)
	err := uri.Parse(nil, []byte(url))
	checkErr(err)
	path := string(uri.Path())
	urlParams := uri.QueryArgs()
	var paramsStr []string
	if urlParams != nil {
		// 应该要忽略以__开头的 key
		urlParams.VisitAll(func(key, value []byte) {
			paramsStr = append(paramsStr, string(key)+"="+string(value))
		})
	}
	if form != nil {
		form.VisitAll(func(key, value []byte) {
			paramsStr = append(paramsStr, string(key)+"="+string(value))
		})
	}
	// 实际上这里应该要比较 key
	natsort.Sort(paramsStr)

	minute := time.Now().Unix() / 60
	randomNum := rand.Int31()
	var nonce int64 = minute | (int64(randomNum) << 32)
	nonceStr := strconv.FormatInt(nonce, 10)

	key, err := base64.StdEncoding.DecodeString(t.SecurityKey)
	checkErr(err)
	needSigned := "POST&" + path + "&" + strings.Join(paramsStr, "&") + "&" + nonceStr
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(needSigned))
	hashed := mac.Sum(nil)

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, nonce)
	checkErr(err)
	signedBytes := buf.Bytes()
	signedBytes = append(signedBytes, hashed...)
	clientSign = base64.RawURLEncoding.EncodeToString(signedBytes)

	return clientSign, nil
}

// FetchKuaiShouAPI 获取快手 API 的响应，测试用
func (ac *AcFunLive) FetchKuaiShouAPI(url string, form *fasthttp.Args, sign bool) (body []byte, e error) {
	return ac.t.fetchKuaiShouAPI(url, form, sign)
}
