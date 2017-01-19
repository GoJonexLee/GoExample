package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

// 这个例子的逻辑有点儿复杂，建议画张图看看
func main() {

	var readOps uint64 = 0			// 总感觉这种声明变量多此一举, 根本不用写后面的' = 0 '
	var writeOps uint64 = 0

	reads := make(chan *readOp)		// 无缓冲channel中存放指针类型
	writes := make(chan *writeOp)

	// 读goroutine，从reads、writes中读取数据
	go func() {

		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:				// 如果读成功，则将state[key]的值放入resp中
				read.resp <- state[read.key]	// 阻塞，因为resp是无缓冲channel
			case write := <-writes:				// 如果成功，则将write中的val写入state中
				state[write.key] = write.val
				write.resp <- true				// 阻塞，直到有人从write.resp中取走数据, 72行代码对应
			}
		}

	}()

	// 启动100个goroutine，往reads中写入数据, 同时丢弃read中channel中的数据（55行）
	for r := 0; r < 100; r++ {

		go func() {
			read := &readOp{
				key: rand.Intn(5),
				resp: make(chan int),		// 不要忘了以','结尾
			}

			reads <- read					// 阻塞
			<-read.resp						// 和36行代码对应

			atomic.AddUint64(&readOps, 1)	// 记录从reads中读取的次数
			time.Sleep(time.Millisecond)
		}()

	}

	// 启动10个goroutine，往writes中写数据, 同时丢弃write中channel的数据（75行）
	for w := 0; w < 10; w++ {

		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}

				writes <- write			// 阻塞，直到writes中的数据被取走
				<-write.resp			// 和39行代码对应

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)		// 原子读取读操作的次数
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)	// 原子读取写操作的次数
	fmt.Println("writeOps:", writeOpsFinal)

}

