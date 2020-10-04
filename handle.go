package acfundanmu

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"sync"

	"github.com/dgrr/fastws"
	"github.com/orzogc/acfundanmu/acproto"
	"github.com/valyala/fastjson"

	"google.golang.org/protobuf/proto"
)

// 处理接受到的数据里的命令
func (dq *DanmuQueue) handleCommand(conn *fastws.Conn, stream *acproto.DownstreamPayload, hb chan<- int64, event bool) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("handleCommand() error: %w", err)
		}
	}()

	if stream == nil {
		panic(fmt.Errorf("stream为nil"))
	}

	switch stream.Command {
	case "Global.ZtLiveInteractive.CsCmd":
		cmd := &acproto.ZtLiveCsCmdAck{}
		err := proto.Unmarshal(stream.PayloadData, cmd)
		checkErr(err)
		switch cmd.CmdAckType {
		case "ZtLiveCsEnterRoomAck":
			enterRoom := &acproto.ZtLiveCsEnterRoomAck{}
			err = proto.Unmarshal(cmd.Payload, enterRoom)
			checkErr(err)
			if enterRoom.HeartbeatIntervalMs > 0 {
				hb <- enterRoom.HeartbeatIntervalMs
			} else {
				hb <- 10000
			}
		case "ZtLiveCsHeartbeatAck":
			heartbeat := &acproto.ZtLiveCsHeartbeatAck{}
			err = proto.Unmarshal(cmd.Payload, heartbeat)
			checkErr(err)
		case "ZtLiveCsUserExitAck":
			//userExit := &acproto.ZtLiveCsUserExitAck{}
			//err = proto.Unmarshal(cmd.Payload, userExit)
			//checkErr(err)
		default:
			log.Printf("未知的cmd.CmdAckType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				cmd.CmdAckType,
				string(cmd.Payload),
				base64.StdEncoding.EncodeToString(cmd.Payload))
		}
	case "Basic.KeepAlive":
		keepalive := &acproto.KeepAliveResponse{}
		err := proto.Unmarshal(stream.PayloadData, keepalive)
		checkErr(err)
	case "Basic.Ping":
		ping := &acproto.PingResponse{}
		err := proto.Unmarshal(stream.PayloadData, ping)
		checkErr(err)
	case "Basic.Unregister":
		unregister := &acproto.UnregisterResponse{}
		err := proto.Unmarshal(stream.PayloadData, unregister)
		checkErr(err)
		conn.CloseString("Unregister")
	case "Push.ZtLiveInteractive.Message":
		_, err := conn.WriteMessage(fastws.ModeBinary, dq.t.pushMessage())
		checkErr(err)
		message := &acproto.ZtLiveScMessage{}
		err = proto.Unmarshal(stream.PayloadData, message)
		checkErr(err)
		payload := message.Payload
		if message.CompressionType == acproto.ZtLiveScMessage_GZIP {
			r, err := gzip.NewReader(bytes.NewReader(message.Payload))
			checkErr(err)
			defer r.Close()
			result, err := ioutil.ReadAll(r)
			checkErr(err)
			payload = result
		}
		switch message.MessageType {
		case "ZtLiveScActionSignal":
			dq.handleActionSignal(&payload, event)
		case "ZtLiveScStateSignal":
			dq.handleStateSignal(&payload, event)
		case "ZtLiveScNotifySignal":
			dq.handleNotifySignal(&payload, event)
		case "ZtLiveScStatusChanged":
			statusChanged := &acproto.ZtLiveScStatusChanged{}
			err := proto.Unmarshal(payload, statusChanged)
			checkErr(err)
			if statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_CLOSED || statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_BANNED {
				dq.t.wsStop(conn, "Live closed")
			}
		case "ZtLiveScTicketInvalid":
			ticketInvalid := &acproto.ZtLiveScTicketInvalid{}
			err := proto.Unmarshal(payload, ticketInvalid)
			checkErr(err)
			dq.t.Lock()
			dq.t.ticketIndex = (dq.t.ticketIndex + 1) % len(dq.t.tickets)
			dq.t.Unlock()
			_, err = conn.WriteMessage(fastws.ModeBinary, dq.t.enterRoom())
			checkErr(err)
		default:
			log.Printf("未知的message.MessageType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				message.MessageType,
				string(payload),
				base64.StdEncoding.EncodeToString(payload))
		}
	case "Push.Message":
		// AcFun帐号收到的私信信息
	case "Push.DataUpdate":
	case "Push.SyncSession":
	case "Push.acfun":
	default:
		if stream.ErrorCode > 0 {
			log.Println("Stream Error:", stream.ErrorCode, stream.ErrorMsg)
			if stream.ErrorCode == 10018 {
				dq.t.wsStop(conn, "Log out")
			}
			log.Println(string(stream.ErrorData))
		} else {
			log.Printf("未知的stream.Command：%s\npayload string:\n%s\npayload base64:\n%s\n",
				stream.Command,
				string(stream.PayloadData),
				base64.StdEncoding.EncodeToString(stream.PayloadData))
		}
	}

	return nil
}

// 处理action signal数据
func (dq *DanmuQueue) handleActionSignal(payload *[]byte, event bool) {
	actionSignal := &acproto.ZtLiveScActionSignal{}
	err := proto.Unmarshal(*payload, actionSignal)
	checkErr(err)

	var danmu []DanmuMessage
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, item := range actionSignal.Item {
		for _, pl := range item.Payload {
			wg.Add(1)
			go func(signalType string, pl []byte) {
				defer wg.Done()
				switch signalType {
				case "CommonActionSignalComment":
					comment := &acproto.CommonActionSignalComment{}
					err := proto.Unmarshal(pl, comment)
					checkErr(err)
					d := &Comment{
						DanmuCommon: DanmuCommon{
							SendTime: comment.SendTimeMs * 1e6,
							UserInfo: UserInfo{
								UserID:   comment.UserInfo.UserId,
								Nickname: comment.UserInfo.Nickname,
							},
						},
						Content: comment.Content,
					}
					dq.t.getMoreInfo(&d.UserInfo, comment.UserInfo)
					if event {
						dq.dispatchEvent(commentDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				case "CommonActionSignalLike":
					like := &acproto.CommonActionSignalLike{}
					err := proto.Unmarshal(pl, like)
					checkErr(err)
					d := &Like{
						SendTime: like.SendTimeMs * 1e6,
						UserInfo: UserInfo{
							UserID:   like.UserInfo.UserId,
							Nickname: like.UserInfo.Nickname,
						},
					}
					dq.t.getMoreInfo(&d.UserInfo, like.UserInfo)
					if event {
						dq.dispatchEvent(likeDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				case "CommonActionSignalUserEnterRoom":
					enter := &acproto.CommonActionSignalUserEnterRoom{}
					err := proto.Unmarshal(pl, enter)
					checkErr(err)
					d := &EnterRoom{
						SendTime: enter.SendTimeMs * 1e6,
						UserInfo: UserInfo{
							UserID:   enter.UserInfo.UserId,
							Nickname: enter.UserInfo.Nickname,
						},
					}
					dq.t.getMoreInfo(&d.UserInfo, enter.UserInfo)
					if event {
						dq.dispatchEvent(enterRoomDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				case "CommonActionSignalUserFollowAuthor":
					follow := &acproto.CommonActionSignalUserFollowAuthor{}
					err := proto.Unmarshal(pl, follow)
					checkErr(err)
					d := &FollowAuthor{
						SendTime: follow.SendTimeMs * 1e6,
						UserInfo: UserInfo{
							UserID:   follow.UserInfo.UserId,
							Nickname: follow.UserInfo.Nickname,
						},
					}
					dq.t.getMoreInfo(&d.UserInfo, follow.UserInfo)
					if event {
						dq.dispatchEvent(followAuthorDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				/*
					case "CommonNotifySignalKickedOut":
						t.handleNotifySignal(signalType, &pl, info)
					case "CommonNotifySignalViolationAlert":
						t.handleNotifySignal(signalType, &pl, info)
					case "CommonNotifySignalLiveManagerState":
						t.handleNotifySignal(signalType, &pl, info)
				*/
				case "AcfunActionSignalThrowBanana":
					banana := &acproto.AcfunActionSignalThrowBanana{}
					err := proto.Unmarshal(pl, banana)
					checkErr(err)
					d := &ThrowBanana{
						DanmuCommon: DanmuCommon{
							SendTime: banana.SendTimeMs * 1e6,
							UserInfo: UserInfo{
								UserID:   banana.Visitor.UserId,
								Nickname: banana.Visitor.Name,
							},
						},
						BananaCount: int(banana.Count),
					}
					if event {
						dq.dispatchEvent(throwBananaDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				case "CommonActionSignalGift":
					gift := &acproto.CommonActionSignalGift{}
					err := proto.Unmarshal(pl, gift)
					checkErr(err)
					// 礼物列表应该不会在直播中途改变，但以防万一
					g, ok := dq.t.gifts[gift.GiftId]
					if !ok {
						g = GiftDetail{
							GiftID:   gift.GiftId,
							GiftName: "未知礼物",
						}
					}
					d := &Gift{
						DanmuCommon: DanmuCommon{
							SendTime: gift.SendTimeMs * 1e6,
							UserInfo: UserInfo{
								UserID:   gift.User.UserId,
								Nickname: gift.User.Nickname,
							},
						},
						GiftDetail:            g,
						Count:                 gift.Count,
						Combo:                 gift.Combo,
						Value:                 gift.Value,
						ComboID:               gift.ComboId,
						SlotDisplayDurationMs: gift.SlotDisplayDurationMs,
						ExpireDurationMs:      gift.ExpireDurationMs,
					}
					dq.t.getMoreInfo(&d.UserInfo, gift.User)
					if gift.DrawGiftInfo != nil {
						d.DrawGiftInfo = DrawGiftInfo{
							ScreenWidth:  gift.DrawGiftInfo.ScreenWidth,
							ScreenHeight: gift.DrawGiftInfo.ScreenHeight,
						}
						d.DrawGiftInfo.DrawPoint = make([]DrawPoint, len(gift.DrawGiftInfo.DrawPoint))
						for i, drawPoint := range gift.DrawGiftInfo.DrawPoint {
							d.DrawGiftInfo.DrawPoint[i] = DrawPoint{
								MarginLeft: drawPoint.MarginLeft,
								MarginTop:  drawPoint.MarginTop,
								ScaleRatio: drawPoint.ScaleRatio,
								Handup:     drawPoint.Handup,
							}
						}
					}
					if event {
						dq.dispatchEvent(giftDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				case "CommonActionSignalRichText":
					richText := &acproto.CommonActionSignalRichText{}
					err := proto.Unmarshal(pl, richText)
					checkErr(err)
					d := &RichText{
						SendTime: richText.SendTimeMs,
					}
					d.Segments = make([]interface{}, len(richText.Segments))
					for i, segment := range richText.Segments {
						switch segment := segment.Segment.(type) {
						case *acproto.CommonActionSignalRichText_RichTextSegment_UserInfo:
							userInfo := RichTextUserInfo{
								UserInfo: UserInfo{
									UserID:   segment.UserInfo.User.UserId,
									Nickname: segment.UserInfo.User.Nickname,
								},
								Color: segment.UserInfo.Color,
							}
							dq.t.getMoreInfo(&userInfo.UserInfo, segment.UserInfo.User)
							d.Segments[i] = userInfo
						case *acproto.CommonActionSignalRichText_RichTextSegment_Plain:
							plain := RichTextPlain{
								Text:  segment.Plain.Text,
								Color: segment.Plain.Color,
							}
							d.Segments[i] = plain
						case *acproto.CommonActionSignalRichText_RichTextSegment_Image:
							image := RichTextImage{
								AlternativeText:  segment.Image.AlternativeText,
								AlternativeColor: segment.Image.AlternativeColor,
							}
							image.Pictures = make([]string, len(segment.Image.Pictures))
							for j, picture := range segment.Image.Pictures {
								image.Pictures[j] = picture.Url
							}
							d.Segments[i] = image
						}
					}
					if event {
						dq.dispatchEvent(richTextDanmu, d)
					} else {
						mu.Lock()
						danmu = append(danmu, d)
						mu.Unlock()
					}
				default:
					log.Printf("未知的Action Signal item.SignalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
						signalType,
						string(pl),
						base64.StdEncoding.EncodeToString(pl))
				}
			}(item.SignalType, pl)
		}
	}
	wg.Wait()

	if !event {
		// 按SendTime大小排序
		sort.Slice(danmu, func(i, j int) bool {
			return danmu[i].GetSendTime() < danmu[j].GetSendTime()
		})

		for _, d := range danmu {
			err = dq.q.Put(d)
			checkErr(err)
		}
	}
}

// 处理state signal数据
func (dq *DanmuQueue) handleStateSignal(payload *[]byte, event bool) {
	stateSignal := &acproto.ZtLiveScStateSignal{}
	err := proto.Unmarshal(*payload, stateSignal)
	checkErr(err)

	var wg sync.WaitGroup
	for _, item := range stateSignal.Item {
		wg.Add(1)
		go func(item *acproto.ZtLiveStateSignalItem) {
			defer wg.Done()
			switch item.SignalType {
			case "AcfunStateSignalDisplayInfo":
				bananaInfo := &acproto.AcfunStateSignalDisplayInfo{}
				err := proto.Unmarshal(item.Payload, bananaInfo)
				checkErr(err)
				dq.info.Lock()
				dq.info.AllBananaCount = bananaInfo.BananaCount
				dq.info.Unlock()
			case "CommonStateSignalDisplayInfo":
				stateInfo := &acproto.CommonStateSignalDisplayInfo{}
				err := proto.Unmarshal(item.Payload, stateInfo)
				checkErr(err)
				dq.info.Lock()
				dq.info.WatchingCount = stateInfo.WatchingCount
				dq.info.LikeCount = stateInfo.LikeCount
				dq.info.LikeDelta = int(stateInfo.LikeDelta)
				dq.info.Unlock()
			case "CommonStateSignalTopUsers":
				topUsers := &acproto.CommonStateSignalTopUsers{}
				err := proto.Unmarshal(item.Payload, topUsers)
				checkErr(err)
				users := make([]TopUser, len(topUsers.User))
				for i, user := range topUsers.User {
					u := TopUser{
						UserInfo: UserInfo{
							UserID:   user.UserInfo.UserId,
							Nickname: user.UserInfo.Nickname,
						},
						AnonymousUser:     user.AnonymousUser,
						DisplaySendAmount: user.DisplaySendAmount,
						CustomData:        user.CustomWatchingListData,
					}
					dq.t.getMoreInfo(&u.UserInfo, user.UserInfo)
					users[i] = u
				}
				dq.info.Lock()
				dq.info.TopUsers = users
				dq.info.Unlock()
			case "CommonStateSignalRecentComment":
				comments := &acproto.CommonStateSignalRecentComment{}
				err := proto.Unmarshal(item.Payload, comments)
				checkErr(err)
				danmu := make([]Comment, len(comments.Comment))
				for i, comment := range comments.Comment {
					d := Comment{
						DanmuCommon: DanmuCommon{
							SendTime: comment.SendTimeMs * 1e6,
							UserInfo: UserInfo{
								UserID:   comment.UserInfo.UserId,
								Nickname: comment.UserInfo.Nickname,
							},
						},
						Content: comment.Content,
					}
					dq.t.getMoreInfo(&d.UserInfo, comment.UserInfo)
					danmu[i] = d
				}
				dq.info.Lock()
				dq.info.RecentComment = danmu
				dq.info.Unlock()
			case "CommonStateSignalChatCall":
				chatCall := &acproto.CommonStateSignalChatCall{}
				err := proto.Unmarshal(item.Payload, chatCall)
				checkErr(err)
				chat := ChatInfo{
					ChatID:          chatCall.ChatId,
					LiveID:          chatCall.LiveId,
					CallTimestampMs: chatCall.CallTimestampMs,
				}
				dq.info.Lock()
				dq.info.Chat = chat
				dq.info.Unlock()
			case "CommonStateSignalChatAccept":
				chatAccept := &acproto.CommonStateSignalChatAccept{}
				err := proto.Unmarshal(item.Payload, chatAccept)
				checkErr(err)
				log.Printf("CommonStateSignalChatAccept: %+v\n", chatAccept)
			case "CommonStateSignalChatReady":
				chatReady := &acproto.CommonStateSignalChatReady{}
				err := proto.Unmarshal(item.Payload, chatReady)
				checkErr(err)
				guest := UserInfo{
					UserID:   chatReady.GuestUserInfo.UserId,
					Nickname: chatReady.GuestUserInfo.Nickname,
				}
				dq.t.getMoreInfo(&guest, chatReady.GuestUserInfo)
				dq.info.Lock()
				dq.info.Chat.ChatID = chatReady.ChatId
				dq.info.Chat.Guest = guest
				dq.info.Chat.MediaType = ChatMediaType(chatReady.MediaType)
				dq.info.Unlock()
			case "CommonStateSignalChatEnd":
				chatEnd := &acproto.CommonStateSignalChatEnd{}
				err := proto.Unmarshal(item.Payload, chatEnd)
				checkErr(err)
				dq.info.Lock()
				dq.info.Chat.ChatID = chatEnd.ChatId
				dq.info.Chat.EndType = ChatEndType(chatEnd.EndType)
				dq.info.Unlock()
			case "CommonStateSignalCurrentRedpackList":
				redpackList := &acproto.CommonStateSignalCurrentRedpackList{}
				err := proto.Unmarshal(item.Payload, redpackList)
				checkErr(err)
				redpacks := make([]Redpack, len(redpackList.Redpacks))
				for i, redpack := range redpackList.Redpacks {
					r := Redpack{
						UserInfo: UserInfo{
							UserID:   redpack.Sender.UserId,
							Nickname: redpack.Sender.Nickname,
						},
						DisplayStatus:        RedpackDisplayStatus(redpack.DisplayStatus),
						GrabBeginTimeMs:      redpack.GrabBeginTimeMs,
						GetTokenLatestTimeMs: redpack.GetTokenLatestTimeMs,
						RedPackID:            redpack.RedPackId,
						RedpackBizUnit:       redpack.RedpackBizUnit,
						RedpackAmount:        redpack.RedpackAmount,
						SettleBeginTime:      redpack.SettleBeginTime,
					}
					dq.t.getMoreInfo(&r.UserInfo, redpack.Sender)
					redpacks[i] = r
				}
				dq.info.Lock()
				dq.info.RedpackList = redpacks
				dq.info.Unlock()
			default:
				log.Printf("未知的State Signal item.SignalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
					item.SignalType,
					string(item.Payload),
					base64.StdEncoding.EncodeToString(item.Payload))
			}
		}(item)
	}
	wg.Wait()
}

// 处理notify signal数据
func (dq *DanmuQueue) handleNotifySignal(payload *[]byte, event bool) {
	notifySignal := &acproto.ZtLiveScNotifySignal{}
	err := proto.Unmarshal(*payload, notifySignal)
	checkErr(err)

	var wg sync.WaitGroup
	for _, item := range notifySignal.Item {
		wg.Add(1)
		go func(item *acproto.ZtLiveNotifySignalItem) {
			defer wg.Done()
			switch item.SignalType {
			case "CommonNotifySignalKickedOut":
				kickedOut := &acproto.CommonNotifySignalKickedOut{}
				err := proto.Unmarshal(item.Payload, kickedOut)
				checkErr(err)
				dq.info.Lock()
				dq.info.KickedOut = kickedOut.Reason
				dq.info.Unlock()
			case "CommonNotifySignalViolationAlert":
				violationAlert := &acproto.CommonNotifySignalViolationAlert{}
				err := proto.Unmarshal(item.Payload, violationAlert)
				checkErr(err)
				dq.info.Lock()
				dq.info.ViolationAlert = violationAlert.ViolationContent
				dq.info.Unlock()
			case "CommonNotifySignalLiveManagerState":
				liveManagerState := &acproto.CommonNotifySignalLiveManagerState{}
				err := proto.Unmarshal(item.Payload, liveManagerState)
				checkErr(err)
				dq.info.Lock()
				dq.info.LiveManagerState = ManagerState(liveManagerState.State)
				dq.info.Unlock()
			default:
				log.Printf("未知的Notify Signal signalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
					item.SignalType,
					string(item.Payload),
					base64.StdEncoding.EncodeToString(item.Payload))
			}
		}(item)
	}
	wg.Wait()
}

// 获取用户的头像、守护徽章和房管类型
func (t *token) getMoreInfo(user *UserInfo, userInfo *acproto.ZtLiveUserInfo) {
	if len(userInfo.Avatar) != 0 {
		user.Avatar = userInfo.Avatar[0].Url
	}

	if userInfo.Badge != "" {
		p := t.medalParser.Get()
		defer t.medalParser.Put(p)
		v, err := p.Parse(userInfo.Badge)
		checkErr(err)
		o := v.GetObject("medalInfo")
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "uperId":
				user.Medal.UperID = v.GetInt64()
			case "clubName":
				user.Medal.ClubName = string(v.GetStringBytes())
			case "level":
				user.Medal.Level = v.GetInt()
			}
		})
	}

	if userInfo.UserIdentity != nil {
		user.ManagerType = ManagerType(userInfo.UserIdentity.ManagerType)
	}
}
