package main

import (
	"fmt"
	"sort"
)

// ByLength和字符串切片具有相同的底层结构,除此之外,没有任何相同点
type ByLength []string

func (b ByLength) Len() int {
	return len(b) // 返回切片的总长度
}

func (b ByLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i] // 交换下标i, j中的元素
}

func (b ByLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j]) // 将短字符放在长字符前面
}

// sort.Sort()函数可以对所有实现Interface接口的函数进行排序操作, Interface一共包含3个函数:
// 1.Len() int: 返回需要排序的容器长度;
// 2.Swap(int, int): 交换容器中不同下标的元素;
// 3.Less(int, int) bool: 返回排序规则，例如20行代码，将短字符排在长字符前面.
func main() {

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

}
