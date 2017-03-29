package main

import (
	"fmt"
	"os"
)

// defer: 延迟操作，当程序执行完或者返回之前的操作。
// 如果方法中存在多个defer语句，则按照优先定义最后调用的顺序执行。
func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // 在15行writeFile()执行完毕后才会执行closeFile()函数

	writeFile(f)

}

func createFile(p string) *os.File {

	fmt.Println("creating")

	f, err := os.Create(p) // Create(): 创建文件, 如果创建失败，则会返回err
	if err != nil {
		panic(err) // 主动引发恐慌，让程序崩溃，标准输出会打印err原因
	}

	return f
}

func writeFile(f *os.File) {

	fmt.Println("writing")
	fmt.Fprintln(f, "data") // Fprintln(): 将"data"输出到f中
}

func closeFile(f *os.File) {

	fmt.Println("closing")
	f.Close() // Close(): 关闭文件
}
