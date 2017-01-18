package main

import "fmt"

// 注意区分rect的普通对象方法和指针对象方法
type rect struct {
	width, height int
}

// rect指针对象绑定了area()方法
func (r *rect) area() int {
	return r.width * r.height
}

// rect普通对象绑定了perim()方法
func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

// 不能改变普通对象r中字段的值
func (r rect) set(w, h int) {
	r.width, r.height = w, h
	fmt.Println(r)
}

// 可以改变指针对象r中的字段值
func (r *rect) set1(w, h int) {
	r.width, r.height = w, h
}

func main() {

	r := rect{width: 10, height: 5}			// 通过{}实例话对象，另外一种实例话方法是通过new()来实现
	fmt.Println("area: ", r.area())			// area()是指针对象的方法，编译器会自动将r转换为rect的指针类型
	fmt.Println("perim: ", r.perim())		// perim()是普通对象的方法，可以直接调用，不需要编译器转换

	rp := &r
	fmt.Println("area: ", rp.area())		// rp已经是rect的指针对象，因此可以直接调用，不需要编译器转换
	fmt.Println("perim: ", rp.perim())		// perim()是普通对象的方法，因此需要编译器将指针类型转换为普通对象类型

	// 总结：1.不论是指针对象方法还是普通对象方法，在调用实例的方法时，编译器都会根据方法的接收器类型自动转换。
	//       2.但在实现接口时需要特别注意，指针对象实例可以调用指针对象的方法以及普通对象的方法，而普通对象只能调用普通对象的方法。
	//       3.只有指针对象的方法才能修改对象实例中的字段值，普通对象不会修改，如下例所示：

	r1 := rect{10, 20}
	r1.set(1, 1)
	fmt.Println(r1)			// r1中的字段值并未被改变

	r2 := &rect{30, 40}
	r2.set(2, 2)
	fmt.Println(r2)			// r2是指针对象，但因为set是普通对象方法，因此编译器会将r2转换为普通对象，因此r2中的字段值也并未被改变

	r2.set1(4,4)
	fmt.Println(r2)			// 通过指针对象方法set1(int, int)则可以改变r字段中的值
}
