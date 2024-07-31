package acfundanmu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/orzogc/acfundanmu/acproto"

	"google.golang.org/protobuf/proto"
)

const maxBytesLength = 4096

var lengthPool = sync.Pool{
	New: func() any {
		b := make([]byte, 4)
		return &b
	},
}

var bytesPool = sync.Pool{
	New: func() any {
		b := make([]byte, maxBytesLength)
		return &b
	},
}

// 生成 ZtLiveCsCmd
func (t *token) genCommand(command string, msg []byte) []byte {
	cmd := &acproto.ZtLiveCsCmd{
		CmdType: command,
		Ticket:  t.tickets[t.ticketIndex.Load()],
		LiveId:  t.liveID,
	}
	if msg != nil {
		cmd.Payload = msg
	}

	cmdBytes, err := proto.Marshal(cmd)
	checkErr(err)

	return cmdBytes
}

// 生成 UpstreamPayload
func (t *token) genPayload(cmd string, msg []byte) []byte {
	payload := &acproto.UpstreamPayload{
		Command:    cmd,
		SeqId:      t.seqID.Load(),
		RetryCount: retryCount,
		SubBiz:     subBiz,
	}
	if msg != nil {
		payload.PayloadData = msg
	}

	body, err := proto.Marshal(payload)
	checkErr(err)

	return body
}

// 生成 PacketHeader
func (t *token) genHeader(length int) (header *acproto.PacketHeader) {
	header = &acproto.PacketHeader{
		AppId:             t.appID,
		Uid:               t.UserID,
		InstanceId:        t.instanceID,
		DecodedPayloadLen: uint32(length),
		EncryptionMode:    acproto.PacketHeader_kEncryptionSessionKey,
		SeqId:             t.seqID.Load(),
		Kpn:               kpn,
	}
	return header
}

func (t *token) handshake() []byte {
	request := &acproto.HandshakeRequest{
		Unknown1: 1,
		Unknown2: 1,
	}
	requestBytes, err := proto.Marshal(request)
	checkErr(err)

	body := t.genPayload("Basic.Handshake", requestBytes)
	header := t.genHeader(len(body))
	header.EncryptionMode = acproto.PacketHeader_kEncryptionServiceToken
	header.TokenInfo = &acproto.TokenInfo{
		TokenType: acproto.TokenInfo_kServiceToken,
		Token:     []byte(t.ServiceToken),
	}

	return t.encode(header, body)
}

// Register 数据
func (t *token) register() []byte {
	request := &acproto.RegisterRequest{
		AppInfo: &acproto.AppInfo{
			SdkVersion:  clientLiveSdkVersion,
			LinkVersion: linkVersion,
		},
		DeviceInfo: &acproto.DeviceInfo{
			PlatformType: acproto.DeviceInfo_H5_WINDOWS,
			DeviceModel:  "h5",
		},
		PresenceStatus:  acproto.RegisterRequest_kPresenceOnline,
		AppActiveStatus: acproto.RegisterRequest_kAppInForeground,
		InstanceId:      t.instanceID,
		ZtCommonInfo: &acproto.ZtCommonInfo{
			Kpn: kpn,
			Kpf: kpf,
			Uid: t.UserID,
		},
	}
	requestBytes, err := proto.Marshal(request)
	checkErr(err)

	body := t.genPayload("Basic.Register", requestBytes)
	header := t.genHeader(len(body))
	header.EncryptionMode = acproto.PacketHeader_kEncryptionServiceToken
	header.TokenInfo = &acproto.TokenInfo{
		TokenType: acproto.TokenInfo_kServiceToken,
		Token:     []byte(t.ServiceToken),
	}
	_ = t.seqID.Inc()

	return t.encode(header, body)
}

// Unregister 数据
func (t *token) unregister() []byte {
	body := t.genPayload("Basic.Unregister", nil)
	header := t.genHeader(len(body))

	return t.encode(header, body)
}

// Ping 数据
/*
func (t *token) ping() []byte {
	ping := &acproto.PingRequest{
		PingType: acproto.PingRequest_kPostRegister,
	}
	pingBytes, err := proto.Marshal(ping)
	checkErr(err)

	body := t.genPayload("Basic.Ping", pingBytes)

	header := t.genHeader(len(body))

	return t.encode(header, body)
}
*/

// EnterRoom 数据
func (t *token) enterRoom() []byte {
	request := &acproto.ZtLiveCsEnterRoom{
		EnterRoomAttach:      t.enterRoomAttach,
		ClientLiveSdkVersion: clientLiveSdkVersion,
	}
	requestBytes, err := proto.Marshal(request)
	checkErr(err)

	cmd := t.genCommand("ZtLiveCsEnterRoom", requestBytes)
	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)
	header := t.genHeader(len(body))
	_ = t.seqID.Inc()

	return t.encode(header, body)
}

// KeepAlive 数据
func (t *token) keepAlive() []byte {
	keepAlive := &acproto.KeepAliveRequest{
		PresenceStatus:  acproto.RegisterRequest_kPresenceOnline,
		AppActiveStatus: acproto.RegisterRequest_kAppInForeground,
	}
	keepAliveBytes, err := proto.Marshal(keepAlive)
	checkErr(err)

	body := t.genPayload("Basic.KeepAlive", keepAliveBytes)
	header := t.genHeader(len(body))
	_ = t.seqID.Inc()

	return t.encode(header, body)
}

// Push Message 数据
func (t *token) pushMessage() []byte {
	body := t.genPayload("Push.ZtLiveInteractive.Message", nil)
	header := t.genHeader((len(body)))
	header.SeqId = t.headerSeqID.Load()

	return t.encode(header, body)
}

// Heartbeat 数据
func (t *token) heartbeat() []byte {
	heartbeat := &acproto.ZtLiveCsHeartbeat{
		ClientTimestampMs: time.Now().UnixNano() / 1e6,
		Sequence:          t.heartbeatSeqID,
	}
	heartbeatBytes, err := proto.Marshal(heartbeat)
	checkErr(err)

	cmd := t.genCommand("ZtLiveCsHeartbeat", heartbeatBytes)
	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)
	header := t.genHeader(len(body))
	t.heartbeatSeqID++
	_ = t.seqID.Inc()

	return t.encode(header, body)
}

// UserExit 数据
func (t *token) userExit() []byte {
	cmd := t.genCommand("ZtLiveCsUserExit", nil)
	body := t.genPayload("Global.ZtLiveInteractive.CsCmd", cmd)
	header := t.genHeader(len(body))
	_ = t.seqID.Inc()

	return t.encode(header, body)
}

// 将 header 和 body 按照格式组合起来
func (t *token) encode(header *acproto.PacketHeader, body []byte) []byte {
	headerBytes, err := proto.Marshal(header)
	checkErr(err)

	// 选择密钥
	key := t.sessionKey
	if header.EncryptionMode == acproto.PacketHeader_kEncryptionServiceToken {
		key, err = base64.StdEncoding.DecodeString(t.SecurityKey)
		checkErr(err)
	}
	encrypted := encrypt(key, body)

	// 具体数据格式看 https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, uint32(0xABCD0001))
	checkErr(err)
	err = binary.Write(buf, binary.BigEndian, uint32(len(headerBytes)))
	checkErr(err)
	err = binary.Write(buf, binary.BigEndian, uint32(len(encrypted)))
	checkErr(err)
	buf.Write(headerBytes)
	buf.Write(encrypted)

	return buf.Bytes()
}

// 根据密钥加密 body，加密方式为 aes-128-cbc
func encrypt(key []byte, body []byte) []byte {
	body = padding(body, aes.BlockSize)

	block, err := aes.NewCipher(key)
	checkErr(err)
	cipherText := make([]byte, len(body))
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	checkErr(err)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, body)

	return append(iv, cipherText...)
}

// aes-128-cbc 的 padding（PKCS #7）
func padding(cipherText []byte, blockSize int) []byte {
	padding := (blockSize - len(cipherText)%blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

// 将 body/payload 从数据中分离出来
func (t *token) decode(b []byte) (downstream *acproto.DownstreamPayload, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("decode() error: %v", err)
		}
	}()

	// 分离 header 和 body/payload
	reader := bytes.NewReader(b)

	// 具体数据格式看 https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu
	length := lengthPool.Get().(*[]byte)
	defer lengthPool.Put(length)
	// 忽略第一个 4 字节数据
	_, err := reader.Read(*length)
	checkErr(err)
	// 读取 header 长度
	_, err = reader.Read(*length)
	checkErr(err)
	headerLength := binary.BigEndian.Uint32(*length)
	// 读取 body/payload 长度
	_, err = reader.Read(*length)
	checkErr(err)
	payloadLength := binary.BigEndian.Uint32(*length)

	// header 数据
	var headerBytes []byte
	if headerLength <= maxBytesLength {
		byt := bytesPool.Get().(*[]byte)
		defer bytesPool.Put(byt)
		headerBytes = (*byt)[:headerLength]
	} else {
		headerBytes = make([]byte, headerLength)
	}
	_, err = reader.Read(headerBytes)
	checkErr(err)

	// body/payload数据
	var payload []byte
	if payloadLength <= maxBytesLength {
		byt := bytesPool.Get().(*[]byte)
		defer bytesPool.Put(byt)
		payload = (*byt)[:payloadLength]
	} else {
		payload = make([]byte, payloadLength)
	}
	_, err = reader.Read(payload)
	checkErr(err)

	if reader.Len() != 0 {
		log.Printf("decode(): reader has more %d bytes", reader.Len())
	}

	header := &acproto.PacketHeader{}
	err = proto.Unmarshal(headerBytes, header)
	checkErr(err)

	if t.appID == 0 && header.AppId != 0 {
		t.appID = header.AppId
	}
	t.headerSeqID.Store(header.SeqId)

	if header.EncryptionMode != acproto.PacketHeader_kEncryptionNone {
		var key []byte
		if header.EncryptionMode == acproto.PacketHeader_kEncryptionServiceToken {
			key, err = base64.StdEncoding.DecodeString(t.SecurityKey)
			checkErr(err)
		} else {
			key = t.sessionKey
		}
		payload = decrypt(payload, key)
	}

	if len(payload) != int(header.DecodedPayloadLen) {
		panic(fmt.Errorf("decode(): the length of body/payload is wrong: payload %d header %d", len(payload), header.DecodedPayloadLen))
	}

	downstream = &acproto.DownstreamPayload{}
	err = proto.Unmarshal(payload, downstream)
	checkErr(err)

	return downstream, nil
}

// 解密数据，解密方式为 aes-128-cbc
func decrypt(ciphertext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	checkErr(err)

	if len(ciphertext) < aes.BlockSize {
		log.Println("decrypt(): the length of ciphertext is less than block size")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	cipherText := ciphertext[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		panic(fmt.Errorf("decrypt(): the length of ciphertext is not a multiple of the block size"))
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	return unpadding(cipherText)
}

// aes-128-cbc 的 unpadding（PKCS #7）
func unpadding(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])

	return text[:(length - unpadding)]
}
