package acfundanmu

import (
	"log"
	"sync"
)

// 弹幕信息类型
type eventType int

const (
	stopDanmu eventType = iota
	commentDanmu
	likeDanmu
	enterRoomDanmu
	followAuthorDanmu
	throwBananaDanmu
	giftDanmu
	richTextDanmu
	joinClubDanmu
	shareLiveDanmu
	bananaCountEvent
	displayEvent
	topUsersEvent
	recentCommentEvent
	chatCallEvent
	chatAcceptEvent
	chatReadyEvent
	chatEndEvent
	authorChatCallEvent
	authorChatAcceptEvent
	authorChatReadyEvent
	authorChatEndEvent
	authorChatChangeSoundConfigEvent
	redpackListEvent
	kickedOutEvent
	violationAlertEvent
	managerStateEvent
)

// 事件handler
type eventHandler func(*AcFunLive, interface{})

// 事件handler的map
type handlerMap struct {
	sync.RWMutex
	listMap map[eventType][]eventHandler
}

// 将f加入到t对应的事件handler列表里
func (h *handlerMap) add(t eventType, f eventHandler) {
	h.Lock()
	defer h.Unlock()
	h.listMap[t] = append(h.listMap[t], f)
}

// 调用事件handler列表里的handler
func (ac *AcFunLive) callEvent(t eventType, i interface{}) {
	ac.handlerMap.RLock()
	defer ac.handlerMap.RUnlock()
	list, ok := ac.handlerMap.listMap[t]
	if ok {
		for _, f := range list {
			go func(f eventHandler) {
				defer func() {
					if err := recover(); err != nil {
						log.Printf("dispatchEvent() %v goroutine error: %v", t, err)
					}
				}()
				f(ac, i)
			}(f)
		}
	}
}

// OnDanmuStop 处理获取弹幕结束，有可能是网络原因导致连接超时无法获取弹幕，直播不一定结束，可以多次调用
func (ac *AcFunLive) OnDanmuStop(handler func(*AcFunLive, error)) {
	ac.handlerMap.add(stopDanmu, func(ac *AcFunLive, i interface{}) {
		if i == nil {
			handler(ac, nil)
		} else {
			handler(ac, i.(error))
		}
	})
}

// OnComment 处理评论弹幕，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnComment(handler func(*AcFunLive, *Comment)) {
	ac.handlerMap.add(commentDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*Comment))
	})
}

// OnLike 处理点赞弹幕，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnLike(handler func(*AcFunLive, *Like)) {
	ac.handlerMap.add(likeDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*Like))
	})
}

// OnEnterRoom 处理用户进场，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnEnterRoom(handler func(*AcFunLive, *EnterRoom)) {
	ac.handlerMap.add(enterRoomDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*EnterRoom))
	})
}

// OnFollowAuthor 处理用户关注主播，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnFollowAuthor(handler func(*AcFunLive, *FollowAuthor)) {
	ac.handlerMap.add(followAuthorDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*FollowAuthor))
	})
}

// OnThrowBanana 处理用户投蕉，现在基本用 OnGift 代替，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnThrowBanana(handler func(*AcFunLive, *ThrowBanana)) {
	ac.handlerMap.add(throwBananaDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ThrowBanana))
	})
}

// OnGift 处理用户赠送礼物，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnGift(handler func(*AcFunLive, *Gift)) {
	ac.handlerMap.add(giftDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*Gift))
	})
}

// OnRichText 处理富文本，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnRichText(handler func(*AcFunLive, *RichText)) {
	ac.handlerMap.add(richTextDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*RichText))
	})
}

// OnJoinClub 处理用户加入主播守护团，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnJoinClub(handler func(*AcFunLive, *JoinClub)) {
	ac.handlerMap.add(joinClubDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*JoinClub))
	})
}

// OnShareLive 处理分享直播间到其他平台的弹幕，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnShareLive(handler func(*AcFunLive, *ShareLive)) {
	ac.handlerMap.add(shareLiveDanmu, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ShareLive))
	})
}

// OnBananaCount 处理直播间获得的香蕉数，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnBananaCount(handler func(ac *AcFunLive, allBananaCount string)) {
	ac.handlerMap.add(bananaCountEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(string))
	})
}

// OnDisplayInfo 处理直播间的一些数据，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnDisplayInfo(handler func(*AcFunLive, *DisplayInfo)) {
	ac.handlerMap.add(displayEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*DisplayInfo))
	})
}

// OnTopUsers 处理直播间礼物榜在线前三的信息，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnTopUsers(handler func(*AcFunLive, []TopUser)) {
	ac.handlerMap.add(topUsersEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.([]TopUser))
	})
}

// OnRecentComment 处理APP进直播间时显示的最近发的弹幕，可以多次调用
func (ac *AcFunLive) OnRecentComment(handler func(*AcFunLive, []Comment)) {
	ac.handlerMap.add(recentCommentEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.([]Comment))
	})
}

// OnChatCall 处理主播发起连麦，可以多次调用
func (ac *AcFunLive) OnChatCall(handler func(*AcFunLive, *ChatCall)) {
	ac.handlerMap.add(chatCallEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ChatCall))
	})
}

// OnChatAccept 处理用户接受连麦，可以多次调用
func (ac *AcFunLive) OnChatAccept(handler func(*AcFunLive, *ChatAccept)) {
	ac.handlerMap.add(chatAcceptEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ChatAccept))
	})
}

// OnChatReady 处理用户接受连麦的信息，可以多次调用
func (ac *AcFunLive) OnChatReady(handler func(*AcFunLive, *ChatReady)) {
	ac.handlerMap.add(chatReadyEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ChatReady))
	})
}

// OnChatEnd 处理连麦结束，可以多次调用
func (ac *AcFunLive) OnChatEnd(handler func(*AcFunLive, *ChatEnd)) {
	ac.handlerMap.add(chatEndEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*ChatEnd))
	})
}

// OnAuthorChatCall 处理主播发起连麦，可以多次调用
func (ac *AcFunLive) OnAuthorChatCall(handler func(*AcFunLive, *AuthorChatCall)) {
	ac.handlerMap.add(authorChatCallEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*AuthorChatCall))
	})
}

// OnAuthorChatAccept 处理主播接受连麦，可以多次调用
func (ac *AcFunLive) OnAuthorChatAccept(handler func(*AcFunLive, *AuthorChatAccept)) {
	ac.handlerMap.add(authorChatAcceptEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*AuthorChatAccept))
	})
}

// OnAuthorChatReady 处理主播接受连麦的信息，可以多次调用
func (ac *AcFunLive) OnAuthorChatReady(handler func(*AcFunLive, *AuthorChatReady)) {
	ac.handlerMap.add(authorChatReadyEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*AuthorChatReady))
	})
}

// OnAuthorChatEnd 处理连麦结束，可以多次调用
func (ac *AcFunLive) OnAuthorChatEnd(handler func(*AcFunLive, *AuthorChatEnd)) {
	ac.handlerMap.add(authorChatEndEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*AuthorChatEnd))
	})
}

// OnAuthorChatChangeSoundConfig 处理主播连麦声音设置更改，可以多次调用
func (ac *AcFunLive) OnAuthorChatChangeSoundConfig(handler func(*AcFunLive, *AuthorChatChangeSoundConfig)) {
	ac.handlerMap.add(authorChatChangeSoundConfigEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(*AuthorChatChangeSoundConfig))
	})
}

// OnRedpackList 处理直播间的红包列表，handler需要支持并行处理，可以多次调用
func (ac *AcFunLive) OnRedpackList(handler func(*AcFunLive, []Redpack)) {
	ac.handlerMap.add(redpackListEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.([]Redpack))
	})
}

// OnKickedOut 处理被踢出直播间，可以多次调用
func (ac *AcFunLive) OnKickedOut(handler func(ac *AcFunLive, kickedOutReason string)) {
	ac.handlerMap.add(kickedOutEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(string))
	})
}

// OnViolationAlert 处理直播间警告，可以多次调用
func (ac *AcFunLive) OnViolationAlert(handler func(ac *AcFunLive, violationContent string)) {
	ac.handlerMap.add(violationAlertEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(string))
	})
}

// OnManagerState 处理登陆帐号的房管状态，可以多次调用
func (ac *AcFunLive) OnManagerState(handler func(*AcFunLive, ManagerState)) {
	ac.handlerMap.add(managerStateEvent, func(ac *AcFunLive, i interface{}) {
		handler(ac, i.(ManagerState))
	})
}
