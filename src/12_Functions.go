package main

import "fmt"

func plus(a int, b int) int { // func定义函数,当函数名首字母大写时，函数可以导出，否则不能导出
	return a + b // plus不能导出，只能在包内使用
}

func Plus(a, b, c int) int { // Plus函数可导出，a、b、c均为int类型
	return a + b + c
}

func main() {

	fmt.Println(plus(1, 2))    // plus函数调用
	fmt.Println(Plus(1, 2, 3)) // Plus函数调用

	var f func(int, int) int // 定义f为需要两个整数的入参，返回一个整数的函数
	f = plus                 // 将plus赋给f
	fmt.Println(f(3, 4))     // 调用f函数

	var f1 func(int, int) int // f1函数内部需要递归，因此必须先定义
	f1 = func(a, b int) int { // 此处的'='绝对不能写为':='
		if a == 0 || b == 0 {
			return a + b
		}

		return f1(a-1, b-1)
	}
	fmt.Println(f1(2, 3))

}
