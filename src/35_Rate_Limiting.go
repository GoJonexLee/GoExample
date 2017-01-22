package main

import (
	"fmt"
	"time"
)

// 通过心跳器进行遍历控制，关键在30行和43行代码
func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests) // 关闭后的requests，除了不能继续写入数据外，已存在requests中的数据可以继续读取

	limiter := time.Tick(time.Millisecond * 1000) // 使用time包中默认的心跳器
	for req := range requests {
		<-limiter                                 // 阻塞1000毫秒
		fmt.Println("request 1", req, time.Now()) // 打印requests中的数据以及当前时间
	}

	burstyLimiter := make(chan time.Time, 3) // 容量为3的Time缓冲channel, 主要用户心跳器的数据写入
	for i := 0; i < 3; i++ {                 // 先写入三个数据，因此42行遍历burstyRequests的前三个数据时不需要等待200毫秒
		burstyLimiter <- time.Now()
	}

	go func() { // 每隔2000毫秒向burstyLimiter中写入新的Time对象
		for t := range time.Tick(time.Millisecond * 2000) {
			burstyLimiter <- t // 阻塞写入burstyLimiter，因此42行的burstyRequests的第四个数据也不用等待
		}
	}()

	burstyRequests := make(chan int, 10) // burstyRequests作为被遍历的channel使用，主要用来消费数据
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests { // burstyRequests中存在5个数据
		<-burstyLimiter                           // burstyLimiter的作用在于：每隔200毫秒遍历burstyRequests中的数据
		fmt.Println("request 2", req, time.Now()) // burstyLimiter中的前4个数据不需要等待
	}

}
