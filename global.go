package acfundanmu

import (
	"github.com/valyala/fastjson"
)

const (
	visitorSt      = "acfun.api.visitor_st"
	midgroundSt    = "acfun.midground.api_st"
	acfunSignInURL = "https://id.app.acfun.cn/rest/web/login/signin"
	liveURL        = "https://live.acfun.cn/live/%d"
	loginURL       = "https://id.app.acfun.cn/rest/app/visitor/login"
	getTokenURL    = "https://id.app.acfun.cn/rest/web/token/get"
	playURL        = "https://api.kuaishouzt.com/rest/zt/live/web/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	giftURL        = "https://api.kuaishouzt.com/rest/zt/live/web/gift/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"

	sid       = "sid"
	visitor   = "acfun.api.visitor"
	midground = "acfun.midground.api"

	wsHost               = "wss://link.xiatou.com/"
	appID                = 13
	appName              = "link-sdk"
	sdkVersion           = "1.2.1"
	kpn                  = "ACFUN_APP"
	kpf                  = "PC_WEB"
	subBiz               = "mainApp"
	clientLiveSdkVersion = "kwai-acfun-live-link"

	retryCount uint32 = 1

	formContentType = "application/x-www-form-urlencoded"
)

const (
	watchingListURL    = "https://api.kuaishouzt.com/rest/zt/live/web/watchingList?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	endSummaryURL      = "https://api.kuaishouzt.com/rest/zt/live/web/endSummary?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	redpackLuckListURL = "https://api.kuaishouzt.com/rest/zt/live/web/redpack/getLuckList?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	getPlayURL         = "https://api.kuaishouzt.com/rest/zt/live/web/getPlayUrls?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	allGiftURL         = "https://api.kuaishouzt.com/rest/zt/live/web/gift/all?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	walletBalanceURL   = "https://api.kuaishouzt.com/rest/zt/live/web/pay/wallet/balance?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	kickHistoryURL     = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/kickHistory?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	managerListURL     = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/manager/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	billboardURL       = "https://api.kuaishouzt.com/rest/zt/live/billboard?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	playbackURL        = "https://api.kuaishouzt.com/rest/zt/live/playBack/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	medalInfoURL       = "https://api-new.app.acfun.cn/rest/app/fansClub/live/medalInfo?uperId=%d"
	userMedalURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/user/info?userId=%d"
	liveListURL        = "https://live.acfun.cn/api/channel/list?count=%d"
)

const (
	managerKickURL   = "https://api.kuaishouzt.com/rest/zt/live/web/manager/kick?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	authorKickURL    = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/kick?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	addManagerURL    = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/manager/add?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	deleteManagerURL = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/manager/delete?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
)

const (
	checkLiveAuthURL = "https://member.acfun.cn/common/api/checkLiveAuth"
	liveTypeListURL  = "https://member.acfun.cn/common/api/getLiveTypeList"
	obsConfigURL     = "https://api.kuaishouzt.com/rest/zt/live/web/obs/config?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
)

type token struct {
	userID          int64  // AcFun帐号uid
	securityKey     string // 第一次发送ws信息时所用密钥
	serviceToken    string
	liveID          string
	enterRoomAttach string
	tickets         []string
	instanceID      int64
	sessionKey      []byte // 除第一次外发送ws信息时所用密钥
	seqID           int64  // 要用原子锁操作
	headerSeqID     int64  // 要用原子锁操作
	heartbeatSeqID  int64
	ticketIndex     uint32 // 要用原子锁操作
	deviceID        string
	gifts           map[int64]GiftDetail
	liverUID        int64 // 主播uid
	livePage        string
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
