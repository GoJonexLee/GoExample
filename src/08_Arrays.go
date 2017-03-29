package main

import "fmt"

func main() {

	var a [5]int // 定义一个长度为5的整数数组,数组中的每个数据默认被初始化为0
	fmt.Println("emp:", a)

	a[4] = 100 // 通过下标修改数组中的元素
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	b := [5]int{1, 2, 3, 4, 5} // 定义以及初始化数组
	fmt.Println("dcl:", b)

	var twoD [2][3]int // 定义一个2x3的二维数组，数组中的每个数据已经被初始化为0
	fmt.Println(twoD)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j // 对二维数组中的数据重新赋值
		}
	}
	fmt.Println(twoD)
}
