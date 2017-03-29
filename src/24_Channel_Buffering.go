package main

import "fmt"

// 有缓冲的channel

func main() {
	message := make(chan string, 2) // 初始化缓冲容量2的channel

	message <- "buffered" // 想channel中放入两个对象
	message <- "channel"

	// message <- "again"				// 千万不能这么干，因为channel已经满了, 因此这一步会阻塞直到channel中有空余位置

	fmt.Println(<-message)
	fmt.Println(<-message)

}
