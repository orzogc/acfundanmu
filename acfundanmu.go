package acfundanmu

import (
	"context"
	"sort"

	"github.com/Workiva/go-datastructures/queue"
)

// 队列长度
const queueLen = 1000

// Comment 就是弹幕的数据
type Comment struct {
	SendTime int64  // 弹幕发送时间，是以纳秒为单位的Unix时间
	UserID   int64  // 用户uid
	Nickname string // 用户名字
	Content  string // 弹幕内容
}

// Queue 就是弹幕的队列
type Queue struct {
	q  *queue.Queue // 弹幕的队列
	ch chan bool    // 用来通知websocket启动
}

// Start 启动websocket获取弹幕，uid是主播的uid，ctx用来结束websocket
func Start(ctx context.Context, uid int) Queue {
	q := queue.New(queueLen)
	ch := make(chan bool, 1)
	qs := Queue{q: q, ch: ch}
	go wsStart(ctx, uid, q, "", "", ch)
	return qs
}

// GetDanmu 返回弹幕数据，返回nil时说明弹幕获取结束（出现错误或者主播可能下播）
func (q Queue) GetDanmu() (comments []Comment) {
	if (*queue.Queue)(q.q).Disposed() {
		return nil
	}
	coms, err := (*queue.Queue)(q.q).Get(queueLen)
	if err != nil {
		return nil
	}

	comments = make([]Comment, len(coms))
	for i, c := range coms {
		comments[i] = c.(Comment)
	}
	// 按SendTime大小排序
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].SendTime < comments[j].SendTime
	})

	return comments
}
