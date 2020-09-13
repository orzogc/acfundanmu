package acfundanmu

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dgrr/fastws"
	"github.com/golang/protobuf/proto"
	"github.com/orzogc/acfundanmu/acproto"
	"github.com/valyala/fasthttp"
)

// 定时发送heartbeat和keepalive数据
func (t *token) wsHeartbeat(ctx context.Context, conn *fastws.Conn, hb chan int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in wsHeartbeat(), the error is: %v", err)
			// 重新启动wsHeartbeat()
			time.Sleep(2 * time.Second)
			hb <- 10000
			t.wsHeartbeat(ctx, conn, hb)
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
			_, err := conn.WriteMessage(fastws.ModeBinary, *t.heartbeat())
			checkErr(err)
			_, err = conn.WriteMessage(fastws.ModeBinary, *t.keepAlive(false))
			checkErr(err)
		}
	}
}

// 启动websocket，uid为主播的uid，cookies是AcFun帐号的cookies，可以调用login()获取，其为nil时使用访客模式，目前登陆模式和访客模式并没有区别
func (dq *DanmuQueue) wsStart(ctx context.Context, uid int, cookies []*fasthttp.Cookie) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("wsStart() error: %v", err)
			log.Println("停止获取弹幕")
		}
	}()
	defer dq.q.Dispose()

	var err error
	for retry := 0; retry < 3; retry++ {
		dq.t, err = initialize(uid, cookies)
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

	conn, err := fastws.Dial(host)
	checkErr(err)

	// 关闭websocket
	wsCtx, wsCancel := context.WithCancel(ctx)
	defer wsCancel()
	go func() {
		<-wsCtx.Done()
		conn.Close()
	}()

	_, err = conn.WriteMessage(fastws.ModeBinary, *dq.t.register())
	checkErr(err)
	var msg []byte
	_, msg, err = conn.ReadMessage(msg[:0])
	checkErr(err)
	registerDown, err := dq.t.decode(&msg)
	checkErr(err)
	regResp := &acproto.RegisterResponse{}
	err = proto.Unmarshal(registerDown.PayloadData, regResp)
	checkErr(err)
	dq.t.instanceID = regResp.InstanceId
	dq.t.sessionKey = base64.StdEncoding.EncodeToString(regResp.SessKey)
	//lz4CompressionThreshold = regResp.SdkOption.Lz4CompressionThresholdBytes

	_, err = conn.WriteMessage(fastws.ModeBinary, *dq.t.keepAlive(true))
	checkErr(err)

	_, err = conn.WriteMessage(fastws.ModeBinary, *dq.t.enterRoom())
	checkErr(err)

	hb := make(chan int64, 10)
	go dq.t.wsHeartbeat(wsCtx, conn, hb)

	msgCh := make(chan []byte, 100)
	payloadCh := make(chan *acproto.DownstreamPayload, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(msgCh)
		var err error
		for {
			// nil是防止data race
			_, msg, err = conn.ReadMessage(nil)
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("websocket接收数据出现错误：%v", err)
					log.Println("停止获取弹幕")
				}
				break
			}
			msgCh <- msg
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(payloadCh)
		for msg := range msgCh {
			stream, err := dq.t.decode(&msg)
			if err != nil {
				log.Printf("解码接收到的数据出现错误：%v", err)
				continue
			}
			payloadCh <- stream
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for stream := range payloadCh {
			err := dq.t.handleCommand(conn, stream, dq.q, dq.info, hb)
			if err != nil {
				log.Printf("处理接收到的数据出现错误：%v", err)
			}
		}
	}()

	wg.Wait()
}

// 停止websocket
func (t *token) wsStop(conn *fastws.Conn, message string) {
	_, err := conn.WriteMessage(fastws.ModeBinary, *t.userExit())
	checkErr(err)
	_, err = conn.WriteMessage(fastws.ModeBinary, *t.unregister())
	checkErr(err)
	conn.CloseString(message)
}
