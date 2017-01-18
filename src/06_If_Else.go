package main

import "fmt"

func main() {

	if 7 % 2 == 0 {					// if后面的条件语句不能用()括起来
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8 % 4 == 0 {
		fmt.Println("8 is devisible by 4")
	}

	if num := 9; num < 0 {					// 可以在if后面进行变量的初始化, 但必须用';'和条件语句分开
		fmt.Println(num, " is negative")	// num的作用域包括之后的else if 和 else, 但不能在if块之外使用
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}

}
