package main

import (
	"fmt"
	"chat/server/model"
	"chat/server/process"
	"net"
	"time"
)

func init() {
	//initialize redis connection pool
	initRedisPool(16, 0, time.Second*300, "127.0.0.1:6379")

	//create userDao to manipulate user information
	model.CurrentUserDao = model.InitUserDao(pool)
}

//communicate and interact with the client
//conn is the connection established between the client and the server
//start a goroutine every time a user logs in
//this groutine is designed to handle client-server communication
func dialogue(conn net.Conn) {
	defer conn.Close()
	processor := process.Processor{Conn: conn}
	processor.MainProcess()
}

func main() {
	fmt.Println("Server is already\n")

	listener, err := net.Listen("tcp", "localhost:8888")
	defer listener.Close()
	if err != nil {
		fmt.Printf("some error when run server, error: %v\n", err)
	}

	for {
		fmt.Println("Waiting fo client ...\n")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("some error when accept server, error: %v\n", err)
		}

		//if the connection is successful, start another goroutine to communicate with the client
		go dialogue(conn)
	}
}