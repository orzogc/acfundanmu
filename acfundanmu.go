package acfundanmu

import (
	"context"
	"fmt"
	"net/http"
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
	// ManagerStateUnknown 未知的房管状态
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

// GiftInfo 就是弹幕里的礼物信息
type GiftInfo struct {
	Giftdetail                   // 礼物详细信息
	Count                 int32  // 礼物数量
	Combo                 int32  // 礼物连击数量
	Value                 int64  // 礼物价值，非免费礼物时单位为ac币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID               string // 礼物连击ID
	SlotDisplayDurationMs int64  // 应该是礼物动画持续的时间，送礼物后在该时间内再送一次可以实现礼物连击
	ExpireDurationMs      int64
	DrawGiftInfo          DrawGiftInfo // 目前好像都没有这部分
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
	GiftInfo // 礼物信息
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

// GetSendTime 获取弹幕发送时间
func (d *Like) GetSendTime() int64 {
	return d.SendTime
}

// GetSendTime 获取弹幕发送时间
func (d *EnterRoom) GetSendTime() int64 {
	return d.SendTime
}

// GetSendTime 获取弹幕发送时间
func (d *FollowAuthor) GetSendTime() int64 {
	return d.SendTime
}

// GetSendTime 获取弹幕发送时间
func (d *ThrowBanana) GetSendTime() int64 {
	return d.SendTime
}

// GetSendTime 获取弹幕发送时间
func (d *Gift) GetSendTime() int64 {
	return d.SendTime
}

// GetSendTime 获取弹幕发送时间
func (d *RichText) GetSendTime() int64 {
	return d.SendTime
}

// WatchingUser 就是观看直播的用户的信息，目前没有Medal和ManagerType
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
	RecentComment    []Comment    // APP进直播间时显示的最近发的弹幕
	RedpackList      []Redpack    // 红包列表
}

// 带锁的LiveInfo
type liveInfo struct {
	sync.Mutex // LiveInfo的锁
	LiveInfo
}

// DanmuQueue 就是弹幕的队列
type DanmuQueue struct {
	q    *queue.Queue // DanmuMessage的队列
	info *liveInfo    // 直播间的相关信息状态
	t    *token       // 令牌相关信息
	ch   chan error   // 用来传递初始化的错误
}

// Login 用帐号邮箱和密码登陆AcFun获取cookies
func Login(username, password string) ([]*http.Cookie, error) {
	if username != "" && password != "" {
		return login(username, password)
	}
	return nil, fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆")
}

// Start 启动websocket获取弹幕，uid是主播的uid，ctx用来结束websocket。
// cookies可以利用Login()获取，为nil时使用访客模式登陆AcFun的弹幕系统，通常使用访客模式即可。
func Start(ctx context.Context, uid int, cookies []*http.Cookie) (dq *DanmuQueue, err error) {
	dq = new(DanmuQueue)
	dq.q = queue.New(queueLen)
	dq.info = new(liveInfo)
	dq.ch = make(chan error, 1)
	go dq.wsStart(ctx, uid, cookies)
	if err = <-dq.ch; err != nil {
		return nil, err
	}
	return dq, nil
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播）
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

// GetInfo 返回直播间的状态信息info
func (dq *DanmuQueue) GetInfo() (info LiveInfo) {
	dq.info.Lock()
	defer dq.info.Unlock()
	info = dq.info.LiveInfo
	return info
}

// GetWatchingList 返回直播间排名前50的在线观众信息列表
func (dq *DanmuQueue) GetWatchingList(cookies []*http.Cookie) (*[]WatchingUser, error) {
	return dq.t.watchingList(cookies)
}
