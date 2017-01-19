package main

// 单向channel, 保证channel在函数内部正确被使用，只读channel不能进行写操作，只写channel不能进行读操作
import "fmt"

// chan<- string: 表明该channel只能往里写string类型，不能读
// 将msg写入只写channel中
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan string: 只读channel，不能写
// 从pings中读取，将读出的数据写入pongs中
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings := make(chan string, 1)			// 不能创建单向channel，编译器会根据函数参数的channel类型进行自动转换
	pongs := make(chan string, 1)			// 因为pings、pongs均为缓冲channel，因此针对channel的读写操作均为异步

	ping(pings, "passed message")			// 编译器会自动将pings转换为只写的单向channel
	pong(pings, pongs)						// 同上，编译器会将pings转换为只读的单向channel，将pangs转换为只写的单向channel
	fmt.Println(<-pongs)					// 特别注意：单向channel不能转换为普通channel

}
