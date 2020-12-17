package acfundanmu

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// LiveTypeList 就是直播分类
type LiveTypeList struct {
	ID           int    `json:"id"`           // 直播次分类ID
	Name         string `json:"name"`         // 直播次分类名字
	CategoryID   int    `json:"categoryID"`   // 直播主分类ID
	CategoryName string `json:"categoryName"` // 直播主分类名字
}

// OBSConfig 就是OBS的推流设置
type OBSConfig struct {
	StreamName        string   `json:"streamName"`        // 直播源名字（ID）
	StreamPullAddress string   `json:"streamPullAddress"` // 拉流地址，也就是直播源地址
	StreamPushAddress []string `json:"streamPushAddress"` // 推流地址，分为阿里云和腾讯云两种
	Panoramic         bool     `json:"panoramic"`         // 是否全景直播
	Interval          int64    `json:"interval"`          // 发送transcodeInfo的时间间隔，单位为毫秒
}

// 检测开播权限
func (t *token) checkLiveAuth() (canLive bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("checkLiveAuth() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("检测开播权限需要登陆AcFun帐号"))
	}

	cookies := make([]*fasthttp.Cookie, len(t.cookies))
	for i, c := range t.cookies {
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)
		err := cookie.Parse(c)
		checkErr(err)
		cookies[i] = cookie
	}
	client := &httpClient{
		url:     checkLiveAuthURL,
		method:  "POST",
		cookies: cookies,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("检测开播权限失败，响应为 %s", string(body)))
	}

	status := v.GetInt("authority", "status")
	if status != 3 {
		return false, nil
	}
	return true, nil
}

// 获取直播分类
func (t *token) getLiveTypeList() (list []LiveTypeList, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveTypeList() error: %w", err)
		}
	}()

	client := &httpClient{
		url:    liveTypeListURL,
		method: "POST",
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取直播分类失败，响应为 %s", string(body)))
	}

	typeList := v.GetArray("typeList")
	list = make([]LiveTypeList, len(typeList))
	for i, l := range typeList {
		o := l.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "id":
				list[i].ID = v.GetInt()
			case "name":
				list[i].Name = string(v.GetStringBytes())
			case "categoryId":
				list[i].CategoryID = v.GetInt()
			case "categoryName":
				list[i].CategoryName = string(v.GetStringBytes())
			}
		})
	}

	return list, nil
}

// 获取OBS推流设置
func (t *token) getOBSConfig() (config *OBSConfig, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getOBSConfig() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("获取OBS推流设置需要登陆AcFun帐号"))
	}

	resp, err := t.fetchKuaiShouAPI(obsConfigURL, nil, false)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取OBS推流设置失败，响应为 %s", string(body)))
	}

	config = &OBSConfig{}
	o := v.GetObject("data")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "streamName":
			config.StreamName = string(v.GetStringBytes())
		case "streamPullAddress":
			config.StreamPullAddress = string(v.GetStringBytes())
		case "streamPushAddress":
			list := v.GetArray()
			config.StreamPushAddress = make([]string, len(list))
			for i, l := range list {
				config.StreamPushAddress[i] = string(l.GetStringBytes())
			}
		case "panoramic":
			config.Panoramic = v.GetBool()
		case "intervalMillis":
			config.Interval = v.GetInt64()
		}
	})

	return config, nil
}

// CheckLiveAuth 检测登陆帐号是否有直播权限，可以调用Init(0, cookies)，不需要调用StartDanmu()
func (ac *AcFunLive) CheckLiveAuth() (bool, error) {
	return ac.t.checkLiveAuth()
}

// GetLiveTypeList 获取直播分类，可以调用Init(0, nil)，不需要调用StartDanmu()
func (ac *AcFunLive) GetLiveTypeList() ([]LiveTypeList, error) {
	return ac.t.getLiveTypeList()
}

// GetOBSConfig 获取OBS推流设置，可以调用Init(0, cookies)，不需要调用StartDanmu()
func (ac *AcFunLive) GetOBSConfig() (*OBSConfig, error) {
	return ac.t.getOBSConfig()
}
