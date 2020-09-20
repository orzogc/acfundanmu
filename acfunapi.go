package acfundanmu

import (
	"fmt"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// WatchingUser 就是观看直播的用户的信息，目前没有Medal
type WatchingUser struct {
	UserInfo                      // 用户信息
	AnonymousUser          bool   // 是否匿名用户
	DisplaySendAmount      string // 赠送的全部礼物的价值，单位是ac币
	CustomWatchingListData string // 用户的一些额外信息，格式为json
}

// Summary 就是直播的总结信息
type Summary struct {
	LiveDurationMs int64 // 直播时长，单位为毫秒
	LikeCount      int   // 点赞总数
	WatchCount     int   // 观看过直播的人数总数
}

// MedalDetail 就是登陆帐号守护徽章的详细信息
type MedalDetail struct {
	MedalInfo
	UperName           string // UP主的名字
	UperAvatar         string // UP主的头像
	WearMedal          bool   // 是否正在佩戴该守护徽章
	FriendshipDegree   int    // 目前守护徽章的亲密度
	JoinClubTime       int64  // 加入守护团的时间，单位为毫秒
	CurrentDegreeLimit int    // 守护徽章目前等级的亲密度的上限
}

// 获取直播间排名前50的在线观众信息列表
func (t *token) watchingList() (watchList []WatchingUser, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("watchingList() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(watchingListURL)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	p := t.watchParser.Get()
	defer t.watchParser.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取在线观众列表失败，响应为 %s", string(body)))
	}

	watchArray := v.GetArray("data", "list")
	watchingUserList := make([]WatchingUser, len(watchArray))
	for i, watch := range watchArray {
		o := watch.GetObject()
		w := WatchingUser{}
		o.Visit(func(k []byte, v *fastjson.Value) {
			switch string(k) {
			case "userId":
				w.UserID = v.GetInt64()
			case "nickname":
				w.Nickname = string(v.GetStringBytes())
			case "avatar":
				w.Avatar = string(v.GetStringBytes("0", "url"))
			case "anonymousUser":
				w.AnonymousUser = v.GetBool()
			case "displaySendAmount":
				w.DisplaySendAmount = string(v.GetStringBytes())
			case "customWatchingListData":
				w.CustomWatchingListData = string(v.GetStringBytes())
			case "managerType":
				w.ManagerType = ManagerType(v.GetInt())
			}
		})
		watchingUserList[i] = w
	}

	return watchingUserList, nil
}

// 获取直播总结信息
func (t *token) getSummary() (summary Summary, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getSummary() error: %w", err)
		}
	}()

	resp, err := t.fetchKuaiShouAPI(endSummaryURL)
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("获取直播总结信息失败，响应为 %s", string(body)))
	}

	v = v.Get("data")
	summary.LiveDurationMs = v.GetInt64("liveDurationMs")
	summary.LikeCount, err = strconv.Atoi(string(v.GetStringBytes("likeCount")))
	checkErr(err)
	summary.WatchCount, err = strconv.Atoi(string(v.GetStringBytes("watchCount")))
	checkErr(err)

	return summary, nil
}

// 获取登陆帐号的守护徽章和指定主播守护徽章的名字
func getMedalInfo(uid int64, cookies []string) (medalList []MedalDetail, clubName string, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("getMedalInfo() error: %w", err)
		}
	}()

	var httpCookies []*fasthttp.Cookie
	for _, c := range cookies {
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)
		err := cookie.Parse(c)
		checkErr(err)
		httpCookies = append(httpCookies, cookie)
	}
	client := &httpClient{
		client: &fasthttp.Client{
			MaxIdleConnDuration: 90 * time.Second,
			ReadTimeout:         10 * time.Second,
			WriteTimeout:        10 * time.Second,
		},
		url:     fmt.Sprintf(medalInfoURL, uid),
		method:  "GET",
		cookies: httpCookies,
	}
	resp, err := client.doRequest()
	checkErr(err)
	defer fasthttp.ReleaseResponse(resp)
	body := resp.Body()

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("获取登陆帐号的守护徽章和指定主播守护徽章的名字失败，响应为 %s", string(body)))
	}

	clubName = string(v.GetStringBytes("clubName"))

	medalArray := v.GetArray("medalList")
	medalList = make([]MedalDetail, len(medalArray))
	for i, medal := range medalArray {
		medalList[i] = MedalDetail{
			MedalInfo: MedalInfo{
				UperID:   medal.GetInt64("uperId"),
				ClubName: string(medal.GetStringBytes("clubName")),
				Level:    medal.GetInt("level"),
			},
			UperName:           string(medal.GetStringBytes("uperName")),
			UperAvatar:         string(medal.GetStringBytes("uperHeadUrl")),
			WearMedal:          medal.GetBool("wearMedal"),
			FriendshipDegree:   medal.GetInt("friendshipDegree"),
			JoinClubTime:       medal.GetInt64("joinClubTime"),
			CurrentDegreeLimit: medal.GetInt("currentDegreeLimit"),
		}
	}

	return medalList, clubName, nil
}

// GetWatchingList 返回直播间排名前50的在线观众信息列表，不需要调用StartDanmu()
func (dq *DanmuQueue) GetWatchingList() ([]WatchingUser, error) {
	return dq.t.watchingList()
}

// GetSummary 返回直播总结信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetSummary() (Summary, error) {
	return dq.t.getSummary()
}

// GetMedalInfo 返回登陆用户的守护徽章列表medalList和uid指定主播的守护徽章的名字clubName
func GetMedalInfo(uid int64, cookies []string) (medalList []MedalDetail, clubName string, e error) {
	return getMedalInfo(uid, cookies)
}
