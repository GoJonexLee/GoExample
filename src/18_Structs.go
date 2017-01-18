package main

import "fmt"

// person是包级私有结构体，因为person的首字母小写
// 如果结构体名称的首字母大写，则可以包导出
type person struct {
	name string		// 字段名称首字母小写，包外不可通过对象+'.'调用
	Age  int		// 字段名称首字母大写，包外可直接通过对象+'.'调用
}

func main() {

	p1 := person{"Bob", 20}					// 定义以及初始化，"Bob"对应name，20对应age，赋值顺序必须和person中的字段声明顺序相同
	fmt.Println(p1, p1.name, p1.Age)

	p2 := person{Age: 30, name: "Job"}		// key-value初始化，字段的赋值顺序随意改变
	fmt.Println(p2, p2.name, p2.Age)

	p3 := person{Age: 48}					// key-value初始化可以选择性赋值，未赋值字段的值则为其类型对应的零值
	fmt.Println(p3, p3.name, p3.Age)

	p4 := new(person)						// new()返回person类型的指针，同时将person中的字段赋值为每个字段对应类型的零值
	fmt.Println(p4, p4.name, p4.Age)

	p4.name, p4.Age = "dawang", 50			// p4对象中的字段赋值
	fmt.Println(p4, p4.name, p4.Age)

	p5 := &person{name: "pp"}				// 通过&+{}的方式返回结构体实例的指针
	p5.name, p5.Age = "p5", 50
}
