package acfundanmu

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// 房管踢人
func (t *token) managerKick(liveID string, kickedUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("managerKick() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("房管踢人需要登陆房管的 AcFun 帐号"))
	}

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	form.Set("kickedUserId", strconv.FormatInt(kickedUID, 10))
	body, err := t.fetchKuaiShouAPI(managerKickURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 || !v.GetBool("data", "kickSucc") {
		panic(fmt.Errorf("房管踢人失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播踢人
func (t *token) authorKick(liveID string, kickedUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("authorKick() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("主播踢人需要登陆主播的 AcFun 帐号"))
	}

	form := t.defaultForm(liveID)
	defer fasthttp.ReleaseArgs(form)
	form.Set("kickedUserId", strconv.FormatInt(kickedUID, 10))
	body, err := t.fetchKuaiShouAPI(authorKickURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 || !v.GetBool("data", "kickSucc") {
		panic(fmt.Errorf("主播踢人失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播添加房管
func (t *token) addManager(managerUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("addManager() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("主播添加房管需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	form.Set("managerUserId", strconv.FormatInt(managerUID, 10))
	body, err := t.fetchKuaiShouAPI(addManagerURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("主播添加房管失败，响应为 %s", string(body)))
	}

	return nil
}

// 主播删除房管
func (t *token) deleteManager(managerUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("deleteManager() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("主播删除房管需要登陆主播的 AcFun 帐号"))
	}

	form := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(form)
	form.Set("visitorId", strconv.FormatInt(t.UserID, 10))
	form.Set("managerUserId", strconv.FormatInt(managerUID, 10))
	body, err := t.fetchKuaiShouAPI(deleteManagerURL, form, false)
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if v.GetInt("result") != 1 {
		panic(fmt.Errorf("主播删除房管失败，响应为 %s", string(body)))
	}

	return nil
}

// 佩戴守护徽章
func (t *token) wearMedal(uid int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("wearMedal() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("佩戴守护徽章需要登陆 AcFun 帐号"))
	}

	client := &httpClient{
		url:      fmt.Sprintf(wearMedalURL, uid),
		method:   "GET",
		cookies:  t.Cookies,
		referer:  t.livePage,
		deviceID: t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("佩戴守护徽章失败，响应为 %s", string(body)))
	}

	return nil
}

// 取消佩戴守护徽章
func (t *token) cancelWearMedal(liverUID int64) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("cancelWearMedal() error: %v", err)
		}
	}()

	if len(t.Cookies) == 0 {
		panic(fmt.Errorf("取消佩戴守护徽章需要登陆 AcFun 帐号"))
	}

	client := &httpClient{
		url:      fmt.Sprintf(cancelWearMedalURL, liverUID),
		method:   "GET",
		cookies:  t.Cookies,
		referer:  t.livePage,
		deviceID: t.DeviceID,
	}
	body, err := client.request()
	checkErr(err)

	p := generalParserPool.Get()
	defer generalParserPool.Put(p)
	v, err := p.ParseBytes(body)
	checkErr(err)
	if !v.Exists("result") || v.GetInt("result") != 0 {
		panic(fmt.Errorf("取消佩戴守护徽章失败，响应为 %s", string(body)))
	}

	return nil
}

// ManagerKick 房管踢人，需要登陆 AcFun 帐号，需要设置主播 uid
func (ac *AcFunLive) ManagerKick(liveID string, kickedUID int64) error {
	return ac.t.managerKick(liveID, kickedUID)
}

// AuthorKick 主播踢人，需要登陆 AcFun 帐号，需要设置主播 uid
func (ac *AcFunLive) AuthorKick(liveID string, kickedUID int64) error {
	return ac.t.authorKick(liveID, kickedUID)
}

// AddManager 主播添加房管，需要登陆 AcFun 帐号
func (ac *AcFunLive) AddManager(managerUID int64) error {
	return ac.t.addManager(managerUID)
}

// DeleteManager 主播删除房管，需要登陆 AcFun 帐号
func (ac *AcFunLive) DeleteManager(managerUID int64) error {
	return ac.t.deleteManager(managerUID)
}

// WearMedal 佩戴 uid 指定的主播的守护徽章，需要登陆 AcFun 帐号，如果登陆帐号没有 uid 指定的主播的守护徽章则会取消佩戴任何徽章
func (ac *AcFunLive) WearMedal(uid int64) error {
	return ac.t.wearMedal(uid)
}

// CancelWearMedalWithLiverUID 取消佩戴守护徽章，需要登陆 AcFun 帐号，liverUID 必须是登陆帐号正在佩戴的守护徽章的主播 uid
func (ac *AcFunLive) CancelWearMedalWithLiverUID(liverUID int64) error {
	return ac.t.cancelWearMedal(liverUID)
}

// CancelWearMedal 取消佩戴守护徽章，需要登陆 AcFun 帐号
func (ac *AcFunLive) CancelWearMedal() error {
	medal, err := getUserMedal(ac.t.UserID, ac.t.DeviceID)
	if err != nil {
		return err
	}

	return ac.t.cancelWearMedal(medal.UperID)
}
