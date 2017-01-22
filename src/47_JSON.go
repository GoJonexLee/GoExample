package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"` // page、fruits: json序列化后的key
	Fruits []string `json:"fruits"`
}

func main() {
	// json.Marshal(): 用于对象值的json序列化
	bolB, _ := json.Marshal(true) // 返回字节切片忽略了第二个返回值error
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1) // 字节切片可以直接通过string()转换为字符串
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34) // 序列化浮点数
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher") // 序列化字符串
	fmt.Println(string(strB))

	slcD, _ := json.Marshal([]string{"apple", "peach", "pear"}) // 序列化字符串切片
	fmt.Println(string(slcD))

	mapD, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7}) // 序列化map
	fmt.Println(string(mapD))

	// 认真感受下Response1和Response2的区别，Response2的字段后面跟有json标签
	res1D := &Response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{Page: 2, Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 反序列化
	bts := []byte(`{"num": 6.13, "strs": ["a", "b"]}`) // ``表示不进行任何转义的字符串，包括换行
	var dat map[string]interface{}                     // interface{}可以接收任何类型的反序列化

	// Unmarshal()只能反序列化字节切片
	if err := json.Unmarshal(bts, &dat); err != nil { // Unmarshal()：第二个参数需要传递地址
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // 类型断言，如果成功，num则为float64对象
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // 类型断言是否为切片
	str1 := strs[0].(string)            // 类型断言切片中的数据是否为字符串
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res) // 看53行，首先需要将字符串转为字节切片
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 与上述序列化不同的是，NewEncode()会将序列化的数据写入到os.Stdout中
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
