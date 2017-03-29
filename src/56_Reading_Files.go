package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 错误处理
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	const fileTmp = "/Users/liqiang/GoExample/src/dat"

	// ioutil.ReadFile()将文件全部读如到内存，如果文件过大，你懂的
	dat, err := ioutil.ReadFile(fileTmp) // 读取文件中的内容,返回字节切片
	check(err)
	fmt.Println(string(dat))

<<<<<<< HEAD
	f, err := os.Open(fileTmp)		// 在20行的ReadFile()函数内部就调用了os.open()函, 该函数以只读模式打开文件
=======
	f, err := os.Open(fileTmp) // 其实20行的ReadFile()函数内部就调用了os.open()函数
>>>>>>> b6deb2bb9e17fc2ee68295cf089a3159d8895196
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // 流式读取文件
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 第一个参数表示: 移动6个字节,第二个参数0: 从文件起始位置; 1:从文件当前位置; 2:从文件末尾
	o2, err := f.Seek(6, 0) // 返回位置的字节下标
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2) // 从f中读取最少2个字节到b3中
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) // 返回5个字节的切片
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // 关闭文件! 文件在这里关闭并不安全，如果发生panic，打开的文件则不会关闭
	// 应该在29行通过 defer的方式关闭
}
