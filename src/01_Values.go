package main

import "fmt"

func main() {

	fmt.Println("go" + "lang") // 字符串拼接，这样干会有性能损耗，因为字符串不能改变，后面会有针对string的章节

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && true)
	fmt.Println(true || false)

	fmt.Println(!true)
}
