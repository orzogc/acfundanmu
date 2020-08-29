package acfundanmu

import "sync"

const visitorSt = "acfun.api.visitor_st"
const midgroundSt = "acfun.midground.api_st"
const acfunHost = "https://live.acfun.cn"
const acfunSignInURL = "https://id.app.acfun.cn/rest/web/login/signin"
const acfunSafetyIDURL = "https://sec-cdn.gifshow.com/safetyid"
const liveMainPage = "https://live.acfun.cn/"
const liveURL = "https://live.acfun.cn/live/"
const loginURL = "https://id.app.acfun.cn/rest/app/visitor/login"
const getTokenURL = "https://id.app.acfun.cn/rest/web/token/get"
const playURL = "https://api.kuaishouzt.com/rest/zt/live/web/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
const giftURL = "https://api.kuaishouzt.com/rest/zt/live/web/gift/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
const watchingURL = "https://api.kuaishouzt.com/rest/zt/live/web/watchingList?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"

const safetyIDContent = "{\"platform\":5,\"app_version\":\"2.0.32\",\"device_id\":\"null\",\"user_id\":\"%d\"}"

const sid = "sid"
const visitor = "acfun.api.visitor"
const midground = "acfun.midground.api"

const host = "wss://link.xiatou.com/"
const appID = 13
const appName = "link-sdk"
const sdkVersion = "1.2.1"
const kpn = "ACFUN_APP"
const kpf = "PC_WEB"
const subBiz = "mainApp"
const clientLiveSdkVersion = "kwai-acfun-live-link"

const retryCount uint32 = 1

type token struct {
	sync.Mutex      // seqID的锁
	userID          int64
	securityKey     string // 第一次发送ws信息时所用密钥
	serviceToken    string
	liveID          string
	enterRoomAttach string
	tickets         []string
	instanceID      int64
	sessionKey      string // 除第一次外发送ws信息时所用密钥
	seqID           int64
	headerSeqID     int64
	heartbeatSeqID  int64
	ticketIndex     int
	deviceID        string
	gifts           map[int]Giftdetail
}
