package acfundanmu

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/orzogc/acfundanmu/acproto"
	"github.com/orzogc/fastws"
)

var msgPool = sync.Pool{
	New: func() any {
		b := make([]byte, maxBytesLength)
		return &b
	},
}

// 从 msgPool 获取 byte slice
func getBytes() *[]byte {
	msg := msgPool.Get().(*[]byte)
	if msg == nil {
		b := make([]byte, maxBytesLength)
		msg = &b
	}

	return msg
}

// 放入 byte slice 到 msgPool
func putBytes(msg *[]byte) {
	if msg != nil {
		msgPool.Put(msg)
	}
}

// DanmuClientType 弹幕客户端类型
type DanmuClientType uint8

const (
	// 弹幕客户端使用 WebSocket 连接
	WebSocketDanmuClientType DanmuClientType = iota
	// 弹幕客户端使用 TCP 连接
	TCPDanmuClientType
)

// DanmuClient 弹幕客户端
type DanmuClient interface {
	// New 返回新的弹幕客户端
	NewDanmuClient() DanmuClient

	// Type 返回弹幕客户端类型
	Type() DanmuClientType

	// Dial 连接弹幕服务器，address 是地址，在调用 Close() 后可重复调用
	Dial(address string) error

	// 读接口
	io.Reader

	// 写接口
	io.Writer

	// Close 关闭连接
	Close(message string) error
}

// WebSocketDanmuClient 使用 WebSocket 连接的弹幕客户端
type WebSocketDanmuClient struct {
	conn *fastws.Conn
}

// NewDanmuClient 返回新的 WebSocketDanmuClient
func (client *WebSocketDanmuClient) NewDanmuClient() DanmuClient {
	return &WebSocketDanmuClient{}
}

// Type 返回弹幕客户端类型 WebSocketDanmuClientType
func (client *WebSocketDanmuClient) Type() DanmuClientType {
	return WebSocketDanmuClientType
}

// Dial 连接弹幕服务器，address 是地址
func (client *WebSocketDanmuClient) Dial(address string) error {
	conn, err := fastws.Dial(address)
	if err != nil {
		client.conn = nil
		return err
	}
	conn.SetReadTimeout(timeout)
	conn.SetWriteTimeout(timeout)
	client.conn = conn

	return nil
}

// Read 读数据
func (client *WebSocketDanmuClient) Read(p []byte) (n int, err error) {
	if client.conn != nil {
		_, p, err = client.conn.ReadMessage(p[:0])

		return len(p), err
	} else {
		return 0, fmt.Errorf("请先调用 Dail() 连接服务器")
	}
}

// Write 写数据
func (client *WebSocketDanmuClient) Write(p []byte) (n int, err error) {
	if client.conn != nil {
		return client.conn.WriteMessage(fastws.ModeBinary, p)
	} else {
		return 0, fmt.Errorf("请先调用 Dail() 连接服务器")
	}
}

// Close 关闭连接
func (client *WebSocketDanmuClient) Close(message string) error {
	if client.conn != nil {
		return client.conn.CloseString(message)
	} else {
		return nil
	}
}

// TCPDanmuClient 使用 TCP 连接的弹幕客户端
type TCPDanmuClient struct {
	conn net.Conn
}

// NewDanmuClient 返回新的 TCPDanmuClient
func (client *TCPDanmuClient) NewDanmuClient() DanmuClient {
	return &TCPDanmuClient{}
}

// Type 返回弹幕客户端类型 TCPDanmuClientType
func (client *TCPDanmuClient) Type() DanmuClientType {
	return TCPDanmuClientType
}

// Dial 连接弹幕服务器，address 是地址
func (client *TCPDanmuClient) Dial(address string) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		client.conn = nil
		return err
	}
	client.conn = conn

	return nil
}

// Read 读数据
func (client *TCPDanmuClient) Read(p []byte) (n int, err error) {
	if client.conn != nil {
		client.conn.SetReadDeadline(time.Now().Add(timeout))
		defer client.conn.SetReadDeadline(time.Time{})

		return client.conn.Read(p)
	} else {
		return 0, fmt.Errorf("请先调用 Dail() 连接服务器")
	}
}

// Write 写数据
func (client *TCPDanmuClient) Write(p []byte) (n int, err error) {
	if client.conn != nil {
		client.conn.SetWriteDeadline(time.Now().Add(timeout))
		defer client.conn.SetWriteDeadline(time.Time{})

		return client.conn.Write(p)
	} else {
		return 0, fmt.Errorf("请先调用 Dail() 连接服务器")
	}
}

// Close 关闭连接
func (client *TCPDanmuClient) Close(message string) error {
	if client.conn != nil {
		return client.conn.Close()
	} else {
		return nil
	}
}

// 弹幕消息
type message struct {
	bytes *[]byte
	len   int
}

// 获取弹幕数据
func (m *message) data() []byte {
	return (*m.bytes)[:m.len]
}

// 定时发送 heartbeat 和 keepalive 数据
func (ac *AcFunLive) clientHeartbeat(ctx context.Context, interval int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in clientHeartbeat(), the error is: %v", err)
			// 重新启动 clientHeartbeat()
			time.Sleep(2 * time.Second)
			ac.clientHeartbeat(ctx, interval)
		}
	}()

	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, err := ac.danmuClient.Write(ac.t.heartbeat())
			checkErr(err)
			if ac.t.heartbeatSeqID%5 == 4 {
				_, err = ac.danmuClient.Write(ac.t.keepAlive())
				checkErr(err)
			}
		}
	}
}

// 启动弹幕 client
func (ac *AcFunLive) clientStart(ctx context.Context, event bool, errCh chan<- error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in clientStart(), the error is:  %v", err)
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

	// 用于关闭 danmuClient
	clientCtx, clientCancel := context.WithCancel(ctx)
	defer clientCancel()

	// WebSocket 连接可以直接发送注册消息，TCP 连接需要先握手
	if ac.danmuClient.Type() == WebSocketDanmuClientType {
		err := ac.danmuClient.Dial(wsHost)
		checkErr(err)
		_, err = ac.danmuClient.Write(ac.t.register())
		checkErr(err)
	} else if ac.danmuClient.Type() == TCPDanmuClientType {
		err := ac.danmuClient.Dial(tcpHost)
		checkErr(err)
		_, err = ac.danmuClient.Write(ac.t.handshake())
		checkErr(err)
	}

	go func() {
		<-clientCtx.Done()
		ac.danmuClient.Close("")
	}()

	msgCh := make(chan message, queueLen)
	payloadCh := make(chan *acproto.DownstreamPayload, queueLen)
	hasError := false
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(msgCh)
		var n int
		var err error
		for {
			msg := getBytes()
			n, err = ac.danmuClient.Read(*msg)
			if err != nil || ac.t.err.Load() != nil {
				acErr := ac.t.err.Load()
				if acErr != nil {
					err = acErr
				}
				if !(errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed)) {
					log.Printf("接收弹幕数据出现错误：%v", err)
					log.Printf("停止获取 uid 为 %d 的主播的直播弹幕", ac.t.liverUID)
					hasError = true
					errCh <- err
					close(errCh)
					if event {
						ac.callEvent(stopDanmu, err)
					}
				}
				putBytes(msg)
				break
			}
			msgCh <- message{bytes: msg, len: n}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(payloadCh)
		remain := []byte{}
		for msg := range msgCh {
			if msg.bytes == nil {
				continue
			}

			if ac.danmuClient.Type() == WebSocketDanmuClientType {
				// WebSocket 连接的数据自动分帧
				stream, err := ac.t.decode(msg.data())
				if err != nil {
					log.Printf("解码接收到的弹幕数据出现错误：%v", err)
					putBytes(msg.bytes)
					continue
				}
				putBytes(msg.bytes)
				payloadCh <- stream
			} else if ac.danmuClient.Type() == TCPDanmuClientType {
				// TCP 连接的数据需要自行分帧
				if remain != nil {
					remain = append(remain, msg.data()...)
				} else {
					remain = append([]byte{}, msg.data()...)
				}

				if len(remain) > 12 {
					var frames [][]byte
					var err error
					frames, remain, err = getFrames(remain)
					if err != nil {
						log.Printf("解码接收到的弹幕数据出现错误：%v", err)
						putBytes(msg.bytes)
						continue
					}

					for _, frame := range frames {
						stream, err := ac.t.decode(frame)
						if err != nil {
							log.Printf("解码接收到的弹幕数据出现错误：%v", err)
							continue
						}
						payloadCh <- stream
					}
				}

				putBytes(msg.bytes)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for stream := range payloadCh {
			go func(stream *acproto.DownstreamPayload) {
				err := ac.handleCommand(clientCtx, stream, event)
				if err != nil {
					log.Printf("处理接收到的弹幕数据出现错误：%v", err)
				}
			}(stream)
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

// 停止弹幕 client
func (ac *AcFunLive) clientStop(message string) {
	_, err := ac.danmuClient.Write(ac.t.userExit())
	checkErr(err)
	_, err = ac.danmuClient.Write(ac.t.unregister())
	checkErr(err)
	_ = ac.danmuClient.Close(message)
}

// 检查魔法数字
func checkMagicNumber(data []byte) bool {
	return data[0] == 0xAB && data[1] == 0xCD && data[2] == 0x00 && data[3] == 0x01
}

// 获取弹幕数据长度
func getDataLength(data []byte) int {
	headerLength := binary.BigEndian.Uint32(data[4:8])
	payloadLength := binary.BigEndian.Uint32(data[8:12])
	return int(headerLength) + int(payloadLength) + 12
}

// 获取弹幕数据
func getFrame(data []byte) ([]byte, []byte) {
	if checkMagicNumber(data) {
		length := getDataLength(data)
		if len(data) < length {
			return nil, data
		} else if len(data) > length {
			return data[:length], data[length:]
		} else {
			return data, nil
		}
	} else {
		return nil, nil
	}
}

// 获取弹幕数据
func getFrames(data []byte) ([][]byte, []byte, error) {
	if len(data) > 12 {
		frames := [][]byte{}
		var frame []byte
		remain := data

		for {
			frame, remain = getFrame(remain)
			if frame != nil && remain != nil {
				frames = append(frames, frame)
				if len(remain) <= 12 {
					return frames, remain, nil
				}
			} else if frame != nil && remain == nil {
				frames = append(frames, frame)
				return frames, nil, nil
			} else if frame == nil && remain != nil {
				if len(frames) > 0 {
					return frames, remain, nil
				} else {
					return nil, remain, nil
				}
			} else {
				return nil, nil, fmt.Errorf("错误的弹幕数据格式")
			}
		}
	} else {
		return nil, data, nil
	}
}
