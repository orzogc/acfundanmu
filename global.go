package acfundanmu

import (
	"sync"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

const (
	visitorSt        = "acfun.api.visitor_st"
	midgroundSt      = "acfun.midground.api_st"
	acfunHost        = "https://live.acfun.cn"
	acfunSignInURL   = "https://id.app.acfun.cn/rest/web/login/signin"
	acfunSafetyIDURL = "https://sec-cdn.gifshow.com/safetyid"
	liveURL          = "https://live.acfun.cn/live/"
	loginURL         = "https://id.app.acfun.cn/rest/app/visitor/login"
	getTokenURL      = "https://id.app.acfun.cn/rest/web/token/get"
	playURL          = "https://api.kuaishouzt.com/rest/zt/live/web/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	giftURL          = "https://api.kuaishouzt.com/rest/zt/live/web/gift/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	watchingURL      = "https://api.kuaishouzt.com/rest/zt/live/web/watchingList?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"

	safetyIDContent = "{\"platform\":5,\"app_version\":\"2.0.32\",\"device_id\":\"null\",\"user_id\":\"%d\"}"

	sid       = "sid"
	visitor   = "acfun.api.visitor"
	midground = "acfun.midground.api"

	host                 = "wss://link.xiatou.com/"
	appID                = 13
	appName              = "link-sdk"
	sdkVersion           = "1.2.1"
	kpn                  = "ACFUN_APP"
	kpf                  = "PC_WEB"
	subBiz               = "mainApp"
	clientLiveSdkVersion = "kwai-acfun-live-link"

	retryCount uint32 = 1
)

type token struct {
	sync.Mutex             // seqID、headerSeqID和ticketIndex的锁
	userID          int64  // AcFun帐号uid
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
	gifts           map[int64]Giftdetail
	uid             int64 // 主播uid
	livePage        string
	client          *fasthttp.Client
	cookies         []string
	medalParser     fastjson.ParserPool
	watchParser     fastjson.ParserPool
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
