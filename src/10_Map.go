package main

import "fmt"

// key-value数据结构，动态增常型

func main() {
	m1 := make(map[string]int)     // 定义方式1: make()，key为string类型，value为int类型，默认长度为0
	m2 := map[string]int{}         // 定义方式2，类型和m1相同，默认长度为0
	m3 := make(map[string]int, 10) // 定义长度为10的map

	m1["k1"] = 7 // 多行赋值
	m1["k2"] = 13
	fmt.Println(m1, len(m1))
	fmt.Println(m1["k3"]) // 如果获取map中不存在的key对应的value，则返回value类型对应的零值

	m2["k1"], m2["k2"] = 15, 16 // 单行赋值
	fmt.Println(m2, len(m2))

	m3["k"], m3["kk"] = 20, 30
	fmt.Println(m3, len(m3))

	delete(m3, "k")   // delete()内置函数删除map中对应的key
	_, ok1 := m3["k"] // 当ok为false时，表明map中不存在key，否则存在
	fmt.Println(ok1)

	_, ok2 := m3["kk"]
	fmt.Println(ok2)

	mp := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	for k, v := range mp { // map只能通过range的方式遍历，并且每次遍历的顺序都不相同
		if k == 1 {
			mp[k*10] = v * 10 // 新添加的k-v，不能假设是否会在之后被遍历到
		} else if k == 5 {
			delete(mp, 6) // 遍历的过程中可以安全删除元素,并且不能假设是否被遍历到
		}
		fmt.Println(k, " ", v)
	}

	fmt.Println(mp)

}
