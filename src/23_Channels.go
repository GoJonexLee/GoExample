package main

import "fmt"

// channel, 用于多个goroutine之间进行通信;
// channel分为有缓冲和无缓冲，有缓冲的channel就像一个循环队列，可以进行异步通信;
// 无缓冲的channel只能进行同步通信;

func main() {

	// 创建无缓冲channel，channel只能存放string类型的对象
	message := make(chan string)	// channel只能通过make()函数创建, 除此之外，其他任何对象都有两种创建方式:object{}, new(object)
									// 如果省略make()函数中的第二个参数，则为无缓冲，否则为有缓冲
	go func() {
		message <- "ping"			// 通过'<-'向无缓冲channel中放入字符串"ping", 该goroutin只能等到其他groutine从message中读取数据后才能推出
	}()								// 千万不能忘记此处的括号:()

	msg := <-message				// 读取message中的数据，注意读和取操作时'<-'和channel的位置
	fmt.Println(msg)

}
