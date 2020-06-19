# acfundanmu
AcFun直播弹幕下载，实现参照 [AcFunDanmaku](https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu)

### 示例代码
##### 获取弹幕
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
q := acfundanmu.Start(ctx, uid)
for {
    if d := q.GetDanmu(); d != nil {
        fmt.Println(d)
    } else {
        fmt.Println("直播结束")
        break
    }
}
```
##### 将弹幕转换成ass字幕文件
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
// uid为主播的uid
q := acfundanmu.Start(ctx, uid)
q.WriteASS(ctx, acfundanmu.SubConfig{
    Title:     "foo",
    PlayResX:  1280, // 直播录播视频的分辨率
    PlayResY:  720,
    FontSize:  40,
    StartTime: time.Now().UnixNano()}, // 这里应该是开始录播的时间
    "foo.ass")
```
