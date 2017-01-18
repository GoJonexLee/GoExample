package main

import "fmt"

func zeroval(val int) {
	val = 20
}

func zeroptr(iptr *int) {
	*iptr = 10
}

func main() {

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)						// 因此函数的参数传递是按照“拷贝传递”，也就是“传值”，因此i是被拷贝了一份传入函数zeroval中
	fmt.Println("zeroval:", i)		// i的值并未被改变

	zeroptr(&i)						// 传入i的地址，因此i的值已经在函数体内被改变
	fmt.Println("zeroptr:", i)

	fmt.Println("opinter:", &i)		// 输出变量i的地址
}
