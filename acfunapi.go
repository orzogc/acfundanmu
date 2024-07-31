package acfundanmu

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// WatchingUser 就是观看直播的用户的信息，目前没有 Medal
type WatchingUser struct {
	UserInfo          `json:"userInfo"`
	AnonymousUser     bool   `json:"anonymousUser"`     // 是否匿名用户
	DisplaySendAmount string `json:"displaySendAmount"` // 赠送的全部礼物的价值，单位是 AC 币，注意不一定是纯数字的字符串
	CustomData        string `json:"customData"`        // 用户的一些额外信息，格式为 json
}

// BillboardUser 就是礼物贡献榜上的用户的信息，没有 AnonymousUser、Medal 和 ManagerType
type BillboardUser WatchingUser

// Summary 就是直播的总结信息
type Summary struct {
	Duration     int64  `json:"duration"`     // 直播时长，单位为毫秒
	LikeCount    string `json:"likeCount"`    // 点赞总数
	WatchCount   string `json:"watchCount"`   // 观看过直播的人数总数
	GiftCount    int    `json:"giftCount"`    // 直播收到的付费礼物数量
	DiamondCount int    `json:"diamondCount"` // 主播收到的实际钻石数量（扣除平台相关费用），100 钻石=1AC 币
	BananaCount  int    `json:"bananaCount"`  // 直播收到的香蕉数量
}

// Medal 就是指定用户的守护徽章信息
type Medal struct {
	MedalInfo          `json:"medalInfo"`
	UperName           string `json:"uperName"`           // UP 主的名字
	UperAvatar         string `json:"uperAvatar"`         // UP 主的头像
	WearMedal          bool   `json:"wearMedal"`          // 是否正在佩戴该守护徽章
	FriendshipDegree   int    `json:"friendshipDegree"`   // 目前守护徽章的亲密度
	JoinClubTime       int64  `json:"joinClubTime"`       // 加入守护团的时间，是以毫秒为单位的 Unix 时间
	CurrentDegreeLimit int    `json:"currentDegreeLimit"` // 守护徽章目前等级的亲密度的上限
	MedalCount         int    `json:"medalCount"`         // 指定用户拥有的守护徽章数量
}

// MedalDegree 就是守护徽章的亲密度信息
type MedalDegree struct {
	UperID               int64 `json:"uperID"`               // UP 主的 uid
	GiftDegree           int   `json:"giftDegree"`           // 本日送直播礼物增加的亲密度
	GiftDegreeLimit      int   `json:"giftDegreeLimit"`      // 本日送直播礼物增加的亲密度上限
	PeachDegree          int   `json:"peachDegree"`          // 本日投桃增加的亲密度
	PeachDegreeLimit     int   `json:"peachDegreeLimit"`     // 本日投桃增加的亲密度上限
	LiveWatchDegree      int   `json:"liveWatchDegree"`      // 本日看直播时长增加的亲密度
	LiveWatchDegreeLimit int   `json:"liveWatchDegreeLimit"` // 本日看直播时长增加的亲密度上限
	BananaDegree         int   `json:"bananaDegree"`         // 本日投蕉增加的亲密度
	BananaDegreeLimit    int   `json:"bananaDegreeLimit"`    // 本日投蕉增加的亲密度上限
}

// MedalDetail 就是登陆用户的守护徽章的详细信息
type MedalDetail struct {
	Medal       Medal       `json:"medal"`       // 守护徽章信息
	MedalDegree MedalDegree `json:"medalDegree"` // 守护徽章亲密度信息
	UserRank    string      `json:"userRank"`    // 登陆用户的主播守护徽章亲密度的排名
}

// MedalList 就是登陆用户拥有的守护徽章列表
//type MedalList struct {
//	MedalList   []Medal     `json:"medalList"`   // 用户拥有的守护徽章列表
//	MedalDetail MedalDetail `json:"medalDetail"` // 指定主播的守护徽章详细信息
//}

// LuckyUser 就是抢到红包的用户，没有 Medal 和 ManagerType
type LuckyUser struct {
	UserInfo   `json:"userInfo"`
	GrabAmount int `json:"grabAmount"` // 抢红包抢到的 AC 币
}

// Playback 就是直播回放的相关信息
type Playback struct {
	Duration  int64  `json:"duration"`  // 录播视频时长，单位是毫秒
	URL       string `json:"url"`       // 录播源链接，目前分为阿里云和腾讯云两种，目前阿里云的下载速度比较快
	BackupURL string `json:"backupURL"` // 备份录播源链接
	M3U8Slice string `json:"m3u8Slice"` // m3u8
	Width     int    `json:"width"`     // 录播视频宽度
	Height    int    `json:"height"`    // 录播视频高度
}

// Manager 就是房管的用户信息，目前没有 Medal 和 ManagerType
type Manager struct {
	UserInfo   `json:"userInfo"`
	CustomData string `json:"customData"` // 用户的一些额外信息，格式为 json
	Online     bool   `json:"online"`     // 是否直播间在线？
}

// UserProfile 就是用户信息
type UserProfile struct {
	UserID          int64  `json:"userID"`          // 用户 uid
	Nickname        string `json:"nickname"`        // 用户名字
	Avatar          string `json:"avatar"`          // 用户头像
	AvatarFrame     string `json:"avatarFrame"`     // 用户头像挂件
	FollowingCount  int    `json:"followingCount"`  // 用户关注数量
	FansCount       int    `json:"fansCount"`       // 用户粉丝数量
	ContributeCount int    `json:"contributeCount"` // 用户投稿数量
	Signature       string `json:"signature"`       // 用户签名
	VerifiedText    string `json:"verifiedText"`    // 用户认证信息
	IsJoinUpCollege bool   `json:"isJoinUpCollege"` // 用户是否加入阿普学院
	IsFollowing     bool   `json:"isFollowing"`     // 登陆帐号是否关注了用户
	IsFollowed      bool   `json:"isFollowed"`      // 用户是否关注了登陆帐号
}

// UserLiveInfo 就是用户直播信息
type UserLiveInfo struct {
	Profile               UserProfile `json:"profile"`               // 用户信息
	LiveType              LiveType    `json:"liveType"`              // 直播分类
	LiveID                string      `json:"liveID"`                // 直播 ID
	StreamName            string      `json:"streamName"`            // 直播源名字（ID）
	Title                 string      `json:"title"`                 // 直播间标题
	LiveStartTime         int64       `json:"liveStartTime"`         // 直播开始的时间，是以毫秒为单位的 Unix 时间
	Portrait              bool        `json:"portrait"`              // 是否手机直播
	Panoramic             bool        `json:"panoramic"`             // 是否全景直播
	LiveCover             string      `json:"liveCover"`             // 直播间封面
	OnlineCount           int         `json:"onlineCount"`           // 直播间在线人数
	LikeCount             int         `json:"likeCount"`             // 直播间点赞总数
	HasFansClub           bool        `json:"hasFansClub"`           // 主播是否有守护团
	DisableDanmakuShow    bool        `json:"disableDanmakuShow"`    // 是否禁止显示弹幕？
	PaidShowUserBuyStatus bool        `json:"paidShowUserBuyStatus"` // 登陆帐号是否购买了付费直播？
}

// UserProfileInfo 就是用户信息
type UserProfileInfo struct {
	UserID          int64  `json:"userID"`          // 用户 uid
	Nickname        string `json:"nickname"`        // 用户名字
	Avatar          string `json:"avatar"`          // 用户头像
	AvatarFrame     string `json:"avatarFrame"`     // 用户头像挂件
	FollowingCount  string `json:"followingCount"`  // 用户关注数量
	FansCount       string `json:"fansCount"`       // 用户粉丝数量
	ContributeCount string `json:"contributeCount"` // 用户投稿数量
	Signature       string `json:"signature"`       // 用户签名
	VerifiedText    string `json:"verifiedText"`    // 用户认证信息
	IsJoinUpCollege bool   `json:"isJoinUpCollege"` // 用户是否加入阿普学院
	IsFollowing     bool   `json:"isFollowing"`     // 登陆帐号是否关注了用户
	IsFollowed      bool   `json:"isFollowed"`      // 用户是否关注了登陆帐号
	LiveID          string `json:"liveID"`          // 直播 ID
	LikeCount       int    `json:"likeCount"`       // 最近一次直播的点赞总数
}

// UserMedalInfo 就是用户的守护徽章信息
type UserMedalInfo struct {
	UserProfile      `json:"profile"`
	FriendshipDegree int `json:"friendshipDegree"` // 用户守护徽章的亲密度
	Level            int `json:"level"`            // 用户守护徽章的等级
}

// MedalRankList 就是主播守护徽章等级列表
type MedalRankList struct {
	HasFansClub          bool            `json:"hasFansClub"`          // 主播是否有守护团
	RankList             []UserMedalInfo `json:"rankList"`             // 主播守护徽章等级列表
	ClubName             string          `json:"clubName"`             // 守护徽章名字
	MedalCount           int             `json:"medalCount"`           // 拥有主播守护徽章的用户的数量
	HasMedal             bool            `json:"hasMedal"`             // 登陆用户是否有主播的守护徽章
	UserFriendshipDegree int             `json:"userFriendshipDegree"` // 登陆用户的主播守护徽章的亲密度
	UserRank             string          `json:"userRank"`             // 登陆用户的主播守护徽章的排名
}

// LiveStat 就是直播统计数据
type LiveStat struct {
	Duration           int64 `json:"duration"` // 直播时长，单位为毫秒
	MaxPopularityValue int   `json:"maxPopularityValue"`
	WatchCount         int   `json:"watchCount"`   // 观看过直播的人数总数
	DiamondCount       int   `json:"diamondCount"` // 直播收到的付费礼物对应的钻石数量，100 钻石=1AC 币
	CommentCount       int   `json:"commentCount"` // 直播弹幕数量
	BananaCount        int   `json:"bananaCount"`  // 直播收到的香蕉数量
}

// LiveDetail 就是单场直播统计数据
type LiveDetail struct {
	LiveStartTime int64 `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的 Unix 时间
	LiveEndTime   int64 `json:"liveEndTime"`   // 直播结束的时间，是以毫秒为单位的 Unix 时间
	LiveStat      `json:"liveStat"`
}

// DailyData 就是单日直播统计数据
type DailyData struct {
	Date      string `json:"date"`      // 直播日期，格式类似"20210206"
	LiveTimes int    `json:"liveTimes"` // 当日直播次数
	LiveStat  `json:"liveStat"`
}

// LiveData 就是直播统计数据
type LiveData struct {
	BeginDate  string                  `json:"beginDate"`  // 统计开始的日期
	EndDate    string                  `json:"endDate"`    // 统计结束的日期
	Overview   LiveStat                `json:"overview"`   // 全部直播的统计概况
	LiveDetail map[string][]LiveDetail `json:"liveDetail"` // 单场直播统计数据，key 为直播日期，格式类似"20210206"
	DailyData  []DailyData             `json:"dailyData"`  // 单日直播统计数据
}

// LiveSchedule 就是直播预告
type LiveSchedule struct {
	ActivityID    int         `json:"activityID"`    // 活动 ID
	Profile       UserProfile `json:"profile"`       // 主播的用户信息
	Title         string      `json:"title"`         // 预告标题
	Cover         string      `json:"cover"`         // 预告封面
	LiveStartTime int64       `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的 Unix 时间
	LiveType      LiveType    `json:"liveType"`      // 直播分类
	Reserve       bool        `json:"reserve"`       // 登陆帐号是否预约了该直播
	ReserveNumber int         `json:"reserveNumber"` // 已预约用户的数量
}

// KickHistory 就是踢人历史记录
type KickHistory struct {
	UserID   int64  `json:"userID"`   // 被踢用户的 uid
	Nickname string `json:"nickname"` // 被踢用户的名字
	KickTime int64  `json:"kickTime"` // 用户被踢的时间，是以毫秒为单位的 Unix 时间
}

// LiveCutInfo 就是直播剪辑信息
type LiveCutInfo struct {
	Status      bool   `json:"status"`      // 是否允许剪辑直播录像（主播允许观众剪辑观众才能剪辑，主播直播时总是能剪辑自己的直播）
	URL         string `json:"url"`         // 剪辑直播的地址，直接访问可能出现登陆问题，需要访问跳转地址
	RedirectURL string `json:"redirectURL"` // 跳转直播剪辑的地址，访问一次后链接里的 token 就会失效
}

// 获取直播间排名前 50 的在线观众信息列表
func (t *token) getWatchingList(liveID string) (watchingList []WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getWatchingList() error: %v", err)
		}
	}()

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(watchingListURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取在线观众列表失败，响应为 %s", string(body)))
	}

	watchArray := v.GetArray("data", "list")
	watchingList = make([]WatchingUser, len(watchArray))
	for i, watch := range watchArray {
		o := watch.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				watchingList[i].UserID = v.GetInt64()
			case "nickname":
				watchingList[i].Nickname = string(v.GetStringBytes())
			case "avatar":
				watchingList[i].Avatar = string(v.GetStringBytes("0", "url"))
			case "anonymousUser":
				watchingList[i].AnonymousUser = v.GetBool()
			case "displaySendAmount":
				watchingList[i].DisplaySendAmount = string(v.GetStringBytes())
			case "customWatchingListData":
				watchingList[i].CustomData = string(v.GetStringBytes())
			case "managerType":
				watchingList[i].ManagerType = ManagerType(v.GetInt())
			default:
				log.Printf("在线观众列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	return watchingList, nil
}

// 获取主播最近七日内的礼物贡献榜前 50 名观众的详细信息
func (t *token) getBillboard(uid int64) (billboard []BillboardUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getBillboard() error: %v", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("authorId", strconv.FormatInt(uid, 10))
	body, err := t.fetchKuaiShouAPI(billboardURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取主播最近七日内的礼物贡献榜失败，响应为 %s", string(body)))
	}

	billboardArray := v.GetArray("data", "list")
	billboard = make([]BillboardUser, len(billboardArray))
	for i, user := range billboardArray {
		o := user.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				billboard[i].UserID = v.GetInt64()
			case "nickname":
				billboard[i].Nickname = string(v.GetStringBytes())
			case "avatar":
				billboard[i].Avatar = string(v.GetStringBytes("0", "url"))
			case "displaySendAmount":
				billboard[i].DisplaySendAmount = string(v.GetStringBytes())
			case "customData":
				billboard[i].CustomData = string(v.GetStringBytes())
			default:
				log.Printf("礼物贡献榜里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	return billboard, nil
}

// 获取直播总结信息
func (t *token) getSummary(liveID string) (summary *Summary, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getSummary() error: %v", err)
		}
	}()

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	body, err := t.fetchKuaiShouAPI(endSummaryURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播总结信息失败，响应为 %s", string(body)))
	}

	summary = new(Summary)
	o := v.GetObject("data")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "liveDurationMs":
			summary.Duration = v.GetInt64()
		case "likeCount":
			summary.LikeCount = string(v.GetStringBytes())
		case "watchCount":
			summary.WatchCount = string(v.GetStringBytes())
		case "payWalletTypeToReceiveCount":
			summary.GiftCount = v.GetInt("1")
		case "payWalletTypeToReceiveCurrency":
			summary.DiamondCount = v.GetInt("1")
			summary.BananaCount = v.GetInt("2")
		default:
			log.Printf("直播总结信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return summary, nil
}

// 获取抢到红包的用户列表
func (t *token) getLuckList(liveID, redpackID, redpackBizUnit string) (luckyList []LuckyUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLuckList() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取抢到红包的用户列表需要登陆 AcFun 帐号"))
	}

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	if redpackBizUnit == "" {
		form.Set("redpackBizUnit", "ztLiveAcfunRedpackGift")
	} else {
		form.Set("redpackBizUnit", redpackBizUnit)
	}
	form.Set("redpackId", redpackID)
	body, err := t.fetchKuaiShouAPI(redpackLuckListURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取抢到红包的用户列表失败，响应为 %s", string(body)))
	}

	luckyArray := v.GetArray("data", "luckyList")
	luckyList = make([]LuckyUser, len(luckyArray))
	for i, user := range luckyArray {
		o := user.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "simpleUserInfo":
				o := v.GetObject()
				o.Visit(func(k []byte, v *fastjson.Value) {
					switch string(k) {
					case "userId":
						luckyList[i].UserID = v.GetInt64()
					case "nickname":
						luckyList[i].Nickname = string(v.GetStringBytes())
					case "headPic":
						luckyList[i].Avatar = string(v.GetStringBytes("0", "url"))
					default:
						log.Printf("抢到红包的用户列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
					}
				})
			case "grabAmount":
				luckyList[i].GrabAmount = v.GetInt()
			default:
				log.Printf("抢到红包的用户列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	return luckyList, nil
}

// 获取直播回放的相关信息
func (t *token) getPlayback(liveID string) (playback *Playback, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getPlayback() error: %v", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("liveId", liveID)
	body, err := t.fetchKuaiShouAPI(playbackURL, form, true)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播回放的相关信息失败，响应为 %s", string(body)))
	}
	adaptiveManifest := v.GetStringBytes("data", "adaptiveManifest")
	v, err = p.ParseBytes(adaptiveManifest)
	checkErr(err)
	if len(v.GetArray("adaptationSet")) > 1 {
		log.Println("adaptationSet 列表长度大于 1，请报告 issue")
	}
	v = v.Get("adaptationSet", "0")
	duration := v.GetInt64("duration")
	if len(v.GetArray("representation")) > 1 {
		log.Println("representation 列表长度大于 1，请报告 issue")
	}
	v = v.Get("representation", "0")
	playback = &Playback{
		Duration:  duration,
		URL:       string(v.GetStringBytes("url")),
		BackupURL: string(v.GetStringBytes("backupUrl", "0")),
		M3U8Slice: string(v.GetStringBytes("m3u8Slice")),
		Width:     v.GetInt("width"),
		Height:    v.GetInt("height"),
	}
	if len(v.GetArray("backupUrl")) > 1 {
		log.Println("backupUrl 列表长度大于 1，请报告 issue")
	}

	return playback, nil
}

// 获取直播源信息，和 getLiveToken() 重复了
/*
func (t *token) getPlayURL() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getPlayURL() error: %w", err)
		}
	}()

	body, err := t.fetchKuaiShouAPI(getPlayURL, nil, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播源信息失败，响应为 %s", string(body)))
	}

	//videoPlayRes := string(v.GetStringBytes("data", "videoPlayRes"))

	return nil
}
*/

// 获取全部礼物的数据
func (t *token) getAllGift() (gifts map[int64]GiftDetail, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getAllGift() error: %v", err)
		}
	}()

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	body, err := t.fetchKuaiShouAPI(allGiftURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取全部礼物的数据失败，响应为 %s", string(body)))
	}

	return updateGiftList(v), nil
}

// 获取钱包里 AC 币和拥有的香蕉的数量
func (t *token) getWalletBalance() (accoins int, bananas int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getWalletBalance() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取钱包里 AC 币和拥有的香蕉的数量需要登陆 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	body, err := t.fetchKuaiShouAPI(walletBalanceURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取拥有的香蕉和钱包里AC币的数量失败，响应为 %s", string(body)))
	}

	o := v.GetObject("data", "payWalletTypeToBalance")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "1":
			accoins = v.GetInt()
		case "2":
			bananas = v.GetInt()
		default:
			log.Printf("用户钱包里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return accoins, bananas, nil
}

// 获取主播踢人的历史记录
func (t *token) getKickHistory(liveID string, count, page int) (list []KickHistory, lastPage bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getKickHistory() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取主播踢人的历史记录需要登陆主播的 AcFun 帐号"))
	}

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	form.Set("limit", strconv.Itoa(count))
	form.Set("pcursor", strconv.Itoa(page))
	body, err := t.fetchKuaiShouAPI(kickHistoryURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取主播踢人的历史记录失败，响应为 %s", string(body)))
	}

	v = v.Get("data")
	kickList := v.GetArray("list")
	list = make([]KickHistory, len(kickList))
	for i, l := range kickList {
		o := l.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				list[i].UserID = v.GetInt64()
			case "nickname":
				list[i].Nickname = string(v.GetStringBytes())
			case "kickTime":
				list[i].KickTime = v.GetInt64()
			default:
				log.Printf("主播踢人的历史记录里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	if string(v.GetStringBytes("pcursor")) == "no_more" {
		lastPage = true
	}

	return list, lastPage, nil
}

// 获取主播的房管列表
func (t *token) getManagerList() (managerList []Manager, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getManagerList() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取主播的房管列表需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	body, err := t.fetchKuaiShouAPI(managerListURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取主播的房管列表失败，响应为 %s", string(body)))
	}

	list := v.GetArray("data", "list")
	managerList = make([]Manager, len(list))
	for i, m := range list {
		o := m.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				managerList[i].UserID = v.GetInt64()
			case "nickname":
				managerList[i].Nickname = string(v.GetStringBytes())
			case "avatar":
				managerList[i].Avatar = string(v.GetStringBytes("0", "url"))
			case "customData":
				managerList[i].CustomData = string(v.GetStringBytes())
			case "online":
				managerList[i].Online = v.GetBool()
			default:
				log.Printf("主播的房管列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	return managerList, nil
}

// 从 json 里获取登陆帐号的守护徽章信息，uid 是登陆帐号的 uid
func getMedalJSON(v *fastjson.Value, uid int64) *Medal {
	medal := new(Medal)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "uperId":
			medal.UperID = v.GetInt64()
		case "clubName":
			medal.ClubName = string(v.GetStringBytes())
		case "level":
			medal.Level = v.GetInt()
		case "uperName":
			medal.UperName = string(v.GetStringBytes())
		case "uperHeadUrl":
			medal.UperAvatar = string(v.GetStringBytes())
		case "wearMedal":
			medal.WearMedal = v.GetBool()
		case "friendshipDegree":
			medal.FriendshipDegree = v.GetInt()
		case "joinClubTime":
			medal.JoinClubTime = v.GetInt64()
		case "currentDegreeLimit":
			medal.CurrentDegreeLimit = v.GetInt()
		default:
			log.Printf("守护徽章信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})
	medal.UserID = uid

	return medal
}

// 从 json 里获取登陆帐号的守护徽章亲密度信息
func getMedalDegreeJSON(v *fastjson.Value) *MedalDegree {
	medal := new(MedalDegree)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "uperId":
			medal.UperID = v.GetInt64()
		case "giftDegree":
			medal.GiftDegree = v.GetInt()
		case "giftDegreeLimit":
			medal.GiftDegreeLimit = v.GetInt()
		case "peachDegree":
			medal.PeachDegree = v.GetInt()
		case "peachDegreeLimit":
			medal.PeachDegreeLimit = v.GetInt()
		case "liveWatchDegree":
			medal.LiveWatchDegree = v.GetInt()
		case "liveWatchDegreeLimit":
			medal.LiveWatchDegreeLimit = v.GetInt()
		case "bananaDegree":
			medal.BananaDegree = v.GetInt()
		case "bananaDegreeLimit":
			medal.BananaDegreeLimit = v.GetInt()
		default:
			log.Printf("守护徽章亲密度信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return medal
}

// 获取登陆帐号拥有的指定主播的守护徽章详细信息
func (t *token) getMedalDetail(uid int64) (medal *MedalDetail, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getMedalDetail() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取登陆帐号拥有的指定主播的守护徽章详细信息需要登陆 AcFun 帐号"))
	}

	client := &httpClient{
		url:      fmt.Sprintf(medalDetailURL, uid),
		method:   "GET",
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
		panic(fmt.Errorf("获取登陆帐号拥有的指定主播的守护徽章详细信息失败，响应为 %s", string(body)))
	}

	medal = new(MedalDetail)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "result":
		case "host-name":
		case "medal":
			medal.Medal = *getMedalJSON(v, t.UserID)
		case "medalDegreeLimit":
			medal.MedalDegree = *getMedalDegreeJSON(v)
		case "rankIndex":
			medal.UserRank = string(v.GetStringBytes())
		default:
			log.Printf("守护徽章详细信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return medal, nil
}

// 获取登陆帐号拥有的守护徽章列表
func (t *token) getMedalList() (medalList []Medal, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getMedalList() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取登陆帐号拥有的守护徽章列表需要登陆 AcFun 帐号"))
	}

	client := &httpClient{
		url:      medalListURL,
		method:   "GET",
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
		panic(fmt.Errorf("获取登陆帐号拥有的守护徽章列表失败，响应为 %s", string(body)))
	}

	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "result":
		case "host-name":
		case "medalList":
			list := v.GetArray()
			num := len(list)
			medalList = make([]Medal, 0, num)
			for _, l := range list {
				medal := getMedalJSON(l, t.UserID)
				medal.MedalCount = num
				medalList = append(medalList, *medal)
			}
		default:
			log.Printf("登陆帐号拥有的守护徽章列表里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return medalList, nil
}

// 获取直播统计数据
func (t *token) getLiveData(days int) (data *LiveData, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveData() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取直播统计数据需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("days", strconv.Itoa(days))
	client := &httpClient{
		url:         liveDataURL,
		body:        form.QueryString(),
		method:      "POST",
		cookies:     t.Cookies,
		contentType: formContentType,
		deviceID:    t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取直播统计数据失败，响应为 %s", string(body)))
	}

	data = new(LiveData)
	data.LiveDetail = make(map[string][]LiveDetail)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "result":
		case "host-name":
		case "days":
		case "beginDate":
			data.BeginDate = string(v.GetStringBytes())
		case "endDate":
			data.EndDate = string(v.GetStringBytes())
		case "overview":
			o := v.GetObject()
			o.Visit(func(k []byte, v *fastjson.Value) {
				switch string(k) {
				case "totalLiveMillisecond":
					data.Overview.Duration = v.GetInt64()
				case "maxPopularityValue":
					data.Overview.MaxPopularityValue = v.GetInt()
				case "totalViewCount":
					data.Overview.WatchCount = v.GetInt()
				case "totalDiamondCount":
					data.Overview.DiamondCount = v.GetInt()
				case "totalCommentCount":
					data.Overview.CommentCount = v.GetInt()
				case "totalBananaCount":
					data.Overview.BananaCount = v.GetInt()
				default:
					log.Printf("直播统计数据里的overview出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
				}
			})
		case "liveDetail":
			o := v.GetObject()
			o.Visit(func(k []byte, v *fastjson.Value) {
				list := v.GetArray()
				d := make([]LiveDetail, len(list))
				for i, l := range list {
					o := l.GetObject()
					o.Visit(func(k []byte, v *fastjson.Value) {
						switch string(k) {
						case "startTime":
							d[i].LiveStartTime = v.GetInt64()
						case "endTime":
							d[i].LiveEndTime = v.GetInt64()
						case "liveMillisecond":
							d[i].Duration = v.GetInt64()
						case "maxPopularityValue":
							d[i].MaxPopularityValue = v.GetInt()
						case "viewCount":
							d[i].WatchCount = v.GetInt()
						case "diamondCount":
							d[i].DiamondCount = v.GetInt()
						case "commentCount":
							d[i].CommentCount = v.GetInt()
						case "bananaCount":
							d[i].BananaCount = v.GetInt()
						default:
							log.Printf("直播统计数据里的liveDetail出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
						}
					})
				}
				data.LiveDetail[string(k)] = d
			})
		case "dailyData":
			list := v.GetArray()
			data.DailyData = make([]DailyData, len(list))
			for i, l := range list {
				o := l.GetObject()
				o.Visit(func(k []byte, v *fastjson.Value) {
					switch string(k) {
					case "date":
						data.DailyData[i].Date = string(v.GetStringBytes())
					case "totalLiveTimes":
						data.DailyData[i].LiveTimes = v.GetInt()
					case "totalLiveMillisecond":
						data.DailyData[i].Duration = v.GetInt64()
					case "maxPopularityValue":
						data.DailyData[i].MaxPopularityValue = v.GetInt()
					case "totalViewCount":
						data.DailyData[i].WatchCount = v.GetInt()
					case "totalDiamondCount":
						data.DailyData[i].DiamondCount = v.GetInt()
					case "totalCommentCount":
						data.DailyData[i].CommentCount = v.GetInt()
					case "totalBananaCount":
						data.DailyData[i].BananaCount = v.GetInt()
					default:
						log.Printf("直播统计数据里的dailyData出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
					}
				})
			}
		default:
			log.Printf("直播统计数据里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return data, nil
}

// 获取直播剪辑信息
func (t *token) getLiveCutInfo(uid int64, liveID string) (info *LiveCutInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveCutInfo() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("获取直播剪辑信息需要登陆 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set(sid, midground)
	client := &httpClient{
		url:         getTokenURL,
		body:        form.QueryString(),
		method:      "POST",
		cookies:     t.Cookies,
		contentType: formContentType,
		referer:     t.livePage,
		deviceID:    t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取AcFun token失败，响应为 %s", string(body)))
	}

	token := string(v.GetStringBytes(midgroundAt))

	client = &httpClient{
		url:      fmt.Sprintf(liveCutInfoURL, uid, liveID),
		method:   "GET",
		cookies:  t.Cookies,
		deviceID: t.DeviceID,
	}
	body, err = client.request()
	checkErr(err)

	v, err = p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取直播剪辑信息失败，响应为：%s", string(body)))
	}

	var status bool
	statusNum := v.GetInt("liveCutStatus")
	if statusNum == 1 {
		status = true
	} else if statusNum == 2 {
		status = false
	} else {
		panic(fmt.Errorf("获取直播剪辑信息失败，响应为：%s", string(body)))
	}
	url := string(v.GetStringBytes("liveCutUrl"))
	info = &LiveCutInfo{
		Status:      status,
		URL:         url,
		RedirectURL: fmt.Sprintf(liveCutRedirectURL, token, url),
	}

	return info, nil
}

// 获取指定用户正在佩戴的守护徽章信息
func getUserMedal(uid int64, deviceID string) (medal *Medal, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getUserMedal() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      fmt.Sprintf(userMedalURL, uid),
		method:   "GET",
		deviceID: deviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取指定用户正在佩戴的守护徽章信息失败，响应为 %s", string(body)))
	}

	medal = new(Medal)
	medal.MedalCount = v.GetInt("medalCount")
	o := v.GetObject("wearMedalInfo")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "uperId":
			medal.UperID = v.GetInt64()
		case "level":
			medal.Level = v.GetInt()
		case "clubName":
			medal.ClubName = string(v.GetStringBytes())
		case "uperName":
			medal.UperName = string(v.GetStringBytes())
		case "uperHeadUrl":
			medal.UperAvatar = string(v.GetStringBytes())
		case "uperHeadImgInfo":
		default:
			log.Printf("指定用户正在佩戴的守护徽章信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})
	medal.UserID = uid
	medal.WearMedal = true

	return medal, nil
}

// 从 json 里获取用户信息，除了 UserID
func getUserProfileJSON(v *fastjson.Value) *UserProfile {
	profile := new(UserProfile)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "name":
			profile.Nickname = string(v.GetStringBytes())
		case "headUrl":
			profile.Avatar = string(v.GetStringBytes())
		case "avatarFrameMobileImg":
			profile.AvatarFrame = string(v.GetStringBytes())
		case "followingCountValue":
			profile.FollowingCount = v.GetInt()
		case "fanCountValue":
			profile.FansCount = v.GetInt()
		case "contributeCountValue":
			profile.ContributeCount = v.GetInt()
		case "signature":
			profile.Signature = string(v.GetStringBytes())
		case "verifiedText":
			profile.VerifiedText = string(v.GetStringBytes())
		case "isJoinUpCollege":
			profile.IsJoinUpCollege = v.GetBool()
		case "isFollowing":
			profile.IsFollowing = v.GetBool()
		case "isFollowed":
			profile.IsFollowed = v.GetBool()
		}
	})

	return profile
}

// 从 json 里获取用户直播信息
func getUserLiveInfoJSON(v *fastjson.Value) *UserLiveInfo {
	info := new(UserLiveInfo)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "user":
			info.Profile = *getUserProfileJSON(v)
		case "authorId":
		case "type":
			info.LiveType = *getLiveType(v)
		case "liveId":
			info.LiveID = string(v.GetStringBytes())
		case "streamName":
			info.StreamName = string(v.GetStringBytes())
		case "title":
			info.Title = string(v.GetStringBytes())
		case "createTime":
			info.LiveStartTime = v.GetInt64()
		case "portrait":
			info.Portrait = v.GetBool()
		case "panoramic":
			info.Panoramic = v.GetBool()
		case "coverUrls":
			info.LiveCover = string(v.GetStringBytes("0"))
		case "onlineCount":
			info.OnlineCount = v.GetInt()
		case "likeCount":
			info.LikeCount = v.GetInt()
		case "hasFansClub":
			info.HasFansClub = v.GetBool()
		case "disableDanmakuShow":
			info.DisableDanmakuShow = v.GetBool()
		case "paidShowUserBuyStatus":
			info.PaidShowUserBuyStatus = v.GetBool()
		}
	})
	info.Profile.UserID = v.GetInt64("authorId")

	return info
}

// 获取指定用户的直播信息
func getUserLiveInfo(uid int64, cookies Cookies, deviceID string) (info *UserLiveInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getUserLiveInfo() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      fmt.Sprintf(liveInfoURL, uid),
		method:   "GET",
		cookies:  cookies,
		deviceID: deviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取指定用户的直播信息失败，响应为 %s", string(body)))
	}

	return getUserLiveInfoJSON(v), nil
}

// 获取指定用户的信息
func getUserInfo(uid int64, cookies Cookies, deviceID string) (info *UserProfileInfo, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getUserProfile() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      fmt.Sprintf(userInfoURL, uid),
		method:   "GET",
		cookies:  cookies,
		deviceID: deviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取指定用户的信息失败，响应为 %s", string(body)))
	}

	info = new(UserProfileInfo)
	o := v.GetObject("profile")
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "userId":
			info.UserID = v.GetInt64()
		case "name":
			info.Nickname = string(v.GetStringBytes())
		case "headUrl":
			info.Avatar = string(v.GetStringBytes())
		case "avatarFrameMobileImg":
			info.AvatarFrame = string(v.GetStringBytes())
		case "following":
			info.FollowingCount = string(v.GetStringBytes())
		case "followed":
			info.FansCount = string(v.GetStringBytes())
		case "contentCount":
			info.ContributeCount = string(v.GetStringBytes())
		case "signature":
			info.Signature = string(v.GetStringBytes())
		case "verifiedText":
			info.VerifiedText = string(v.GetStringBytes())
		case "isContractUp":
			info.IsJoinUpCollege = v.GetBool()
		case "isFollowing":
			info.IsFollowing = v.GetBool()
		case "isFollowed":
			info.IsFollowed = v.GetBool()
		case "liveId":
			info.LiveID = string(v.GetStringBytes())
		case "likeCount":
			info.LikeCount = v.GetInt()
		}
	})

	return info, nil
}

// 获取指定主播的守护榜
func getMedalRankList(uid int64, cookies Cookies, deviceID string) (medalRankList *MedalRankList, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getMedalRankList() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      fmt.Sprintf(medalRankURL, uid),
		method:   "GET",
		cookies:  cookies,
		deviceID: deviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取指定主播的守护榜失败，响应为 %s", string(body)))
	}

	medalRankList = new(MedalRankList)
	o := v.GetObject()
	o.Visit(func(k []byte, v *fastjson.Value) {
		switch string(k) {
		case "result":
		case "host-name":
		case "hasFansClub":
			medalRankList.HasFansClub = v.GetBool()
		case "friendshipDegreeRank":
			list := v.GetArray()
			medalRankList.RankList = make([]UserMedalInfo, len(list))
			for i, l := range list {
				o := l.GetObject()
				o.Visit(func(k []byte, v *fastjson.Value) {
					switch string(k) {
					case "userId":
					case "userInfo":
						medalRankList.RankList[i].UserProfile = *getUserProfileJSON(v)
					case "friendshipDegree":
						medalRankList.RankList[i].FriendshipDegree = v.GetInt()
					case "medalLevel":
						medalRankList.RankList[i].Level = v.GetInt()
					default:
						log.Printf("用户的守护徽章信息里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
					}
				})
				medalRankList.RankList[i].UserID = l.GetInt64("userId")
			}
		case "clubName":
			medalRankList.ClubName = string(v.GetStringBytes())
		case "fansTotalCount":
			medalRankList.MedalCount = v.GetInt()
		case "isInFansClub":
			medalRankList.HasMedal = v.GetBool()
		case "curUserFriendshipDegree":
			medalRankList.UserFriendshipDegree = v.GetInt()
		case "curUserRankIndex":
			medalRankList.UserRank = string(v.GetStringBytes())
		default:
			log.Printf("指定主播的守护榜里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
		}
	})

	return medalRankList, nil
}

// 获取正在直播的直播间列表
func getLiveList(count, page int, cookies Cookies, deviceID string) (liveList []UserLiveInfo, lastPage bool, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getLiveList() error: %v", err)
		}
	}()

	client := &httpClient{
		url:      fmt.Sprintf(liveListURL, count, page),
		method:   "GET",
		cookies:  cookies,
		deviceID: deviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	v = v.Get("channelListData")
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取正在直播的直播间列表失败，响应为：%s", string(body)))
	}

	list := v.GetArray("liveList")
	liveList = make([]UserLiveInfo, 0, len(list))
	for _, l := range list {
		liveList = append(liveList, *getUserLiveInfoJSON(l))
	}

	if string(v.GetStringBytes("pcursor")) == "no_more" {
		lastPage = true
	}

	return liveList, lastPage, nil
}

// 获取全部正在直播的直播间列表
func getAllLiveList(cookies Cookies, deviceID string) ([]UserLiveInfo, error) {
	list, _, err := getLiveList(1000000, 0, cookies, deviceID)
	return list, err
}

// 获取直播预告列表
/*
func getScheduleList(cookies Cookies) (scheduleList []LiveSchedule, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getScheduleList() error: %w", err)
		}
	}()

	client := &httpClient{
		url:     scheduleListURL,
		method:  "POST",
		cookies: cookies,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取直播预告列表失败，响应为 %s", string(body)))
	}

	list := v.GetArray("liveScheduleList")
	scheduleList = make([]LiveSchedule, len(list))
	for i, l := range list {
		o := l.GetObject()
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "requestId":
			case "groupId":
			case "action":
			case "href":
			case "authorId":
			case "activityId":
				scheduleList[i].ActivityID = v.GetInt()
			case "user":
				scheduleList[i].Profile = *getUserProfileJSON(v)
			case "title":
				scheduleList[i].Title = string(v.GetStringBytes())
			case "cover":
				scheduleList[i].Cover = string(v.GetStringBytes())
			case "startTime":
				scheduleList[i].LiveStartTime = v.GetInt64()
			case "type":
				scheduleList[i].LiveType = *getLiveType(v)
			case "reserve":
				scheduleList[i].Reserve = v.GetBool()
			case "reserveNumber":
				scheduleList[i].ReserveNumber = v.GetInt()
			default:
				log.Printf("直播预告列表里出现未处理的 key 和 value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
		scheduleList[i].Profile.UserID = l.GetInt64("authorId")
	}

	return scheduleList, nil
}
*/

// Distinguish 返回阿里云和腾讯云链接，目前阿里云的下载速度比较快
func (pb *Playback) Distinguish() (aliURL, txURL string) {
	switch {
	case strings.Contains(pb.URL, "alivod"):
		aliURL = pb.URL
	case strings.Contains(pb.URL, "txvod"):
		txURL = pb.URL
	default:
		log.Printf("未能识别的录播链接：%s", pb.URL)
	}

	switch {
	case strings.Contains(pb.BackupURL, "alivod"):
		aliURL = pb.BackupURL
	case strings.Contains(pb.BackupURL, "txvod"):
		txURL = pb.BackupURL
	default:
		log.Printf("未能识别的录播链接：%s", pb.BackupURL)
	}

	return aliURL, txURL
}

// GetWatchingList 返回直播间排名前 50 的在线观众信息列表
func (ac *AcFunLive) GetWatchingList(liveID string) ([]WatchingUser, error) {
	return ac.t.getWatchingList(liveID)
}

// GetBillboard 返回指定 uid 的主播最近七日内的礼物贡献榜前 50 名观众的详细信息
func (ac *AcFunLive) GetBillboard(uid int64) ([]BillboardUser, error) {
	return ac.t.getBillboard(uid)
}

// GetSummary 返回直播总结信息
func (ac *AcFunLive) GetSummary(liveID string) (*Summary, error) {
	return ac.t.getSummary(liveID)
}

// GetLuckList 返回抢到红包的用户列表，需要登陆 AcFun 帐号，redpackBizUnit 为空时默认为 ztLiveAcfunRedpackGift
func (ac *AcFunLive) GetLuckList(liveID, redpackID, redpackBizUnit string) ([]LuckyUser, error) {
	return ac.t.getLuckList(liveID, redpackID, redpackBizUnit)
}

// GetPlayback 返回直播回放的相关信息，目前部分直播没有回放
func (ac *AcFunLive) GetPlayback(liveID string) (*Playback, error) {
	return ac.t.getPlayback(liveID)
}

// GetGiftList 返回指定主播直播间的礼物数据
func (ac *AcFunLive) GetGiftList(liveID string) (map[int64]GiftDetail, error) {
	return ac.t.getGiftList(liveID)
}

// GetAllGiftList 返回全部礼物的数据
func (ac *AcFunLive) GetAllGiftList() (map[int64]GiftDetail, error) {
	return ac.t.getAllGift()
}

// GetWalletBalance 返回钱包里 AC 币和拥有的香蕉的数量，需要登陆 AcFun 帐号
func (ac *AcFunLive) GetWalletBalance() (accoins int, bananas int, e error) {
	return ac.t.getWalletBalance()
}

// GetKickHistory 返回主播正在直播的那一场踢人的历史记录，count 为每页的数量，page 为第几页（从 0 开始数起），lastPage 说明是否最后一页，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetKickHistory(liveID string, count, page int) (list []KickHistory, lastPage bool, e error) {
	return ac.t.getKickHistory(liveID, count, page)
}

// GetAllKickHistory 返回主播正在直播的那一场踢人的全部历史记录，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetAllKickHistory(liveID string) ([]KickHistory, error) {
	list, _, err := ac.t.getKickHistory(liveID, 1000000, 0)
	return list, err
}

// GetManagerList 返回主播的房管列表，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetManagerList() ([]Manager, error) {
	return ac.t.getManagerList()
}

// GetMedalDetail 返回登陆帐号拥有的指定主播的守护徽章详细信息，需要登陆 AcFun 帐号
func (ac *AcFunLive) GetMedalDetail(uid int64) (*MedalDetail, error) {
	return ac.t.getMedalDetail(uid)
}

// GetMedalList 返回登陆用户拥有的守护徽章列表，最多返回亲密度最高的 300 个，需要登陆 AcFun 帐号
func (ac *AcFunLive) GetMedalList() ([]Medal, error) {
	return ac.t.getMedalList()
}

// GetLiveData 返回前 days 日到目前为止所有直播的统计数据，需要登陆主播的 AcFun 帐号
func (ac *AcFunLive) GetLiveData(days int) (*LiveData, error) {
	return ac.t.getLiveData(days)
}

// GetLiveCutInfo 获取 uid 指定主播的直播剪辑信息，只在主播直播时才能请求，需要直播的 liveID，需要登陆 AcFun 帐号
func (ac *AcFunLive) GetLiveCutInfo(uid int64, liveID string) (*LiveCutInfo, error) {
	return ac.t.getLiveCutInfo(uid, liveID)
}

// GetUserLiveInfo 返回 uid 指定用户的直播信息，可能会出现超时等各种网络原因的错误
func (ac *AcFunLive) GetUserLiveInfo(uid int64) (*UserLiveInfo, error) {
	return getUserLiveInfo(uid, ac.t.Cookies, ac.t.DeviceID)
}

// GetUserInfo 返回 uid 指定用户的信息
func (ac *AcFunLive) GetUserInfo(uid int64) (*UserProfileInfo, error) {
	return getUserInfo(uid, ac.t.Cookies, ac.t.DeviceID)
}

// GetMedalRankList 返回 uid 指定主播的守护榜（守护徽章亲密度排名前 50 名的用户），可用于获取指定主播的守护徽章名字
func (ac *AcFunLive) GetMedalRankList(uid int64) (medalRankList *MedalRankList, e error) {
	return getMedalRankList(uid, ac.t.Cookies, ac.t.DeviceID)
}

// GetLiveList 返回正在直播的直播间列表，count 为每页的直播间数量，page 为第几页（从 0 开始数起），lastPage 说明是否最后一页
func (ac *AcFunLive) GetLiveList(count, page int) (liveList []UserLiveInfo, lastPage bool, err error) {
	return getLiveList(count, page, ac.t.Cookies, ac.t.DeviceID)
}

// GetAllLiveList 返回全部正在直播的直播间列表
func (ac *AcFunLive) GetAllLiveList() ([]UserLiveInfo, error) {
	return getAllLiveList(ac.t.Cookies, ac.t.DeviceID)
}

// GetScheduleList 返回直播预告列表，目前有问题不可用
//func (ac *AcFunLive) GetScheduleList() ([]LiveSchedule, error) {
//	return getScheduleList(ac.t.Cookies)
//}

// GetDeviceID 获取设备 ID
func GetDeviceID() (string, error) {
	return getDeviceID()
}

// GetUserMedal 返回 uid 指定用户正在佩戴的守护徽章信息，没有 FriendshipDegree、JoinClubTime 和 CurrentDegreeLimit
func GetUserMedal(uid int64, deviceID string) (medal *Medal, e error) {
	return getUserMedal(uid, deviceID)
}

// GetUserLiveInfo 返回 uid 指定用户的直播信息，可能会出现超时等各种网络原因的错误
func GetUserLiveInfo(uid int64, deviceID string) (*UserLiveInfo, error) {
	return getUserLiveInfo(uid, nil, deviceID)
}

// GetUserInfo 返回 uid 指定用户的信息
func GetUserInfo(uid int64, deviceID string) (*UserProfileInfo, error) {
	return getUserInfo(uid, nil, deviceID)
}

// GetMedalRankList 返回 uid 指定主播的守护榜（守护徽章亲密度排名前 50 名的用户），可用于获取指定主播的守护徽章名字
func GetMedalRankList(uid int64, deviceID string) (medalRankList *MedalRankList, e error) {
	return getMedalRankList(uid, nil, deviceID)
}

// GetLiveList 返回正在直播的直播间列表，count 为每页的直播间数量，page 为第几页（从 0 开始数起），lastPage 说明是否最后一页
func GetLiveList(count, page int, deviceID string) (liveList []UserLiveInfo, lastPage bool, err error) {
	return getLiveList(count, page, nil, deviceID)
}

// GetAllLiveList 返回全部正在直播的直播间列表
func GetAllLiveList(deviceID string) ([]UserLiveInfo, error) {
	return getAllLiveList(nil, deviceID)
}

// GetScheduleList 返回直播预告列表，目前有问题不可用
//func GetScheduleList() ([]LiveSchedule, error) {
//	return getScheduleList(nil)
//}
