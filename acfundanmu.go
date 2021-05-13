package acfundanmu

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/segmentio/encoding/json"
	"github.com/valyala/fasthttp"
)

// 弹幕队列长度
const queueLen = 1024

// ManagerType 就是房管类型
type ManagerType int32

const (
	// NotManager 不是房管
	NotManager ManagerType = iota
	// NormalManager 是房管
	NormalManager
)

// ManagerState 就是房管状态
type ManagerState int32

const (
	// ManagerStateUnknown 未知的房管状态，通常说明登陆用户不是房管
	ManagerStateUnknown ManagerState = iota
	// ManagerAdded 登陆用户被添加房管权限？
	ManagerAdded
	// ManagerRemoved 登陆用户被移除房管权限？
	ManagerRemoved
	// IsManager 登陆用户是房管
	IsManager
)

// RedpackDisplayStatus 红包状态
type RedpackDisplayStatus int32

const (
	// RedpackShow 红包出现？
	RedpackShow RedpackDisplayStatus = iota
	// RedpackGetToken 可以获取红包token？
	RedpackGetToken
	// RedpackGrab 可以抢红包
	RedpackGrab
)

// ChatMediaType 连麦类型
type ChatMediaType int32

const (
	// ChatMediaUnknown 未知的连麦类型
	ChatMediaUnknown ChatMediaType = iota
	// ChatMediaAudio 语音连麦
	ChatMediaAudio
	// ChatMediaVideo 视频连麦
	ChatMediaVideo
)

// ChatEndType 连麦结束类型
type ChatEndType int32

const (
	// ChatEndUnknown 未知的连麦结束类型
	ChatEndUnknown ChatEndType = iota
	// ChatEndCancelByAuthor 连麦发起者（主播）取消连麦
	ChatEndCancelByAuthor
	// ChatEndByAuthor 连麦发起者（主播）结束连麦
	ChatEndByAuthor
	// ChatEndByGuest 被连麦的人结束连麦
	ChatEndByGuest
	// ChatEndGuestReject 被连麦的人拒绝连麦
	ChatEndGuestReject
	// ChatEndGuestTimeout 等待连麦超时
	ChatEndGuestTimeout
	// ChatEndGuestHeartbeatTimeout 被连麦的人Heartbeat超时
	ChatEndGuestHeartbeatTimeout
	// ChatEndAuthorHeartbeatTimeout 连麦发起者（主播）Heartbeat超时
	ChatEndAuthorHeartbeatTimeout
	// ChatEndPeerLiveStopped 直播下播？
	ChatEndPeerLiveStopped
)

// SoundConfigChangeType 声音设置更改的类型
type SoundConfigChangeType int32

const (
	// SoundConfigChangeUnknown 未知的声音设置更改的类型
	SoundConfigChangeUnknown SoundConfigChangeType = iota
	// SoundConfigChangeOpenSound 打开声音
	SoundConfigChangeOpenSound
	// SoundConfigChangeCloseSound 关闭声音
	SoundConfigChangeCloseSound
)

// GiftDetail 就是礼物的详细信息
type GiftDetail struct {
	GiftID                 int64  `json:"giftID"`                 // 礼物ID
	GiftName               string `json:"giftName"`               // 礼物名字
	ARLiveName             string `json:"arLiveName"`             // 不为空时礼物属于虚拟偶像区的特殊礼物
	PayWalletType          int    `json:"payWalletType"`          // 1为付费礼物，2为免费礼物
	Price                  int    `json:"price"`                  // 礼物价格，付费礼物时单位为AC币，免费礼物（香蕉）时为1
	WebpPic                string `json:"webpPic"`                // 礼物的webp格式图片（动图）
	PngPic                 string `json:"pngPic"`                 // 礼物的png格式图片（大）
	SmallPngPic            string `json:"smallPngPic"`            // 礼物的png格式图片（小）
	AllowBatchSendSizeList []int  `json:"allowBatchSendSizeList"` // 网页或APP单次能够赠送的礼物数量列表
	CanCombo               bool   `json:"canCombo"`               // 是否能连击
	CanDraw                bool   `json:"canDraw"`                // 是否能涂鸦
	MagicFaceID            int    `json:"magicFaceID"`
	VupArID                int    `json:"vupArID"`
	Description            string `json:"description"`  // 礼物的描述
	RedpackPrice           int    `json:"redpackPrice"` // 礼物红包价格总额，单位为AC币
}

// DrawPoint 单个涂鸦礼物的位置
type DrawPoint struct {
	MarginLeft int64   `json:"marginLeft"` // 到手机屏幕左边的距离
	MarginTop  int64   `json:"marginTop"`  // 到手机屏幕顶部的距离
	ScaleRatio float64 `json:"scaleRatio"` // 放大倍数？
	Handup     bool    `json:"handup"`
}

// DrawGiftInfo 涂鸦礼物信息
type DrawGiftInfo struct {
	ScreenWidth  int64       `json:"screenWidth"`  // 手机屏幕宽度
	ScreenHeight int64       `json:"screenHeight"` // 手机屏幕高度
	DrawPoint    []DrawPoint `json:"drawPoint"`
}

// UserInfo 就是用户信息
type UserInfo struct {
	UserID      int64       `json:"userID"`      // 用户uid
	Nickname    string      `json:"nickname"`    // 用户名字
	Avatar      string      `json:"avatar"`      // 用户头像
	Medal       MedalInfo   `json:"medal"`       // 用户正在佩戴的守护徽章
	ManagerType ManagerType `json:"managerType"` // 用户是否房管
}

// MedalInfo 就是守护徽章信息
type MedalInfo struct {
	UperID   int64  `json:"uperID"`   // UP主的uid
	UserID   int64  `json:"userID"`   // 用户的uid
	ClubName string `json:"clubName"` // 守护徽章名字
	Level    int    `json:"level"`    // 守护徽章等级
}

// RichTextSegment 副文本片段的接口
type RichTextSegment interface {
	RichTextType() string
}

// RichTextUserInfo 富文本里的用户信息
type RichTextUserInfo struct {
	UserInfo `json:"userInfo"`
	Color    string `json:"color"` // 用户信息的颜色
}

// RichTextPlain 富文本里的文字
type RichTextPlain struct {
	Text  string `json:"text"`  // 文字
	Color string `json:"color"` // 文字的颜色
}

// RichTextImage 富文本里的图片
type RichTextImage struct {
	Pictures         []string `json:"pictures"`         // 图片
	AlternativeText  string   `json:"alternativeText"`  // 可选的文本？
	AlternativeColor string   `json:"alternativeColor"` // 可选的文本颜色？
}

// DanmuMessage 弹幕的接口
type DanmuMessage interface {
	GetSendTime() int64     // 获取弹幕发送时间
	GetUserInfo() *UserInfo // 获取UserInfo
}

// DanmuCommon 弹幕通用部分
type DanmuCommon struct {
	SendTime int64 `json:"sendTime"` // 弹幕发送时间，是以毫秒为单位的Unix时间
	UserInfo `json:"userInfo"`
}

// Comment 用户发的弹幕
type Comment struct {
	DanmuCommon `json:"danmuInfo"`
	Content     string `json:"content"` // 弹幕文字内容
}

// Like 用户点赞的弹幕
type Like DanmuCommon

// EnterRoom 用户进入直播间的弹幕
type EnterRoom DanmuCommon

// FollowAuthor 用户关注主播的弹幕
type FollowAuthor DanmuCommon

// ThrowBanana 用户投蕉的弹幕，没有Avatar、Medal和ManagerType，现在基本不用这个，通常用Gift代替
type ThrowBanana struct {
	DanmuCommon `json:"danmuInfo"`
	BananaCount int `json:"bananaCount"` // 投蕉数量
}

// Gift 用户赠送礼物的弹幕
type Gift struct {
	DanmuCommon         `json:"danmuInfo"`
	GiftDetail          `json:"giftDetail"`
	Count               int32        `json:"count"`               // 礼物单次赠送的数量，礼物总数是Count * Combo
	Combo               int32        `json:"combo"`               // 礼物连击数量，礼物总数是Count * Combo
	Value               int64        `json:"value"`               // 礼物价值，付费礼物时单位为AC币*1000，免费礼物（香蕉）时单位为礼物数量
	ComboID             string       `json:"comboID"`             // 礼物连击ID
	SlotDisplayDuration int64        `json:"slotDisplayDuration"` // 应该是礼物动画持续的时间，单位为毫秒，送礼物后在该时间内再送一次可以实现礼物连击
	ExpireDuration      int64        `json:"ExpireDuration"`
	DrawGiftInfo        DrawGiftInfo `json:"drawGiftInfo"` // 礼物涂鸦
}

// RichText 富文本，目前是用于发红包和抢红包的相关消息
type RichText struct {
	SendTime int64             `json:"sendTime"` // 富文本的发送时间，是以毫秒为单位的Unix时间，可能为0
	Segments []RichTextSegment `json:"segments"` // 富文本各部分，类型是RichTextUserInfo、RichTextPlain或RichTextImage
}

// JoinClub 用户加入主播的守护团，FansInfo和UperInfo都没有Avatar、Medal和ManagerType
type JoinClub struct {
	JoinTime int64    `json:"joinTime"` // 用户加入守护团的时间，是以毫秒为单位的Unix时间
	FansInfo UserInfo `json:"fansInfo"` // 用户的信息
	UperInfo UserInfo `json:"uperInfo"` // 主播的信息
}

// TopUser 就是礼物榜在线前三，目前没有Medal和ManagerType
type TopUser WatchingUser

// Redpack 红包信息
type Redpack struct {
	UserInfo           `json:"userInfo"`    // 发红包的用户
	DisplayStatus      RedpackDisplayStatus `json:"displayStatus"`      // 红包的状态
	GrabBeginTime      int64                `json:"grabBeginTime"`      // 开始抢红包的时间，是以毫秒为单位的Unix时间
	GetTokenLatestTime int64                `json:"getTokenLatestTime"` // 抢红包的用户获得token的最晚时间？是以毫秒为单位的Unix时间
	RedpackID          string               `json:"redpackID"`          // 红包ID
	RedpackBizUnit     string               `json:"redpackBizUnit"`     // "ztLiveAcfunRedpackGift"代表的是观众，"ztLiveAcfunRedpackAuthor"代表的是主播？
	RedpackAmount      int64                `json:"redpackAmount"`      // 红包的总价值，单位是AC币
	SettleBeginTime    int64                `json:"settleBeginTime"`    // 抢红包的结束时间，是以毫秒为单位的Unix时间
}

// ChatCall 主播发起连麦
type ChatCall struct {
	ChatID   string `json:"chatID"`   // 连麦ID
	LiveID   string `json:"liveID"`   // 直播ID
	CallTime int64  `json:"callTime"` // 连麦发起时间，是以毫秒为单位的Unix时间
}

// ChatAccept 用户接受连麦
type ChatAccept struct {
	ChatID     string        `json:"chatID"`    // 连麦ID
	MediaType  ChatMediaType `json:"mediaType"` // 连麦类型
	SignalInfo string        `json:"signalInfo"`
}

// ChatReady 用户接受连麦的信息
type ChatReady struct {
	ChatID    string        `json:"chatID"`    // 连麦ID
	Guest     UserInfo      `json:"guest"`     // 被连麦的帐号信息，目前没有房管类型
	MediaType ChatMediaType `json:"mediaType"` // 连麦类型
}

// ChatEnd 连麦结束
type ChatEnd struct {
	ChatID  string      `json:"chatID"`  // 连麦ID
	EndType ChatEndType `json:"endType"` // 连麦结束类型
}

// AuthorChatPlayerInfo 主播之间连麦的主播信息
type AuthorChatPlayerInfo struct {
	UserInfo               `json:"userInfo"`
	LiveID                 string `json:"liveID"`                 // 直播ID
	EnableJumpPeerLiveRoom bool   `json:"enableJumpPeerLiveRoom"` // 允许跳转到连麦的主播直播间？
}

// AuthorChatCall 主播发起连麦
type AuthorChatCall struct {
	Inviter  AuthorChatPlayerInfo `json:"inviter"`  // 发起连麦的主播的用户信息
	ChatID   string               `json:"chatID"`   // 连麦ID
	CallTime int64                `json:"callTime"` // 连麦发起时间，是以毫秒为单位的Unix时间
}

// AuthorChatAccept 主播接受连麦
type AuthorChatAccept struct {
	ChatID     string `json:"chatID"` // 连麦ID
	SignalInfo string `json:"signalInfo"`
}

// AuthorChatReady 主播接受连麦的信息
type AuthorChatReady struct {
	Inviter AuthorChatPlayerInfo `json:"inviter"` // 发起连麦的主播的用户信息
	Invitee AuthorChatPlayerInfo `json:"invitee"` // 接受连麦的主播的用户信息
	ChatID  string               `json:"chatID"`  // 连麦ID
}

// AuthorChatEnd 主播连麦结束
type AuthorChatEnd struct {
	ChatID    string      `json:"chatID"`    // 连麦ID
	EndType   ChatEndType `json:"endType"`   // 连麦结束类型
	EndLiveID string      `json:"endLiveID"` // 结束连麦的直播ID
}

// AuthorChatChangeSoundConfig 主播连麦更改声音设置
type AuthorChatChangeSoundConfig struct {
	ChatID                string                `json:"chatID"`                // 连麦ID
	SoundConfigChangeType SoundConfigChangeType `json:"soundConfigChangeType"` // 声音设置更改的类型
}

// Cookies 就是AcFun帐号的cookies
type Cookies []*fasthttp.Cookie

// TokenInfo 就是AcFun直播的token相关信息
type TokenInfo struct {
	UserID       int64   `json:"userID"`       // 登陆模式或游客模式的uid
	SecurityKey  string  `json:"securityKey"`  // 密钥，第一次发送ws信息时用
	ServiceToken string  `json:"serviceToken"` // 令牌
	DeviceID     string  `json:"deviceID"`     // 设备ID
	Cookies      Cookies `json:"cookies"`      // AcFun帐号的cookies
}

// StreamURL 就是直播源相关信息
type StreamURL struct {
	URL         string `json:"url"`         // 直播源链接
	Bitrate     int    `json:"bitrate"`     // 直播源码率，不一定是实际码率
	QualityType string `json:"qualityType"` // 直播源类型，一般是"SMOOTH"、"STANDARD"、"HIGH"、"SUPER"、"BLUE_RAY"
	QualityName string `json:"qualityName"` // 直播源类型的中文名字，一般是"流畅"、"高清"、"超清"、"蓝光 4M"、"蓝光 5M"、"蓝光 6M"、"蓝光 7M"、"蓝光 8M"
}

// StreamInfo 就是直播的直播源信息
type StreamInfo struct {
	LiveID        string      `json:"liveID"`        // 直播ID
	Title         string      `json:"title"`         // 直播间标题
	LiveStartTime int64       `json:"liveStartTime"` // 直播开始的时间，是以毫秒为单位的Unix时间
	Panoramic     bool        `json:"panoramic"`     // 是否全景直播
	StreamList    []StreamURL `json:"streamList"`    // 直播源列表
	StreamName    string      `json:"streamName"`    // 直播源名字（ID）
}

// DisplayInfo 就是直播间的一些数据
type DisplayInfo struct {
	WatchingCount string `json:"watchingCount"` // 直播间在线观众数量
	LikeCount     string `json:"likeCount"`     // 直播间点赞总数
	LikeDelta     int    `json:"likeDelta"`     // 点赞增加数量
}

// LiveInfo 就是直播间的相关状态信息
type LiveInfo struct {
	KickedOut        string       `json:"kickedOut"`        // 被踢理由？
	ViolationAlert   string       `json:"violationAlert"`   // 直播间警告？
	LiveManagerState ManagerState `json:"liveManagerState"` // 登陆帐号的房管状态
	AllBananaCount   string       `json:"allBananaCount"`   // 直播间香蕉总数
	DisplayInfo      `json:"displayInfo"`
	TopUsers         []TopUser `json:"topUsers"`      // 礼物榜在线前三
	RecentComment    []Comment `json:"recentComment"` // APP进直播间时显示的最近发的弹幕
	RedpackList      []Redpack `json:"redpackList"`   // 红包列表
}

// 带锁的LiveInfo
type liveInfo struct {
	sync.RWMutex // LiveInfo的锁
	LiveInfo
	StreamInfo
}

// AcFunLive 就是直播间弹幕系统相关信息，支持并行
type AcFunLive struct {
	q          *queue.Queue // DanmuMessage的队列
	info       *liveInfo    // 直播间的相关信息状态
	t          *token       // 令牌相关信息
	handlerMap *handlerMap  // 事件handler的map
}

// Option 就是AcFunLive的选项
type Option func(*AcFunLive)

// SetLiverUID 设置主播uid
func SetLiverUID(uid int64) Option {
	if uid <= 0 {
		return func(ac *AcFunLive) {}
	}
	return func(ac *AcFunLive) {
		ac.t.liverUID = uid
		ac.t.livePage = fmt.Sprintf(liveURL, uid)
	}
}

// SetCookies 设置AcFun帐号的cookies
func SetCookies(cookies Cookies) Option {
	return func(ac *AcFunLive) {
		ac.t.Cookies = append([]*fasthttp.Cookie{}, cookies...)
	}
}

// SetTokenInfo 设置TokenInfo
func SetTokenInfo(tokenInfo *TokenInfo) Option {
	if tokenInfo == nil {
		return func(ac *AcFunLive) {}
	}
	return func(ac *AcFunLive) {
		ac.t.TokenInfo = TokenInfo{
			UserID:       tokenInfo.UserID,
			SecurityKey:  tokenInfo.SecurityKey,
			ServiceToken: tokenInfo.ServiceToken,
			DeviceID:     tokenInfo.DeviceID,
			Cookies:      append([]*fasthttp.Cookie{}, tokenInfo.Cookies...),
		}
	}
}

// MarshalJSON 实现json的Marshaler接口
func (c Cookies) MarshalJSON() ([]byte, error) {
	cookies := make([]string, 0, len(c))
	for _, cookie := range c {
		cookies = append(cookies, cookie.String())
	}

	return json.Marshal(cookies)
}

// UnmarshalJSON 实现json的Unmarshaler接口
func (c *Cookies) UnmarshalJSON(b []byte) error {
	cookies := new([]string)
	err := json.Unmarshal(b, cookies)
	if err != nil {
		return err
	}
	*c = make(Cookies, 0, len(*cookies))
	for _, cookie := range *cookies {
		co := fasthttp.AcquireCookie()
		err = co.Parse(cookie)
		if err != nil {
			return err
		}
		*c = append(*c, co)
	}

	return nil
}

// Login 登陆AcFun帐号，account为帐号邮箱或手机号，password为帐号密码
func Login(account, password string) (cookies Cookies, err error) {
	if account == "" || password == "" {
		return nil, fmt.Errorf("AcFun帐号邮箱或密码为空，无法登陆")
	}

	for retry := 0; retry < 3; retry++ {
		cookies, err = login(account, password)
		if err != nil {
			if retry == 2 {
				log.Printf("登陆AcFun帐号失败：%v", err)
				return nil, fmt.Errorf("Login() error: 登陆AcFun帐号失败：%w", err)
			}
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	return cookies, nil
}

// NewAcFunLive 新建一个 *AcFunLive
func NewAcFunLive(options ...Option) (ac *AcFunLive, err error) {
	ac = new(AcFunLive)
	ac.info = new(liveInfo)
	ac.t = new(token)
	ac.t.livePage = liveHost
	ac.handlerMap = new(handlerMap)
	ac.handlerMap.listMap = make(map[eventType][]eventHandler)

	for _, option := range options {
		option(ac)
	}

	for retry := 0; retry < 3; retry++ {
		if ac.t.UserID == 0 {
			ac.info.StreamInfo, err = ac.t.getToken()
		} else {
			ac.info.StreamInfo, err = ac.t.getLiveToken()
		}
		if err != nil {
			if retry == 2 {
				log.Printf("初始化失败：%v", err)
				return nil, fmt.Errorf("NewAcFunLive() error: 初始化失败，主播可能不在直播：%w", err)
			}
		} else {
			break
		}
		time.Sleep(10 * time.Second)
	}

	return ac, nil
}

// SetLiverUID 设置主播uid，返回一个新的 *AcFunLive，不会复制弹幕获取采用事件响应模式时的事件处理函数
func (ac *AcFunLive) SetLiverUID(uid int64) (newAC *AcFunLive, err error) {
	tokenInfo := ac.GetTokenInfo()
	newAC, err = NewAcFunLive(SetLiverUID(uid), SetTokenInfo(tokenInfo))
	if err != nil {
		return nil, err
	}

	return newAC, nil
}

// CopyEventHandlers 弹幕获取采用事件响应模式时复制 anotherAC 的事件处理函数到 ac
func (ac *AcFunLive) CopyEventHandlers(anotherAC *AcFunLive) {
	anotherAC.handlerMap.RLock()
	defer anotherAC.handlerMap.RUnlock()
	ac.handlerMap.Lock()
	defer ac.handlerMap.Unlock()
	for k, v := range anotherAC.handlerMap.listMap {
		ac.handlerMap.listMap[k] = v
	}
}

// StartDanmu 启动websocket获取弹幕，ctx用来结束websocket，event为true时采用事件响应模式。
// event为false时最好调用GetDanmu()或WriteASS()以清空弹幕队列。
func (ac *AcFunLive) StartDanmu(ctx context.Context, event bool) <-chan error {
	ch := make(chan error, 1)
	if ac.t.liverUID == 0 {
		err := fmt.Errorf("主播uid不能为0")
		log.Println(err)
		ch <- err
		return ch
	}
	if !event {
		ac.q = queue.New(queueLen)
	}
	go ac.wsStart(ctx, event, ch)
	return ch
}

// GetDanmu 返回弹幕数据danmu，danmu为nil时说明弹幕获取结束（出现错误或者主播下播），需要先调用StartDanmu(ctx, false)
func (ac *AcFunLive) GetDanmu() (danmu []DanmuMessage) {
	if ac.q == nil {
		log.Println("需要先调用StartDanmu()，event不能为true")
		return nil
	}
	if ac.t.liverUID == 0 {
		log.Println("主播uid不能为0")
		return nil
	}
	if (*queue.Queue)(ac.q).Disposed() {
		return nil
	}
	ds, err := ac.q.Get(queueLen)
	if err != nil {
		return nil
	}

	danmu = make([]DanmuMessage, len(ds))
	for i, d := range ds {
		danmu[i] = d.(DanmuMessage)
	}

	return danmu
}

// GetLiveInfo 返回直播间的状态信息，需要先调用StartDanmu(ctx, false)
func (ac *AcFunLive) GetLiveInfo() *LiveInfo {
	ac.info.RLock()
	defer ac.info.RUnlock()
	info := ac.info.LiveInfo
	info.TopUsers = append([]TopUser{}, ac.info.TopUsers...)
	info.RecentComment = append([]Comment{}, ac.info.RecentComment...)
	info.RedpackList = append([]Redpack{}, ac.info.RedpackList...)
	return &info
}

// GetTokenInfo 返回直播间token相关信息，不需要调用StartDanmu()
func (ac *AcFunLive) GetTokenInfo() *TokenInfo {
	info := ac.t.TokenInfo
	info.Cookies = append([]*fasthttp.Cookie{}, ac.t.Cookies...)
	return &info
}

// GetStreamInfo 返回直播的直播源信息，不需要调用StartDanmu()
func (ac *AcFunLive) GetStreamInfo() *StreamInfo {
	info := ac.info.StreamInfo
	info.StreamList = append([]StreamURL{}, ac.info.StreamList...)
	return &info
}

// GetUserID 返回AcFun帐号的uid
func (ac *AcFunLive) GetUserID() int64 {
	return ac.t.UserID
}

// GetLiverUID 返回主播的uid，有可能是0
func (ac *AcFunLive) GetLiverUID() int64 {
	return ac.t.liverUID
}

// GetLiveID 返回liveID，有可能为空
func (ac *AcFunLive) GetLiveID() string {
	return ac.t.liveID
}

// GetTokenInfo 返回TokenInfo，cookies可以利用Login()获取，为nil时为游客模式
func GetTokenInfo(cookies Cookies) (*TokenInfo, error) {
	ac, err := NewAcFunLive(SetCookies(cookies))
	if err != nil {
		return nil, err
	}
	return ac.GetTokenInfo(), nil
}
