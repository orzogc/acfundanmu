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

	"github.com/orzogc/acfundanmu/acproto"

	"github.com/Workiva/go-datastructures/queue"
	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"
)

// 处理接受到的数据里的命令
func (t *token) handleCommand(ctx context.Context, c *websocket.Conn, stream *acproto.DownstreamPayload, q *queue.Queue, info *liveInfo, hb chan<- int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("handleCommand() error: %w", err)
		}
	}()

	if stream == nil {
		return
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
		c.Close(websocket.StatusNormalClosure, "Unregister")
	case "Push.ZtLiveInteractive.Message":
		err := c.Write(ctx, websocket.MessageBinary, *t.pushMessage())
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
			t.handleMsgAct(&payload, q, info)
		case "ZtLiveScStateSignal":
			t.handleMsgState(&payload, info)
		case "ZtLiveScStatusChanged":
			statusChanged := &acproto.ZtLiveScStatusChanged{}
			err := proto.Unmarshal(payload, statusChanged)
			checkErr(err)
			if statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_CLOSED || statusChanged.Type == acproto.ZtLiveScStatusChanged_LIVE_BANNED {
				t.wsStop(ctx, c, "直播已经结束")
			}
		case "ZtLiveScTicketInvalid":
			ticketInvalid := &acproto.ZtLiveScTicketInvalid{}
			err := proto.Unmarshal(payload, ticketInvalid)
			checkErr(err)
			t.ticketIndex = (t.ticketIndex + 1) % len(t.tickets)
			err = c.Write(ctx, websocket.MessageBinary, *t.enterRoom())
			checkErr(err)
		default:
			log.Printf("未知的message.MessageType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				message.MessageType,
				string(payload),
				base64.StdEncoding.EncodeToString(payload))
		}
	case "Push.Message":
		msg := &acproto.Message_Message{}
		err := proto.Unmarshal(stream.PayloadData, msg)
		checkErr(err)
		switch msg.ContentType {
		case int32(acproto.Cloud_Message_TEXT):
			txt := &acproto.Cloud_Message_Text{}
			err = proto.Unmarshal(msg.Content, txt)
		default:
			log.Println("未知的IM Push.Message：", msg.ContentType)
		}
	default:
		if stream.ErrorCode > 0 {
			log.Println("Error: ", stream.ErrorCode, stream.ErrorMsg)
			if stream.ErrorCode == 10018 {
				t.wsStop(ctx, c, "Log out")
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
func (t *token) handleMsgAct(payload *[]byte, q *queue.Queue, info *liveInfo) {
	actionSignal := &acproto.ZtLiveScActionSignal{}
	err := proto.Unmarshal(*payload, actionSignal)
	checkErr(err)

	var danmu []DanmuMessage
	for _, item := range actionSignal.Item {
		for _, pl := range item.Payload {
			switch item.SingalType {
			case "CommonActionSignalComment":
				comment := &acproto.CommonActionSignalComment{}
				err = proto.Unmarshal(pl, comment)
				checkErr(err)
				d := DanmuMessage{
					Type:     Comment,
					SendTime: comment.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   comment.UserInfo.UserId,
						Nickname: comment.UserInfo.Nickname,
					},
					Comment: comment.Content,
				}
				if len(comment.UserInfo.Avatar) != 0 {
					d.Avatar = comment.UserInfo.Avatar[0].Url
				}
				danmu = append(danmu, d)
			case "CommonActionSignalLike":
				like := &acproto.CommonActionSignalLike{}
				err = proto.Unmarshal(pl, like)
				checkErr(err)
				d := DanmuMessage{
					Type:     Like,
					SendTime: like.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   like.UserInfo.UserId,
						Nickname: like.UserInfo.Nickname,
					},
				}
				if len(like.UserInfo.Avatar) != 0 {
					d.Avatar = like.UserInfo.Avatar[0].Url
				}
				danmu = append(danmu, d)
			case "CommonActionSignalUserEnterRoom":
				enter := &acproto.CommonActionSignalUserEnterRoom{}
				err = proto.Unmarshal(pl, enter)
				checkErr(err)
				d := DanmuMessage{
					Type:     EnterRoom,
					SendTime: enter.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   enter.UserInfo.UserId,
						Nickname: enter.UserInfo.Nickname,
					},
				}
				if len(enter.UserInfo.Avatar) != 0 {
					d.Avatar = enter.UserInfo.Avatar[0].Url
				}
				danmu = append(danmu, d)
			case "CommonActionSignalUserFollowAuthor":
				follow := &acproto.CommonActionSignalUserFollowAuthor{}
				err = proto.Unmarshal(pl, follow)
				checkErr(err)
				d := DanmuMessage{
					Type:     FollowAuthor,
					SendTime: follow.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   follow.UserInfo.UserId,
						Nickname: follow.UserInfo.Nickname,
					},
				}
				if len(follow.UserInfo.Avatar) != 0 {
					d.Avatar = follow.UserInfo.Avatar[0].Url
				}
				danmu = append(danmu, d)
			case "CommonNotifySignalKickedOut":
				kickedOut := &acproto.CommonNotifySignalKickedOut{}
				err = proto.Unmarshal(pl, kickedOut)
				checkErr(err)
				info.Lock()
				info.KickedOut = kickedOut.Reason
				info.Unlock()
			case "CommonNotifySignalViolationAlert":
				violationAlert := &acproto.CommonNotifySignalViolationAlert{}
				err = proto.Unmarshal(pl, violationAlert)
				checkErr(err)
				info.Lock()
				info.ViolationAlert = violationAlert.ViolationContent
				info.Unlock()
			case "AcfunActionSignalThrowBanana":
				banana := &acproto.AcfunActionSignalThrowBanana{}
				err = proto.Unmarshal(pl, banana)
				checkErr(err)
				d := DanmuMessage{
					Type:     ThrowBanana,
					SendTime: banana.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   banana.Visitor.UserId,
						Nickname: banana.Visitor.Name,
					},
					BananaCount: int(banana.Count),
				}
				danmu = append(danmu, d)
			case "CommonActionSignalGift":
				gift := &acproto.CommonActionSignalGift{}
				err = proto.Unmarshal(pl, gift)
				checkErr(err)
				// 礼物列表应该不会在直播中途改变，但以防万一
				g, ok := t.gifts[int(gift.GiftId)]
				if !ok {
					g = Giftdetail{
						ID:   int(gift.GiftId),
						Name: "未知礼物",
					}
				}
				d := DanmuMessage{
					Type:     Gift,
					SendTime: gift.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   gift.User.UserId,
						Nickname: gift.User.Nickname,
					},
					Gift: GiftInfo{
						Giftdetail:            g,
						Count:                 int(gift.Count),
						Combo:                 int(gift.Combo),
						Value:                 int(gift.Value),
						ComboID:               gift.ComboId,
						SlotDisplayDurationMs: int(gift.SlotDisplayDurationMs),
						ExpireDurationMs:      int(gift.ExpireDurationMs),
					},
				}
				if len(gift.User.Avatar) != 0 {
					d.Avatar = gift.User.Avatar[0].Url
				}
				danmu = append(danmu, d)
			case "CommonActionSignalRichText":
				/*
					richText := &acproto.CommonActionSignalRichText{}
					err = proto.Unmarshal(pl, richText)
					checkErr(err)
					log.Printf("CommonActionSignalRichText: \n%+v\n", richText)
					log.Printf("CommonActionSignalRichText payload base64: \n%s\n", base64.StdEncoding.EncodeToString(pl))
				*/
			default:
				log.Printf("未知的Action Signal item.SingalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
					item.SingalType,
					string(pl),
					base64.StdEncoding.EncodeToString(pl))
			}
		}
	}

	// 按SendTime大小排序
	sort.Slice(danmu, func(i, j int) bool {
		return danmu[i].SendTime < danmu[j].SendTime
	})

	for _, d := range danmu {
		err = q.Put(d)
		checkErr(err)
	}
}

// 处理state signal数据
func (t *token) handleMsgState(payload *[]byte, info *liveInfo) {
	signal := &acproto.ZtLiveScStateSignal{}
	err := proto.Unmarshal(*payload, signal)
	checkErr(err)

	for _, item := range signal.Item {
		switch item.SingalType {
		case "AcfunStateSignalDisplayInfo":
			bananaInfo := &acproto.AcfunStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, bananaInfo)
			checkErr(err)
			info.Lock()
			info.AllBananaCount = bananaInfo.BananaCount
			info.Unlock()
		case "CommonStateSignalDisplayInfo":
			stateInfo := &acproto.CommonStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, stateInfo)
			checkErr(err)
			info.Lock()
			info.WatchingCount = stateInfo.WatchingCount
			info.LikeCount = stateInfo.LikeCount
			info.LikeDelta = int(stateInfo.LikeDelta)
			info.Unlock()
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
					AnonymousUser:          user.AnonymousUser,
					DisplaySendAmount:      user.CustomWatchingListData, // proto里应该是写反了
					CustomWatchingListData: user.DisplaySendAmount,
				}
				if len(user.UserInfo.Avatar) != 0 {
					u.Avatar = user.UserInfo.Avatar[0].Url
				}
				users[i] = u
			}
			info.Lock()
			info.TopUsers = users
			info.Unlock()
		case "CommonStateSignalRecentComment":
			comments := &acproto.CommonStateSignalRecentComment{}
			err = proto.Unmarshal(item.Payload, comments)
			checkErr(err)
			danmu := make([]DanmuMessage, len(comments.Comment))
			for i, comment := range comments.Comment {
				d := DanmuMessage{
					Type:     Comment,
					SendTime: comment.SendTimeMs * 1e6,
					UserInfo: UserInfo{
						UserID:   comment.UserInfo.UserId,
						Nickname: comment.UserInfo.Nickname,
					},
					Comment: comment.Content,
				}
				if len(comment.UserInfo.Avatar) != 0 {
					d.Avatar = comment.UserInfo.Avatar[0].Url
				}
				danmu[i] = d
			}
			info.Lock()
			info.RecentComment = danmu
			info.Unlock()
		case "CommonStateSignalChatCall":
			//chatCall := &acproto.CommonStateSignalChatCall{}
			//err = proto.Unmarshal(item.Payload, chatCall)
			//checkErr(err)
		case "CommonStateSignalChatAccept":
			//chatAccept := &acproto.CommonStateSignalChatAccept{}
			//err = proto.Unmarshal(item.Payload, chatAccept)
			//checkErr(err)
		case "CommonStateSignalChatReady":
			//chatReady := &acproto.CommonStateSignalChatReady{}
			//err = proto.Unmarshal(item.Payload, chatReady)
			//checkErr(err)
		case "CommonStateSignalChatEnd":
			//chatEnd := &acproto.CommonStateSignalChatEnd{}
			//err = proto.Unmarshal(item.Payload, chatEnd)
			//checkErr(err)
		case "CommonStateSignalCurrentRedpackList":
			//redpackList := &acproto.CommonStateSignalCurrentRedpackList{}
			//err = proto.Unmarshal(item.Payload, redpackList)
			//checkErr(err)
		default:
			log.Printf("未知的State Signal item.SingalType：%s\npayload string:\n%s\npayload base64:\n%s\n",
				item.SingalType,
				string(item.Payload),
				base64.StdEncoding.EncodeToString(item.Payload))
		}
	}
}
