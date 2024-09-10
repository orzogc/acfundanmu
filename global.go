package acfundanmu

import (
	"sync"

	"github.com/valyala/fastjson"
	"go.uber.org/atomic"
)

const (
	visitorSt      = "acfun.api.visitor_st"
	midgroundSt    = "acfun.midground.api_st"
	midgroundAt    = "acfun.midground.api.at"
	acfunSignInURL = "https://id.app.acfun.cn/rest/web/login/signin"
	liveHost       = "https://live.acfun.cn/"
	liveURL        = "https://live.acfun.cn/live/%d"
	loginURL       = "https://id.app.acfun.cn/rest/app/visitor/login"
	getTokenURL    = "https://id.app.acfun.cn/rest/web/token/get"

	sid       = "sid"
	visitor   = "acfun.api.visitor"
	midground = "acfun.midground.api"

	wsHost               = "wss://klink-newproduct-ws3.kwaizt.com/"
	tcpHost              = "slink.gifshow.com:14000"
	appName              = "link-sdk"
	sdkVersion           = "1.2.1"
	kpn                  = "ACFUN_APP"
	kpf                  = "PC_WEB"
	subBiz               = "mainApp"
	clientLiveSdkVersion = "kwai-acfun-live-link"
	linkVersion          = "2.13.8"

	retryCount uint32 = 1

	formContentType = "application/x-www-form-urlencoded"
	jsonContentType = "application/json"
	pushType        = `{"typeId":%d,"type":[%d,%d]}`
	liveCutStatus   = `{"status":%d}`
)

const (
	startScanQRURL    = "https://scan.acfun.cn/rest/pc-direct/qr/start?type=WEB_LOGIN&_=%d"
	scanQRResultURL   = "https://scan.acfun.cn/rest/pc-direct/qr/scanResult?qrLoginToken=%s&qrLoginSignature=%s&_=%d"
	acceptQRResultURL = "https://scan.acfun.cn/rest/pc-direct/qr/acceptResult?qrLoginToken=%s&qrLoginSignature=%s&_=%d"
)

const (
	playURL            = "https://api.kuaishouzt.com/rest/zt/live/web/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	giftURL            = "https://api.kuaishouzt.com/rest/zt/live/web/gift/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
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
	userMedalURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/user/info?userId=%d"
	medalRankURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/friendshipDegreeRankInfo?uperId=%d"
	medalDetailURL     = "https://live.acfun.cn/rest/pc-direct/fansClub/fans/medal/detail?uperId=%d"
	medalListURL       = "https://www.acfun.cn/rest/pc-direct/fansClub/fans/medal/list"
	liveInfoURL        = "https://live.acfun.cn/api/live/info?authorId=%d"
	userInfoURL        = "https://www.acfun.cn/rest/pc-direct/user/userInfo?userId=%d"
	liveListURL        = "https://live.acfun.cn/api/channel/list?count=%d&pcursor=%d"
	liveCutInfoURL     = "https://live.acfun.cn/rest/pc-direct/live/getLiveCutInfo?authorId=%d&liveId=%s"
	liveCutRedirectURL = "https://onvideoapi.kuaishou.com/rest/infra/sts?authToken=%s&sid=acfun.midground.api&followUrl=%s"
	//scheduleListURL    = "https://api-new.app.acfun.cn/rest/app/live/schedule/list"
)

const (
	managerKickURL     = "https://api.kuaishouzt.com/rest/zt/live/web/manager/kick?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	authorKickURL      = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/kick?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	addManagerURL      = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/manager/add?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	deleteManagerURL   = "https://api.kuaishouzt.com/rest/zt/live/web/author/action/manager/delete?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	wearMedalURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/fans/medal/wear?uperId=%d"
	cancelWearMedalURL = "https://live.acfun.cn/rest/pc-direct/fansClub/fans/medal/cancelWear?uperId=%d"
)

const (
	checkLiveAuthURL = "https://member.acfun.cn/common/api/checkLiveAuth"
	liveTypeListURL  = "https://member.acfun.cn/common/api/getLiveTypeList"
	liveDataURL      = "https://member.acfun.cn/dataCenter/api/liveData"
	liveCutStatusURL = "https://member.acfun.cn/liveToll/api/getUserLiveCut"
	updateLiveCutURL = "https://member.acfun.cn/liveTool/api/updateLiveCut"
	obsConfigURL     = "https://api.kuaishouzt.com/rest/zt/live/web/obs/config?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	obsStatusURL     = "https://api.kuaishouzt.com/rest/zt/live/web/obs/status?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	transcodeInfoURL = "https://api.kuaishouzt.com/rest/zt/live/web/obs/transcodeInfo?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	startPushURL     = `https://api.kuaishouzt.com/rest/zt/live/web/obs/startPush?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&acfun.midground.api_st=%s&videoPushReq=&streamName=%s&portrait=%v&isPanoramic=%v`
	stopPushURL      = "https://api.kuaishouzt.com/rest/zt/live/web/obs/stopPush?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	changeCoverURL   = `https://api.kuaishouzt.com/rest/zt/live/web/changeCover?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&acfun.midground.api_st=%s&videoPushReq=&liveId=%s`
)

var (
	generalParserPool fastjson.ParserPool
	medalParserPool   fastjson.ParserPool
)

type token struct {
	TokenInfo
	liveID          string
	enterRoomAttach string
	tickets         []string
	appID           int32
	instanceID      int64
	sessionKey      []byte        // 除第一次外发送 ws 信息时所用密钥
	seqID           *atomic.Int64 // 要用原子锁操作
	headerSeqID     *atomic.Int64 // 要用原子锁操作
	heartbeatSeqID  int64
	ticketIndex     *atomic.Uint32 // 要用原子锁操作
	giftsMutex      sync.RWMutex
	gifts           map[int64]GiftDetail
	liverUID        int64 // 主播 uid
	livePage        string
	err             *atomic.Error
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
