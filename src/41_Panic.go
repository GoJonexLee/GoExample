package main

import "os"

// panic: go中独有的错误处理方式，会使程序崩溃，通常将defer和recover()结合使用;
// go将异常处理机制分为error 和 panic, 在所有会发生错误的地方几乎都会有error返回,
// panic不太常用，除非你确定程序需要被崩溃，当然defer+recover()可以从panic中恢复。
func main() {

	panic("a proble")					// 程序运行到此处就崩溃了，后面的程序不会执行

	_, err := os.Create("/tmp/file")	// 创建文件，因为10行引发panic，因此这行不会被执行
	if err != nil {
		panic(err)
	}

}
