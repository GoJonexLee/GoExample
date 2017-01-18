package main

// 常量的定义和使用

import (
	"fmt"
	"math"
)

/*
多包导入的时候，也可以分开, 但是不建议这么干，如下所示：

import "fmt"
import "math"

*/

const s string = "LuoJiLab"		// 定义常量s
const t = "LuoJi"				// go会根据"LuoJi"推断出t的类型

const (							// 可以将常量组织在const()中
	a = 1
	b = 2
	c = 2.0
	d = "double"
)

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))		// 整数之间可以直接通过type()的方式进行类型转换

	fmt.Println(math.Sin(n))
}
