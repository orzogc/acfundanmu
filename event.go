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
	joinClubDanmu
	bananaCountInfo
	displayInfo
	topUsersInfo
	recentCommentInfo
	chatCallInfo
	chatAcceptInfo
	chatReadyInfo
	chatEndInfo
	redpackListInfo
	kickedOutInfo
	violationAlertInfo
	managerStateInfo
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

// OnLiveOff 处理直播结束信号，有可能是网络原因导致连接超时，直播不一定真的结束，可以多次调用
func (dq *DanmuQueue) OnLiveOff(handler func(*DanmuQueue, error)) {
	dq.handlerMap.add(liveOff, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(error))
	})
}

// OnComment 处理评论弹幕，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnComment(handler func(*DanmuQueue, *Comment)) {
	dq.handlerMap.add(commentDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Comment))
	})
}

// OnLike 处理点赞弹幕，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnLike(handler func(*DanmuQueue, *Like)) {
	dq.handlerMap.add(likeDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Like))
	})
}

// OnEnterRoom 处理用户进场，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnEnterRoom(handler func(*DanmuQueue, *EnterRoom)) {
	dq.handlerMap.add(enterRoomDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*EnterRoom))
	})
}

// OnFollowAuthor 处理用户关注主播，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnFollowAuthor(handler func(*DanmuQueue, *FollowAuthor)) {
	dq.handlerMap.add(followAuthorDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*FollowAuthor))
	})
}

// OnThrowBanana 处理用户投蕉，现在基本用 OnGift 代替，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnThrowBanana(handler func(*DanmuQueue, *ThrowBanana)) {
	dq.handlerMap.add(throwBananaDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ThrowBanana))
	})
}

// OnGift 处理用户赠送礼物，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnGift(handler func(*DanmuQueue, *Gift)) {
	dq.handlerMap.add(giftDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*Gift))
	})
}

// OnRichText 处理富文本，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnRichText(handler func(*DanmuQueue, *RichText)) {
	dq.handlerMap.add(richTextDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*RichText))
	})
}

// OnJoinClub 处理用户加入主播守护团，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnJoinClub(handler func(*DanmuQueue, *JoinClub)) {
	dq.handlerMap.add(joinClubDanmu, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*JoinClub))
	})
}

// OnBananaCount 处理直播间获得的香蕉数，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnBananaCount(handler func(dq *DanmuQueue, allBananaCount string)) {
	dq.handlerMap.add(bananaCountInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(string))
	})
}

// OnDisplayInfo 处理直播间的一些数据，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnDisplayInfo(handler func(*DanmuQueue, *DisplayInfo)) {
	dq.handlerMap.add(displayInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*DisplayInfo))
	})
}

// OnTopUsers 处理直播间礼物榜在线前三的信息，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnTopUsers(handler func(*DanmuQueue, []TopUser)) {
	dq.handlerMap.add(topUsersInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.([]TopUser))
	})
}

// OnRecentComment 处理APP进直播间时显示的最近发的弹幕，可以多次调用
func (dq *DanmuQueue) OnRecentComment(handler func(*DanmuQueue, []Comment)) {
	dq.handlerMap.add(recentCommentInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.([]Comment))
	})
}

// OnChatCall 处理主播发起连麦，可以多次调用
func (dq *DanmuQueue) OnChatCall(handler func(*DanmuQueue, *ChatCall)) {
	dq.handlerMap.add(chatCallInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ChatCall))
	})
}

// OnChatAccept 处理用户接受连麦？一般不会出现这个信号，可以多次调用
func (dq *DanmuQueue) OnChatAccept(handler func(*DanmuQueue, *ChatAccept)) {
	dq.handlerMap.add(chatAcceptInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ChatAccept))
	})
}

// OnChatReady 处理用户接受连麦的信息，可以多次调用
func (dq *DanmuQueue) OnChatReady(handler func(*DanmuQueue, *ChatReady)) {
	dq.handlerMap.add(chatReadyInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ChatReady))
	})
}

// OnChatEnd 处理连麦结束，可以多次调用
func (dq *DanmuQueue) OnChatEnd(handler func(*DanmuQueue, *ChatEnd)) {
	dq.handlerMap.add(chatEndInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(*ChatEnd))
	})
}

// OnRedpackList 处理直播间的红包列表，handler需要支持并行处理，可以多次调用
func (dq *DanmuQueue) OnRedpackList(handler func(*DanmuQueue, []Redpack)) {
	dq.handlerMap.add(redpackListInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.([]Redpack))
	})
}

// OnKickedOut 处理被踢出直播间，可以多次调用
func (dq *DanmuQueue) OnKickedOut(handler func(dq *DanmuQueue, kickedOutReason string)) {
	dq.handlerMap.add(kickedOutInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(string))
	})
}

// OnViolationAlert 处理直播间警告，可以多次调用
func (dq *DanmuQueue) OnViolationAlert(handler func(dq *DanmuQueue, violationContent string)) {
	dq.handlerMap.add(violationAlertInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(string))
	})
}

// OnManagerState 处理登陆帐号的房管状态，可以多次调用
func (dq *DanmuQueue) OnManagerState(handler func(*DanmuQueue, ManagerState)) {
	dq.handlerMap.add(managerStateInfo, func(dq *DanmuQueue, i interface{}) {
		handler(dq, i.(ManagerState))
	})
}
