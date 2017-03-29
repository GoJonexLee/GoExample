package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// 协议、认证信息、host、端口、路径、查询参数以及'#'后面的锚点
	s := "https://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // url的协议

	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host, " port:", port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
