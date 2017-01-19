package main

import (
	"fmt"
	"time"
)

// select:多路复用, 只针对channel的特殊操作语句, 并且case后只能跟针对channel的操作语句;
// select会随机对case后的channel进行操作，哪个成功则执行对应case块的逻辑, 但会顺序针对case后的语句进行求值操作

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 特别注意：如果每个case后的channel都能成功执行，则会随机选取case逻辑块，而不是选择第一个, 这点和switch完全不同
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:						// 如果成功从c1中读取成功，则打印c1中的数据, 执行完case后的逻辑块后，自动跳出select{}逻辑块
			fmt.Println("received", msg1)
		case msg2 := <-c2:						// 如果成功从c2中读取成功，则打印c2中的数据
			fmt.Println("received", msg2)
		default:								// 如果上述case都不成功，则执行default后的逻辑
			fmt.Println("default case")
		}
		time.Sleep(time.Second * 1)
	}

	c := make(chan bool, 1)
	for i := 0; i < 3; i++ {
		select {
		case c <- true:							// 第一次执行select时，c <- true、c <- false都可以执行成功，但并不是一定执行 c <- true，而是随机执行 c <- true或者 c <- false
			fmt.Println("true")
			break								// 此处的break并不能跳出for{}循环块，如果想跳出for{}块，则需要在for{}逻辑块的开头加tag，然后通过：break tag的方式跳出循环
		case c <- false:
			fmt.Println("false")
		default:
			fmt.Println("default")
		}
	}

}
