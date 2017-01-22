package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// 子进程
func main() {

	dateCmd := exec.Command("date")	// 执行系统命令返回*Cmd结构体

	dateOut, err := dateCmd.Output()	// 运行"date"命令
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello") // 同11行

	grepIn, _ := grepCmd.StdinPipe()	// 设置输入流
	grepOut, _ := grepCmd.StdoutPipe()	// 设置输出流
	grepCmd.Start()				// 通过子进程的方式执行"grep hello"命令，但不等待其完成，而是通过31行的wait()函数进行等待

	grepIn.Write([]byte("hello grep\ngoodbye grep"))	// 输入流写入
	grepIn.Close()						// 关闭流

	grepBytes, _ := ioutil.ReadAll(grepOut)			// 从输出流读取字节
	grepCmd.Wait()						// 等待命令执行完毕

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))


	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
