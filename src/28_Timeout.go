package main

import (
	"fmt"
	"time"
)

// 通过select{}以及time.After()进行超时控制

func main() {

	// 超时: c1中2秒后才会有数据，time.After()在1秒后就会有数据
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2) // 睡2秒的觉
		c1 <- "result 1"            // 向通道c1中放入字符串
	}()

	select { // 阻塞，直到有case可以执行，如果多个case都可以执行，则随机选取一个
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1): // time.After(): 先产生一个chan Time的缓冲channel，1秒后向该channel中放入Time对象，最后将该channel当成只读channel返回
		fmt.Println("timeout 1")
	}

	// 未超时: 因为c2中2秒后就会有数据，而time.After()在3秒后才会产生数据
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
