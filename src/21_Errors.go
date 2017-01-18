package main

import (
	"errors"
	"fmt"
)

// 在需要错误返回类型时，error一般作为最后一个参数返回，如果程序正确，则error类型为nil
// 通过errors.New()可以自定义错误类型
// error是接口类型，其中只包含了Error() string 这一个方法
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)		//Sprintf()函数根据格式化以及参数列表返回字符串
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}	// 由于*argError实现了error接口，因此可以被当作error类型返回
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {		// 类型断言,即e是否是*argError的实现，如果是ok为true，否则ok为false
		fmt.Println(ae.arg)					// 如果类型断言成功，ae则为argError的指针对象
		fmt.Println(ae.prob)
	}
}
