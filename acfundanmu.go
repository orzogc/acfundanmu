package acfundanmu

import (
	"context"

	"github.com/Workiva/go-datastructures/queue"
)

// 队列长度
const queueLen = 1000

// DanmuType 就是弹幕信息的类型
type DanmuType int

const (
	// Comment 弹幕
	Comment DanmuType = iota
	// Like 点赞
	Like
	// EnterRoom 用户进入直播间
	EnterRoom
	// FollowAuthor 用户关注了主播
	FollowAuthor
	// ThrowBanana 投蕉
	ThrowBanana
	// Gift 礼物
	Gift
)

// Giftdetail 就是礼物的详细信息
type Giftdetail struct {
	ID          int    // 礼物ID
	Name        string // 礼物名字
	Price       int    // 礼物价格，非免费礼物时单位为AC币，免费礼物（香蕉）时为1
	WebpPic     string // 礼物的webp格式图片（动图）
	PngPic      string // 礼物的png格式图片（大）
	SmallPngPic string // 礼物的png格式图片（小）
	Description string // 礼物的描述
}

// GiftInfo 就是礼物信息
type GiftInfo struct {
	Gift                  Giftdetail // 礼物详细信息
	Count                 int        // 礼物数量
	Combo                 int        // 礼物连击数量
	Value                 int        // 礼物价值，非免费礼物时单位为AC币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID               string     // 礼物连击ID
	SlotDisplayDurationMs int
	ExpireDurationMs      int
}

// DanmuMessage 就是websocket接受到的弹幕相关信息。
// 不论是哪种Type都会有SendTime、UserID、Nickname。
// Type为Comment时，Comment就是弹幕文字。
// Type为Gift时，GiftInfo就是礼物信息。
// Type为Like、EnterRoom和FollowAuthor时没有多余的数据。
// Type为ThrowBanana时，BananaCount就是投蕉数量，不过现在好像都是用Gift代替。
type DanmuMessage struct {
	Type        DanmuType // 弹幕类型
	SendTime    int64     // 弹幕发送时间，是以纳秒为单位的Unix时间
	UserID      int64     // 用户uid
	Nickname    string    // 用户名字
	Comment     string    // 弹幕内容
	BananaCount int       // 投蕉数量
	GiftInfo    GiftInfo  // 礼物信息
}

//TopUser 就是礼物榜在线前三
type TopUser struct {
	UserID                 int64  // 用户uid
	Nickname               string // 用户名字
	CustomWatchingListData string
	DisplaySendAmount      string
	AnonymousUser          bool
}

// LiveInfo 就是直播间的相关信息状态
type LiveInfo struct {
	KickedOut      string         // 被踢理由？
	ViolationAlert string         // 直播间警告？
	AllBananaCount string         // 直播间香蕉总数
	WatchingCount  string         // 直播间在线观众数量
	LikeCount      string         // 直播间点赞总数
	LikeDelta      int            // 点赞增加数量？
	TopUsers       []TopUser      // 礼物榜在线前三
	RecentComment  []DanmuMessage // 进直播间时显示的最近发的弹幕
}

// DanmuQueue 就是弹幕的队列
type DanmuQueue struct {
	q    *queue.Queue // DanmuMessage的队列
	info *LiveInfo    // 直播间的相关信息状态
	ch   chan bool    // 用来通知websocket启动
}

// Start 启动websocket获取弹幕，uid是主播的uid，ctx用来结束websocket
func Start(ctx context.Context, uid int) (dq DanmuQueue) {
	q := queue.New(queueLen)
	info := new(LiveInfo)
	ch := make(chan bool, 1)
	dq = DanmuQueue{q: q, info: info, ch: ch}
	go dq.wsStart(ctx, uid, "", "")
	return dq
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播）
func (dq DanmuQueue) GetDanmu() (danmu []DanmuMessage) {
	if (*queue.Queue)(dq.q).Disposed() {
		return nil
	}
	ds, err := (*queue.Queue)(dq.q).Get(queueLen)
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
func (dq DanmuQueue) GetInfo() (info LiveInfo) {
	return *dq.info
}
