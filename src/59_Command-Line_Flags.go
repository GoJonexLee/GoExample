package main

import (
	"flag"
	"fmt"
)

// 命令行标签参数
// String()、Int()、Bool() 和 StringVar()、IntVar()、BoolVar()的区别：
// 前三个函数返回的是指针，后三个函数返回的是值，但后三个函数需要将变量的地址传递到方法中

// 感受下面中方式的区别，第二个启动命令将会使用标签的默认值。
// go run 59_Command-Line_Flag.go -word hehe -numb 1 -fork true
// go run 59_Command-Line_Flag.go
func main() {

	// 参数1：名称， 参数2：默认值， 参数3：作用
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string		// flag().StringVar()需要单独声明变量，但flag.String()不用
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()		// 这条语句必须，目的是让系统的参数重新解析

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
