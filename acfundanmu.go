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
	// Like 点亮爱心（点赞）
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

// Comment 就是弹幕的数据
/*
type Comment struct {
	SendTime int64  // 弹幕发送时间，是以纳秒为单位的Unix时间
	UserID   int64  // 用户uid
	Nickname string // 用户名字
	Content  string // 弹幕内容
}
*/

// GiftInfo 就是礼物信息
type GiftInfo struct {
	GiftID                int
	Count                 int
	Combo                 int
	Value                 int
	ComboID               string
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
	Gift        GiftInfo  // 礼物信息
}

// DanmuQueue 就是弹幕的队列
type DanmuQueue struct {
	q  *queue.Queue // DanmuMessage的队列
	ch chan bool    // 用来通知websocket启动
}

// Start 启动websocket获取弹幕，uid是主播的uid，ctx用来结束websocket
func Start(ctx context.Context, uid int) (dq DanmuQueue) {
	q := queue.New(queueLen)
	ch := make(chan bool, 1)
	dq = DanmuQueue{q: q, ch: ch}
	go dq.wsStart(ctx, uid, "", "")
	return dq
}

// GetDanmu 返回弹幕数据，返回nil时说明弹幕获取结束（出现错误或者主播可能下播）
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
