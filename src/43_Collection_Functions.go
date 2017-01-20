package main

import (
	"strings"			// 专门用于处理字符串的包包
	"fmt"
)

// 在字符串切片中查找指定的字符串，如果找到则返回下标，否则返回-1
func Index(vs []string, t string) int {

	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 判断字符串切片中是否包含目标字符串
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 遍历字符串切片，判断是否有满足func(string) bool签名函数的字符串
func Any(vs []string, f func(string) bool) bool {

	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 遍历字符串切片，判断是否所有字符串都满足func(string) bool签名函数
func All(vs []string, f func(string) bool) bool {

	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//根据func(string) bool签名的函数过滤字符串切片
func Filter(vs []string, f func(string) bool) []string {

	res := []string{}				// 和 res := make([]string, 0) 效果相同；后面有优化
	for _, v := range vs {
		if f(v) {
			res = append(res, v)	// 将满足函数f()函数的字符串切片追加到结果切片中
		}
	}
	return res

/* 不用49行代码那个额外的切片开销
	k := 0
	for i, v := range vs {
		if f(v) {
			vs[k] = vs[i]
		}
	}

	return vs[:k]
*/
}

// 针对切片中的每一个字符串进行函数操作
func Map(vs []string, f func(string) string) []string {

	res := make([]string, len(vs))
	for i, v := range vs {
		res[i] = f(v)
	}
	return res
}


func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")			// 判断是否以"P"开头
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.Contains(v, "e")				// 判断字符串v是否包含"e"
	}))

	fmt.Println(Map(strs, strings.ToUpper))			// 将字符串转化为大写
}

