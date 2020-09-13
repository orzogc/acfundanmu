package acfundanmu

import (
	"context"
	"sync"

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

// Giftdetail 就是礼物的详细信息
type Giftdetail struct {
	GiftID        int64  // 礼物ID
	GiftName      string // 礼物名字
	ARLiveName    string // 不为空时礼物属于虚拟偶像区的特殊礼物
	PayWalletType int    // 1为非免费礼物，2为免费礼物
	Price         int    // 礼物价格，非免费礼物时单位为ac币，免费礼物（香蕉）时为1
	WebpPic       string // 礼物的webp格式图片（动图）
	PngPic        string // 礼物的png格式图片（大）
	SmallPngPic   string // 礼物的png格式图片（小）
	CanCombo      bool   // 是否能连击，一般免费礼物（香蕉）不能连击，其余能连击
	MagicFaceID   int
	Description   string // 礼物的描述
	RedpackPrice  int    // 礼物红包价格总额，单位为AC币
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
	Medal       MedalInfo   // 粉丝牌
	ManagerType ManagerType // 用户是否房管
}

// MedalInfo 就是粉丝牌信息
type MedalInfo struct {
	UperID   int64  // UP主的uid
	ClubName string // 粉丝牌名字
	Level    int    // 粉丝牌等级
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
	GetSendTime() int64 // 获取弹幕发送时间
	GetUserInfo() UserInfo
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
	Giftdetail                   // 礼物详细信息
	Count                 int32  // 礼物数量
	Combo                 int32  // 礼物连击数量
	Value                 int64  // 礼物价值，非免费礼物时单位为ac币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID               string // 礼物连击ID
	SlotDisplayDurationMs int64  // 应该是礼物动画持续的时间，送礼物后在该时间内再送一次可以实现礼物连击
	ExpireDurationMs      int64
	DrawGiftInfo          DrawGiftInfo // 目前好像都没有这部分
}

// RichText 富文本，目前是用于发红包和抢红包的相关消息
type RichText struct {
	SendTime int64         // 弹幕发送时间，是以纳秒为单位的Unix时间
	Segments []interface{} // 富文本各部分，类型是RichTextUserInfo、RichTextPlain或RichTextImage
}

// GetSendTime 获取弹幕发送时间
func (d *Comment) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Comment) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *Like) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Like) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *EnterRoom) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *EnterRoom) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *FollowAuthor) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *FollowAuthor) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *ThrowBanana) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *ThrowBanana) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *Gift) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Gift) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *RichText) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息，返回第一个RichTextUserInfo的UserInfo，否则返回空的UserInfo
func (d *RichText) GetUserInfo() UserInfo {
	for _, segment := range d.Segments {
		if u, ok := segment.(RichTextUserInfo); ok {
			return u.UserInfo
		}
	}
	return UserInfo{}
}

// WatchingUser 就是观看直播的用户的信息，目前没有Medal
type WatchingUser struct {
	UserInfo                      // 用户信息
	AnonymousUser          bool   // 是否匿名用户
	DisplaySendAmount      string // 赠送的全部礼物的价值，单位是ac币
	CustomWatchingListData string // 用户的一些额外信息，格式为json
}

// TopUser 就是礼物榜在线前三，目前没有Medal和ManagerType
type TopUser WatchingUser

// Redpack 红包信息
type Redpack struct {
	UserInfo                                  // 发红包的用户
	DisplayStatus        RedpackDisplayStatus // 红包的状态
	GrabBeginTimeMs      int64                // 抢红包的开始时间
	GetTokenLatestTimeMs int64                // 抢红包的用户获得token的最晚时间？
	RedPackID            string               // 红包ID
	RedpackBizUnit       string               // 一般是"ztLiveAcfunRedpackGift"
	RedpackAmount        int64                // 红包的总价值，单位是ac币
	SettleBeginTime      int64                // 抢红包的结束时间
}

// ChatInfo 连麦信息
type ChatInfo struct {
	ChatID          string        // 连麦ID
	LiveID          string        // 直播ID
	CallTimestampMs int64         // 连麦发起时间
	Guest           UserInfo      // 被连麦的帐号信息，目前没有房管类型
	MediaType       ChatMediaType // 连麦类型
	EndType         ChatEndType   // 连麦结束类型
}

// LiveInfo 就是直播间的相关状态信息
type LiveInfo struct {
	KickedOut        string       // 被踢理由？
	ViolationAlert   string       // 直播间警告？
	LiveManagerState ManagerState // 登陆帐号的房管状态？
	AllBananaCount   string       // 直播间香蕉总数
	WatchingCount    string       // 直播间在线观众数量
	LikeCount        string       // 直播间点赞总数
	LikeDelta        int          // 点赞增加数量？
	TopUsers         []TopUser    // 礼物榜在线前三
	RedpackList      []Redpack    // 红包列表
	Chat             ChatInfo     // 连麦信息
}

// 带锁的LiveInfo
type liveInfo struct {
	sync.Mutex // liveInfo的锁
	LiveInfo
	RecentComment []Comment // APP进直播间时显示的最近发的弹幕
}

// DanmuQueue 就是弹幕的队列
type DanmuQueue struct {
	q    *queue.Queue // DanmuMessage的队列
	info *liveInfo    // 直播间的相关信息状态
	t    *token       // 令牌相关信息
}

// Init 初始化，uid为主播的uid，usernameAndPassword参数可以依次传递AcFun帐号邮箱和密码以登陆AcFun，没有时为游客模式，目前登陆模式和游客模式并没有区别
func Init(uid int64, usernameAndPassword ...string) (dq *DanmuQueue, e error) {
	dq = new(DanmuQueue)
	dq.t = new(token)
	dq.t.uid = uid
	if len(usernameAndPassword) == 2 && usernameAndPassword[0] != "" && usernameAndPassword[1] != "" {
		cookies, err := login(usernameAndPassword[0], usernameAndPassword[1])
		if err != nil {
			return nil, err
		}
		dq.t.cookies = cookies
	}
	err := dq.t.initialize()
	if err != nil {
		return nil, err
	}
	return dq, nil
}

// StartDanmu 启动websocket获取弹幕，ctx用来结束websocket
func (dq *DanmuQueue) StartDanmu(ctx context.Context) {
	dq.q = queue.New(queueLen)
	dq.info = new(liveInfo)
	go dq.wsStart(ctx)
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播），需要先调用StartDanmu()
func (dq *DanmuQueue) GetDanmu() (danmu []DanmuMessage) {
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

// GetInfo 返回直播间的状态信息，需要先调用StartDanmu()
func (dq *DanmuQueue) GetInfo() (info LiveInfo) {
	dq.info.Lock()
	defer dq.info.Unlock()
	info = dq.info.LiveInfo
	return info
}

// GetRecentComment 返回进直播间时直播间里最近发的十条弹幕，需要先调用StartDanmu()
func (dq *DanmuQueue) GetRecentComment() (comments []Comment) {
	dq.info.Lock()
	defer dq.info.Unlock()
	comments = dq.info.RecentComment
	return comments
}

// GetWatchingList 返回直播间排名前50的在线观众信息列表，不需要调用StartDanmu()
func (dq *DanmuQueue) GetWatchingList() (*[]WatchingUser, error) {
	return dq.t.watchingList()
}
