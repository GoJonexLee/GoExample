package main

import "fmt"

// 函数的多返回值
func vals() (int, int) {
	return 3, 4
}

// a、b在返回列表中被初始化为0
func vals1() (a, b int) {
	return
}

// a、b在函数内被重新赋值1,2
func vals2() (a, b int) {
	return 1, 2
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	c, d := vals1()
	fmt.Println(c, d)

	e, f := vals2()
	fmt.Println(e, f)

	_, h := vals2()		// 如果只需要一个返回值，则另外一个可以使用'_'过滤，但'_'符号不能省略
	fmt.Println(h)
}
