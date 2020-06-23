package acfundanmu

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/orzogc/acfundanmu/acproto"

	"github.com/Workiva/go-datastructures/queue"
	"google.golang.org/protobuf/proto"
	"nhooyr.io/websocket"
)

// 处理接受到的数据里的命令
func (t *token) handleCommand(ctx context.Context, c *websocket.Conn, stream *acproto.DownstreamPayload, q *queue.Queue, hb chan<- int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from panic in handleCommand(), the error is:", err)
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
		default:
			log.Printf("未知的cmd.CmdAckType：%s, payload string: %s, payload base64: %s\n",
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
			result, err := ioutil.ReadAll(r)
			checkErr(err)
			payload = result
		}
		switch message.MessageType {
		case "ZtLiveScActionSignal":
			handleMsgAct(&payload, q)
		case "ZtLiveScStateSignal":
			handleMsgState(&payload)
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
			log.Printf("未知的message.MessageType：%s, payload string: %s, payload base64: %s\n",
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
			log.Printf("未知的stream.Command：%s, payload string: %s, payload base64: %s\n",
				stream.Command,
				string(stream.PayloadData),
				base64.StdEncoding.EncodeToString(stream.PayloadData))
		}
	}
}

// 处理action signal数据
func handleMsgAct(payload *[]byte, q *queue.Queue) {
	actionSignal := &acproto.ZtLiveScActionSignal{}
	err := proto.Unmarshal(*payload, actionSignal)
	checkErr(err)

	for _, item := range actionSignal.Item {
		for _, pl := range item.Payload {
			switch item.SingalType {
			case "CommonActionSignalComment":
				comment := &acproto.CommonActionSignalComment{}
				err = proto.Unmarshal(pl, comment)
				checkErr(err)
				//fmt.Println(comment.UserInfo.Nickname, "：", comment.Content)
				c := Comment{
					SendTime: comment.SendTimeMs * 1e6,
					UserID:   comment.UserInfo.UserId,
					Nickname: comment.UserInfo.Nickname,
					Content:  comment.Content}
				err = q.Put(c)
				checkErr(err)
			case "CommonActionSignalLike":
				like := &acproto.CommonActionSignalLike{}
				err = proto.Unmarshal(pl, like)
				checkErr(err)
				//fmt.Println(like.UserInfo.Nickname, "点赞")
			case "CommonActionSignalUserEnterRoom":
				enter := &acproto.CommonActionSignalUserEnterRoom{}
				err = proto.Unmarshal(pl, enter)
				checkErr(err)
				//fmt.Println(enter.UserInfo.Nickname, "进入房间")
			case "CommonActionSignalUserFollowAuthor":
				follower := &acproto.CommonActionSignalUserFollowAuthor{}
				err = proto.Unmarshal(pl, follower)
				checkErr(err)
				//fmt.Println(follower.UserInfo.Nickname, "关注了主播")
			case "CommonNotifySignalKickedOut":
				kickedOut := &acproto.CommonNotifySignalKickedOut{}
				err = proto.Unmarshal(pl, kickedOut)
				checkErr(err)
				//fmt.Println("被踢信息：", kickedOut.Reason)
			case "CommonNotifySignalViolationAlert":
				violationAlert := &acproto.CommonNotifySignalViolationAlert{}
				err = proto.Unmarshal(pl, violationAlert)
				checkErr(err)
				//fmt.Println("警告信息：", violationAlert.ViolationContent)
			case "AcfunActionSignalThrowBanana":
				banana := &acproto.AcfunActionSignalThrowBanana{}
				err = proto.Unmarshal(pl, banana)
				checkErr(err)
				//fmt.Println(banana.Visitor.Name, "送香蕉")
			case "CommonActionSignalGift":
				gift := &acproto.CommonActionSignalGift{}
				err = proto.Unmarshal(pl, gift)
				checkErr(err)
				//fmt.Println(gift.User.Name, "送出礼物：", gifts[int(gift.ItemId)], "数量：", gift.Count, "连击总数：", gift.Combo, "单个价值：", gift.Value)
			default:
				log.Printf("未知的Action Signal item.SingalType：%s, payload string: %s, payload base64: %s\n",
					item.SingalType,
					string(pl),
					base64.StdEncoding.EncodeToString(pl))
			}
		}
	}
}

// 处理state signal数据
func handleMsgState(payload *[]byte) {
	signal := &acproto.ZtLiveScStateSignal{}
	err := proto.Unmarshal(*payload, signal)
	checkErr(err)

	for _, item := range signal.Item {
		switch item.SingalType {
		case "AcfunStateSignalDisplayInfo":
			bananaInfo := &acproto.AcfunStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, bananaInfo)
			checkErr(err)
			//fmt.Println("香蕉总数：", bananaInfo.BananaCount)
		case "CommonStateSignalDisplayInfo":
			stateInfo := &acproto.CommonStateSignalDisplayInfo{}
			err = proto.Unmarshal(item.Payload, stateInfo)
			checkErr(err)
			//fmt.Println("观看人数和点赞总数：", stateInfo.WatchingCount, stateInfo.LikeCount)
		case "CommonStateSignalTopUsers":
			topUsers := &acproto.CommonStateSignalTopUsers{}
			err = proto.Unmarshal(item.Payload, topUsers)
			checkErr(err)
			//for _, user := range topUsers.User {
			//	fmt.Println("老板", user.Detail.Name)
			//}
		case "CommonStateSignalRecentComment":
			comments := &acproto.CommonStateSignalRecentComment{}
			err = proto.Unmarshal(item.Payload, comments)
			checkErr(err)
			//for _, comment := range comments.Comment {
			//	fmt.Println(comment.UserInfo.Nickname, "：", comment.Content)
			//}
		default:
			log.Printf("未知的State Signal item.SingalType：%s, payload string: %s, payload base64: %s\n",
				item.SingalType,
				string(item.Payload),
				base64.StdEncoding.EncodeToString(item.Payload))
		}
	}
}
