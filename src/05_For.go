package main

// for 语句的使用
import "fmt"

/*
go中只存在for一种循环方式，while在go中不存在
*/

func main() {
	i := 1

	for i <= 3 {					// go中规定，'{'必须放在一行的末尾，不能另起一行
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {							// 无限循环，通过break可以跳出循环体
		fmt.Println("loop")
		break
	}

	Lee:							// 标签，用于break, continue, goto跳转使用
	for {							// 双无限循环
		for {
			fmt.Println("double loop")
			break Lee				// 通过标签的方式一次跳出双重循环; continue Lee 则为从外层循环继续
		}
	}

									// 此处有坑，大家一定要小心非引用类型的拷贝代价
	dt := [...]int{1, 2, 3}			// 数组, 非引用类型, 后面会单独讲解
	for index, value := range dt {	// range循环, index为数据的下表，value为下标对应的值
		if index == 0 {				// range循环会对dt数据进行拷贝，因此对dt的操作不会影响value的值
			dt[0] += 10
			dt[1] += 20
			dt[2] += 30
		}
		fmt.Println(value)			// value的值仍为1、2、3, 不会应为对dt的修改而改变
	}
}
