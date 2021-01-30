package acfundanmu

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// LiveType 就是直播分类
type LiveType struct {
	CategoryID      int    `json:"categoryID"`      // 直播主分类ID
	CategoryName    string `json:"categoryName"`    // 直播主分类名字
	SubCategoryID   int    `json:"subCategoryID"`   // 直播次分类ID
	SubCategoryName string `json:"subCategoryName"` // 直播次分类名字
}

// OBSConfig 就是OBS的推流设置
type OBSConfig struct {
	StreamName        string   `json:"streamName"`        // 直播源名字（ID）
	StreamPullAddress string   `json:"streamPullAddress"` // 拉流地址，也就是直播源地址
	StreamPushAddress []string `json:"streamPushAddress"` // 推流地址，目前分为阿里云和腾讯云两种
	Panoramic         bool     `json:"panoramic"`         // 是否全景直播
	Interval          int64    `json:"interval"`          // 发送TranscodeInfo的时间间隔，单位为毫秒
	RTMPServer        string   `json:"rtmpServer"`        // RTMP服务器
	StreamKey         string   `json:"streamKey"`         // 直播码/串流密钥
}

// OBSStatus 就是OBS直播状态
type OBSStatus struct {
	LiveID        string `json:"liveID"`        // 直播ID
	StreamName    string `json:"streamName"`    // 直播源名字
	Title         string `json:"title"`         // 直播间标题
	LiveCover     string `json:"liveCover"`     // 直播间封面
	LiveStartTime int64  `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的Unix时间
	Panoramic     bool   `json:"panoramic"`     // 是否全景直播
	BizUnit       string `json:"bizUnit"`       // 通常是"acfun"
	BizCustomData string `json:"bizCustomData"` // 直播分类，格式是json
}

// QiniuToken 就是七牛云上传的token
type QiniuToken struct {
	URL         string `json:"url"`         // 图片上传后图片地址的域名
	UploadToken string `json:"uploadToken"` // 七牛云上传token
}

// TranscodeInfo 就是转码信息
type TranscodeInfo struct {
	StreamURL  `json:"streamURL"`
	Resolution string `json:"resolution"` // 直播视频分辨率
	FrameRate  int    `json:"frameRate"`  // 直播视频FPS？
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
			e = fmt.Errorf("checkLiveAuth() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("检测开播权限需要登陆AcFun帐号"))
	}

	client := &httpClient{
		url:     checkLiveAuthURL,
		method:  "POST",
		cookies: t.Cookies,
	}
	body, err := client.request()
	checkErr(err)

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
			e = fmt.Errorf("getLiveTypeList() error: %w", err)
		}
	}()

	client := &httpClient{
		url:    liveTypeListURL,
		method: "POST",
	}
	body, err := client.request()
	checkErr(err)

	var p fastjson.Parser
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

// 获取OBS推流设置
func (t *token) getOBSConfig() (config *OBSConfig, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getOBSConfig() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取OBS推流设置需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(obsConfigURL, form, false)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取OBS推流设置失败，响应为 %s", string(body)))
	}

	config = new(OBSConfig)
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
			log.Printf("OBS推流设置里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	i := strings.LastIndex(config.StreamPushAddress[0], `/`)
	config.RTMPServer = config.StreamPushAddress[0][:i]
	config.StreamKey = config.StreamPushAddress[0][i+1:]

	return config, nil
}

// 获取OBS直播状态
func (t *token) getOBSStatus() (status *OBSStatus, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getOBSStatus() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取OBS直播状态需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(obsStatusURL, form, false)
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取OBS直播状态失败，响应为 %s", string(body)))
	}

	status = new(OBSStatus)
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
			log.Printf("OBS直播状态里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return status, nil
}

// 获取七牛云上传token
func (t *token) getQiniuToken() (token *QiniuToken, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getQiniuToken() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取七牛云上传token需要登陆AcFun帐号"))
	}

	client := &httpClient{
		url:     getQiniuTokenURL,
		method:  "POST",
		cookies: t.Cookies,
	}
	body, err := client.request()
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取七牛云上传token失败，响应为 %s", string(body)))
	}

	token = new(QiniuToken)
	o := v.GetObject("info")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "url":
			token.URL = string(v.GetStringBytes())
		case "upToken":
			src := v.GetStringBytes()
			dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
			n, err := base64.StdEncoding.Decode(dst, src)
			checkErr(err)
			s := string(dst[:n])
			i := strings.Index(s, ":")
			token.UploadToken = s[i+1:]
		default:
			log.Printf("七牛云上传token里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return token, nil
}

func (token *QiniuToken) uploadImage(file string) (fileURL string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("uploadImage() error: %w", err)
		}
	}()

	key := uuid.NewString() + filepath.Ext(file)
	cfg := &storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      true,
		UseCdnDomains: true,
	}
	formUploader := storage.NewFormUploader(cfg)
	ret := &storage.PutRet{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := formUploader.PutFile(ctx, ret, token.UploadToken, key, file, nil)
	checkErr(err)

	return token.URL + `/` + ret.Key, nil
}

// 获取转码信息
func (t *token) getTranscodeInfo(streamName string) (info []TranscodeInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getTranscodeInfo() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取转码信息需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("streamName", streamName)
	body, err := t.fetchKuaiShouAPI(transcodeInfoURL, form, false)
	checkErr(err)

	var p fastjson.Parser
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

func fileFormData(file string) (data []byte, contentType string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("fileFormData() error: %w", err)
		}
	}()

	f, err := os.Open(file)
	checkErr(err)
	defer f.Close()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	//defer w.Close()
	fw, err := w.CreateFormFile("cover", filepath.Base(file))
	checkErr(err)
	_, err = io.Copy(fw, f)
	checkErr(err)
	err = w.Close()
	checkErr(err)

	return buf.Bytes(), w.FormDataContentType(), nil
}

func pushQuery(title string, liveType *LiveType) (query string) {
	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)
	if title != "" {
		args.Set("caption", title)
	}
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
func (t *token) startLive(title, coverFile, streamName string, isPanoramic bool, liveType *LiveType) (liveID string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("startLive() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("启动直播需要登陆AcFun帐号"))
	}

	var data []byte
	contentType := formContentType
	var err error
	if coverFile != "" {
		data, contentType, err = fileFormData(coverFile)
		checkErr(err)
	}
	query := pushQuery(title, liveType)

	uri := fmt.Sprintf(startPushURL, t.UserID, t.DeviceID, t.ServiceToken, streamName, isPanoramic) + query
	client := &httpClient{
		url:         uri,
		body:        data,
		method:      "POST",
		contentType: contentType,
		referer:     t.livePage,
	}
	body, err := client.request()
	checkErr(err)

	var p fastjson.Parser
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
			e = fmt.Errorf("stopLive() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("停止直播需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("liveId", liveID)
	body, err := t.fetchKuaiShouAPI(stopPushURL, form, false)
	checkErr(err)

	var p fastjson.Parser
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
			e = fmt.Errorf("changeTitleAndCover() error: %w", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("更改直播间标题和封面需要登陆AcFun帐号"))
	}

	var data []byte
	contentType := formContentType
	var err error
	if coverFile != "" {
		data, contentType, err = fileFormData(coverFile)
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
	}
	body, err := client.request()
	checkErr(err)

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("更改直播间标题和封面失败，响应为 %s", string(body)))
	}

	return nil
}

// CheckLiveAuth 检测登陆帐号是否有直播权限，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) CheckLiveAuth() (bool, error) {
	return ac.t.checkLiveAuth()
}

// GetLiveTypeList 返回直播分类列表，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) GetLiveTypeList() ([]LiveType, error) {
	return ac.t.getLiveTypeList()
}

// GetOBSConfig 返回OBS推流设置，需要登陆主播的AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) GetOBSConfig() (*OBSConfig, error) {
	return ac.t.getOBSConfig()
}

// GetOBSStatus 返回OBS直播状态，需要登陆主播的AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) GetOBSStatus() (*OBSStatus, error) {
	return ac.t.getOBSStatus()
}

// GetQiniuToken 返回七牛云上传token，需要登陆AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) GetQiniuToken() (*QiniuToken, error) {
	return ac.t.getQiniuToken()
}

// UploadImage 上传图片，file为图片的路径，返回图片链接fileURL
func (token *QiniuToken) UploadImage(file string) (fileURL string, e error) {
	return token.uploadImage(file)
}

// GetTranscodeInfo 返回转码信息，OBS推流后调用，返回的info长度不为0说明推流成功，需要登陆AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) GetTranscodeInfo(streamName string) ([]TranscodeInfo, error) {
	return ac.t.getTranscodeInfo(streamName)
}

// StartLive 启动直播，title为直播间标题，coverFile为直播间封面图片（可以是gif）的路径，isPanoramic为是否全景直播，推流成功服务器开始转码后调用
// 需要登陆AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) StartLive(title, coverFile, streamName string, isPanoramic bool, liveType *LiveType) (liveID string, e error) {
	return ac.t.startLive(title, coverFile, streamName, isPanoramic, liveType)
}

// StopLive 停止直播，需要登陆AcFun帐号，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) StopLive(liveID string) (*StopPushInfo, error) {
	return ac.t.stopLive(liveID)
}

// ChangeTitleAndCover 更改直播间标题和封面，coverFile为直播间封面图片（可以是gif）的路径，不需要设置主播uid，不需要调用StartDanmu()
func (ac *AcFunLive) ChangeTitleAndCover(title, coverFile, liveID string) error {
	return ac.t.changeTitleAndCover(title, coverFile, liveID)
}
