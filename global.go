package acfundanmu

import (
	"github.com/valyala/fastjson"
	"go.uber.org/atomic"
)

const (
	visitorSt      = "acfun.api.visitor_st"
	midgroundSt    = "acfun.midground.api_st"
	acfunSignInURL = "https://id.app.acfun.cn/rest/web/login/signin"
	liveHost       = "https://live.acfun.cn"
	liveURL        = "https://live.acfun.cn/live/%d"
	loginURL       = "https://id.app.acfun.cn/rest/app/visitor/login"
	getTokenURL    = "https://id.app.acfun.cn/rest/web/token/get"
	playURL        = "https://api.kuaishouzt.com/rest/zt/live/web/startPlay?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"
	giftURL        = "https://api.kuaishouzt.com/rest/zt/live/web/gift/list?subBiz=mainApp&kpn=ACFUN_APP&kpf=PC_WEB&userId=%d&did=%s&%s=%s"

	sid       = "sid"
	visitor   = "acfun.api.visitor"
	midground = "acfun.midground.api"

	wsHost               = "wss://link.xiatou.com/"
	appName              = "link-sdk"
	sdkVersion           = "1.2.1"
	kpn                  = "ACFUN_APP"
	kpf                  = "PC_WEB"
	subBiz               = "mainApp"
	clientLiveSdkVersion = "kwai-acfun-live-link"
	linkVersion          = "2.13.8"

	retryCount uint32 = 1

	formContentType = "application/x-www-form-urlencoded"
	pushType        = `{"typeId":%d,"type":[%d,%d]}`
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
	userMedalURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/user/info?userId=%d"
	medalRankURL       = "https://live.acfun.cn/rest/pc-direct/fansClub/friendshipDegreeRankInfo?uperId=%d"
	medalDetailURL     = "https://api-new.app.acfun.cn/rest/app/fansClub/fans/medal/detail?uperId=%d"
	medalListURL       = "https://api-new.app.acfun.cn/rest/app/fansClub/live/medalInfo?uperId=%d"
	liveInfoURL        = "https://api-new.app.acfun.cn/rest/app/live/info?authorId=%d"
	liveListURL        = "https://api-new.app.acfun.cn/rest/app/live/channel"
	scheduleListURL    = "https://api-new.app.acfun.cn/rest/app/live/schedule/list"
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
	getQiniuTokenURL = "https://member.acfun.cn/common/api/getQiniuToken"
	liveDataURL      = "https://member.acfun.cn/dataCenter/api/liveData"
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
	sessionKey      []byte        // 除第一次外发送ws信息时所用密钥
	seqID           *atomic.Int64 // 要用原子锁操作
	headerSeqID     *atomic.Int64 // 要用原子锁操作
	heartbeatSeqID  int64
	ticketIndex     *atomic.Uint32 // 要用原子锁操作
	gifts           map[int64]GiftDetail
	liverUID        int64 // 主播uid
	livePage        string
}

// 检查错误
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
