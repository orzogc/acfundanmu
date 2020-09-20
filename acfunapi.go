package acfundanmu

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

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

// 获取直播结束后的总结信息
func (t *token) getEndSummary() (summary EndSummary, e error) {
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

// GetEndSummary 返回直播结束时的总结信息，需要在直播结束前调用Init()，直播结束后调用本函数，不需要调用StartDanmu()
func (dq *DanmuQueue) GetEndSummary() (EndSummary, error) {
	return dq.t.getEndSummary()
}
