package main

import (
	"fmt"
	"time"
)

// 心跳器：可以定期从心跳器的channel中读取数据
func main() {

	ticker := time.NewTicker(time.Millisecond * 500) // 产生一个心跳器，每隔500毫秒向心跳器的缓冲channel中放入Time对象
	go func() {
		for t := range ticker.C { // 从心跳器的channel中读取数据，每隔500毫秒则会读取一个Time对象
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 3600) // 睡眠3600毫秒
	ticker.Stop()                       // 停止心跳，停止后则不能再继续读取心跳器channel中的数据
	fmt.Println("Ticker stopped")

}
