package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println // 将Println()函数赋给f

	now := time.Now() // 获取当前时间
	p(now)            // 注意打印的时间格式，并不是时间戳的格式
	p(now.Location()) // local: CST

	then := time.Date(2017, 01, 20, 23, 32, 10, 651387237, time.Local)
	p(then)

	p(then.Year())       // 年:2017
	p(then.Month())      // January
	p(then.Day())        // 20
	p(then.Hour())       // 23
	p(then.Minute())     // 32
	p(then.Second())     // 10
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	p(then.Weekday()) // Friday

	p(then.Before(now)) // 是否在now时间之前:   before(then) -- now -- after
	p(then.After(now))  // 是否在now之后
	p(then.Equal(now))  // 是否等于now

	diff := now.Sub(then) // 从now到then的时间窗口
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	now = time.Now() // now在12行已经定义过，因此不需要':='符号重新定义
	p(now)

	// 时间转字符串, 记忆方式：06(年)， 1(月)， 2(日)， 3(小时)， 4(分钟)， 5(秒)
	const datetime = "2006-01-02 15:04:05" // 这个点据说是go的诞生时间，并且格式化字符串必须是这个
	snow := now.Format(datetime)           // 时分秒
	p(snow)

	const time = "15:04:05" // 时间
	tnow := now.Format(time)
	p(tnow)

	const date = "2006-01-02" // 日期
	dnow := now.Format(date)
	p(dnow)
}
