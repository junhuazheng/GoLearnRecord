package utils

import (
	"encoding/json"
	"encoding/binary"
	"errors"
	"fmt"
	common "chat/common/message"
	"net"
)

tupe Dispatcher struct {
	Conn net.Conn
	Buf [1024*10]byte
}

func (dispatcher Dispatcher) ReadData() (message common.Message, err error) {
	buf := dispatcher.Buf

	//read message length
	n, err := dispatcher.Conn.Read(buf[:4])
	if err != nil {
		return
	}
	var dataLen uint32
	dataLen = binary.BigEndian.Unit32(buf[:4])

	//read message itself
	i, err := dispatcher.Conn.Read(buf[:dataLen])
	if err != nil {
		fmt.Printf("server read data login data error: %v\n", err)
		return
	}

		//Compare the length of the message to the expected length
		if i != int(dataLen) {
			err = errors.New("login message length error")
			return
		}
	
		//Parse the message from conn and place it in msg, the msg address must be passed here
		err = json.Unmarshal(buf[:dataLen], &msg)
		if err != nil {
			fmt.Printf("json.Unmarshal error: %v\n", err)
		}
		return
	}
}
	
func (dispatcher Dispatcher) SendData(data []byte) (err error) {
	//First send the length of the data to server
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4],dataLen)
	
	//client sends the message length
	_, err := dispatcher.Conn.Write(bytes[:])
	if err != nil {
		fmt.Printf("send data length to server error: %v\n", err)
		return
	}
	
	//client sends the message
	writeLen, err = dispatcher.Conn.Write(data)
	if err != nil {
		fmt.Printf("send data length to server error: %v\n", err)
		return
	}
	return
}