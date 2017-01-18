package main

import "fmt"

// 闭包：返回的函数包含原函数的上下文环境
// 例如: intSeq()函数返回的函数包含了intSeq中的i变量

func intSeq() func() int {
	i := 0
	return func() int {
		i++						// 和i+=1相同效果
		return i
	}
}

func main() {

	nextInt := intSeq()			// nextInt函数是由intSeq()函数创建，因此会捕获intSeq中的i变量，于是在每次对nextInt调用时，都会对同一个i变量进行+1操作

	fmt.Println(nextInt())		// 1
	fmt.Println(nextInt())		// 2
	fmt.Println(nextInt())		// 3

	newInt := intSeq()			// 重新返回的函数值
	fmt.Println(newInt())		// 1

}
