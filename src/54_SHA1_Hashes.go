package main

import (
	"fmt"
	"crypto/sha1"
)


func main() {

	s := "sha1 this string"	// 需要哈希处理的字符串

	h := sha1.New()		// 返回哈希结构体

	h.Write([]byte(s))	// 写入需要hash处理的字节切片

	bs := h.Sum(nil)	// 计算哈希值

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
