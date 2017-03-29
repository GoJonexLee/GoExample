package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("Hello\tgo\n")
	err := ioutil.WriteFile("/Users/liqiang/GoExample/src/write.txt", d1, 0644)
	check(err)			// check()函数已经在第56的读文件中定义

	f, err := os.Create("/Users/liqiang/GoExample/src/write1.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}		// 字节切片
	n2, err := f.Write(d2)						// 将字节切片写入f文件
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")		// 返回的第一个参数为写入文件的字节数
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()									// 将将内存中的数据写入硬盘

	w := bufio.NewWriter(f)						// 带缓冲的文件写入, 缓冲区的大小为4096字节
	n4, err := w.WriteString("buffered\n")		// 写入字符串
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()									// 如果缓冲区满，则会自动刷新，为了确保最后一次写入文件，必须手动刷新
}
