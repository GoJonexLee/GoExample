package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 按行读取数据
func main() {

	scanner := bufio.NewScanner(os.Stdin) // 通过scanner按行读取数据。os.Stdin已经实现了io.Reader接口

	for scanner.Scan() {	// 判断是否有下一行数据，如果有，则读到token，同时返回true。scanner.Scan()每次读取的最大字节数为64M

		ucl := strings.ToUpper(scanner.Text()) // ToUpper(),小写转大写; scanner.Text():将token以字符串形式返回
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
