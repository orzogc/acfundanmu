package acfundanmu

import (
	"log"
	"sync"
)

// 弹幕信息类型
type danmuType int

const (
	commentType danmuType = iota
	likeType
	enterRoomType
	followAuthorType
	throwBananaType
	giftType
	richTextType
)

// 事件handler
type eventHandler func(*DanmuQueue, interface{})

// 事件handler的map
type handlerMap struct {
	sync.Mutex
	listMap map[danmuType][]eventHandler
}

// 将f加入到t对应的事件handler列表里
func (h *handlerMap) add(t danmuType, f eventHandler) {
	h.Lock()
	defer h.Unlock()
	h.listMap[t] = append(h.listMap[t], f)
}

// 调用事件handler列表里的handler
func (dq *DanmuQueue) dispatchEvent(t danmuType, i interface{}) {
	dq.handlerMap.Lock()
	list, ok := dq.handlerMap.listMap[t]
	dq.handlerMap.Unlock()
	if ok {
		for _, f := range list {
			go func(f eventHandler) {
				defer func() {
					if err := recover(); err != nil {
						log.Printf("dispatchEvent() %v goroutine error: %v", t, err)
					}
				}()
				f(dq, i)
			}(f)
		}
	}
}

// OnComment 处理评论弹幕
func (dq *DanmuQueue) OnComment(handler func(*DanmuQueue, *Comment)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Comment))
	}
	dq.handlerMap.add(commentType, f)
}

// OnLike 处理点赞弹幕
func (dq *DanmuQueue) OnLike(handler func(*DanmuQueue, *Like)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Like))
	}
	dq.handlerMap.add(likeType, f)
}

// OnEnterRoom 处理用户进场
func (dq *DanmuQueue) OnEnterRoom(handler func(*DanmuQueue, *EnterRoom)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*EnterRoom))
	}
	dq.handlerMap.add(enterRoomType, f)
}

// OnFollowAuthor 处理用户关注主播
func (dq *DanmuQueue) OnFollowAuthor(handler func(*DanmuQueue, *FollowAuthor)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*FollowAuthor))
	}
	dq.handlerMap.add(followAuthorType, f)
}

// OnThrowBanana 处理用户投蕉，现在基本用 OnGift 代替
func (dq *DanmuQueue) OnThrowBanana(handler func(*DanmuQueue, *ThrowBanana)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ThrowBanana))
	}
	dq.handlerMap.add(throwBananaType, f)
}

// OnGift 处理用户赠送礼物
func (dq *DanmuQueue) OnGift(handler func(*DanmuQueue, *Gift)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Gift))
	}
	dq.handlerMap.add(giftType, f)
}

// OnRichText 处理富文本
func (dq *DanmuQueue) OnRichText(handler func(*DanmuQueue, *RichText)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*RichText))
	}
	dq.handlerMap.add(richTextType, f)
}
