package main

import (
	"fmt"
	"os"
)


func main() {

	defer fmt.Println("Over!")

	os.Exit(3)		// os.Exit()是唯一个在推出后不执行defer的函数

}