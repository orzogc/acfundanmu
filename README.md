# acfundanmu

[![PkgGoDev](https://pkg.go.dev/badge/github.com/orzogc/acfundanmu)](https://pkg.go.dev/github.com/orzogc/acfundanmu)

AcFun 直播 API，弹幕实现参照 [AcFunDanmaku](https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu)

### 示例代码

#### 获取弹幕（非事件响应模式）

```go
// uid 为主播的 uid
ac, err := acfundanmu.NewAcFunLive(acfundanmu.SetLiverUID(uid))
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ch := ac.StartDanmu(ctx, false)
for {
    if danmu := ac.GetDanmu(); danmu != nil {
        for _, d := range danmu {
            switch d := d.(type) {
            case *acfundanmu.Comment:
                log.Printf("%s（%d）：%s\n", d.Nickname, d.UserID, d.Content)
            case *acfundanmu.Like:
                log.Printf("%s（%d）点赞\n", d.Nickname, d.UserID)
            case *acfundanmu.EnterRoom:
                log.Printf("%s（%d）进入直播间\n", d.Nickname, d.UserID)
            case *acfundanmu.FollowAuthor:
                log.Printf("%s（%d）关注了主播\n", d.Nickname, d.UserID)
            case *acfundanmu.ThrowBanana:
                log.Printf("%s（%d）送出香蕉 * %d\n", d.Nickname, d.UserID, d.BananaCount)
            case *acfundanmu.Gift:
                log.Printf("%s（%d）送出礼物 %s * %d，连击数：%d\n", d.Nickname, d.UserID, d.GiftName, d.Count, d.Combo)
            case *acfundanmu.RichText:
                for _, r := range d.Segments {
                    switch r := r.(type) {
                    case *acfundanmu.RichTextUserInfo:
                        log.Printf("富文本用户信息：%+v\n", *r)
                    case *acfundanmu.RichTextPlain:
                        log.Printf("富文本文字：%s，颜色：%s\n", r.Text, r.Color)
                    case *acfundanmu.RichTextImage:
                        for _, image := range r.Pictures {
                            log.Printf("富文本图片：%s\n", image)
                        }
                        log.Printf("富文本图片另外的文字：%s，颜色：%s\n", r.AlternativeText, r.AlternativeColor)
                    }
                }
            case *acfundanmu.JoinClub:
                log.Printf("%s（%d）加入主播%s（%d）的守护团", d.FansInfo.Nickname, d.FansInfo.UserID, d.UperInfo.Nickname, d.UperInfo.UserID)
            }
            case *acfundanmu.ShareLive:
                log.Printf("%s（%d）分享直播间到 %d %s", d.Nickname, d.UserID, d.SharePlatform, d.SharePlatformIcon)
        }
    } else {
        if err = <-ch; err != nil {
            log.Panicln(err)
        } else {
            log.Println("直播结束")
        }
        break
    }
}
```

#### 采用事件响应模式

```go
// uid 为主播的 uid
ac, err := acfundanmu.NewAcFunLive(acfundanmu.SetLiverUID(uid))
if err != nil {
    log.Panicln(err)
}
ac.OnDanmuStop(func(ac *acfundanmu.AcFunLive, err error) {
    if err != nil {
        log.Println(err)
    } else {
        log.Println("直播结束")
    }
})
ac.OnComment(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Comment) {
    log.Printf("%s（%d）：%s\n", d.Nickname, d.UserID, d.Content)
})
ac.OnLike(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Like) {
    log.Printf("%s（%d）点赞\n", d.Nickname, d.UserID)
})
ac.OnEnterRoom(func(ac *acfundanmu.AcFunLive, d *acfundanmu.EnterRoom) {
    log.Printf("%s（%d）进入直播间\n", d.Nickname, d.UserID)
})
ac.OnFollowAuthor(func(ac *acfundanmu.AcFunLive, d *acfundanmu.FollowAuthor) {
    log.Printf("%s（%d）关注了主播\n", d.Nickname, d.UserID)
})
ac.OnThrowBanana(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ThrowBanana) {
    log.Printf("%s（%d）送出香蕉 * %d\n", d.Nickname, d.UserID, d.BananaCount)
})
ac.OnGift(func(ac *acfundanmu.AcFunLive, d *acfundanmu.Gift) {
    log.Printf("%s（%d）送出礼物 %s * %d，连击数：%d\n", d.Nickname, d.UserID, d.GiftName, d.Count, d.Combo)
})
ac.OnJoinClub(func(ac *acfundanmu.AcFunLive, d *acfundanmu.JoinClub) {
    log.Printf("%s（%d）加入主播%s（%d）的守护团", d.FansInfo.Nickname, d.FansInfo.UserID, d.UperInfo.Nickname, d.UperInfo.UserID)
})
ac.OnShareLive(func(ac *acfundanmu.AcFunLive, d *acfundanmu.ShareLive) {
    log.Printf("%s（%d）分享直播间到 %d %s", d.Nickname, d.UserID, d.SharePlatform, d.SharePlatformIcon)
})
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
_ = ac.StartDanmu(ctx, true)
// 做其他事情
```

#### 获取直播间状态信息（非事件模式）

```go
// uid 为主播的 uid
ac, err := acfundanmu.NewAcFunLive(acfundanmu.SetLiverUID(uid))
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ch := ac.StartDanmu(ctx, false)
for {
    select {
    case <-ctx.Done():
        return
    default:
        // 循环获取 info 并处理
        time.Sleep(5 * time.Second)
        info := ac.GetLiveInfo()
        log.Printf("%+v\n", info)
    }
}
if err = <-ch; err != nil {
    log.Panicln(err)
} else {
    log.Println("直播结束")
}
```

#### 获取直播间排名前 50 的在线观众信息列表

```go
// uid 为主播的 uid
ac, err := acfundanmu.NewAcFunLive(acfundanmu.SetLiverUID(uid))
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
liveID := ac.GetLiveID()
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 循环获取 watchingList 并处理
            watchingList, err := ac.GetWatchingList(liveID)
            if err != nil {
                log.Panicln(err)
            }
            log.Printf("%+v\n", *watchingList)
            time.Sleep(30 * time.Second)
        }
    }
}()
// 做其他事情
```

#### 将弹幕转换成 ass 字幕文件

```go
// uid 为主播的 uid
ac, err := acfundanmu.NewAcFunLive(acfundanmu.SetLiverUID(uid))
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
ch := ac.StartDanmu(ctx, false)
ac.WriteASS(ctx, acfundanmu.SubConfig{
    Title:     "foo",
    PlayResX:  1280, // 直播录播视频的分辨率
    PlayResY:  720,
    FontSize:  40,
    StartTime: time.Now().UnixNano()}, // 这里应该是开始录播的时间
    "foo.ass", true)
if err = <-ch; err != nil {
    log.Panicln(err)
} else {
    log.Println("直播结束")
}
```
