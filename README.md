# acfundanmu
[![GoDoc](https://godoc.org/github.com/orzogc/acfundanmu?status.png)](https://godoc.org/github.com/orzogc/acfundanmu) [![PkgGoDev](https://pkg.go.dev/badge/github.com/orzogc/acfundanmu)](https://pkg.go.dev/github.com/orzogc/acfundanmu)

AcFun直播API，弹幕实现参照 [AcFunDanmaku](https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu)

### 示例代码
#### 获取弹幕
```go
// uid为主播的uid
dq, err := acfundanmu.Init(uid, nil)
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
dq.StartDanmu(ctx, false)
for {
    if danmu := dq.GetDanmu(); danmu != nil {
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
                    case acfundanmu.RichTextUserInfo:
                        log.Printf("富文本用户信息：%+v\n", r)
                    case acfundanmu.RichTextPlain:
                        log.Printf("富文本文字：%s，颜色：%s\n", r.Text, r.Color)
                    case acfundanmu.RichTextImage:
                        for _, image := range r.Pictures {
                            log.Printf("富文本图片：%s\n", image)
                        }
                        log.Printf("富文本图片另外的文字：%s，颜色：%s\n", r.AlternativeText, r.AlternativeColor)
                    }
                }
            }
        }
    } else {
        log.Println("直播结束")
        break
    }
}
```
#### 获取直播间状态信息
```go
// uid为主播的uid
dq, err := acfundanmu.Init(uid, nil)
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
dq.StartDanmu(ctx, false)
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 循环获取info并处理
            time.Sleep(5 * time.Second)
            info := dq.GetLiveInfo()
            log.Printf("%+v\n", info)
        }
    }
}()
// 做其他事情
```
#### 获取直播间排名前50的在线观众信息列表
```go
// uid为主播的uid
dq, err := acfundanmu.Init(uid, nil)
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 循环获取watchingList并处理
            watchingList, err := dq.GetWatchingList()
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
#### 将弹幕转换成ass字幕文件
```go
// uid为主播的uid
dq, err := acfundanmu.Init(uid, nil)
if err != nil {
    log.Panicln(err)
}
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
dq.StartDanmu(ctx, false)
dq.WriteASS(ctx, acfundanmu.SubConfig{
    Title:     "foo",
    PlayResX:  1280, // 直播录播视频的分辨率
    PlayResY:  720,
    FontSize:  40,
    StartTime: time.Now().UnixNano()}, // 这里应该是开始录播的时间
    "foo.ass", true)
```
