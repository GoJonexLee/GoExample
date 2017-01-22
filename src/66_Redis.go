package main

import (
	"github.com/garyburd/redigo/redis"	// redis包, 通过:go get github.com/garyburd/redigo/redis  命令就可以自己下载安装
	"time"
	"fmt"
)

var rdPool = redis.Pool		// redis连接池

func init() {			// 初始化函数， 会在全局变量初始化之后，main()函数之前执行

	rdPool = &redis.Pool {
		MaxIdel: 10,	// 设置最大链接数
		IdelTimeout: 240 * time.Second,	// 设置超时
		Dial: func() (redis.Comm, error) {		// 设置链接函数
			c, err := redis.Dial("tcp", "host")
			if err != nil {
				return nil, err
			}

			_, err = c.Do("AUTH", "password")	// 验证密码
			if err != nil {
				c.Close()			// 返回时需要关闭链接
				return nil, err
			}

			_, err = c.Do("SELECT", "database number")	// 选择库
			if err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
	}
}

func main() {

	cn := rdPool.Get()	// 从连接池中获取一个链接
	defer cn.Close()	// 函数返回前关闭链接

	cn.Do("COMMAND", "Parames")	// 第一个参数是要执行的命令，后面的参数列表你懂的

	ok, err := redis.Bool(cn.Do("EXISTS", "id"))	// redis.Bool()：将Do()函数的结果转换为bool类型
	if err != nil {
		panic(err)
	}

	fmt.Println("your id ", ok)
}