package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")					// 执行完case逻辑块后，会自动跳出switch语句，不用使用break
	case 2:
		fmt.Println("two")
		fallthrough							// switch专属关键字，表明不跳出switch，直接执行下一个case逻辑块，而不用判断是否匹配下一个case后的语句
	case 3:
		fmt.Println("three")				// 因为case 2逻辑块有fallthrough关键字，所以该行代码一定会执行，而不用判断case 3这条语句
	default:								// default的位置可以任意摆放，而不用放在最后。
		fmt.Println("hehe")
	}

	switch time.Now().Weekday() {
	default:								// 虽然default放在第一位，但default仍然是在所有的case不满足时才会执行
		fmt.Println("It's a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	}

	t := time.Now()
	switch {								// switch后也可不用跟任何变量
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {		// switch后可跟赋值语句
		switch t := i.(type) {				// 类型断言，t为i接口的具体类型
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	switch i := 10; {						// 如果switch后有赋值语句，则必须以';'符号结尾
	case i > 0 && i < 10:
		fmt.Println("small")
	case i > 10:
		fmt.Println("equal")
	default:
		fmt.Println("ten")
	}
}
