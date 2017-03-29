package main

/*
导入fmt包，因为Println()函数来源于fmt

ps：如果想直接使用Println()函数，而不是fmt.Println()，则在导入fmt包时使用：. "fmt"
	如果想使用别名调用Println()函数，则导入fmt的方式为：tmp "fmt", tmp为fmt的别名，
	别名调用方式为：tmp.Println()
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Hello PHPer!")
}

// 运行命令：go run 01_Hello Word.go
// 或者：go build -o hello 01_Hello word.go && ./hello
// 参数-o：后面为编译后的项目名称，详情请参照go build --help
