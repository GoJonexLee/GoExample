package main

import (
	"syscall"
	"os"
	"os/exec"
)

func main() {

	binary, err := exec.LookPath("ls")		// 查找'ls'的命令文件
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}	// 设置命令和参数

	env := os.Environ()				// 获取环境变量

	err = syscall.Exec(binary, args, env)		// 系统调用
	if err != nil {
		panic(err)
	}
}
