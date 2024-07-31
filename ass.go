package acfundanmu

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Workiva/go-datastructures/queue"
)

// ass 文件的 Script Info
const scriptInfo = `[Script Info]
; LiveID: %s
; StreamName: %s
Title: %s
ScriptType: v4.00+
Collisions: Normal
PlayResX: %d
PlayResY: %d

`

// ass 文件的 V4+ Styles
const sytles = `[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: Danmu,Microsoft YaHei,%d,&H00FFFFFF,&H00FFFFFF,&H00000000,&H00000000,0,0,0,0,100,100,0,0,1,1,0,2,20,20,2,0

`

// ass 文件的 Events
const events = `[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`

// 弹幕字幕
const dialogue = `Dialogue: 0,%s,%s,Danmu,%s(%d),20,20,2,,{\move(%d,%d,%d,%d)}%s
`

// ass 文件里弹幕出现或消失的时间格式
const startEnd = `%d:%02d:%02d.%02d`

// 弹幕持续时间，单位为纳秒
const duration = int64(10 * time.Second)

// 弹幕在视频里出现和消失的时间，单位为纳秒
type danmuTime int64

// SubConfig 是字幕的详细设置
type SubConfig struct {
	Title     string `json:"title"`     // 字幕标题
	PlayResX  int    `json:"playResX"`  // 视频分辨率
	PlayResY  int    `json:"playResY"`  // 视频分辨率
	FontSize  int    `json:"fontSize"`  // 字体大小
	StartTime int64  `json:"startTime"` // 直播录播开始的时间，是以纳秒为单位的 Unix 时间
}

// dTime 就是计算弹幕碰撞需要的数据
type dTime struct {
	appear    int64 // 弹幕出现的时间
	emerge    int64 // 整个弹幕完全出现在视频右边的时间
	disappear int64 // 弹幕消失的时间
}

// 将指定时间转换为 ass 字幕特定格式
func (d danmuTime) String() string {
	if d < 0 {
		return fmt.Sprintf(startEnd, 0, 0, 0, 0)
	}
	t := time.Unix(0, int64(d)).UTC()
	return fmt.Sprintf(
		startEnd,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1e7,
	)
}

// 不能使用","，需要转换用户昵称
func convert(name string) string {
	return strings.ReplaceAll(name, ",", " ")
}

// WriteASS 将 ass 字幕写入到 file 里，s 为字幕的设置，ctx 用来结束写入 ass 字幕，需要先调用 StartDanmu(ctx, false)。
// newFile 为 true 时覆盖写入，为 false 时不覆盖写入且只写入 Dialogue 字幕。
// 一个 AcFunLive 只能同时调用 WriteASS() 一次。
func (ac *AcFunLive) WriteASS(ctx context.Context, s SubConfig, file string, newFile bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Recovering from panic in WriteASS(), the error is: %v", err)
			log.Println("停止写入 ass 字幕")
		}
	}()

	if ac.q == nil {
		log.Println("需要先调用 StartDanmu()，event 不能为 true")
		return
	}
	if ac.t.liverUID == 0 {
		log.Println("主播 uid 不能为 0")
		return
	}
	if (*queue.Queue)(ac.q).Disposed() {
		return
	}

	var f *os.File
	var err error
	if newFile {
		f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		checkErr(err)
		defer f.Close()

		info := fmt.Sprintf(scriptInfo, ac.info.LiveID, ac.info.StreamName, s.Title, s.PlayResX, s.PlayResY)
		style := fmt.Sprintf(sytles, s.FontSize)

		_, err = f.WriteString(info)
		checkErr(err)
		_, err = f.WriteString(style)
		checkErr(err)
		_, err = f.WriteString(events)
		checkErr(err)
	} else {
		f, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		checkErr(err)
		defer f.Close()
	}

	// lastTime 存放每一行最后的弹幕的 dTime
	lastTime := make([]dTime, queueLen)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			danmu := ac.GetDanmu()
			if danmu == nil {
				return
			}

			for _, d := range danmu {
				c, ok := d.(*Comment)
				if !ok {
					continue
				}

				length := utf8.RuneCountInString(c.Content) * s.FontSize
				sendTime := c.SendTime * 1e6
				// leftTime 就是弹幕运动到视频左边的时间
				leftTime := sendTime - s.StartTime + (int64(s.PlayResX)*duration)/int64(s.PlayResX+length)
				dt := dTime{
					appear:    sendTime - s.StartTime,
					emerge:    sendTime - s.StartTime + (int64(length)*duration)/int64(s.PlayResX+length),
					disappear: sendTime - s.StartTime + duration}
				for i, t := range lastTime {
					// 防止弹幕发生碰撞重叠
					if dt.appear > t.emerge && leftTime > t.disappear {
						lastTime[i] = dt
						s := fmt.Sprintf(dialogue,
							danmuTime(dt.appear),
							danmuTime(dt.disappear),
							convert(c.Nickname),
							c.UserID,
							s.PlayResX+length/2,
							s.FontSize*(i+1),
							-length/2,
							s.FontSize*(i+1),
							c.Content,
						)
						_, err = f.WriteString(s)
						checkErr(err)
						break
					}
				}
			}
		}
	}
}
