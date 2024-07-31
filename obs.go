package acfundanmu

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// LiveType 就是直播分类
type LiveType struct {
	CategoryID      int    `json:"categoryID"`      // 直播主分类 ID
	CategoryName    string `json:"categoryName"`    // 直播主分类名字
	SubCategoryID   int    `json:"subCategoryID"`   // 直播次分类 ID
	SubCategoryName string `json:"subCategoryName"` // 直播次分类名字
}

// PushConfig 就是推流设置
type PushConfig struct {
	StreamName        string   `json:"streamName"`        // 直播源名字（ID）
	StreamPullAddress string   `json:"streamPullAddress"` // 拉流地址，也就是直播源地址
	StreamPushAddress []string `json:"streamPushAddress"` // 推流地址，目前分为阿里云和腾讯云两种
	Panoramic         bool     `json:"panoramic"`         // 是否全景直播
	Interval          int64    `json:"interval"`          // 查询转码信息的时间间隔，单位为毫秒
	RTMPServer        string   `json:"rtmpServer"`        // RTMP 服务器
	StreamKey         string   `json:"streamKey"`         // 直播码/串流密钥
}

// LiveStatus 就是直播状态
type LiveStatus struct {
	LiveID        string `json:"liveID"`        // 直播 ID
	StreamName    string `json:"streamName"`    // 直播源名字
	Title         string `json:"title"`         // 直播间标题
	LiveCover     string `json:"liveCover"`     // 直播间封面
	LiveStartTime int64  `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的 Unix 时间
	Panoramic     bool   `json:"panoramic"`     // 是否全景直播
	BizUnit       string `json:"bizUnit"`       // 通常是"acfun"
	BizCustomData string `json:"bizCustomData"` // 直播分类，格式是 json
}

// TranscodeInfo 就是转码信息
type TranscodeInfo struct {
	StreamURL  `json:"streamURL"`
	Resolution string `json:"resolution"` // 直播视频分辨率
	FrameRate  int    `json:"frameRate"`  // 直播视频 FPS？
	Template   string `json:"template"`   // 直播模板？
}

// StopPushInfo 停止直播返回的信息
type StopPushInfo struct {
	Duration  int64  `json:"duration"`  // 直播时长，单位为毫秒
	EndReason string `json:"endReason"` // 停止直播的原因
}

// 检测开播权限
func (t *token) checkLiveAuth() (canLive bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("checkLiveAuth() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("检测开播权限需要登陆主播的 AcFun 帐号"))
	}

	client := &httpClient{
		url:      checkLiveAuthURL,
		method:   "POST",
		cookies:  t.Cookies,
		deviceID: t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
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
func getLiveType(v *fastjson.Value) *LiveType {
	liveType := new(LiveType)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "id":
			liveType.SubCategoryID = v.GetInt()
		case "name":
			liveType.SubCategoryName = string(v.GetStringBytes())
		case "categoryId":
			liveType.CategoryID = v.GetInt()
		case "categoryName":
			liveType.CategoryName = string(v.GetStringBytes())
		default:
			log.Printf("直播分类里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return liveType
}

// 获取直播分类列表
func (t *token) getLiveTypeList() (list []LiveType, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveTypeList() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      liveTypeListURL,
		method:   "POST",
		deviceID: t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取直播分类失败，响应为 %s", string(body)))
	}

	typeList := v.GetArray("typeList")
	list = make([]LiveType, 0, len(typeList))
	for _, l := range typeList {
		list = append(list, *getLiveType(l))
	}

	return list, nil
}

// 获取推流设置
func (t *token) getPushConfig() (config *PushConfig, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getPushConfig() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取推流设置需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(obsConfigURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取推流设置失败，响应为 %s", string(body)))
	}

	config = new(PushConfig)
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
		default:
			log.Printf("推流设置里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	i := strings.LastIndex(config.StreamPushAddress[0], `/`)
	config.RTMPServer = config.StreamPushAddress[0][:i]
	config.StreamKey = config.StreamPushAddress[0][i+1:]

	return config, nil
}

// 获取直播状态
func (t *token) getLiveStatus() (status *LiveStatus, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveStatus() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取直播状态需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(obsStatusURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播状态失败，响应为 %s", string(body)))
	}

	status = new(LiveStatus)
	o := v.GetObject("data")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "liveId":
			status.LiveID = string(v.GetStringBytes())
		case "streamName":
			status.StreamName = string(v.GetStringBytes())
		case "caption":
			status.Title = string(v.GetStringBytes())
		case "cover":
			status.LiveCover = string(v.GetStringBytes("0", "url"))
		case "createTime":
			status.LiveStartTime = v.GetInt64()
		case "panoramic":
			status.Panoramic = v.GetBool()
		case "bizUnit":
			status.BizUnit = string(v.GetStringBytes())
		case "bizCustomData":
			status.BizCustomData = string(v.GetStringBytes())
		default:
			log.Printf("直播状态里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return status, nil
}

// 获取转码信息
func (t *token) getTranscodeInfo(streamName string) (info []TranscodeInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getTranscodeInfo() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取转码信息需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("streamName", streamName)
	body, err := t.fetchKuaiShouAPI(transcodeInfoURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取转码信息状态失败，响应为 %s", string(body)))
	}

	list := v.GetArray("data", "transcodeInfoList")
	info = make([]TranscodeInfo, len(list))
	for i, l := range list {
		o := l.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "pullUrl":
				info[i].URL = string(v.GetStringBytes())
			case "bitRate":
				info[i].Bitrate = v.GetInt()
			case "qualityType":
				info[i].QualityType = string(v.GetStringBytes())
			case "qualityTypeName":
				info[i].QualityName = string(v.GetStringBytes())
			case "resolution":
				info[i].Resolution = string(v.GetStringBytes())
			case "frameRate":
				info[i].FrameRate = v.GetInt()
			case "template":
				info[i].Template = string(v.GetStringBytes())
			default:
				log.Printf("转码信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	return info, nil
}

// 读取文件
func loadFile(file string) (data []byte, contentType string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("loadFile() error: %v", err)
		}
	}()

	var fileData []byte
	if u, err := url.Parse(file); err == nil && u.Scheme != "" && u.Host != "" {
		client := &httpClient{
			url:     file,
			method:  "GET",
			noReqID: true,
		}
		fileData, err = client.request()
		checkErr(err)
	} else {
		fileData, err = os.ReadFile(file)
		checkErr(err)
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	//defer w.Close()
	fw, err := w.CreateFormFile("cover", filepath.Base(file))
	checkErr(err)
	_, err = fw.Write(fileData)
	checkErr(err)
	err = w.Close()
	checkErr(err)

	return buf.Bytes(), w.FormDataContentType(), nil
}

// 推流地址的 query
func pushQuery(title string, liveType *LiveType) (query string) {
	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)
	args.Set("caption", title)
	if liveType != nil {
		args.Set("bizCustomData", fmt.Sprintf(pushType, liveType.SubCategoryID, liveType.CategoryID, liveType.SubCategoryID))
	}
	query = args.String()
	if query != "" {
		query = "&" + query
	}

	return query
}

// 启动直播
func (t *token) startLive(title, coverFile, streamName string, portrait, panoramic bool, liveType *LiveType) (liveID string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("startLive() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("启动直播需要登陆主播的 AcFun 帐号"))
	}

	var data []byte
	contentType := formContentType
	var err error
	if coverFile != "" {
		data, contentType, err = loadFile(coverFile)
		checkErr(err)
	}
	query := pushQuery(title, liveType)

	uri := fmt.Sprintf(startPushURL, t.UserID, t.DeviceID, t.ServiceToken, streamName, portrait, panoramic) + query
	client := &httpClient{
		url:         uri,
		body:        data,
		method:      "POST",
		contentType: contentType,
		referer:     t.livePage,
		noReqID:     true,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("启动直播失败，响应为 %s", string(body)))
	}

	liveID = string(v.GetStringBytes("data", "liveId"))

	return liveID, nil
}

// 停止直播
func (t *token) stopLive(liveID string) (info *StopPushInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("stopLive() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("停止直播需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("liveId", liveID)
	body, err := t.fetchKuaiShouAPI(stopPushURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("停止直播失败，响应为 %s", string(body)))
	}

	info = new(StopPushInfo)
	o := v.GetObject("data")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "durationMs":
			info.Duration = v.GetInt64()
		case "endReason":
			info.EndReason = string(v.GetStringBytes())
		default:
			log.Printf("停止直播信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return info, nil
}

// 更改直播间标题和封面
func (t *token) changeTitleAndCover(title, coverFile, liveID string) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("changeTitleAndCover() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("更改直播间标题和封面需要登陆主播的 AcFun 帐号"))
	}

	var data []byte
	contentType := formContentType
	var err error
	if coverFile != "" {
		data, contentType, err = loadFile(coverFile)
		checkErr(err)
	}
	query := pushQuery(title, nil)

	uri := fmt.Sprintf(changeCoverURL, t.UserID, t.DeviceID, t.ServiceToken, liveID) + query
	client := &httpClient{
		url:         uri,
		body:        data,
		method:      "POST",
		contentType: contentType,
		referer:     t.livePage,
		noReqID:     true,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("更改直播间标题和封面失败，响应为 %s", string(body)))
	}

	return nil
}

// 查询是否允许观众剪辑直播录像
func (t *token) getLiveCutStatus() (canCut bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveCutStatus() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("查询是否允许观众剪辑直播录像需要登陆主播的 AcFun 帐号"))
	}

	client := &httpClient{
		url:      liveCutStatusURL,
		method:   "POST",
		cookies:  t.Cookies,
		deviceID: t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("查询是否允许观众剪辑直播录像失败，响应为 %s", string(body)))
	}

	status := v.GetInt("liveCutStatus")
	if status == 1 {
		return true, nil
	}
	if status == 2 {
		return false, nil
	}
	panic(fmt.Errorf("查询是否允许观众剪辑直播录像失败，响应为 %s", string(body)))
}

// 设置是否允许观众剪辑直播录像
func (t *token) setLiveCutStatus(canCut bool) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("setLiveCutStatus() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("设置是否允许观众剪辑直播录像需要登陆主播的 AcFun 帐号"))
	}

	status := 1
	if !canCut {
		status = 2
	}

	client := &httpClient{
		url:         updateLiveCutURL,
		body:        []byte(fmt.Sprintf(liveCutStatus, status)),
		method:      "POST",
		cookies:     t.Cookies,
		contentType: jsonContentType,
		deviceID:    t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("设置是否允许观众剪辑直播录像失败，响应为 %s", string(body)))
	}

	returnStatus := v.GetInt("liveCutStatus")
	if status != returnStatus {
		panic(fmt.Errorf("设置是否允许观众剪辑直播录像失败，响应为 %s", string(body)))
	}

	return nil
}

// CheckLiveAuth 检测登陆帐号是否有直播权限，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) CheckLiveAuth() (bool, error) {
	return ac.t.checkLiveAuth()
}

// GetLiveTypeList 返回直播分类列表
func (ac *AcFunLive) GetLiveTypeList() ([]LiveType, error) {
	return ac.t.getLiveTypeList()
}

// GetPushConfig 返回推流设置，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetPushConfig() (*PushConfig, error) {
	return ac.t.getPushConfig()
}

// GetLiveStatus 返回直播状态，需要登陆主播的 AcFun 帐号并启动直播后调用
func (ac *AcFunLive) GetLiveStatus() (*LiveStatus, error) {
	return ac.t.getLiveStatus()
}

// GetTranscodeInfo 返回转码信息，推流后调用，返回的 info 长度不为 0 说明推流成功，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetTranscodeInfo(streamName string) ([]TranscodeInfo, error) {
	return ac.t.getTranscodeInfo(streamName)
}

// StartLive 启动直播，title 为直播间标题，coverFile 为直播间封面图片（可以是 gif）的本地路径或网络链接，portrait 为是否手机直播，panoramic 为是否全景直播。
// 推流成功服务器开始转码（用 GetTranscodeInfo() 判断）后调用，title 和 coverFile 不能为空，需要登陆主播的 AcFun 帐号。
func (ac *AcFunLive) StartLive(title, coverFile, streamName string, portrait, panoramic bool, liveType *LiveType) (liveID string, e error) {
	return ac.t.startLive(title, coverFile, streamName, portrait, panoramic, liveType)
}

// StopLive 停止直播，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) StopLive(liveID string) (*StopPushInfo, error) {
	return ac.t.stopLive(liveID)
}

// ChangeTitleAndCover 更改直播间标题和封面，coverFile 为直播间封面图片（可以是 gif）的本地路径或网络链接。
// title 为空时会没有标题，coverFile 为空时只更改标题，需要登陆主播的 AcFun 帐号。
func (ac *AcFunLive) ChangeTitleAndCover(title, coverFile, liveID string) error {
	return ac.t.changeTitleAndCover(title, coverFile, liveID)
}

// GetLiveCutStatus 查询是否允许观众剪辑直播录像，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetLiveCutStatus() (bool, error) {
	return ac.t.getLiveCutStatus()
}

// SetLiveCutStatus 设置是否允许观众剪辑直播录像，需要登陆主播的 AcFun 帐号，主播直播时无法设置
func (ac *AcFunLive) SetLiveCutStatus(canCut bool) error {
	return ac.t.setLiveCutStatus(canCut)
}
