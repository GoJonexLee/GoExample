package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	secs := now.Unix()			// 秒
	fmt.Println(secs)

	nanos := now.UnixNano()		// 纳秒
	fmt.Println(nanos)

	millis := nanos / 1000000
	fmt.Println(millis)			// 毫秒

	fmt.Println(time.Unix(secs, 0))		//秒转时间，0为纳秒
	fmt.Println(time.Unix(0, nanos))	//纳秒转时间

}
