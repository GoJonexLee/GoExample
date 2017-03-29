package main

import (
	"os"
	"fmt"
)

// 从命令行获取参数
func main() {

	argsWithProg := os.Args			// 获取命令行所有参数，并以切片形式返回
	argsWithoutProg := os.Args[1:]		// 第一个参数为当前文件编译后的二进制文件名称，其余为从命令行输入的参数

	args := os.Args[3]			// 获取第三个参数

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(args)

}
