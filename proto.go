package acfundanmu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/orzogc/acfundanmu/acproto"

	"google.golang.org/protobuf/proto"
)

// 生成ZtLiveCsCmd
func (t *token) genCommand(command string, msg *[]byte) *[]byte {
	cmd := &acproto.ZtLiveCsCmd{
		CmdType: command,
		Ticket:  t.tickets[t.ticketIndex],
		LiveId:  t.liveID,
	}
	if msg != nil {
		cmd.Payload = *msg
	}

	cmdBytes, err := proto.Marshal(cmd)
	checkErr(err)

	return &cmdBytes
}

// 生成UpstreamPayload
func (t *token) genPayload(cmd string, msg *[]byte) *[]byte {
	t.Lock()
	payload := &acproto.UpstreamPayload{
		Command:    cmd,
		SeqId:      t.seqID,
		RetryCount: retryCount,
		SubBiz:     subBiz,
	}
	t.Unlock()
	if msg != nil {
		payload.PayloadData = *msg
	}

	body, err := proto.Marshal(payload)
	checkErr(err)

	return &body
}

// 生成PacketHeader
func (t *token) genHeader(length int) (header *acproto.PacketHeader) {
	t.Lock()
	header = &acproto.PacketHeader{
		AppId:             appID,
		Uid:               t.userID,
		InstanceId:        t.instanceID,
		DecodedPayloadLen: int32(length),
		EncryptionMode:    acproto.PacketHeader_kEncryptionSessionKey,
		SeqId:             t.seqID,
		Kpn:               kpn,
	}
	t.Unlock()
	return header
}

// register数据
func (t *token) register() *[]byte {
	request := &acproto.RegisterRequest{
		AppInfo: &acproto.AppInfo{
			AppName:    appName,
			SdkVersion: sdkVersion,
		},
		DeviceInfo: &acproto.DeviceInfo{
			PlatformType: acproto.DeviceInfo_H5,
			DeviceModel:  "h5",
		},
		PresenceStatus:  acproto.RegisterRequest_kPresenceOnline,
		AppActiveStatus: acproto.RegisterRequest_kAppInForeground,
		InstanceId:      t.instanceID,
		ZtCommonInfo: &acproto.ZtCommonInfo{
			Kpn: kpn,
			Kpf: kpf,
			Uid: t.userID,
		},
	}

	requestBytes, err := proto.Marshal(request)
	checkErr(err)

	body := t.genPayload("Basic.Register", &requestBytes)

	header := t.genHeader(len(*body))
	header.EncryptionMode = acproto.PacketHeader_kEncryptionServiceToken
	header.TokenInfo = &acproto.TokenInfo{
		TokenType: acproto.TokenInfo_kServiceToken,
		Token:     []byte(t.serviceToken),
	}
	t.seqID++

	return t.encode(header, body)
}

// unregister数据
func (t *token) unregister() *[]byte {
	//unregister := &acproto.UnregisterRequest{}
	//unregisterBytes, err := proto.Marshal(unregister)
	//checkErr(err)

	body := t.genPayload("Basic.Ping", nil)

	header := t.genHeader(len(*body))

	return t.encode(header, body)
}

// ping数据
func (t *token) ping() *[]byte {
	ping := &acproto.PingRequest{
		PingType: acproto.PingRequest_kPostRegister,
	}
	pingBytes, err := proto.Marshal(ping)
	checkErr(err)

	body := t.genPayload("Basic.Ping", &pingBytes)

	header := t.genHeader(len(*body))

	return t.encode(header, body)
}

// enter room数据
func (t *token) enterRoom() *[]byte {
	request := &acproto.ZtLiveCsEnterRoom{
		EnterRoomAttach:      t.enterRoomAttach,
		ClientLiveSdkVersion: clientLiveSdkVersion,
	}
	requestBytes, err := proto.Marshal(request)
	checkErr(err)

	cmd := t.genCommand("ZtLiveCsEnterRoom", &requestBytes)

	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)

	header := t.genHeader(len(*body))
	t.Lock()
	t.seqID++
	t.Unlock()

	return t.encode(header, body)
}

// keep alive数据
func (t *token) keepAlive(increase bool) *[]byte {
	keepAlive := &acproto.KeepAliveRequest{
		PresenceStatus:  acproto.RegisterRequest_kPresenceOnline,
		AppActiveStatus: acproto.RegisterRequest_kAppInForeground,
	}
	keepAliveBytes, err := proto.Marshal(keepAlive)
	checkErr(err)

	body := t.genPayload("Basic.KeepAlive", &keepAliveBytes)

	header := t.genHeader(len(*body))

	if increase {
		t.seqID++
	}

	return t.encode(header, body)
}

// push message数据
func (t *token) pushMessage() *[]byte {
	body := t.genPayload("Push.ZtLiveInteractive.Message", nil)

	header := t.genHeader((len(*body)))
	header.SeqId = t.headerSeqID

	return t.encode(header, body)
}

// heartbeat数据
func (t *token) heartbeat() *[]byte {
	heartbeat := &acproto.ZtLiveCsHeartbeat{
		ClientTimestampMs: time.Now().UnixNano() / 1e6,
		Sequence:          t.heartbeatSeqID,
	}
	heartbeatBytes, err := proto.Marshal(heartbeat)
	checkErr(err)

	cmd := t.genCommand("ZtLiveCsHeartbeat", &heartbeatBytes)

	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)

	header := t.genHeader(len(*body))
	t.heartbeatSeqID++
	t.Lock()
	t.seqID++
	t.Unlock()

	return t.encode(header, body)
}

// user exit数据
func (t *token) userExit() *[]byte {
	cmd := t.genCommand("ZtLiveCsUserExit", nil)

	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)

	header := t.genHeader(len(*body))
	t.Lock()
	t.seqID++
	t.Unlock()

	return t.encode(header, body)
}

// 将header和body按照格式组合起来
func (t *token) encode(header *acproto.PacketHeader, body *[]byte) *[]byte {
	headerBytes, err := proto.Marshal(header)
	checkErr(err)

	// 选择密钥
	key := t.sessionKey
	if header.EncryptionMode == acproto.PacketHeader_kEncryptionServiceToken {
		key = t.securityKey
	}
	encrypted := encrypt(key, body)

	// 具体数据格式看https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, uint32(0xABCD0001))
	checkErr(err)
	err = binary.Write(buf, binary.BigEndian, uint32(len(headerBytes)))
	checkErr(err)
	err = binary.Write(buf, binary.BigEndian, uint32(len(*encrypted)))
	checkErr(err)
	buf.Write(headerBytes)
	buf.Write(*encrypted)

	b := buf.Bytes()
	return &b
}

// 根据密钥加密body，加密方式为AES的CBC模式
func encrypt(key string, body *[]byte) *[]byte {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	checkErr(err)
	body = padding(body, aes.BlockSize)

	block, err := aes.NewCipher(keyBytes)
	checkErr(err)
	cipherText := make([]byte, len(*body))
	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	checkErr(err)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, *body)

	encrypted := append(iv, cipherText...)
	return &encrypted
}

// AES的CBC模式的padding
func padding(cipherText *[]byte, blockSize int) *[]byte {
	padding := (blockSize - len(*cipherText)%blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	text := append(*cipherText, padText...)
	return &text
}

// 将body/payload从数据中分离出来
func (t *token) decode(byt *[]byte) (downstream *acproto.DownstreamPayload, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("decode() error: %w", err)
		}
	}()

	header, payloadBytes := decodeResponse(byt)
	t.headerSeqID = header.SeqId

	payload := payloadBytes
	if header.EncryptionMode != acproto.PacketHeader_kEncryptionNone {
		key := t.sessionKey
		if header.EncryptionMode == acproto.PacketHeader_kEncryptionServiceToken {
			key = t.securityKey
		}
		payload = decrypt(payloadBytes, key)
	}

	if len(*payload) != int(header.DecodedPayloadLen) {
		panicln(fmt.Errorf("decode(): the length of body/payload is wrong: payload %d header %d", len(*payload), header.DecodedPayloadLen))
	}

	//payload = payload[:header.DecodedPayloadLen]

	downstream = &acproto.DownstreamPayload{}
	err := proto.Unmarshal(*payload, downstream)
	checkErr(err)

	return downstream, nil
}

// 分离header和body/payload
func decodeResponse(byt *[]byte) (*acproto.PacketHeader, *[]byte) {
	reader := bytes.NewReader(*byt)

	// 具体数据格式看https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu
	length := make([]byte, 4)
	// 忽略第一个4字节数据
	_, err := reader.Read(length)
	checkErr(err)
	// 读取header长度
	_, err = reader.Read(length)
	checkErr(err)
	headerLength := binary.BigEndian.Uint32(length)
	// 读取body/payload长度
	_, err = reader.Read(length)
	checkErr(err)
	payloadLength := binary.BigEndian.Uint32(length)

	// header数据
	headerBytes := make([]byte, headerLength)
	_, err = reader.Read(headerBytes)
	checkErr(err)

	// body/payload数据
	payloadBytes := make([]byte, payloadLength)
	_, err = reader.Read(payloadBytes)
	checkErr(err)

	header := &acproto.PacketHeader{}
	err = proto.Unmarshal(headerBytes, header)
	checkErr(err)

	return header, &payloadBytes
}

// 解密数据，解密方式为AES的CBC模式
func decrypt(byt *[]byte, key string) *[]byte {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	checkErr(err)
	block, err := aes.NewCipher(keyBytes)
	checkErr(err)

	if len(*byt) < aes.BlockSize {
		log.Println("decrypt(): Ciphertext block size is too short!")
		return nil
	}

	iv := (*byt)[:aes.BlockSize]
	cipherText := (*byt)[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		panicln(fmt.Errorf("decrypt(): cipherText is not a multiple of the block size"))
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return unpadding(&cipherText)
}

// AES的CBC模式的unpadding
func unpadding(text *[]byte) *[]byte {
	length := len(*text)
	unpadding := int((*text)[length-1])
	t := (*text)[:(length - unpadding)]
	return &t
}
