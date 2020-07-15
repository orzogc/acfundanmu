package acfundanmu

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/orzogc/acfundanmu/acproto"

	"github.com/golang/protobuf/proto"
	"nhooyr.io/websocket"
)

// 定时发送heartbeat和keepalive数据
func (t *token) wsHeartbeat(ctx context.Context, c *websocket.Conn, hb chan int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in wsHeartbeat(), the error is: %v", err)
			// 重新启动wsHeartbeat()
			time.Sleep(2 * time.Second)
			hb <- 10000
			t.wsHeartbeat(ctx, c, hb)
		}
	}()

	b := <-hb
	ticker := time.NewTicker(time.Duration(b) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := c.Write(ctx, websocket.MessageBinary, *t.heartbeat())
			checkErr(err)
			err = c.Write(ctx, websocket.MessageBinary, *t.keepAlive(false))
			checkErr(err)
		}
	}
}

// 启动websocket，username（邮箱）和password用来登陆AcFun，其为空串时启动访客模式，目前登陆模式和访客模式并没有区别
func (q Queue) wsStart(ctx context.Context, uid int, username, password string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("wsStart() error: %v", err)
			log.Println("websocket停止运行，如果要获取弹幕请重新启动websocket")
		}
	}()
	defer q.q.Dispose()

	var cookieContainer []*http.Cookie = nil
	var err error
	if username != "" && password != "" {
		cookieContainer, err = login(username, password)
		if err != nil {
			q.ch <- false
			log.Panicln(err)
		}
	}

	var t *token
	for retry := 0; retry < 3; retry++ {
		t, err = initialize(uid, cookieContainer)
		if err != nil {
			if retry == 2 {
				q.ch <- false
				log.Println("获取token失败，主播可能不在直播")
				log.Panicln(err)
			} else {
				log.Printf("初始化出现错误：%v", err)
				log.Println("尝试重新初始化")
			}
		} else {
			break
		}
	}

	q.ch <- true

	c, _, err := websocket.Dial(ctx, host, nil)
	checkErr(err)
	defer c.Close(websocket.StatusInternalError, "可能出现错误")

	err = c.Write(ctx, websocket.MessageBinary, *t.register())
	checkErr(err)
	_, resp, err := c.Read(ctx)
	checkErr(err)
	registerDown, err := t.decode(&resp)
	checkErr(err)
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

	// 让websocket关闭时能马上结束wsHeartbeat()
	hbCtx, hbCancel := context.WithCancel(ctx)
	defer hbCancel()
	hb := make(chan int64, 20)
	go t.wsHeartbeat(hbCtx, c, hb)

	for {
		_, buffer, err := c.Read(ctx)
		if err != nil {
			break
		}

		stream, err := t.decode(&buffer)
		if err != nil {
			log.Printf("解码接受到的数据出现错误：%v", err)
			continue
		}

		err = t.handleCommand(ctx, c, stream, q.q, hb)
		if err != nil {
			log.Printf("处理接受到的数据出现错误：%v", err)
		}
	}

	return
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
