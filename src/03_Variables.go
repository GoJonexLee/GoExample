package main

// 初始化变量的几种方式
import "fmt"

/*
ps：go中不存在未初始化的变量
	未使用的局部变量编译时会报错
	未使用的全局变量会通过编译
*/

func main() {
	var a string = "initial"		// 定义字符串变量:a, 同时初始化; 字符串在go中很特殊，后面会单独讲解
	fmt.Println(a)

	var x, b int = 1, 2				// 定义int类型的整数变量:a, b
	fmt.Println(x, b)

	var c, d = 3, 4					// 和上一步相同，go会根据1，2推断出c，d的类型
	fmt.Println(c, d)

	e, f := 5, 6					// 该初始化方式只能用于函数内部; e,f分别为整数5，6
	fmt.Println(e, f)

	l := "short"					// 同上，l的值为字符串"short"
	fmt.Println(l)

	var s string					// go会将s初始化为空字符串
	fmt.Println(s)

	var m, n int					// s, t会被自动初始化为0
	fmt.Println(m, n)
}
