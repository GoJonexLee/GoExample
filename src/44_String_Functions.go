package main

import (
	"fmt"
	. "strings" // '.'符号可以在调用包函数的时候省略包名
)

var p = fmt.Println // 函数赋值

func main() {

	p("Contains:", Contains("test", "es")) // Contain()函数来自strings包，因为在导入包的时候加了'.'符号，因此调用不需要再写包名,例如: strings.Contains()

	p("Count:", Count("test", "t")) // 统计子字符串出现的此时

	p("HasPrefix:", HasPrefix("test", "te")) // 是否以前缀"te"开头

	p("HasSuffix:", HasSuffix("test", "st")) // 是否以后缀"st"结尾

	p("Index:", Index("test", "e")) // 查找下标

	p("Join:", Join([]string{"a", "b", "c"}, "-")) // 字符串拼接

	p("Repeat:", Repeat("a", 5)) // 重复字符串"a"5次

	p("Replace:", Replace("foo", "o", "0", -1)) // 将o替换为0，-1表示全部替换，如果把-1换做正整数，则表明执行几次替换操作

	p("Split:", Split("a-b-c-d", "-")) // 切割，返回字符串切片

	p("ToLower:", ToLower("TEST")) // 小写转换

	p("ToUpper:", ToUpper("test")) // 大写转换

	p() // 换行

	p("Len:", len("hello")) // 输出字符串长度

	p("Char:", "hello"[1]) // 输出字符

}
