# acfundanmu
AcFun直播弹幕下载，实现参照 [AcFunDanmaku](https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu)

### 示例代码
#### 获取弹幕
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
dq, err := acfundanmu.Start(ctx, uid)
if err != nil {
	log.Panicln(err)
}
for {
    if danmu := dq.GetDanmu(); danmu != nil {
        for _, d := range danmu {
            // 根据Type处理弹幕
            switch d.Type {
            case acfundanmu.Comment:
                fmt.Printf("%s（%d）：%s\n", d.Nickname, d.UserID, d.Comment)
            case acfundanmu.Like:
                fmt.Printf("%s（%d）点赞\n", d.Nickname, d.UserID)
            case acfundanmu.EnterRoom:
                fmt.Printf("%s（%d）进入直播间\n", d.Nickname, d.UserID)
            case acfundanmu.FollowAuthor:
                fmt.Printf("%s（%d）关注了主播\n", d.Nickname, d.UserID)
            case acfundanmu.ThrowBanana:
                fmt.Printf("%s（%d）送出香蕉 * %d\n", d.Nickname, d.UserID, d.BananaCount)
            case acfundanmu.Gift:
                fmt.Printf("%s（%d）送出礼物 %s * %d，连击数：%d\n", d.Nickname, d.UserID, d.Gift.Name, d.Gift.Count, d.Gift.Combo)
            }
        }
    } else {
        fmt.Println("直播结束")
        break
    }
}
```
#### 获取直播间状态信息
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
dq, err := acfundanmu.Start(ctx, uid)
if err != nil {
	log.Panicln(err)
}
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 循环获取info并处理
            time.Sleep(5 * time.Second)
            info := dq.GetInfo()
            fmt.Printf("%+v\n", info)
        }
    }
}()
// 做其他事情
```
#### 获取直播间排名前50的在线观众信息列表
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
dq, err := acfundanmu.Start(ctx, uid)
if err != nil {
	log.Panicln(err)
}
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 循环获取watchingList并处理
            watchingList := dq.GetWatchingList()
            fmt.Printf("%+v\n", *watchingList)
            time.Sleep(30 * time.Second)
        }
    }
}()
// 做其他事情
```
#### 将弹幕转换成ass字幕文件
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
dq, err := acfundanmu.Start(ctx, uid)
if err != nil {
	log.Panicln(err)
}
dq.WriteASS(ctx, acfundanmu.SubConfig{
    Title:     "foo",
    PlayResX:  1280, // 直播录播视频的分辨率
    PlayResY:  720,
    FontSize:  40,
    StartTime: time.Now().UnixNano()}, // 这里应该是开始录播的时间
    "foo.ass", true)
```
