package main

import (
	"fmt"
	"time"
)

// 通过无缓冲channel进行同步
func worker(done chan bool) {

	fmt.Println("working...")
	time.Sleep(time.Second) // time.Sleep()：睡眠函数，参数为睡眠时间；time.Second：1秒
	fmt.Println("done")

	done <- true // 当该函数运行完时，通过向done中放入bool值，表明该函数已经完成
	fmt.Println("go done")
}

func main() {
	done := make(chan bool, 1)

	go worker(done) // 开启一个异步的goroutine

	<-done // main()函数阻塞，直到worker函数将bool值放入channel中
	fmt.Println("main done")
	// 注意：main()函数和go 函数到底谁先执行完，谁后执行完
	// 答案：谁先阻塞，谁后执行完
}
