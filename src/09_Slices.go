package main

import "fmt"

func main() {

	s1 := make([]string, 3)					// 初始化长度为3，容量为3的字符串切片，切片中的3个值已经被初始化为""
	s2 := make([]string, 0, 3)				// 初始化长度为0，容量为3的字符串切片，切片为空
	s3 := []string{"x", "y", "z"}			// 切片第二种初始化方法，长度为3，容量为3

	fmt.Println(s1, len(s1), cap(s1))		// 注意s1和s2的不同
	fmt.Println(s2, len(s2), cap(s2))		// len()返回切片的长度，cap()返回切片的容量
	fmt.Println(s3, len(s3), cap(s3))

	s1[0], s1[1], s1[2] = "a", "b", "c"		// 通过下标更改切片中的值
	s1 = append(s1, "e")					// 追加"e"到s1中
	s1 = append(s1, "f", "g")				// 追加"f", "g"到s1中
	s1 = append(s1, s3...)					// 将s3中的元素追加到s1中, s3后面的三个'.’可以理解为将s3打碎
	fmt.Println(s1)

	// s2[0], s2[1], s2[2] = "e", "f", "g"	// 引发panic，因为s2的长度为0，因此不能通过下标更改切片中的值
	s2 = append(s2, "e", "f", "g")			// 只能通过append()函数向s2中追加字符串
	fmt.Println(s2)

	c := make([]string, len(s1))
	copy(c, s1)								// copy():将s1拷贝到c中
	fmt.Println(c)

	f := s1[2:5]							// f为s1的子切片, 从s1的下标2开始到下标5之前，对f的修改会影响s1的值
	fmt.Println("sl1:", f)

	t := s1[2:]
	fmt.Println("sl2:", t)					// 从s1的下标开始，到s1的末尾

	twoD := make([][]int, 3)				// 3行二维切片，需要手动初始化
	fmt.Println(twoD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)		// 每一纬度的长度不相同，切片中的原始值为0
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j				// 更新切片中的值 
		}
	}
	fmt.Println("2d:", twoD)

}
