package main

import (
	"strconv"		// 经常使用该包进行字符串和数字之间的转换
	. "fmt"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)		// 将字符串1.234转换为64位的小数
	Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)		// 将字符串123，转换为64位的整数
	Println(i)									// 第二个参数0表示根据字符串判断几进制表达,这里是10进制

	u, _ := strconv.ParseUint("789", 0, 64)		// 无符号10进制64位整数
	Println(u)

	k, _ := strconv.Atoi("123")		// string转int
	Println(k)

	_, e := strconv.Atoi("wat")		// 如果不能转为int型，则e != nil
	Println(e)

	Println(strconv.Itoa(123))		// 整数转为字符串
	Println(strconv.ParseBool("true"))	// 字符串转bool类型

}
