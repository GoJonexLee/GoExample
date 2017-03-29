package main

import (
	"fmt"
	"sort"
)

// 通过sort包对字符串切片以及整数切片进行排序
func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints) // 判断是否有序
	fmt.Println("Sorted:", s)

}
