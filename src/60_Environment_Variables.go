package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {

	os.Setenv("FOO", "1")		// 设置环境变量
	fmt.Println("Foo:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {		// Environ():获取所有环境变量
		pair := strings.Split(e, "=")
		fmt.Println("key:", pair[0], "\tvalue:", pair[1:])
	}

}
