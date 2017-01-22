package main

import "fmt"

func main() {

	nums := []int{2, 3, 4} // 整数切片, 和数组不同的是：数组定长、切片变长
	sum := 0

	for _, num := range nums { // 用'_'忽略下标，因为切片是引用类型，因此拷贝代价很小
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums { // 使用下标
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // 字典:key-value数据类型
	for k, v := range kvs {                               // map只能使用range遍历，并且每次的遍历顺序都不相同, k为键，v为值
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // 省略value的遍历方式
		fmt.Println("key:", k)
	}

	for _, v := range kvs { // 省略key的遍历方式
		fmt.Println("value:", v)
	}
}
