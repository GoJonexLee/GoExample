package main

import "fmt"

// 建议通过range的方式遍历channel，否则需要对关闭的channel进行特殊判断
func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 当channel被关闭后，已经存在channel中的数据仍然可以被读取

	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("close data:", <-queue) // 关闭的channel仍然可以读取数据，但读取到的是channel中数据类型的零值，即""

	dt := make(chan int, 1)
	dt <- 10
	close(dt)

	fmt.Println(<-dt) // 10可以正常读取
	fmt.Println(<-dt) // 因为dt应被关闭，因此读取到的是int类型的零值0

	k, ok := <-dt // 通过ok可以判断channel是否已经被关闭, true:未关闭, false:已关闭
	fmt.Println("value:", k, " close:", ok)

}
