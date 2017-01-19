package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"sync"
)

/*
为什么需要原子操作：
当整数被重新赋值时，例如：a += 10，该过程分为三个步骤：
1.加载变量a到寄存器；
2.给寄存器中的变量+10；
3.将寄存器中的变量写入内存中；

当多个goroutine对同一个变量进行赋值时，则需要锁或者原子操作对同一个变量进行处理，
以保证变量会得到理想状态的值
*/

func main() {

	var ops uint64 = 0
	done := make(chan bool)

	for i := 0; i < 50; i++ {					// 启动50个goroutine
		go func() {
			for {
				select {
				case <-done:					// 当done被关闭时，退出ops的+1操作
					return
				default:
					atomic.AddUint64(&ops, 1)		// 循环对ops进行原子+1操作
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	time.Sleep(time.Second)
	close(done)						// 关闭done，用于通知所有16行启动的goroutine退出针对ops的原子+1操作

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)



	// 通过WaitGroup确保所有启动的goroutine执行完毕, 主要用到:Add(), Done(), Wait()这三个函数进行控制
	var wg sync.WaitGroup
	de := make(chan bool)

	for i := 0; i < 50; i++ {
		wg.Add(1)								// 每启动一个goroutine，Add()会将wg的计数器+1
		go func() {
			defer wg.Done()						// Done()操作将wg的计数器-1, defer为延迟操作，后面会降到
			for {
				select {
				case <-de:
					return
				default:
					atomic.AddUint64(&ops, 1)
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	time.Sleep(time.Second)
	close(de)									// 和29行的作用相同

	wg.Wait()									// 等待所有42行启动的goroutine完成
	fmt.Println("ops:", ops)
}
