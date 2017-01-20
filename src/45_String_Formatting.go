package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)			// 只输出结构体的字段值
	fmt.Printf("%+v\n", p)			// 输出结构体的字段和值
	fmt.Printf("%#v\n", p)			// 输出结构体的来自哪个包以及结构体的字段和值
	fmt.Printf("%T\n", p)			// 输出结构体的来源
	fmt.Printf("%p\n", &p)			// 输出结构体的内存地址

	fmt.Printf("%t\n", true)		// 输出bool型
	fmt.Printf("%d\n", 123)			// 输出十进制整数
	fmt.Printf("%b\n", 14)			// 输出二进制
	fmt.Printf("%c\n", 33)			// 输出对应字符
	fmt.Printf("%x\n", 456)			// 十六进制
	fmt.Printf("%f\n", 78.9)		// 浮点数

	fmt.Printf("%e\n", 123400000.0)		// 小e表示
	fmt.Printf("%E\n", 123400000.0)		// 大E表示

	fmt.Printf("%s\n", "\"string\"")	// 转义输出
	fmt.Printf("%q\n", "\"string\"")	// 非转义输出

	fmt.Printf("%x\n", "hex this")		// 将字符串以16进制的字节形式输出

	fmt.Printf("|%6d|%6d|\n", 12, 345)			// 宽度为6，右对其
	fmt.Printf("|%6.2f|%-6.2f|\n", 1.2, 3.45)	// 宽度为6，保留2位有效数字，-为左对齐

	s := fmt.Sprintf("a %s", "string")			// 以字符串形式返回
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")	// 输出重定向到os.Stderr

}
