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
	ID          int
	Name        string
	Price       int
	WebpPic     string
	PngPic      string
	SmallPngPic string
	Description string
}

// GiftInfo 就是礼物信息
type GiftInfo struct {
	Gift                  *Giftdetail // 礼物详细信息
	Count                 int         // 礼物数量
	Combo                 int         // 礼物连击数量？
	Value                 int         // 礼物价值
	ComboID               string      // 礼物连击ID？
	SlotDisplayDurationMs int
	ExpireDurationMs      int
}

// DanmuMessage 就是websocket接受到的弹幕相关信息
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
	DisplaySendAmount      string // 送的礼物价值总数？
	AnonymousUser          bool
}

// LiveInfo 就是直播间的相关信息状态
type LiveInfo struct {
	KickedOut      string         // 被踢理由？
	ViolationAlert string         // 直播间警告？
	AllBananaCount string         // 直播间香蕉总数
	WatchingCount  string         // 观众数量
	LikeCount      string         // 点赞数量
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

// GetDanmu 返回弹幕数据，返回danmu为nil时说明弹幕获取结束（出现错误或者主播可能下播）
func (dq DanmuQueue) GetDanmu() (danmu []DanmuMessage, info LiveInfo) {
	if (*queue.Queue)(dq.q).Disposed() {
		return nil, info
	}
	ds, err := (*queue.Queue)(dq.q).Get(queueLen)
	if err != nil {
		return nil, info
	}

	danmu = make([]DanmuMessage, len(ds))
	for i, d := range ds {
		danmu[i] = d.(DanmuMessage)
	}

	return danmu, *dq.info
}
