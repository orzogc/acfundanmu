package acfundanmu

import (
	"log"
	"sync"
)

// 弹幕信息类型
type eventType int

const (
	liveOff eventType = iota
	commentDanmu
	likeDanmu
	enterRoomDanmu
	followAuthorDanmu
	throwBananaDanmu
	giftDanmu
	richTextDanmu
)

// 事件handler
type eventHandler func(*DanmuQueue, interface{})

// 事件handler的map
type handlerMap struct {
	sync.Mutex
	listMap map[eventType][]eventHandler
}

// 将f加入到t对应的事件handler列表里
func (h *handlerMap) add(t eventType, f eventHandler) {
	h.Lock()
	defer h.Unlock()
	h.listMap[t] = append(h.listMap[t], f)
}

// 调用事件handler列表里的handler
func (dq *DanmuQueue) dispatchEvent(t eventType, i interface{}) {
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

// OnLiveOff 处理直播结束信号，handler需要支持并行处理
func (dq *DanmuQueue) OnLiveOff(handler func(*DanmuQueue, string)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(string))
	}
	dq.handlerMap.add(liveOff, f)
}

// OnComment 处理评论弹幕，handler需要支持并行处理
func (dq *DanmuQueue) OnComment(handler func(*DanmuQueue, *Comment)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Comment))
	}
	dq.handlerMap.add(commentDanmu, f)
}

// OnLike 处理点赞弹幕，handler需要支持并行处理
func (dq *DanmuQueue) OnLike(handler func(*DanmuQueue, *Like)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Like))
	}
	dq.handlerMap.add(likeDanmu, f)
}

// OnEnterRoom 处理用户进场，handler需要支持并行处理
func (dq *DanmuQueue) OnEnterRoom(handler func(*DanmuQueue, *EnterRoom)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*EnterRoom))
	}
	dq.handlerMap.add(enterRoomDanmu, f)
}

// OnFollowAuthor 处理用户关注主播，handler需要支持并行处理
func (dq *DanmuQueue) OnFollowAuthor(handler func(*DanmuQueue, *FollowAuthor)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*FollowAuthor))
	}
	dq.handlerMap.add(followAuthorDanmu, f)
}

// OnThrowBanana 处理用户投蕉，现在基本用 OnGift 代替，handler需要支持并行处理
func (dq *DanmuQueue) OnThrowBanana(handler func(*DanmuQueue, *ThrowBanana)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ThrowBanana))
	}
	dq.handlerMap.add(throwBananaDanmu, f)
}

// OnGift 处理用户赠送礼物，handler需要支持并行处理
func (dq *DanmuQueue) OnGift(handler func(*DanmuQueue, *Gift)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Gift))
	}
	dq.handlerMap.add(giftDanmu, f)
}

// OnRichText 处理富文本，handler需要支持并行处理
func (dq *DanmuQueue) OnRichText(handler func(*DanmuQueue, *RichText)) {
	f := func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*RichText))
	}
	dq.handlerMap.add(richTextDanmu, f)
}
