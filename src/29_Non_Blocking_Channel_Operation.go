package main

import "fmt"

// 无阻塞channel操作，说白了就是在select{}中加入default逻辑块
func main() {

	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message: // 因为message是无缓冲channel，因此这个case一直会被阻塞
		fmt.Println("received message", msg)
	default: // 当所有case都不能立刻执行时，就会执行default逻辑块
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg: // 由于12行并未从message读取成功，因此12行的读操作已经取消，因此本行的写操作也不能执行成功，因为没有接收者
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message: // <-message阻塞
		fmt.Println("received message", msg)
	case sig := <-signals: // <-signals阻塞
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
