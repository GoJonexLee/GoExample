package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?@#$()'-=&~"		// 编解码处理数据

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))		// 用标准的方式进行编码处理
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)		// 标准解码，忽略error
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))		// URLEncoding: 针对urls和文件名称的编解码器
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(uDec)

}
