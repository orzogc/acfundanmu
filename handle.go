package acfundanmu

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/dgrr/fastws"
	"github.com/orzogc/acfundanmu/acproto"
	"github.com/valyala/fastjson"

	"google.golang.org/protobuf/proto"
)

// 处理接受到的数据里的命令
func (ac *AcFunLive) handleCommand(ctx context.Context, conn *fastws.Conn, stream *acproto.DownstreamPayload, event bool) (e error) {
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
				go ac.t.wsHeartbeat(ctx, conn, enterRoom.HeartbeatIntervalMs)
			} else {
				go ac.t.wsHeartbeat(ctx, conn, 10000)
			}
		case "ZtLiveCsHeartbeatAck":
			//heartbeat := &acproto.ZtLiveCsHeartbeatAck{}
			//err = proto.Unmarshal(cmd.Payload, heartbeat)
			//checkErr(err)
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
		//keepalive := &acproto.KeepAliveResponse{}
		//err := proto.Unmarshal(stream.PayloadData, keepalive)
		//checkErr(err)
	case "Basic.Ping":
		//ping := &acproto.PingResponse{}
		//err := proto.Unmarshal(stream.PayloadData, ping)
		//checkErr(err)
	case "Basic.Unregister":
		unregister := &acproto.UnregisterResponse{}
		err := proto.Unmarshal(stream.PayloadData, unregister)
		checkErr(err)
		_ = conn.CloseString("Unregister")
	case "Push.ZtLiveInteractive.Message":
		_, err := conn.WriteMessage(fastws.ModeBinary, ac.t.pushMessage())
		checkErr(err)
		message := &acproto.ZtLiveScMessage{}
		err = proto.Unmarshal(stream.PayloadData, message)
		checkErr(err)
		payload := message.Payload
		if message.CompressionType == acproto.ZtLiveScMessage_GZIP {
			r, err := gzip.NewReader(bytes.NewReader(message.Payload))
			checkErr(err)
			defer r.Close()
			payload, err = ioutil.ReadAll(r)
			checkErr(err)
		}
		switch message.MessageType {
		case "ZtLiveScActionSignal":
			ac.handleActionSignal(&payload, event)
		case "ZtLiveScStateSignal":
			ac.handleStateSignal(&payload, event)
		case "ZtLiveScNotifySignal":
			ac.handleNotifySignal(&payload, event)
		case "ZtLiveScStatusChanged":
			statusChanged := &acproto.ZtLiveScStatusChanged{}
			err = proto.Unmarshal(payload, statusChanged)
			checkErr(err)
			if statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_CLOSED || statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_BANNED {
				ac.t.wsStop(conn, "Live closed")
			}
		case "ZtLiveScTicketInvalid":
			ticketInvalid := &acproto.ZtLiveScTicketInvalid{}
			err = proto.Unmarshal(payload, ticketInvalid)
			checkErr(err)
			index := ac.t.ticketIndex.Load()
			_ = ac.t.ticketIndex.CAS(index, (index+1)%uint32(len(ac.t.tickets)))
			_, err = conn.WriteMessage(fastws.ModeBinary, ac.t.enterRoom())
			checkErr(err)
		default:
			log.Printf("未知的message.MessageType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				message.MessageType,
				string(payload),
				base64.StdEncoding.EncodeToString(payload))
		}
	// AcFun帐号收到的私信信息
	case "Push.Message":
	case "Push.DataUpdate":
	case "Push.SyncSession":
	case "Push.acfun":
	default:
		if stream.ErrorCode > 0 {
			log.Println("DownstreamPayload error:", stream.ErrorCode, stream.ErrorMsg)
			if stream.ErrorCode == 10018 {
				ac.t.wsStop(conn, "Log out")
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
func (ac *AcFunLive) handleActionSignal(payload *[]byte, event bool) {
	actionSignal := &acproto.ZtLiveScActionSignal{}
	err := proto.Unmarshal(*payload, actionSignal)
	checkErr(err)

	var danmu []DanmuMessage
	for _, item := range actionSignal.Item {
		for _, pl := range item.Payload {
			switch item.SignalType {
			case "CommonActionSignalComment":
				comment := &acproto.CommonActionSignalComment{}
				err = proto.Unmarshal(pl, comment)
				checkErr(err)
				d := &Comment{
					DanmuCommon: DanmuCommon{
						SendTime: comment.SendTimeMs,
						UserInfo: UserInfo{
							UserID:   comment.UserInfo.UserId,
							Nickname: comment.UserInfo.Nickname,
						},
					},
					Content: comment.Content,
				}
				ac.t.getMoreInfo(&d.UserInfo, comment.UserInfo)
				danmu = append(danmu, d)
			case "CommonActionSignalLike":
				like := &acproto.CommonActionSignalLike{}
				err = proto.Unmarshal(pl, like)
				checkErr(err)
				d := &Like{
					SendTime: like.SendTimeMs,
					UserInfo: UserInfo{
						UserID:   like.UserInfo.UserId,
						Nickname: like.UserInfo.Nickname,
					},
				}
				ac.t.getMoreInfo(&d.UserInfo, like.UserInfo)
				danmu = append(danmu, d)
			case "CommonActionSignalUserEnterRoom":
				enter := &acproto.CommonActionSignalUserEnterRoom{}
				err = proto.Unmarshal(pl, enter)
				checkErr(err)
				d := &EnterRoom{
					SendTime: enter.SendTimeMs,
					UserInfo: UserInfo{
						UserID:   enter.UserInfo.UserId,
						Nickname: enter.UserInfo.Nickname,
					},
				}
				ac.t.getMoreInfo(&d.UserInfo, enter.UserInfo)
				danmu = append(danmu, d)
			case "CommonActionSignalUserFollowAuthor":
				follow := &acproto.CommonActionSignalUserFollowAuthor{}
				err = proto.Unmarshal(pl, follow)
				checkErr(err)
				d := &FollowAuthor{
					SendTime: follow.SendTimeMs,
					UserInfo: UserInfo{
						UserID:   follow.UserInfo.UserId,
						Nickname: follow.UserInfo.Nickname,
					},
				}
				ac.t.getMoreInfo(&d.UserInfo, follow.UserInfo)
				danmu = append(danmu, d)
			case "AcfunActionSignalThrowBanana":
				banana := &acproto.AcfunActionSignalThrowBanana{}
				err = proto.Unmarshal(pl, banana)
				checkErr(err)
				d := &ThrowBanana{
					DanmuCommon: DanmuCommon{
						SendTime: banana.SendTimeMs,
						UserInfo: UserInfo{
							UserID:   banana.Visitor.UserId,
							Nickname: banana.Visitor.Name,
						},
					},
					BananaCount: int(banana.Count),
				}
				danmu = append(danmu, d)
			case "CommonActionSignalGift":
				gift := &acproto.CommonActionSignalGift{}
				err = proto.Unmarshal(pl, gift)
				checkErr(err)
				// 礼物列表应该不会在直播中途改变，但以防万一
				g, ok := ac.t.gifts[gift.GiftId]
				if !ok {
					g = GiftDetail{
						GiftID:   gift.GiftId,
						GiftName: "未知礼物",
					}
				}
				d := &Gift{
					DanmuCommon: DanmuCommon{
						SendTime: gift.SendTimeMs,
						UserInfo: UserInfo{
							UserID:   gift.User.UserId,
							Nickname: gift.User.Nickname,
						},
					},
					GiftDetail:          g,
					Count:               gift.Count,
					Combo:               gift.Combo,
					Value:               gift.Value,
					ComboID:             gift.ComboId,
					SlotDisplayDuration: gift.SlotDisplayDurationMs,
					ExpireDuration:      gift.ExpireDurationMs,
				}
				ac.t.getMoreInfo(&d.UserInfo, gift.User)
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
				danmu = append(danmu, d)
			case "CommonActionSignalRichText":
				richText := &acproto.CommonActionSignalRichText{}
				err = proto.Unmarshal(pl, richText)
				checkErr(err)
				d := &RichText{
					SendTime: richText.SendTimeMs,
				}
				d.Segments = make([]RichTextSegment, len(richText.Segments))
				for i, segment := range richText.Segments {
					switch segment := segment.Segment.(type) {
					case *acproto.CommonActionSignalRichText_RichTextSegment_UserInfo:
						userInfo := &RichTextUserInfo{
							UserInfo: UserInfo{
								UserID:   segment.UserInfo.User.UserId,
								Nickname: segment.UserInfo.User.Nickname,
							},
							Color: segment.UserInfo.Color,
						}
						ac.t.getMoreInfo(&userInfo.UserInfo, segment.UserInfo.User)
						d.Segments[i] = userInfo
					case *acproto.CommonActionSignalRichText_RichTextSegment_Plain:
						plain := &RichTextPlain{
							Text:  segment.Plain.Text,
							Color: segment.Plain.Color,
						}
						d.Segments[i] = plain
					case *acproto.CommonActionSignalRichText_RichTextSegment_Image:
						image := &RichTextImage{
							AlternativeText:  segment.Image.AlternativeText,
							AlternativeColor: segment.Image.AlternativeColor,
						}
						image.Pictures = make([]string, len(segment.Image.Pictures))
						for j, picture := range segment.Image.Pictures {
							image.Pictures[j] = picture.Url
						}
						d.Segments[i] = image
					default:
						log.Println("出现未处理的RichText Segment")
					}
				}
				danmu = append(danmu, d)
			case "AcfunActionSignalJoinClub":
				join := &acproto.AcfunActionSignalJoinClub{}
				err = proto.Unmarshal(pl, join)
				checkErr(err)
				d := &JoinClub{
					JoinTime: join.JoinTimeMs,
					FansInfo: UserInfo{
						UserID:   join.FansInfo.UserId,
						Nickname: join.FansInfo.Name,
					},
					UperInfo: UserInfo{
						UserID:   join.UperInfo.UserId,
						Nickname: join.UperInfo.Name,
					},
				}
				danmu = append(danmu, d)
			default:
				log.Printf("未知的Action Signal item.SignalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
					item.SignalType,
					string(pl),
					base64.StdEncoding.EncodeToString(pl))
			}
		}
	}

	// 按SendTime大小排序
	sort.Slice(danmu, func(i, j int) bool {
		return danmu[i].GetSendTime() < danmu[j].GetSendTime()
	})

	for _, d := range danmu {
		if event {
			switch d := d.(type) {
			case *Comment:
				ac.callEvent(commentDanmu, d)
			case *Like:
				ac.callEvent(likeDanmu, d)
			case *EnterRoom:
				ac.callEvent(enterRoomDanmu, d)
			case *FollowAuthor:
				ac.callEvent(followAuthorDanmu, d)
			case *ThrowBanana:
				ac.callEvent(throwBananaDanmu, d)
			case *Gift:
				ac.callEvent(giftDanmu, d)
			case *RichText:
				ac.callEvent(richTextDanmu, d)
			case *JoinClub:
				ac.callEvent(joinClubDanmu, d)
			default:
				log.Println("出现未处理的DanmuMessage")
			}
		} else {
			err = ac.q.Put(d)
			checkErr(err)
		}
	}
}

// 处理state signal数据
func (ac *AcFunLive) handleStateSignal(payload *[]byte, event bool) {
	stateSignal := &acproto.ZtLiveScStateSignal{}
	err := proto.Unmarshal(*payload, stateSignal)
	checkErr(err)

	for _, item := range stateSignal.Item {
		switch item.SignalType {
		case "AcfunStateSignalDisplayInfo":
			bananaInfo := &acproto.AcfunStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, bananaInfo)
			checkErr(err)
			if event {
				ac.callEvent(bananaCountEvent, bananaInfo.BananaCount)
			} else {
				ac.info.Lock()
				ac.info.AllBananaCount = bananaInfo.BananaCount
				ac.info.Unlock()
			}
		case "CommonStateSignalDisplayInfo":
			stateInfo := &acproto.CommonStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, stateInfo)
			checkErr(err)
			info := DisplayInfo{
				WatchingCount: stateInfo.WatchingCount,
				LikeCount:     stateInfo.LikeCount,
				LikeDelta:     int(stateInfo.LikeDelta),
			}
			if event {
				ac.callEvent(displayEvent, &info)
			} else {
				ac.info.Lock()
				ac.info.DisplayInfo = info
				ac.info.Unlock()
			}
		case "CommonStateSignalTopUsers":
			topUsers := &acproto.CommonStateSignalTopUsers{}
			err = proto.Unmarshal(item.Payload, topUsers)
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
				ac.t.getMoreInfo(&u.UserInfo, user.UserInfo)
				users[i] = u
			}
			if event {
				ac.callEvent(topUsersEvent, users)
			} else {
				ac.info.Lock()
				ac.info.TopUsers = users
				ac.info.Unlock()
			}
		case "CommonStateSignalRecentComment":
			comments := &acproto.CommonStateSignalRecentComment{}
			err = proto.Unmarshal(item.Payload, comments)
			checkErr(err)
			danmu := make([]Comment, len(comments.Comment))
			for i, comment := range comments.Comment {
				d := Comment{
					DanmuCommon: DanmuCommon{
						SendTime: comment.SendTimeMs,
						UserInfo: UserInfo{
							UserID:   comment.UserInfo.UserId,
							Nickname: comment.UserInfo.Nickname,
						},
					},
					Content: comment.Content,
				}
				ac.t.getMoreInfo(&d.UserInfo, comment.UserInfo)
				danmu[i] = d
			}
			if event {
				ac.callEvent(recentCommentEvent, danmu)
			} else {
				ac.info.Lock()
				ac.info.RecentComment = danmu
				ac.info.Unlock()
			}
		case "CommonStateSignalCurrentRedpackList":
			redpackList := &acproto.CommonStateSignalCurrentRedpackList{}
			err = proto.Unmarshal(item.Payload, redpackList)
			checkErr(err)
			redpacks := make([]Redpack, len(redpackList.Redpacks))
			for i, redpack := range redpackList.Redpacks {
				r := Redpack{
					UserInfo: UserInfo{
						UserID:   redpack.Sender.UserId,
						Nickname: redpack.Sender.Nickname,
					},
					DisplayStatus:      RedpackDisplayStatus(redpack.DisplayStatus),
					GrabBeginTime:      redpack.GrabBeginTimeMs,
					GetTokenLatestTime: redpack.GetTokenLatestTimeMs,
					RedpackID:          redpack.RedPackId,
					RedpackBizUnit:     redpack.RedpackBizUnit,
					RedpackAmount:      redpack.RedpackAmount,
					SettleBeginTime:    redpack.SettleBeginTime,
				}
				ac.t.getMoreInfo(&r.UserInfo, redpack.Sender)
				redpacks[i] = r
			}
			if event {
				ac.callEvent(redpackListEvent, redpacks)
			} else {
				ac.info.Lock()
				ac.info.RedpackList = redpacks
				ac.info.Unlock()
			}
		case "CommonStateSignalChatCall":
			chatCall := &acproto.CommonStateSignalChatCall{}
			err = proto.Unmarshal(item.Payload, chatCall)
			checkErr(err)
			if event {
				ac.callEvent(chatCallEvent, &ChatCall{
					ChatID:   chatCall.ChatId,
					LiveID:   chatCall.LiveId,
					CallTime: chatCall.CallTimestampMs,
				})
			}
		case "CommonStateSignalChatAccept":
			chatAccept := &acproto.CommonStateSignalChatAccept{}
			err = proto.Unmarshal(item.Payload, chatAccept)
			checkErr(err)
			if event {
				ac.callEvent(chatAcceptEvent, &ChatAccept{
					ChatID:     chatAccept.ChatId,
					MediaType:  ChatMediaType(chatAccept.MediaType),
					SignalInfo: chatAccept.AryaSignalInfo,
				})
			}
		case "CommonStateSignalChatReady":
			chatReady := &acproto.CommonStateSignalChatReady{}
			err = proto.Unmarshal(item.Payload, chatReady)
			checkErr(err)
			guest := UserInfo{
				UserID:   chatReady.GuestUserInfo.UserId,
				Nickname: chatReady.GuestUserInfo.Nickname,
			}
			ac.t.getMoreInfo(&guest, chatReady.GuestUserInfo)
			if event {
				ac.callEvent(chatReadyEvent, &ChatReady{
					ChatID:    chatReady.ChatId,
					Guest:     guest,
					MediaType: ChatMediaType(chatReady.MediaType),
				})
			}
		case "CommonStateSignalChatEnd":
			chatEnd := &acproto.CommonStateSignalChatEnd{}
			err = proto.Unmarshal(item.Payload, chatEnd)
			checkErr(err)
			if event {
				ac.callEvent(chatEndEvent, &ChatEnd{
					ChatID:  chatEnd.ChatId,
					EndType: ChatEndType(chatEnd.EndType),
				})
			}
		//case "AuthorChatPlayerInfo":
		case "CommonStateSignalAuthorChatCall":
			chatCall := &acproto.CommonStateSignalAuthorChatCall{}
			err = proto.Unmarshal(item.Payload, chatCall)
			checkErr(err)
			inviter := UserInfo{
				UserID:   chatCall.InviterUserInfo.Player.UserId,
				Nickname: chatCall.InviterUserInfo.Player.Nickname,
			}
			ac.t.getMoreInfo(&inviter, chatCall.InviterUserInfo.Player)
			if event {
				ac.callEvent(authorChatCallEvent, &AuthorChatCall{
					Inviter: AuthorChatPlayerInfo{
						UserInfo:               inviter,
						LiveID:                 chatCall.InviterUserInfo.LiveId,
						EnableJumpPeerLiveRoom: chatCall.InviterUserInfo.EnableJumpPeerLiveRoom,
					},
					ChatID:   chatCall.AuthorChatId,
					CallTime: chatCall.CallTimestampMs,
				})
			}
		case "CommonStateSignalAuthorChatAccept":
			chatAccept := &acproto.CommonStateSignalAuthorChatAccept{}
			err = proto.Unmarshal(item.Payload, chatAccept)
			checkErr(err)
			if event {
				ac.callEvent(authorChatAcceptEvent, &AuthorChatAccept{
					ChatID:     chatAccept.AuthorChatId,
					SignalInfo: chatAccept.AryaSignalInfo,
				})
			}
		case "CommonStateSignalAuthorChatReady":
			chatReady := &acproto.CommonStateSignalAuthorChatReady{}
			err = proto.Unmarshal(item.Payload, chatReady)
			checkErr(err)
			inviter := UserInfo{
				UserID:   chatReady.InviterUserInfo.Player.UserId,
				Nickname: chatReady.InviterUserInfo.Player.Nickname,
			}
			ac.t.getMoreInfo(&inviter, chatReady.InviterUserInfo.Player)
			invitee := UserInfo{
				UserID:   chatReady.InviteeUserInfo.Player.UserId,
				Nickname: chatReady.InviteeUserInfo.Player.Nickname,
			}
			ac.t.getMoreInfo(&invitee, chatReady.InviteeUserInfo.Player)
			if event {
				ac.callEvent(authorChatReadyEvent, &AuthorChatReady{
					Inviter: AuthorChatPlayerInfo{
						UserInfo:               inviter,
						LiveID:                 chatReady.InviterUserInfo.LiveId,
						EnableJumpPeerLiveRoom: chatReady.InviterUserInfo.EnableJumpPeerLiveRoom,
					},
					Invitee: AuthorChatPlayerInfo{
						UserInfo:               invitee,
						LiveID:                 chatReady.InviteeUserInfo.LiveId,
						EnableJumpPeerLiveRoom: chatReady.InviteeUserInfo.EnableJumpPeerLiveRoom,
					},
					ChatID: chatReady.AuthorChatId,
				})
			}
		case "CommonStateSignalAuthorChatEnd":
			chatEnd := &acproto.CommonStateSignalAuthorChatEnd{}
			err = proto.Unmarshal(item.Payload, chatEnd)
			checkErr(err)
			if event {
				ac.callEvent(authorChatEndEvent, &AuthorChatEnd{
					ChatID:    chatEnd.AuthorChatId,
					EndType:   ChatEndType(chatEnd.EndType),
					EndLiveID: chatEnd.EndLiveId,
				})
			}
		case "CommonStateSignalAuthorChatChangeSoundConfig":
			soundConfig := &acproto.CommonStateSignalAuthorChatChangeSoundConfig{}
			err = proto.Unmarshal(item.Payload, soundConfig)
			checkErr(err)
			if event {
				ac.callEvent(authorChatChangeSoundConfigEvent, &AuthorChatChangeSoundConfig{
					ChatID:                soundConfig.AuthorChatId,
					SoundConfigChangeType: SoundConfigChangeType(soundConfig.SoundConfigChangeType),
				})
			}
		case "CommonStateSignalLiveState":
		default:
			log.Printf("未知的State Signal item.SignalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				item.SignalType,
				string(item.Payload),
				base64.StdEncoding.EncodeToString(item.Payload))
		}
	}
}

// 处理notify signal数据
func (ac *AcFunLive) handleNotifySignal(payload *[]byte, event bool) {
	notifySignal := &acproto.ZtLiveScNotifySignal{}
	err := proto.Unmarshal(*payload, notifySignal)
	checkErr(err)

	for _, item := range notifySignal.Item {
		switch item.SignalType {
		case "CommonNotifySignalKickedOut":
			kickedOut := &acproto.CommonNotifySignalKickedOut{}
			err = proto.Unmarshal(item.Payload, kickedOut)
			checkErr(err)
			if event {
				ac.callEvent(kickedOutEvent, kickedOut.Reason)
			} else {
				ac.info.Lock()
				ac.info.KickedOut = kickedOut.Reason
				ac.info.Unlock()
			}
		case "CommonNotifySignalViolationAlert":
			violationAlert := &acproto.CommonNotifySignalViolationAlert{}
			err = proto.Unmarshal(item.Payload, violationAlert)
			checkErr(err)
			if event {
				ac.callEvent(violationAlertEvent, violationAlert.ViolationContent)
			} else {
				ac.info.Lock()
				ac.info.ViolationAlert = violationAlert.ViolationContent
				ac.info.Unlock()
			}
		case "CommonNotifySignalLiveManagerState":
			liveManagerState := &acproto.CommonNotifySignalLiveManagerState{}
			err = proto.Unmarshal(item.Payload, liveManagerState)
			checkErr(err)
			if event {
				ac.callEvent(managerStateEvent, ManagerState(liveManagerState.State))
			} else {
				ac.info.Lock()
				ac.info.LiveManagerState = ManagerState(liveManagerState.State)
				ac.info.Unlock()
			}
		default:
			log.Printf("未知的Notify Signal signalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				item.SignalType,
				string(item.Payload),
				base64.StdEncoding.EncodeToString(item.Payload))
		}
	}
}

// 获取用户的头像、守护徽章和房管类型
func (t *token) getMoreInfo(user *UserInfo, userInfo *acproto.ZtLiveUserInfo) {
	if len(userInfo.Avatar) != 0 {
		user.Avatar = userInfo.Avatar[0].Url
	}

	if userInfo.Badge != "" {
		p := medalParserPool.Get()
		defer medalParserPool.Put(p)
		v, err := p.Parse(userInfo.Badge)
		checkErr(err)
		o := v.GetObject("medalInfo")
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "uperId":
				user.Medal.UperID = v.GetInt64()
			case "userId":
				user.Medal.UserID = v.GetInt64()
			case "clubName":
				user.Medal.ClubName = string(v.GetStringBytes())
			case "level":
				user.Medal.Level = v.GetInt()
			default:
				log.Printf("守护徽章里出现未处理的key和value：%s %s", string(k), string(v.MarshalTo([]byte{})))
			}
		})
	}

	if userInfo.UserIdentity != nil {
		user.ManagerType = ManagerType(userInfo.UserIdentity.ManagerType)
	}
}
