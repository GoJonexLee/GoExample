package main

import "fmt"

// 函数的递归

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(10))

	// 如果在函数体内定义递归函数，则需要特别注意
	var fact1 func (n int) int		// 函数声明
	fact1 = func(n int) int {		// 函数定义
		if n == 0 {
			return 1
		}

		return n * fact1(n-1)		// 递归调用
	}

	fmt.Println(fact1(10))

	// 函数体内的递归函数定义绝对不能写成下列方式；函数的声明和定义必须分开
	/*
	fact1 := func(n int) int {
		if n == 0 {
			return 1
		}

		return n * fact1(n-1)
	}
	*/
}
