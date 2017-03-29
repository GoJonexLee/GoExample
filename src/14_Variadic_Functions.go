package main

import "fmt"

// 函数的变长参数, 在函数体内，把nums当作切片处理就行
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1)

	sum(1, 2)

	sum(1, 2, 3)

	nums := []int{4, 5, 6, 7} // 切片必须通过'...'打碎处理后才能传入sum函数中
	sum(nums...)

	nms := [...]int{1, 2, 3, 4, 5} // nms是数组，因此不能直接传入，必须使用'...'打碎传入
	sum(nms[:]...)                 // 切片传入

}
