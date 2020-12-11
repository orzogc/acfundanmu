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
	GiftID                 int64  // 礼物ID
	GiftName               string // 礼物名字
	ARLiveName             string // 不为空时礼物属于虚拟偶像区的特殊礼物
	PayWalletType          int    // 1为非免费礼物，2为免费礼物
	Price                  int    // 礼物价格，非免费礼物时单位为AC币，免费礼物（香蕉）时为1
	WebpPic                string // 礼物的webp格式图片（动图）
	PngPic                 string // 礼物的png格式图片（大）
	SmallPngPic            string // 礼物的png格式图片（小）
	AllowBatchSendSizeList []int  // 网页或APP单次能够赠送的礼物数量列表
	CanCombo               bool   // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
	CanDraw                bool   // 是否能涂鸦？
	MagicFaceID            int
	Description            string // 礼物的描述
	RedpackPrice           int    // 礼物红包价格总额，单位为AC币
}

// DrawPoint 绘制礼物的点？
type DrawPoint struct {
	MarginLeft int64
	MarginTop  int64
	ScaleRatio float64
	Handup     bool
}

// DrawGiftInfo 绘制礼物的信息？
type DrawGiftInfo struct {
	ScreenWidth  int64
	ScreenHeight int64
	DrawPoint    []DrawPoint
}

// UserInfo 就是用户信息
type UserInfo struct {
	UserID      int64       // 用户uid
	Nickname    string      // 用户名字
	Avatar      string      // 用户头像
	Medal       MedalInfo   // 守护徽章
	ManagerType ManagerType // 用户是否房管
}

// MedalInfo 就是守护徽章信息
type MedalInfo struct {
	UperID   int64  // UP主的uid
	UserID   int64  // 用户的uid
	ClubName string // 守护徽章名字
	Level    int    // 守护徽章等级
}

// RichTextSegment 副文本片段的接口
type RichTextSegment interface {
	RichTextType() string
}

// RichTextUserInfo 富文本里的用户信息
type RichTextUserInfo struct {
	UserInfo
	Color string // 用户信息的颜色
}

// RichTextPlain 富文本里的文字
type RichTextPlain struct {
	Text  string // 文字
	Color string // 文字的颜色
}

// RichTextImage 富文本里的图片
type RichTextImage struct {
	Pictures         []string // 图片
	AlternativeText  string   // 可选的文本？
	AlternativeColor string   // 可选的文本颜色？
}

// DanmuMessage 弹幕的接口
type DanmuMessage interface {
	GetSendTime() int64     // 获取弹幕发送时间
	GetUserInfo() *UserInfo // 获取UserInfo
}

// DanmuCommon 弹幕通用部分
type DanmuCommon struct {
	SendTime int64 // 弹幕发送时间，是以纳秒为单位的Unix时间
	UserInfo       // 用户信息
}

// Comment 用户发的弹幕
type Comment struct {
	DanmuCommon
	Content string // 弹幕内容
}

// Like 用户点赞的弹幕
type Like DanmuCommon

// EnterRoom 用户进入直播间的弹幕
type EnterRoom DanmuCommon

// FollowAuthor 用户关注主播的弹幕
type FollowAuthor DanmuCommon

// ThrowBanana 用户投蕉的弹幕，没有Avatar、Medal和ManagerType，现在基本不用这个，通常用Gift代替
type ThrowBanana struct {
	DanmuCommon
	BananaCount int // 投蕉数量
}

// Gift 用户赠送礼物的弹幕
type Gift struct {
	DanmuCommon
	GiftDetail                   // 礼物详细信息
	Count                 int32  // 礼物单次赠送的数量，礼物总数是Count * Combo
	Combo                 int32  // 礼物连击数量，礼物总数是Count * Combo
	Value                 int64  // 礼物价值，非免费礼物时单位为AC币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID               string // 礼物连击ID
	SlotDisplayDurationMs int64  // 应该是礼物动画持续的时间，送礼物后在该时间内再送一次可以实现礼物连击
	ExpireDurationMs      int64
	DrawGiftInfo          DrawGiftInfo // 礼物涂鸦
}

// RichText 富文本，目前是用于发红包和抢红包的相关消息
type RichText struct {
	SendTime int64             // 弹幕发送时间，是以纳秒为单位的Unix时间
	Segments []RichTextSegment // 富文本各部分，类型是RichTextUserInfo、RichTextPlain或RichTextImage
}

// JoinClub 用户加入主播的守护团，FansInfo和UperInfo都没有Avatar、Medal和ManagerType
type JoinClub struct {
	JoinTime int64    // 用户加入守护团的时间，是以纳秒为单位的Unix时间
	FansInfo UserInfo // 用户的信息
	UperInfo UserInfo // 主播的信息
}

// TopUser 就是礼物榜在线前三，目前没有Medal和ManagerType
type TopUser WatchingUser

// Redpack 红包信息
type Redpack struct {
	UserInfo                                // 发红包的用户
	DisplayStatus      RedpackDisplayStatus // 红包的状态
	GrabBeginTime      int64                // 抢红包的开始时间，是以纳秒为单位的Unix时间
	GetTokenLatestTime int64                // 抢红包的用户获得token的最晚时间？是以纳秒为单位的Unix时间
	RedPackID          string               // 红包ID
	RedpackBizUnit     string               // 一般是"ztLiveAcfunRedpackGift"
	RedpackAmount      int64                // 红包的总价值，单位是AC币
	SettleBeginTime    int64                // 抢红包的结束时间，是以纳秒为单位的Unix时间
}

// ChatInfo 连麦信息
/*
type ChatInfo struct {
	ChatID          string        // 连麦ID
	LiveID          string        // 直播ID
	CallTimestampMs int64         // 连麦发起时间
	Guest           UserInfo      // 被连麦的帐号信息，目前没有房管类型
	MediaType       ChatMediaType // 连麦类型
	EndType         ChatEndType   // 连麦结束类型
}
*/

// ChatCall 主播发起连麦
type ChatCall struct {
	ChatID   string // 连麦ID
	LiveID   string // 直播ID
	CallTime int64  // 连麦发起时间，是以纳秒为单位的Unix时间
}

// ChatAccept 用户接受连麦？一般不会出现这个信号
type ChatAccept struct {
	ChatID          string        // 连麦ID
	MediaType       ChatMediaType // 连麦类型
	ArraySignalInfo string
}

// ChatReady 用户接受连麦的信息
type ChatReady struct {
	ChatID    string        // 连麦ID
	Guest     UserInfo      // 被连麦的帐号信息，目前没有房管类型
	MediaType ChatMediaType // 连麦类型
}

// ChatEnd 连麦结束
type ChatEnd struct {
	ChatID  string      // 连麦ID
	EndType ChatEndType // 连麦结束类型
}

// TokenInfo 就是AcFun直播的token相关信息
type TokenInfo struct {
	UserID       int64    // 登陆模式或游客模式的uid
	SecurityKey  string   // 密钥
	ServiceToken string   // token
	DeviceID     string   // 设备ID
	Cookies      []string // AcFun帐号的cookies
}

// StreamURL 就是直播源相关信息
type StreamURL struct {
	URL         string // 直播源链接
	Bitrate     int    // 直播源码率，不一定是实际码率
	QualityType string // 直播源类型，一般是"STANDARD"、"HIGH"、"SUPER"、"BLUE_RAY"
	QualityName string // 直播源类型的中文名字，一般是"高清"、"超清"、"蓝光 4M"、"蓝光 5M"、"蓝光 6M"、"蓝光 7M"、"蓝光 8M"
}

// StreamInfo 就是直播的一部分信息
type StreamInfo struct {
	LiveID        string      // 直播ID
	Title         string      // 直播间标题
	LiveStartTime int64       // 直播开始的时间，是以毫秒为单位的Unix time
	Panoramic     bool        // 是否全景直播
	StreamList    []StreamURL // 直播源列表
	StreamName    string      // 直播源名字（ID）
}

// DisplayInfo 就是直播间的一些数据
type DisplayInfo struct {
	WatchingCount string // 直播间在线观众数量
	LikeCount     string // 直播间点赞总数
	LikeDelta     int    // 点赞增加数量
}

// LiveInfo 就是直播间的相关状态信息
type LiveInfo struct {
	KickedOut        string       // 被踢理由？
	ViolationAlert   string       // 直播间警告？
	LiveManagerState ManagerState // 登陆帐号的房管状态
	AllBananaCount   string       // 直播间香蕉总数
	DisplayInfo
	TopUsers      []TopUser // 礼物榜在线前三
	RecentComment []Comment // APP进直播间时显示的最近发的弹幕
	RedpackList   []Redpack // 红包列表
}

// 带锁的LiveInfo
type liveInfo struct {
	sync.Mutex // LiveInfo和RecentComment的锁
	LiveInfo
	TokenInfo
	StreamInfo
}

// DanmuQueue 就是直播间弹幕系统相关信息，支持并行
type DanmuQueue struct {
	q          *queue.Queue // DanmuMessage的队列
	info       *liveInfo    // 直播间的相关信息状态
	t          *token       // 令牌相关信息
	handlerMap *handlerMap  // 事件handler的map
}

// Login 登陆AcFun帐号，username为帐号邮箱或手机号，password为帐号密码
func Login(username, password string) (cookies []string, err error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆")
	}

	for retry := 0; retry < 3; retry++ {
		cookies, err = login(username, password)
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
// 应该尽可能复用返回的 *DanmuQueue 。
func Init(uid int64, cookies []string) (dq *DanmuQueue, err error) {
	dq = new(DanmuQueue)
	dq.t = &token{
		uid:      uid,
		livePage: fmt.Sprintf(liveURL, uid),
	}
	if len(cookies) != 0 {
		dq.t.cookies = append([]string{}, cookies...)
	}
	dq.info = new(liveInfo)
	dq.handlerMap = new(handlerMap)
	dq.handlerMap.listMap = make(map[eventType][]eventHandler)

	for retry := 0; retry < 3; retry++ {
		dq.info.StreamInfo, err = dq.t.getToken()
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

	dq.info.TokenInfo = TokenInfo{
		UserID:       dq.t.userID,
		SecurityKey:  dq.t.securityKey,
		ServiceToken: dq.t.serviceToken,
		DeviceID:     dq.t.deviceID,
		Cookies:      dq.t.cookies,
	}

	return dq, nil
}

// InitWithToken 利用tokenInfo初始化，uid为主播的uid
func InitWithToken(uid int64, tokenInfo TokenInfo) (dq *DanmuQueue, err error) {
	dq = new(DanmuQueue)
	dq.t = &token{
		uid:          uid,
		livePage:     fmt.Sprintf(liveURL, uid),
		userID:       tokenInfo.UserID,
		securityKey:  tokenInfo.SecurityKey,
		serviceToken: tokenInfo.ServiceToken,
		deviceID:     tokenInfo.DeviceID,
		cookies:      append([]string{}, tokenInfo.Cookies...),
	}
	dq.info = new(liveInfo)
	dq.info.TokenInfo = tokenInfo
	dq.info.TokenInfo.Cookies = append([]string{}, tokenInfo.Cookies...)
	dq.handlerMap = new(handlerMap)
	dq.handlerMap.listMap = make(map[eventType][]eventHandler)

	for retry := 0; retry < 3; retry++ {
		dq.info.StreamInfo, err = dq.t.getLiveToken()
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

	return dq, nil
}

// ReInit 利用已有的 *DanmuQueue 重新初始化，返回新的 *DanmuQueue，事件模式下clearHandlers为true时需要重新调用OnComment等函数
func (dq *DanmuQueue) ReInit(uid int64, clearHandlers bool) (newDQ *DanmuQueue, err error) {
	tokenInfo := dq.GetTokenInfo()
	newDQ, err = InitWithToken(uid, *tokenInfo)
	if err != nil {
		return nil, err
	}
	if !clearHandlers {
		for k, v := range dq.handlerMap.listMap {
			newDQ.handlerMap.listMap[k] = v
		}
	}
	return newDQ, nil
}

// StartDanmu 启动websocket获取弹幕，ctx用来结束websocket，event为true时采用事件模式。
// event为false时最好调用GetDanmu()或WriteASS()以清空弹幕队列。
func (dq *DanmuQueue) StartDanmu(ctx context.Context, event bool) <-chan error {
	ch := make(chan error, 1)
	if dq.t.uid == 0 {
		err := fmt.Errorf("主播uid不能为0")
		log.Println(err)
		ch <- err
		return ch
	}
	if !event {
		dq.q = queue.New(queueLen)
	}
	go dq.wsStart(ctx, event, ch)
	return ch
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播），需要先调用StartDanmu(ctx, false)
func (dq *DanmuQueue) GetDanmu() (danmu []DanmuMessage) {
	if dq.q == nil {
		log.Println("需要先调用StartDanmu()，event不能为true")
		return nil
	}
	if dq.t.uid == 0 {
		log.Println("主播uid不能为0")
		return nil
	}
	if (*queue.Queue)(dq.q).Disposed() {
		return nil
	}
	ds, err := dq.q.Get(queueLen)
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
func (dq *DanmuQueue) GetLiveInfo() *LiveInfo {
	dq.info.Lock()
	defer dq.info.Unlock()
	info := dq.info.LiveInfo
	info.TopUsers = append([]TopUser{}, dq.info.TopUsers...)
	info.RecentComment = append([]Comment{}, dq.info.RecentComment...)
	info.RedpackList = append([]Redpack{}, dq.info.RedpackList...)
	return &info
}

// GetTokenInfo 返回直播间token相关信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetTokenInfo() *TokenInfo {
	info := dq.info.TokenInfo
	info.Cookies = append([]string{}, dq.info.Cookies...)
	return &info
}

// GetStreamInfo 返回直播的一些信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetStreamInfo() *StreamInfo {
	info := dq.info.StreamInfo
	info.StreamList = append([]StreamURL{}, dq.info.StreamList...)
	return &info
}

// GetUID 返回主播的uid，有可能是0
func (dq *DanmuQueue) GetUID() int64 {
	return dq.t.uid
}

// GetTokenInfo 返回TokenInfo，相当于调用 Init(0, cookies) 后返回对应的TokenInfo，cookies可以利用Login()获取，为nil时为游客模式
func GetTokenInfo(cookies []string) (*TokenInfo, error) {
	dq, err := Init(0, cookies)
	if err != nil {
		return nil, err
	}
	return dq.GetTokenInfo(), nil
}
