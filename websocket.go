package acfundanmu

import (
	"context"
	"encoding/base64"
	"fmt"
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

// 启动websocket，username（邮箱）和password用来登陆AcFun，其为空串时使用访客模式，目前登陆模式和访客模式并没有区别
func (dq *DanmuQueue) wsStart(ctx context.Context, uid int, username, password string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("wsStart() error: %v", err)
			log.Println("websocket停止运行，如果要获取弹幕请重新启动websocket")
		}
	}()
	defer dq.q.Dispose()

	var cookieContainer []*http.Cookie = nil
	var err error
	if username != "" && password != "" {
		cookieContainer, err = login(username, password)
		if err != nil {
			dq.ch <- err
			panicln(err)
		}
	}

	for retry := 0; retry < 3; retry++ {
		dq.t, err = initialize(uid, cookieContainer)
		if err != nil {
			if retry == 2 {
				e := fmt.Errorf("获取token失败，主播可能不在直播：%w", err)
				log.Println("停止获取弹幕")
				dq.ch <- e
				panicln(e)
			}
			log.Printf("初始化出现错误：%v", err)
			log.Println("尝试重新初始化")
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	dq.ch <- nil

	c, _, err := websocket.Dial(ctx, host, nil)
	checkErr(err)
	defer c.Close(websocket.StatusInternalError, "可能出现错误")

	err = c.Write(ctx, websocket.MessageBinary, *dq.t.register())
	checkErr(err)
	_, resp, err := c.Read(ctx)
	checkErr(err)
	registerDown, err := dq.t.decode(&resp)
	checkErr(err)
	regResp := &acproto.RegisterResponse{}
	err = proto.Unmarshal(registerDown.PayloadData, regResp)
	checkErr(err)
	dq.t.instanceID = regResp.InstanceId
	dq.t.sessionKey = base64.StdEncoding.EncodeToString(regResp.SessKey)
	//lz4CompressionThreshold = regResp.SdkOption.Lz4CompressionThresholdBytes

	err = c.Write(ctx, websocket.MessageBinary, *dq.t.keepAlive(true))
	checkErr(err)

	err = c.Write(ctx, websocket.MessageBinary, *dq.t.enterRoom())
	checkErr(err)

	// 让websocket关闭时能马上结束wsHeartbeat()
	hbCtx, hbCancel := context.WithCancel(ctx)
	defer hbCancel()
	hb := make(chan int64, 20)
	go dq.t.wsHeartbeat(hbCtx, c, hb)

	for {
		_, buffer, err := c.Read(ctx)
		if err != nil {
			break
		}

		stream, err := dq.t.decode(&buffer)
		if err != nil {
			log.Printf("解码接受到的数据出现错误：%v", err)
			continue
		}

		err = dq.t.handleCommand(ctx, c, stream, dq.q, dq.info, hb)
		if err != nil {
			log.Printf("处理接受到的数据出现错误：%v", err)
		}
	}
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
