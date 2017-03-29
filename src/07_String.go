package main

import (
	"fmt"
	"reflect"
)

// go中的字符串有多让人蛋疼的地方，需要特别注意;
// 遍历字符串时，底层实际便利的是字符串对应的字节数组;
// byte和uint8采用相同的底层结构
// rune和int32采用相同的底层结构
func main() {

	// 中文在字符串中占用3个字节
	s1 := "abc中国"
	fmt.Println(s1)

	// 1. 字符串不能修改,如下操作是会让程序崩溃的
	//s1[0] = "e"	// 将第一个字符串修改为e

	// 2. 便利字符串的方式不同，结果也不同
	for i := 0; i < len(s1); i++ {		// 单字节遍历，便利到中文时，肯定不会正常打印中文
		// 需要将字节通过string()转换为字符串
		fmt.Println(string(s1[i]), "\t", reflect.TypeOf(s1[i]))	// 字符串、类型
	}
	fmt.Println()

	// 3. 正常遍历包含中文的字符串, data实际是rune类型，rune和int32是相同的底层结构
	for index, data := range s1 {
		fmt.Println(index, string(data), reflect.TypeOf(data))	// 下标、字符串、类型
	}


	// 4. 如果需要修改字符串的处理方式
	s2 := "abcdeft1234"

	b1 := []byte(s2)	// 通过[]byte()先把字符串转换为字节数组
	b1[0] = 'x'		// 修改下表为0的字节，字节用单引号

	fmt.Println(string(b1))	// 将字节切片转换为字符串输出

}
