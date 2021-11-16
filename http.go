package acfundanmu

import (
	"bytes"
	"crypto/hmac"
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
const timeOut = 10 * time.Second

type httpClient struct {
	url         string
	body        []byte
	method      string
	cookies     Cookies
	contentType string
	referer     string
}

var defaultClient = &fasthttp.Client{
	MaxIdleConnDuration: maxIdleConnDuration,
	ReadTimeout:         timeOut,
	WriteTimeout:        timeOut,
}

// 完成http请求，调用后需要 defer fasthttp.ReleaseResponse(resp)
func (c *httpClient) doRequest() (resp *fasthttp.Response, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("doRequest() error: %v", err)
			fasthttp.ReleaseResponse(resp)
		}
	}()

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp = fasthttp.AcquireResponse()

	if c.url != "" {
		req.SetRequestURI(c.url)
	} else {
		panic(fmt.Errorf("请求的url不能为空"))
	}

	if len(c.body) != 0 {
		req.SetBody(c.body)
		//req.Header.SetContentLength(len(c.body))
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

	req.Header.Set("Accept-Encoding", "gzip")

	err := defaultClient.Do(req, resp)
	checkErr(err)

	return resp, nil
}

// http请求，返回响应body
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

// http请求，返回响应body和cookies
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

// 通过快手API获取数据，form为nil时采用默认form，sign为true时会对请求签名
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
	}

	return client.request()
}

// 默认form，调用后需要 defer fasthttp.ReleaseArgs(form)
func (t *token) defaultForm(liveID string) *fasthttp.Args {
	form := fasthttp.AcquireArgs()
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	form.Set("liveId", liveID)
	return form
}

// 获取响应body
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

// 生成client sign
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
		// 应该要忽略以__开头的key
		urlParams.VisitAll(func(key, value []byte) {
			paramsStr = append(paramsStr, string(key)+"="+string(value))
		})
	}
	if form != nil {
		form.VisitAll(func(key, value []byte) {
			paramsStr = append(paramsStr, string(key)+"="+string(value))
		})
	}
	// 实际上这里应该要比较key
	natsort.Sort(paramsStr)

	minute := time.Now().Unix() / 60
	rand.Seed(time.Now().UnixNano())
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

// FetchKuaiShouAPI 获取快手API的响应，测试用
func (ac *AcFunLive) FetchKuaiShouAPI(url string, form *fasthttp.Args, sign bool) (body []byte, e error) {
	return ac.t.fetchKuaiShouAPI(url, form, sign)
}
