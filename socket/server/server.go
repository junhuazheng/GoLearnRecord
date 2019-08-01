package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("Listen failed err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed err: ", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Println("read: ", string(buf))
	}
}