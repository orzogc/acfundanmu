package acfundanmu

import (
	"fmt"
	"strconv"

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
	LiveDurationMs int64 // 直播时长，以毫秒为单位
	LikeCount      int   // 点赞总数
	WatchCount     int   // 观看直播的人数总数
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
			e = fmt.Errorf("getEndSummary() error: %w", err)
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

// GetWatchingList 返回直播间排名前50的在线观众信息列表，不需要调用StartDanmu()
func (dq *DanmuQueue) GetWatchingList() ([]WatchingUser, error) {
	return dq.t.watchingList()
}

// GetSummary 返回直播总结信息，不需要调用StartDanmu()
func (dq *DanmuQueue) GetSummary() (Summary, error) {
	return dq.t.getSummary()
}
