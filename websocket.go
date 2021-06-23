package acfundanmu

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/dgrr/fastws"
	"github.com/orzogc/acfundanmu/acproto"
	"google.golang.org/protobuf/proto"
)

var msgPool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, maxBytesLength)
		return &b
	},
}

// 定时发送heartbeat和keepalive数据
func (t *token) wsHeartbeat(ctx context.Context, conn *fastws.Conn, interval int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in wsHeartbeat(), the error is: %v", err)
			// 重新启动wsHeartbeat()
			time.Sleep(2 * time.Second)
			t.wsHeartbeat(ctx, conn, interval)
		}
	}()

	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, err := conn.WriteMessage(fastws.ModeBinary, t.heartbeat())
			checkErr(err)
			if t.heartbeatSeqID%5 == 4 {
				_, err = conn.WriteMessage(fastws.ModeBinary, t.keepAlive())
				checkErr(err)
			}
		}
	}
}

// 启动websocket
func (ac *AcFunLive) wsStart(ctx context.Context, event bool, errCh chan<- error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in wsStart(), the error is:  %v", err)
			log.Println("停止获取弹幕")
			errCh <- err.(error)
			close(errCh)
			if event {
				ac.callEvent(stopDanmu, err.(error))
			}
		}
	}()

	if !event {
		defer ac.q.Dispose()
	}

	// 关闭websocket
	wsCtx, wsCancel := context.WithCancel(ctx)
	defer wsCancel()

	conn, err := fastws.Dial(wsHost)
	checkErr(err)
	go func() {
		<-wsCtx.Done()
		_ = conn.Close()
	}()

	_, err = conn.WriteMessage(fastws.ModeBinary, ac.t.register())
	checkErr(err)
	_, msg, err := conn.ReadMessage(nil)
	checkErr(err)
	registerDown, err := ac.t.decode(msg)
	checkErr(err)
	regResp := &acproto.RegisterResponse{}
	err = proto.Unmarshal(registerDown.PayloadData, regResp)
	checkErr(err)
	ac.t.instanceID = regResp.InstanceId
	ac.t.sessionKey = regResp.SessKey
	//lz4CompressionThreshold = regResp.SdkOption.Lz4CompressionThresholdBytes

	_, err = conn.WriteMessage(fastws.ModeBinary, ac.t.keepAlive())
	checkErr(err)

	_, err = conn.WriteMessage(fastws.ModeBinary, ac.t.enterRoom())
	checkErr(err)

	msgCh := make(chan *[]byte, queueLen)
	payloadCh := make(chan *acproto.DownstreamPayload, queueLen)
	hasError := false
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(msgCh)
		var err error
		for {
			msg := msgPool.Get().(*[]byte)
			if msg == nil {
				b := make([]byte, maxBytesLength)
				msg = &b
			}
			_, *msg, err = conn.ReadMessage((*msg)[:0])
			if err != nil {
				if !errors.Is(err, fastws.EOF) {
					log.Printf("websocket接收数据出现错误：%v", err)
					log.Printf("停止获取uid为%d的主播的直播弹幕", ac.t.liverUID)
					hasError = true
					errCh <- err
					close(errCh)
					if event {
						ac.callEvent(stopDanmu, err)
					}
				}
				if msg != nil {
					msgPool.Put(msg)
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
			if msg == nil {
				continue
			}
			stream, err := ac.t.decode(*msg)
			if err != nil {
				log.Printf("解码接收到的数据出现错误：%v", err)
				msgPool.Put(msg)
				continue
			}
			payloadCh <- stream
			msgPool.Put(msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for stream := range payloadCh {
			err := ac.handleCommand(wsCtx, conn, stream, event)
			if err != nil {
				log.Printf("处理接收到的数据出现错误：%v", err)
			}
		}
	}()

	wg.Wait()
	if !hasError {
		errCh <- nil
		close(errCh)
		if event {
			ac.callEvent(stopDanmu, nil)
		}
	}
}

// 停止websocket
func (t *token) wsStop(conn *fastws.Conn, message string) {
	_, err := conn.WriteMessage(fastws.ModeBinary, t.userExit())
	checkErr(err)
	_, err = conn.WriteMessage(fastws.ModeBinary, t.unregister())
	checkErr(err)
	_ = conn.CloseString(message)
}
