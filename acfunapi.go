package acfundanmu

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"facette.io/natsort"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// WatchingUser 就是观看直播的用户的信息，目前没有Medal
type WatchingUser struct {
	UserInfo                 // 用户信息
	AnonymousUser     bool   // 是否匿名用户
	DisplaySendAmount string // 赠送的全部礼物的价值，单位是AC币，注意不一定是纯数字的字符串
	CustomData        string // 用户的一些额外信息，格式为json
}

// BillboardUser 就是礼物贡献榜上的用户的信息，没有AnonymousUser、Medal和ManagerType
type BillboardUser WatchingUser

// Summary 就是直播的总结信息
type Summary struct {
	LiveDurationMs int64 // 直播时长，单位为毫秒
	LikeCount      int   // 点赞总数
	WatchCount     int   // 观看过直播的人数总数
}

// MedalDetail 就是登陆帐号守护徽章的详细信息
type MedalDetail struct {
	MedalInfo
	UperName           string // UP主的名字
	UperAvatar         string // UP主的头像
	WearMedal          bool   // 是否正在佩戴该守护徽章
	FriendshipDegree   int    // 目前守护徽章的亲密度
	JoinClubTime       int64  // 加入守护团的时间，单位为毫秒
	CurrentDegreeLimit int    // 守护徽章目前等级的亲密度的上限
}

// LuckyUser 就是抢到红包的用户，没有Medal和ManagerType
type LuckyUser struct {
	UserInfo
	GrabAmount int // 抢红包抢到的AC币
}

// Playback 就是直播回放的相关信息
type Playback struct {
	Duration  int64  // 录播视频时长，单位是毫秒
	URL       string // 录播源链接，目前分为阿里云和腾讯云两种，目前阿里云的下载速度比较快
	BackupURL string // 备份录播源链接
	M3U8Slice string // m3u8
	Width     int    // 录播视频宽度
	Height    int    // 录播视频高度
}

var liveListParser fastjson.ParserPool

// 获取直播间排名前50的在线观众信息列表
func (t *token) getWatchingList() (watchingList []WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(watchingListURL, nil)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	p := t.watchParser.Get()
	defer t.watchParser.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取在线观众列表失败，响应为 %s", string(body)))
	}

	watchArray := v.GetArray("data", "list")
	watchingList = make([]WatchingUser, len(watchArray))
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
				w.CustomData = string(v.GetStringBytes())
			case "managerType":
				w.ManagerType = ManagerType(v.GetInt())
			}
		})
		watchingList[i] = w
	}

	return watchingList, nil
}

// 获取直播间最近七日内的礼物贡献榜前50名观众的详细信息
func (t *token) getBillboard() (billboard []BillboardUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getBillboard() error: %w", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("authorId", strconv.FormatInt(t.uid, 10))
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)
	if len(t.cookies) != 0 {
		cookie.SetKey(midgroundSt)
	} else {
		cookie.SetKey(visitorSt)
	}
	cookie.SetValue(t.serviceToken)
	client := &httpClient{
		url:         fmt.Sprintf(billboardURL, t.userID, t.deviceID),
		body:        form.QueryString(),
		method:      "POST",
		cookies:     []*fasthttp.Cookie{cookie},
		contentType: formContentType,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	p := t.watchParser.Get()
	defer t.watchParser.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取最近七日内的礼物贡献榜失败，响应为 %s", string(body)))
	}

	billboardArray := v.GetArray("data", "list")
	billboard = make([]BillboardUser, len(billboardArray))
	for i, user := range billboardArray {
		o := user.GetObject()
		b := BillboardUser{}
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				b.UserID = v.GetInt64()
			case "nickname":
				b.Nickname = string(v.GetStringBytes())
			case "avatar":
				b.Avatar = string(v.GetStringBytes("0", "url"))
			case "displaySendAmount":
				b.DisplaySendAmount = string(v.GetStringBytes())
			case "customData":
				b.CustomData = string(v.GetStringBytes())
			}
		})
		billboard[i] = b
	}

	return billboard, nil
}

// 获取直播总结信息
func (t *token) getSummary() (summary Summary, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getSummary() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(endSummaryURL, nil)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播总结信息失败，响应为 %s", string(body)))
	}

	v = v.Get("data")
	summary.LiveDurationMs = v.GetInt64("liveDurationMs")
	summary.LikeCount, err = strconv.Atoi(string(v.GetStringBytes("likeCount")))
	checkErr(err)
	summary.WatchCount, err = strconv.Atoi(string(v.GetStringBytes("watchCount")))
	checkErr(err)

	return summary, nil
}

// 获取抢到红包的用户列表
func (t *token) getLuckList(redpack Redpack) (luckyList []LuckyUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLuckList() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("获取抢到红包的用户列表需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	form.Set("liveId", t.liveID)
	form.Set("redpackBizUnit", redpack.RedpackBizUnit)
	form.Set("redpackId", redpack.RedPackID)
	resp, err := t.fetchKuaiShouAPI(redpackLuckListURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	p := t.watchParser.Get()
	defer t.watchParser.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取抢到红包的用户列表失败，响应为 %s", string(body)))
	}

	luckyArray := v.GetArray("data", "luckyList")
	luckyList = make([]LuckyUser, len(luckyArray))
	for i, user := range luckyArray {
		v := user.Get("simpleUserInfo")
		luckyList[i] = LuckyUser{
			UserInfo: UserInfo{
				UserID:   v.GetInt64("userId"),
				Nickname: string(v.GetStringBytes("nickname")),
				Avatar:   string(v.GetStringBytes("headPic", "0", "url")),
			},
			GrabAmount: user.GetInt("grabAmount"),
		}
	}

	return luckyList, nil
}

// 获取直播回放的相关信息
func (t *token) getPlayback(liveID string) (playback Playback, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getPlayback() error: %w", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("liveId", liveID)
	url := fmt.Sprintf(playbackURL, t.userID, t.deviceID)
	clientSign, err := t.genClientSign(url, form)
	checkErr(err)
	form.Set("__clientSign", clientSign)
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)
	if len(t.cookies) != 0 {
		cookie.SetKey(midgroundSt)
	} else {
		cookie.SetKey(visitorSt)
	}
	cookie.SetValue(t.serviceToken)
	client := &httpClient{
		url:         url,
		body:        form.QueryString(),
		method:      "POST",
		cookies:     []*fasthttp.Cookie{cookie},
		contentType: formContentType,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播回放的相关信息失败，响应为 %s", string(body)))
	}
	adaptiveManifest := v.GetStringBytes("data", "adaptiveManifest")
	v, err = p.ParseBytes(adaptiveManifest)
	checkErr(err)
	v = v.Get("adaptationSet", "0")
	duration := v.GetInt64("duration")
	v = v.Get("representation", "0")
	playback = Playback{
		Duration:  duration,
		URL:       string(v.GetStringBytes("url")),
		BackupURL: string(v.GetStringBytes("backupUrl", "0")),
		M3U8Slice: string(v.GetStringBytes("m3u8Slice")),
		Width:     v.GetInt("width"),
		Height:    v.GetInt("height"),
	}

	return playback, nil
}

// 获取直播源信息，和getLiveToken()重复了
func (t *token) getPlayURL() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getPlayURL() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(getPlayURL, nil)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播源信息失败，响应为 %s", string(body)))
	}

	//videoPlayRes := string(v.GetStringBytes("data", "videoPlayRes"))

	return nil
}

// 获取全部礼物的数据
func (t *token) getAllGift() (gifts map[int64]GiftDetail, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getAllGift() error: %w", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	resp, err := t.fetchKuaiShouAPI(allGiftURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取全部礼物的数据失败，响应为 %s", string(body)))
	}

	gifts = updateGiftList(v)

	return gifts, nil
}

// 获取钱包里AC币和拥有的香蕉的数量
func (t *token) getWalletBalance() (accoins int, bananas int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getWalletBalance() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("获取钱包里AC币和拥有的香蕉的数量需要登陆AcFun帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.userID, 10))
	resp, err := t.fetchKuaiShouAPI(walletBalanceURL, form)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取拥有的香蕉和钱包里AC币的数量失败，响应为 %s", string(body)))
	}

	v = v.Get("data", "payWalletTypeToBalance")
	accoins = v.GetInt("1")
	bananas = v.GetInt("2")

	return accoins, bananas, nil
}

// 获取主播踢人的历史记录
func (t *token) getAuthorKickHistory() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getAuthorKickHistory() error: %w", err)
		}
	}()

	if len(t.cookies) == 0 {
		panic(fmt.Errorf("获取主播踢人的历史记录需要登陆AcFun帐号"))
	}

	resp, err := t.fetchKuaiShouAPI(authorKickHistoryURL, nil)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取主播踢人的历史记录失败，响应为 %s", string(body)))
	} else {
		log.Printf("获取主播踢人的历史记录的响应为 %s", string(body))
	}

	return nil
}

// 生成client sign
func (t *token) genClientSign(url string, form *fasthttp.Args) (clientSign string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("genClientSign() error: %w", err)
		}
	}()

	uri := fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(uri)
	err := uri.Parse(nil, []byte(url))
	checkErr(err)
	path := string(uri.Path())
	urlParams := uri.QueryArgs()
	var paramsStr []string
	// 应该要忽略以__开头的key
	urlParams.VisitAll(func(key, value []byte) {
		paramsStr = append(paramsStr, string(key)+"="+string(value))
	})
	form.VisitAll(func(key, value []byte) {
		paramsStr = append(paramsStr, string(key)+"="+string(value))
	})
	// 实际上这里应该要比较key
	natsort.Sort(paramsStr)

	minute := time.Now().Unix() / 60
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int31()
	var nonce int64 = minute | (int64(randomNum) << 32)
	nonceStr := strconv.FormatInt(nonce, 10)

	key, err := base64.StdEncoding.DecodeString(t.securityKey)
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

// 获取登陆帐号的守护徽章和指定主播守护徽章的名字
func getMedalInfo(uid int64, cookies []string) (medalList []MedalDetail, clubName string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getMedalInfo() error: %w", err)
		}
	}()

	var httpCookies []*fasthttp.Cookie
	for _, c := range cookies {
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)
		err := cookie.Parse(c)
		checkErr(err)
		httpCookies = append(httpCookies, cookie)
	}
	client := &httpClient{
		url:     fmt.Sprintf(medalInfoURL, uid),
		method:  "GET",
		cookies: httpCookies,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取登陆帐号的守护徽章和指定主播守护徽章的名字失败，响应为 %s", string(body)))
	}

	clubName = string(v.GetStringBytes("clubName"))

	medalArray := v.GetArray("medalList")
	medalList = make([]MedalDetail, len(medalArray))
	for i, medal := range medalArray {
		medalList[i] = MedalDetail{
			MedalInfo: MedalInfo{
				UperID:   medal.GetInt64("uperId"),
				ClubName: string(medal.GetStringBytes("clubName")),
				Level:    medal.GetInt("level"),
			},
			UperName:           string(medal.GetStringBytes("uperName")),
			UperAvatar:         string(medal.GetStringBytes("uperHeadUrl")),
			WearMedal:          medal.GetBool("wearMedal"),
			FriendshipDegree:   medal.GetInt("friendshipDegree"),
			JoinClubTime:       medal.GetInt64("joinClubTime"),
			CurrentDegreeLimit: medal.GetInt("currentDegreeLimit"),
		}
	}

	return medalList, clubName, nil
}

// 获取正在直播的直播间列表
func getLiveList() (body string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveList() error: %w", err)
		}
	}()

	for count := 1000; count < 1e8; count *= 10 {
		client := httpClient{
			url:    fmt.Sprintf(liveListURL, count),
			method: "GET",
		}
		resp, err := client.doRequest()
		checkErr(err)
		defer fasthttp.ReleaseResponse(resp)
		respBody := resp.Body()

		p := liveListParser.Get()
		defer liveListParser.Put(p)
		v, err := p.ParseBytes(respBody)
		checkErr(err)
		v = v.Get("channelListData")
		if !v.Exists("result") || v.GetInt("result") != 0 {
			panic(fmt.Errorf("获取正在直播的直播间列表失败"))
		}
		cursor := string(v.GetStringBytes("pcursor"))
		if cursor == "no_more" {
			body = string(respBody)
			break
		}
	}

	return body, nil
}

// Distinguish 返回阿里云和腾讯云链接，目前阿里云的下载速度比较快
func (pb *Playback) Distinguish() (aliURL, txURL string) {
	switch {
	case strings.Contains(pb.URL, "alivod"):
		aliURL = pb.URL
	case strings.Contains(pb.URL, "txvod"):
		txURL = pb.URL
	default:
	}

	switch {
	case strings.Contains(pb.BackupURL, "alivod"):
		aliURL = pb.BackupURL
	case strings.Contains(pb.BackupURL, "txvod"):
		txURL = pb.BackupURL
	default:
	}

	return aliURL, txURL
}

// GetWatchingList 返回直播间排名前50的在线观众信息列表，不需要调用StartDanmu()
func (dq *DanmuQueue) GetWatchingList() ([]WatchingUser, error) {
	return dq.t.getWatchingList()
}

// GetBillboard 返回直播间最近七日内的礼物贡献榜前50名观众的详细信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetBillboard() ([]BillboardUser, error) {
	return dq.t.getBillboard()
}

// GetSummary 返回直播总结信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetSummary() (Summary, error) {
	return dq.t.getSummary()
}

// GetLuckList 返回抢到红包的用户列表，需要调用Login()登陆AcFun帐号，不需要调用StartDanmu()
func (dq *DanmuQueue) GetLuckList(redpack Redpack) ([]LuckyUser, error) {
	return dq.t.getLuckList(redpack)
}

// GetPlayback 返回直播回放的相关信息，需要liveID，可以调用Init(0, nil)，不需要调用StartDanmu()，目前部分直播没有回放
func (dq *DanmuQueue) GetPlayback(liveID string) (Playback, error) {
	return dq.t.getPlayback(liveID)
}

// GetAllGift 返回全部礼物的数据，可以调用Init(0, nil)，不需要调用StartDanmu()
func (dq *DanmuQueue) GetAllGift() (map[int64]GiftDetail, error) {
	return dq.t.getAllGift()
}

// GetWalletBalance 返回钱包里AC币和拥有的香蕉的数量，需要调用Login()登陆AcFun帐号，可以调用Init(0, cookies)，不需要调用StartDanmu()
func (dq *DanmuQueue) GetWalletBalance() (accoins int, bananas int, e error) {
	return dq.t.getWalletBalance()
}

// GetAuthorKickHistory 返回主播踢人的历史记录，不需要调用StartDanmu()，未测试
func (dq *DanmuQueue) GetAuthorKickHistory() (e error) {
	return dq.t.getAuthorKickHistory()
}

// GetMedalInfo 返回登陆用户的守护徽章列表medalList和uid指定主播的守护徽章的名字clubName，利用Login()获取AcFun帐号的cookies
func GetMedalInfo(uid int64, cookies []string) (medalList []MedalDetail, clubName string, err error) {
	return getMedalInfo(uid, cookies)
}
