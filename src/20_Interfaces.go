package main

import (
	"fmt"
	"math"
)

// 包含三个方法的接口
type geometry interface {
	area()  float64			// 无入参，返回float64类型
	perim() float64			// 同上
	name(string)			// 入参string类型，无返回值
}

// 结构体rect，通过对象实例实现geometry接口
type rect struct {
	width, height float64
}

// 实现geometry接口的area方法
func (r rect) area() float64 {
	return r.width * r.height
}

// 实现geometry接口的perim方法
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 实现geometry接口的name方法
func (rect) name(st string) {
	fmt.Println(st)
}

// 结构体circle，通过对象实例的指针实现geometry接口
type circle struct {
	radius float64
}

// circle对象实例的指针类型为接收体, 即*circle
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (*circle) name(st string) {
	fmt.Println(st)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	g.name("i am not a interface")
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(&r)

	//measure(c)		// 不能通过circle的实例对象进行接口的赋值，因为circle是通过指针实现接口的
	measure(&c)		// 正确方式是这样

	// 总结: 如果是结构体的对象实例实现接口，均可通过对象、指针传入interface的参数中, 例如rect；
	//		 如果是结构体的指针实例实现接口，则只能通过指针传入interface的参数中，例如circle。
}

