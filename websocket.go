package acfundanmu

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/orzogc/acfundanmu/acproto"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/golang/protobuf/proto"
	"nhooyr.io/websocket"
)

func (t *token) wsHeartbeat(hbctx context.Context, c *websocket.Conn, hb <-chan int64) {
	b := <-hb
	ticker := time.NewTicker(time.Duration(b) * time.Millisecond)
	defer ticker.Stop()
	ctx, cancel := context.WithCancel(hbctx)
	defer cancel()
	for {
		select {
		case <-hbctx.Done():
			return
		case <-ticker.C:
			err := c.Write(ctx, websocket.MessageBinary, *t.heartbeat())
			checkErr(err)
			err = c.Write(ctx, websocket.MessageBinary, *t.keepAlive(false))
			checkErr(err)
		default:
		}
	}
}

// 启动websocket，username（邮箱）和password用来登陆AcFun，其为空串时启动访客模式，目前登陆模式和访客模式并没有区别
func wsStart(ctx context.Context, uid int, q *queue.Queue, username, password string) (e error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from panic in wsStart(), the error is:", err)
			q.Dispose()
			e = errors.New(err.(string))
		}
	}()

	var cookieContainer []*http.Cookie = nil

	if username != "" && password != "" {
		if cookieContainer = login(username, password); cookieContainer == nil {
			return errors.New("登陆AcFun失败")
		}
	}
	deviceID, t := initialize(uid, cookieContainer)

	if t == nil {
		log.Println("获取token失败，主播可能不在直播")
		return errors.New("获取token失败，主播可能不在直播")
	}

	t.gifts = updateGiftList(cookieContainer, deviceID, t)

	c, _, err := websocket.Dial(ctx, host, nil)
	checkErr(err)
	defer c.Close(websocket.StatusInternalError, "出现错误")

	err = c.Write(ctx, websocket.MessageBinary, *t.register())
	checkErr(err)
	_, resp, err := c.Read(ctx)
	checkErr(err)
	registerDown := t.decode(&resp)
	regResp := &acproto.RegisterResponse{}
	err = proto.Unmarshal(registerDown.PayloadData, regResp)
	checkErr(err)
	t.instanceID = regResp.InstanceId
	t.sessionKey = base64.StdEncoding.EncodeToString(regResp.SessKey)
	//lz4CompressionThreshold = regResp.SdkOption.Lz4CompressionThresholdBytes

	err = c.Write(ctx, websocket.MessageBinary, *t.keepAlive(true))
	checkErr(err)

	err = c.Write(ctx, websocket.MessageBinary, *t.enterRoom())
	checkErr(err)

	hb := make(chan int64, 5)
	hbCtx, hbCancel := context.WithCancel(ctx)
	defer hbCancel()
	go t.wsHeartbeat(hbCtx, c, hb)

	for {
		_, buffer, err := c.Read(ctx)
		/*
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				break
			} else {
				checkErr(err)
			}
		*/
		if err != nil {
			break
		}
		stream := t.decode(&buffer)
		t.handleCommand(ctx, c, stream, q, hb)
	}

	q.Dispose()

	return nil
}

// 停止websocket
func (t *token) wsStop(ctx context.Context, c *websocket.Conn, message string) {
	err := c.Write(ctx, websocket.MessageBinary, *t.userExit())
	checkErr(err)
	err = c.Write(ctx, websocket.MessageBinary, *t.unregister())
	checkErr(err)
	c.Close(websocket.StatusNormalClosure, message)
	//fmt.Println(message)
}
