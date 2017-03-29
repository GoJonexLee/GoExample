package main

import (
	"fmt"
	"time"
)

// 消费jobs中的数据，向results中产生数据; jobs为只读channel，results为只写channel
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs { // 从jobs中读取数据
		fmt.Println("worker", id, "started job", j)  // 打印goroutine的id以及jobs中的数据
		time.Sleep(time.Second)                      // 休息1秒
		fmt.Println("worker", id, "finished job", j) // 打印
		results <- j * 2                             // 将jobs中的数据翻倍后放入results中
	}

}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ { // 启动3个工作goroutine
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ { // 向jobs中写入5个数据
		jobs <- j
	}
	close(jobs) // 关闭jobs通道，以便通知工作routine停止工作

	for a := 1; a <= 5; a++ { // 消费results中的数据
		fmt.Println(<-results)
	}
}
