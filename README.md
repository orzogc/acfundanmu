# acfundanmu
AcFun直播弹幕下载，实现参照 [AcFunDanmaku](https://github.com/wpscott/AcFunDanmaku/tree/master/AcFunDanmu)

### 示例代码
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
