package acfundanmu

import (
	"context"

	"github.com/Workiva/go-datastructures/queue"
)

const queueLen = 1000

// Comment 就是弹幕的数据
type Comment struct {
	SendTime int64  // 弹幕发送时间，单位为毫秒
	UserID   int64  // 用户uid
	Nickname string // 用户名字
	Content  string // 弹幕内容
}

// Queue 就是弹幕的队列
type Queue queue.Queue

// Start 启动websocket获取弹幕，uid是主播的uid，ctx用来结束websocket
func Start(ctx context.Context, uid int) *Queue {
	q := queue.New(queueLen)
	go wsStart(ctx, uid, q, "", "")
	return (*Queue)(q)
}

// GetDanmu 返回弹幕数据，返回nil时说明弹幕获取结束（主播可能下播）
func (q *Queue) GetDanmu() (comments []Comment) {
	if (*queue.Queue)(q).Disposed() {
		return nil
	}
	coms, err := (*queue.Queue)(q).Get(queueLen)
	if err != nil {
		return nil
	}

	for _, c := range coms {
		comments = append(comments, c.(Comment))
	}
	return comments
}

/*
func main() {
	if len(os.Args) != 2 {
		return
	}
	uid, err := strconv.Atoi(os.Args[1])
	checkErr(err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	q := Start(ctx, uid)
	for {
		if d := q.GetDanmu(); d != nil {
			fmt.Println(d)
		} else {
			fmt.Println("直播结束")
			break
		}
	}
}
*/
