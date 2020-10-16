package acfundanmu

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dgrr/fastws"
	"github.com/golang/protobuf/proto"
	"github.com/orzogc/acfundanmu/acproto"
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
			_, err := conn.WriteMessage(fastws.ModeBinary, t.heartbeat())
			checkErr(err)
			_, err = conn.WriteMessage(fastws.ModeBinary, t.keepAlive(false))
			checkErr(err)
		}
	}
}

// 启动websocket
func (dq *DanmuQueue) wsStart(ctx context.Context, event bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in wsStart(), the error is:  %v", err)
			log.Println("停止获取弹幕")
		}
	}()

	if event {
		defer dq.dispatchEvent(liveOff, "停止获取弹幕")
	} else {
		defer dq.q.Dispose()
	}

	conn, err := fastws.Dial(host)
	checkErr(err)

	// 关闭websocket
	wsCtx, wsCancel := context.WithCancel(ctx)
	defer wsCancel()
	go func() {
		<-wsCtx.Done()
		conn.Close()
	}()

	_, err = conn.WriteMessage(fastws.ModeBinary, dq.t.register())
	checkErr(err)
	var msg []byte
	_, msg, err = conn.ReadMessage(msg[:0])
	checkErr(err)
	registerDown, err := dq.t.decode(msg)
	checkErr(err)
	regResp := &acproto.RegisterResponse{}
	err = proto.Unmarshal(registerDown.PayloadData, regResp)
	checkErr(err)
	dq.t.instanceID = regResp.InstanceId
	dq.t.sessionKey = base64.StdEncoding.EncodeToString(regResp.SessKey)
	//lz4CompressionThreshold = regResp.SdkOption.Lz4CompressionThresholdBytes

	_, err = conn.WriteMessage(fastws.ModeBinary, dq.t.keepAlive(true))
	checkErr(err)

	_, err = conn.WriteMessage(fastws.ModeBinary, dq.t.enterRoom())
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
					log.Printf("停止获取uid为%d的主播的直播弹幕", dq.t.uid)
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
			stream, err := dq.t.decode(msg)
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
			err := dq.handleCommand(conn, stream, hb, event)
			if err != nil {
				log.Printf("处理接收到的数据出现错误：%v", err)
			}
		}
	}()

	wg.Wait()
}

// 停止websocket
func (t *token) wsStop(conn *fastws.Conn, message string) {
	_, err := conn.WriteMessage(fastws.ModeBinary, t.userExit())
	checkErr(err)
	_, err = conn.WriteMessage(fastws.ModeBinary, t.unregister())
	checkErr(err)
	conn.CloseString(message)
}
