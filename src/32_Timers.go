package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2) // 产生一个计时器，2秒后向该计时器的channel中写入Time对象

	<-timer1.C // 阻塞，直到2秒后timer1中的C有数据，同时读取并且抛弃计时器timer1中channel的Time对象
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C // 同10行代码
		fmt.Println("Timer 2 expired")
	}()

	time.Sleep(time.Second * 2) // 注释该行，可以发生未超时的stop操作
	stop2 := timer2.Stop()      // 停止计时，如果计时器未过期，则返回true，否则返回false
	if stop2 {
		fmt.Println("Timer 2 stopped")
	} else {
		fmt.Println("Timer has expired")
	}

}
