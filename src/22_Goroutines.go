package main

import "fmt"

// 并发模型goroutine，通过关键字go启动并行逻辑

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") // 串行执行f函

	go f("goroutine") // 启动一个goroutine，和main()函数并行执行

	go func(msg string) { // 启动一个匿名goroutine, 和main()以及上步中的goroutine并行
		fmt.Println(msg)
	}("going") // 匿名goroutine的调用绝对不能忘记（）以及括号里面的参数列表

	var input string
	fmt.Scanln(&input) // 阻塞式等待标准输入，回车后输入结束
	fmt.Println("done")

}
