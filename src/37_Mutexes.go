package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作只能针对变量进行操作
// 其他数据类型需要通过mutex进行互斥操作
// ps:下面的例子针对整体map进行加解锁处理，粒度有点儿大，可以通过针对单个key进行加解锁优化
func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}					// 互斥量，保证多个goroutine互斥访问

	var readOps uint64 = 0
	var writeOps uint64 = 0

	for r := 0; r < 100; r++ {					// 启动100个goroutine

		go func() {
			total := 0
			for {
				key := rand.Intn(5)				// 在[0,1,2,3,4]中随机选取一个数
				mutex.Lock()					// 加锁; 此处可以使用读写锁进行优化
				total += state[key]
				mutex.Unlock()					// 释放锁
				atomic.AddUint64(&readOps, 1)	// 记录读操作的次数

				time.Sleep(time.Millisecond)
			}
		}()

	}

	for w := 0; w < 10; w++ {					// 启动10个goroutine

		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()					// 加锁
				state[key] = val				// 设置key-value，有可能重复设置
				mutex.Unlock()					// 解锁
				atomic.AddUint64(&writeOps, 1)	// 记录写操作的次数

				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}

