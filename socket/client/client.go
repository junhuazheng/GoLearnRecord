package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error dialing: ", err.Error())
		return
	}

	defer conn.Close()

	/*
	inputReader是一个指向bufio.Reader的指针。
	inputReader := bufio.NewReader(os.Stdin)，将会创建一个读取器，并将其与标准输入绑定
	*/
	inputReader := bufio.NewReader(os.Stdin)

	for {
		/*
		读取器对象提供了一个方法 ReadString(delim byte)
		该方法从输入中读取内容，直到碰到delim指定的字符
		然后将读取到的内容连同 delim 字符一起放到缓冲区

		在这个例子中，我们会读取键盘输入，直到回车键(\n)被按下
		*/
		input, _ := inputReader.ReadString('\n')

		/*
		func Trim(s string, cutest string) string
		返回将 s 前后端所有 cutest 包含的 utf-8 码值都去掉的字符串
		*/
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			return
		}
	}
}