package main

import "fmt"

// 通常会通过关闭channel操作进行广播；
// 使用完毕的channel记得要尽心关闭操作，否则会造成channel泄漏，即垃圾回收机制不会对未关闭的channel进行回收
func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for { // 一般不建议通过这种方式遍历channel，因为逻辑比较复杂；建议通过range遍历channel，例如下一个goroutine
			j, ok := <-jobs // ok的作用：当jobs被关闭后，仍然能从其中读取值，但读取的是channel中类型的零值，因此需要ok进行判断
			if ok {         // 当ok为false时，表明channel已经被关闭，true表明jobs未被关闭
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // 当jobs被关闭时，给done发送数据，表明jobs中的数据已经被接收完
				return
			}
		}
	}()

	go func() {
		for job := range jobs { // 可以使用range的方式遍历channel，当jobs被关闭后，该循环体会自动退出, 避免了使用ok返回值判断
			fmt.Println("range receive", job)
		}
		fmt.Println("finish receive")
	}()

	for j := 1; j <= 3; j++ { // 向jobs中写入数据
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 通常在发送数据方关闭channel，因为向已经关闭的channel中发送数据会引发panic，而从已经关闭的channel中接收数据则不会引发panic
	fmt.Println("sent all jobs")

	<-done // 阻塞，直到done中有数据
}
