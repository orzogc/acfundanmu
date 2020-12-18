package acfundanmu

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Workiva/go-datastructures/queue"
)

// 队列长度
const queueLen = 1000

// ManagerType 就是房管类型
type ManagerType int32

const (
	// NotManager 不是房管
	NotManager ManagerType = iota
	// NormalManager 是房管
	NormalManager
)

// ManagerState 就是房管状态
type ManagerState int32

const (
	// ManagerStateUnknown 未知的房管状态，通常说明登陆用户不是房管
	ManagerStateUnknown ManagerState = iota
	// ManagerAdded 登陆用户被添加房管权限？
	ManagerAdded
	// ManagerRemoved 登陆用户被移除房管权限？
	ManagerRemoved
	// IsManager 登陆用户是房管
	IsManager
)

// RedpackDisplayStatus 红包状态
type RedpackDisplayStatus int32

const (
	// RedpackShow 红包出现？
	RedpackShow RedpackDisplayStatus = iota
	// RedpackGetToken 可以获取红包token？
	RedpackGetToken
	// RedpackGrab 可以抢红包
	RedpackGrab
)

// ChatMediaType 连麦类型
type ChatMediaType int32

const (
	// ChatMediaUnknown 未知的连麦类型
	ChatMediaUnknown ChatMediaType = iota
	// ChatMediaAudio 语音连麦
	ChatMediaAudio
	// ChatMediaVideo 视频连麦
	ChatMediaVideo
)

// ChatEndType 连麦结束类型
type ChatEndType int32

const (
	// ChatEndUnknown 未知的连麦结束类型
	ChatEndUnknown ChatEndType = iota
	// ChatEndCancelByAuthor 连麦发起者（主播）取消连麦
	ChatEndCancelByAuthor
	// ChatEndByAuthor 连麦发起者（主播）结束连麦
	ChatEndByAuthor
	// ChatEndByGuest 被连麦的人结束连麦
	ChatEndByGuest
	// ChatEndGuestReject 被连麦的人拒绝连麦
	ChatEndGuestReject
	// ChatEndGuestTimeout 等待连麦超时
	ChatEndGuestTimeout
	// ChatEndGuestHeartbeatTimeout 被连麦的人Heartbeat超时
	ChatEndGuestHeartbeatTimeout
	// ChatEndAuthorHeartbeatTimeout 连麦发起者（主播）Heartbeat超时
	ChatEndAuthorHeartbeatTimeout
)

// GiftDetail 就是礼物的详细信息
type GiftDetail struct {
	GiftID                 int64  `json:"giftID"`                 // 礼物ID
	GiftName               string `json:"giftName"`               // 礼物名字
	ARLiveName             string `json:"arLiveName"`             // 不为空时礼物属于虚拟偶像区的特殊礼物
	PayWalletType          int    `json:"payWalletType"`          // 1为非免费礼物，2为免费礼物
	Price                  int    `json:"price"`                  // 礼物价格，非免费礼物时单位为AC币，免费礼物（香蕉）时为1
	WebpPic                string `json:"webpPic"`                // 礼物的webp格式图片（动图）
	PngPic                 string `json:"pngPic"`                 // 礼物的png格式图片（大）
	SmallPngPic            string `json:"smallPngPic"`            // 礼物的png格式图片（小）
	AllowBatchSendSizeList []int  `json:"allowBatchSendSizeList"` // 网页或APP单次能够赠送的礼物数量列表
	CanCombo               bool   `json:"canCombo"`               // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
	CanDraw                bool   `json:"canDraw"`                // 是否能涂鸦？
	MagicFaceID            int    `json:"magicFaceID"`
	Description            string `json:"description"`  // 礼物的描述
	RedpackPrice           int    `json:"redpackPrice"` // 礼物红包价格总额，单位为AC币
}

// DrawPoint 绘制礼物的点？
type DrawPoint struct {
	MarginLeft int64   `json:"marginLeft"`
	MarginTop  int64   `json:"marginTop"`
	ScaleRatio float64 `json:"scaleRatio"`
	Handup     bool    `json:"handup"`
}

// DrawGiftInfo 绘制礼物的信息？
type DrawGiftInfo struct {
	ScreenWidth  int64       `json:"screenWidth"`
	ScreenHeight int64       `json:"screenHeight"`
	DrawPoint    []DrawPoint `json:"drawPoint"`
}

// UserInfo 就是用户信息
type UserInfo struct {
	UserID      int64       `json:"userID"`      // 用户uid
	Nickname    string      `json:"nickname"`    // 用户名字
	Avatar      string      `json:"avatar"`      // 用户头像
	Medal       MedalInfo   `json:"medal"`       // 守护徽章
	ManagerType ManagerType `json:"managerType"` // 用户是否房管
}

// MedalInfo 就是守护徽章信息
type MedalInfo struct {
	UperID   int64  `json:"uperID"`   // UP主的uid
	UserID   int64  `json:"userID"`   // 用户的uid
	ClubName string `json:"clubName"` // 守护徽章名字
	Level    int    `json:"level"`    // 守护徽章等级
}

// RichTextSegment 副文本片段的接口
type RichTextSegment interface {
	RichTextType() string
}

// RichTextUserInfo 富文本里的用户信息
type RichTextUserInfo struct {
	UserInfo `json:"userInfo"`
	Color    string `json:"color"` // 用户信息的颜色
}

// RichTextPlain 富文本里的文字
type RichTextPlain struct {
	Text  string `json:"text"`  // 文字
	Color string `json:"color"` // 文字的颜色
}

// RichTextImage 富文本里的图片
type RichTextImage struct {
	Pictures         []string `json:"pictures"`         // 图片
	AlternativeText  string   `json:"alternativeText"`  // 可选的文本？
	AlternativeColor string   `json:"alternativeColor"` // 可选的文本颜色？
}

// DanmuMessage 弹幕的接口
type DanmuMessage interface {
	GetSendTime() int64     // 获取弹幕发送时间
	GetUserInfo() *UserInfo // 获取UserInfo
}

// DanmuCommon 弹幕通用部分
type DanmuCommon struct {
	SendTime int64             `json:"sendTime"` // 弹幕发送时间，是以毫秒为单位的Unix时间
	UserInfo `json:"userInfo"` // 用户信息
}

// Comment 用户发的弹幕
type Comment struct {
	DanmuCommon `json:"danmuInfo"`
	Content     string `json:"content"` // 弹幕文字内容
}

// Like 用户点赞的弹幕
type Like DanmuCommon

// EnterRoom 用户进入直播间的弹幕
type EnterRoom DanmuCommon

// FollowAuthor 用户关注主播的弹幕
type FollowAuthor DanmuCommon

// ThrowBanana 用户投蕉的弹幕，没有Avatar、Medal和ManagerType，现在基本不用这个，通常用Gift代替
type ThrowBanana struct {
	DanmuCommon `json:"danmuInfo"`
	BananaCount int `json:"bananaCount"` // 投蕉数量
}

// Gift 用户赠送礼物的弹幕
type Gift struct {
	DanmuCommon         `json:"danmuInfo"`
	GiftDetail          `json:"giftDetail"`
	Count               int32        `json:"count"`               // 礼物单次赠送的数量，礼物总数是Count * Combo
	Combo               int32        `json:"combo"`               // 礼物连击数量，礼物总数是Count * Combo
	Value               int64        `json:"value"`               // 礼物价值，非免费礼物时单位为AC币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID             string       `json:"comboID"`             // 礼物连击ID
	SlotDisplayDuration int64        `json:"slotDisplayDuration"` // 应该是礼物动画持续的时间，单位为毫秒，送礼物后在该时间内再送一次可以实现礼物连击
	ExpireDuration      int64        `json:"ExpireDuration"`
	DrawGiftInfo        DrawGiftInfo `json:"drawGiftInfo"` // 礼物涂鸦
}

// RichText 富文本，目前是用于发红包和抢红包的相关消息
type RichText struct {
	SendTime int64             `json:"sendTime"` // 弹幕发送时间，是以毫秒为单位的Unix时间
	Segments []RichTextSegment `json:"segments"` // 富文本各部分，类型是RichTextUserInfo、RichTextPlain或RichTextImage
}

// JoinClub 用户加入主播的守护团，FansInfo和UperInfo都没有Avatar、Medal和ManagerType
type JoinClub struct {
	JoinTime int64    `json:"joinTime"` // 用户加入守护团的时间，是以毫秒为单位的Unix时间
	FansInfo UserInfo `json:"fansInfo"` // 用户的信息
	UperInfo UserInfo `json:"uperInfo"` // 主播的信息
}

// TopUser 就是礼物榜在线前三，目前没有Medal和ManagerType
type TopUser WatchingUser

// Redpack 红包信息
type Redpack struct {
	UserInfo           `json:"userInfo"`    // 发红包的用户
	DisplayStatus      RedpackDisplayStatus `json:"displayStatus"`      // 红包的状态
	GrabBeginTime      int64                `json:"grabBeginTime"`      // 抢红包的开始时间，是以毫秒为单位的Unix时间
	GetTokenLatestTime int64                `json:"getTokenLatestTime"` // 抢红包的用户获得token的最晚时间？是以毫秒为单位的Unix时间
	RedPackID          string               `json:"redPackID"`          // 红包ID
	RedpackBizUnit     string               `json:"redpackBizUnit"`     // 一般是"ztLiveAcfunRedpackGift"
	RedpackAmount      int64                `json:"redpackAmount"`      // 红包的总价值，单位是AC币
	SettleBeginTime    int64                `json:"settleBeginTime"`    // 抢红包的结束时间，是以毫秒为单位的Unix时间
}

// ChatCall 主播发起连麦
type ChatCall struct {
	ChatID   string `json:"chatID"`   // 连麦ID
	LiveID   string `json:"liveID"`   // 直播ID
	CallTime int64  `json:"callTime"` // 连麦发起时间，是以毫秒为单位的Unix时间
}

// ChatAccept 用户接受连麦？一般不会出现这个信号
type ChatAccept struct {
	ChatID          string        `json:"chatID"`    // 连麦ID
	MediaType       ChatMediaType `json:"mediaType"` // 连麦类型
	ArraySignalInfo string        `json:"arraySignalInfo"`
}

// ChatReady 用户接受连麦的信息
type ChatReady struct {
	ChatID    string        `json:"chatID"`    // 连麦ID
	Guest     UserInfo      `json:"guest"`     // 被连麦的帐号信息，目前没有房管类型
	MediaType ChatMediaType `json:"mediaType"` // 连麦类型
}

// ChatEnd 连麦结束
type ChatEnd struct {
	ChatID  string      `json:"chatID"`  // 连麦ID
	EndType ChatEndType `json:"endType"` // 连麦结束类型
}

// TokenInfo 就是AcFun直播的token相关信息
type TokenInfo struct {
	UserID       int64    `json:"userID"`       // 登陆模式或游客模式的uid
	SecurityKey  string   `json:"securityKey"`  // 密钥
	ServiceToken string   `json:"serviceToken"` // token
	DeviceID     string   `json:"deviceID"`     // 设备ID
	Cookies      []string `json:"cookies"`      // AcFun帐号的cookies
}

// StreamURL 就是直播源相关信息
type StreamURL struct {
	URL         string `json:"url"`         // 直播源链接
	Bitrate     int    `json:"bitrate"`     // 直播源码率，不一定是实际码率
	QualityType string `json:"qualityType"` // 直播源类型，一般是"STANDARD"、"HIGH"、"SUPER"、"BLUE_RAY"
	QualityName string `json:"qualityName"` // 直播源类型的中文名字，一般是"高清"、"超清"、"蓝光 4M"、"蓝光 5M"、"蓝光 6M"、"蓝光 7M"、"蓝光 8M"
}

// StreamInfo 就是直播的一部分信息
type StreamInfo struct {
	LiveID        string      `json:"liveID"`        // 直播ID
	Title         string      `json:"title"`         // 直播间标题
	LiveStartTime int64       `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的Unix time
	Panoramic     bool        `json:"panoramic"`     // 是否全景直播
	StreamList    []StreamURL `json:"streamList"`    // 直播源列表
	StreamName    string      `json:"streamName"`    // 直播源名字（ID）
}

// DisplayInfo 就是直播间的一些数据
type DisplayInfo struct {
	WatchingCount string `json:"watchingCount"` // 直播间在线观众数量
	LikeCount     string `json:"likeCount"`     // 直播间点赞总数
	LikeDelta     int    `json:"likeDelta"`     // 点赞增加数量
}

// LiveInfo 就是直播间的相关状态信息
type LiveInfo struct {
	KickedOut        string       `json:"kickedOut"`        // 被踢理由？
	ViolationAlert   string       `json:"violationAlert"`   // 直播间警告？
	LiveManagerState ManagerState `json:"liveManagerState"` // 登陆帐号的房管状态
	AllBananaCount   string       `json:"allBananaCount"`   // 直播间香蕉总数
	DisplayInfo      `json:"displayInfo"`
	TopUsers         []TopUser `json:"topUsers"`      // 礼物榜在线前三
	RecentComment    []Comment `json:"recentComment"` // APP进直播间时显示的最近发的弹幕
	RedpackList      []Redpack `json:"redpackList"`   // 红包列表
}

// 带锁的LiveInfo
type liveInfo struct {
	sync.Mutex // LiveInfo和RecentComment的锁
	LiveInfo
	TokenInfo
	StreamInfo
}

// AcFunLive 就是直播间弹幕系统相关信息，支持并行
type AcFunLive struct {
	q          *queue.Queue // DanmuMessage的队列
	info       *liveInfo    // 直播间的相关信息状态
	t          *token       // 令牌相关信息
	handlerMap *handlerMap  // 事件handler的map
}

// Login 登陆AcFun帐号，account为帐号邮箱或手机号，password为帐号密码
func Login(account, password string) (cookies []string, err error) {
	if account == "" || password == "" {
		return nil, fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆")
	}

	for retry := 0; retry < 3; retry++ {
		cookies, err = login(account, password)
		if err != nil {
			if retry == 2 {
				log.Printf("登陆AcFun帐号失败：%v", err)
				return nil, fmt.Errorf("Login() error: 登陆AcFun帐号失败：%w", err)
			}
			//log.Printf("登陆AcFun帐号出现错误：%v", err)
			//log.Println("尝试重新登陆AcFun帐号")
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	return cookies, nil
}

// Init 初始化，uid为主播的uid，cookies可以利用Login()获取，为nil时为游客模式，目前登陆模式和游客模式并没有太大区别。
// uid为0时仅获取TokenInfo，可以调用GetTokenInfo()获取。
// 应该尽可能复用返回的 *AcFunLive 。
func Init(uid int64, cookies []string) (ac *AcFunLive, err error) {
	ac = new(AcFunLive)
	ac.t = &token{
		liverID:  uid,
		livePage: fmt.Sprintf(liveURL, uid),
	}
	if len(cookies) != 0 {
		ac.t.cookies = append([]string{}, cookies...)
	}
	ac.info = new(liveInfo)
	ac.handlerMap = new(handlerMap)
	ac.handlerMap.listMap = make(map[eventType][]eventHandler)

	for retry := 0; retry < 3; retry++ {
		ac.info.StreamInfo, err = ac.t.getToken()
		if err != nil {
			if retry == 2 {
				log.Printf("初始化失败：%v", err)
				return nil, fmt.Errorf("Init() error: 初始化失败，主播可能不在直播：%w", err)
			}
			//log.Printf("初始化出现错误：%v", err)
			//log.Println("尝试重新初始化")
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	ac.info.TokenInfo = TokenInfo{
		UserID:       ac.t.userID,
		SecurityKey:  ac.t.securityKey,
		ServiceToken: ac.t.serviceToken,
		DeviceID:     ac.t.deviceID,
		Cookies:      ac.t.cookies,
	}

	return ac, nil
}

// InitWithToken 利用tokenInfo初始化，uid为主播的uid
func InitWithToken(uid int64, tokenInfo TokenInfo) (ac *AcFunLive, err error) {
	ac = new(AcFunLive)
	ac.t = &token{
		liverID:      uid,
		livePage:     fmt.Sprintf(liveURL, uid),
		userID:       tokenInfo.UserID,
		securityKey:  tokenInfo.SecurityKey,
		serviceToken: tokenInfo.ServiceToken,
		deviceID:     tokenInfo.DeviceID,
		cookies:      append([]string{}, tokenInfo.Cookies...),
	}
	ac.info = new(liveInfo)
	ac.info.TokenInfo = tokenInfo
	ac.info.TokenInfo.Cookies = append([]string{}, tokenInfo.Cookies...)
	ac.handlerMap = new(handlerMap)
	ac.handlerMap.listMap = make(map[eventType][]eventHandler)

	for retry := 0; retry < 3; retry++ {
		ac.info.StreamInfo, err = ac.t.getLiveToken()
		if err != nil {
			if retry == 2 {
				log.Printf("初始化失败：%v", err)
				return nil, fmt.Errorf("InitWithToken() error: 初始化失败，主播可能不在直播：%w", err)
			}
			//log.Printf("初始化出现错误：%v", err)
			//log.Println("尝试重新初始化")
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	return ac, nil
}

// ReInit 利用已有的 *AcFunLive 重新初始化，返回新的 *AcFunLive，事件模式下clearHandlers为true时需要重新调用OnComment等函数
func (ac *AcFunLive) ReInit(uid int64, clearHandlers bool) (newAC *AcFunLive, err error) {
	tokenInfo := ac.GetTokenInfo()
	newAC, err = InitWithToken(uid, *tokenInfo)
	if err != nil {
		return nil, err
	}
	if !clearHandlers {
		for k, v := range ac.handlerMap.listMap {
			newAC.handlerMap.listMap[k] = v
		}
	}
	return newAC, nil
}

// StartDanmu 启动websocket获取弹幕，ctx用来结束websocket，event为true时采用事件模式。
// event为false时最好调用GetDanmu()或WriteASS()以清空弹幕队列。
func (ac *AcFunLive) StartDanmu(ctx context.Context, event bool) <-chan error {
	ch := make(chan error, 1)
	if ac.t.liverID == 0 {
		err := fmt.Errorf("主播uid不能为0")
		log.Println(err)
		ch <- err
		return ch
	}
	if !event {
		ac.q = queue.New(queueLen)
	}
	go ac.wsStart(ctx, event, ch)
	return ch
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播），需要先调用StartDanmu(ctx, false)
func (ac *AcFunLive) GetDanmu() (danmu []DanmuMessage) {
	if ac.q == nil {
		log.Println("需要先调用StartDanmu()，event不能为true")
		return nil
	}
	if ac.t.liverID == 0 {
		log.Println("主播uid不能为0")
		return nil
	}
	if (*queue.Queue)(ac.q).Disposed() {
		return nil
	}
	ds, err := ac.q.Get(queueLen)
	if err != nil {
		return nil
	}

	danmu = make([]DanmuMessage, len(ds))
	for i, d := range ds {
		danmu[i] = d.(DanmuMessage)
	}

	return danmu
}

// GetLiveInfo 返回直播间的状态信息，需要先调用StartDanmu(ctx, false)
func (ac *AcFunLive) GetLiveInfo() *LiveInfo {
	ac.info.Lock()
	defer ac.info.Unlock()
	info := ac.info.LiveInfo
	info.TopUsers = append([]TopUser{}, ac.info.TopUsers...)
	info.RecentComment = append([]Comment{}, ac.info.RecentComment...)
	info.RedpackList = append([]Redpack{}, ac.info.RedpackList...)
	return &info
}

// GetTokenInfo 返回直播间token相关信息，不需要调用StartDanmu()
func (ac *AcFunLive) GetTokenInfo() *TokenInfo {
	info := ac.info.TokenInfo
	info.Cookies = append([]string{}, ac.info.Cookies...)
	return &info
}

// GetStreamInfo 返回直播的一些信息，不需要调用StartDanmu()
func (ac *AcFunLive) GetStreamInfo() *StreamInfo {
	info := ac.info.StreamInfo
	info.StreamList = append([]StreamURL{}, ac.info.StreamList...)
	return &info
}

// GetUserID 返回AcFun帐号的uid
func (ac *AcFunLive) GetUserID() int64 {
	return ac.t.userID
}

// GetLiverID 返回主播的uid，有可能是0
func (ac *AcFunLive) GetLiverID() int64 {
	return ac.t.liverID
}

// GetLiveID 返回liveID，有可能为空
func (ac *AcFunLive) GetLiveID() string {
	return ac.t.liveID
}

// GetTokenInfo 返回TokenInfo，相当于调用 Init(0, cookies) 后返回对应的TokenInfo，cookies可以利用Login()获取，为nil时为游客模式
func GetTokenInfo(cookies []string) (*TokenInfo, error) {
	ac, err := Init(0, cookies)
	if err != nil {
		return nil, err
	}
	return ac.GetTokenInfo(), nil
}
