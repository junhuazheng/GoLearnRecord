package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"chat/client/logger"
	common "chat/common/message"
	"net"
)

type Dispatcher struct {
	Conn net.Conn
	Buf [1024 * 10]byte
}

func (dispatcher Dispatcher) ReadDate() (msg common.ResponseMessage, err error) {
	buf := make([]byte, 1024*10)

	//Read message length
	n, err := dispatcher.Conn.Read(buf[:4])
	if err != nil {
		return
	}
	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(buf[:4])

	//Read the message itself
	n, err = dispatcher.Conn.Read(buf[:dataLen])
	if err != nil {
		logger.Error("Server read data login data error: %v\n", err)
	}

	//Compare the length of the message to the expected length
	if n != int(dataLen) {
		err = errors.New("login message length error")
		return
	}

	//Parse the message from conn and place it in msg, the msg address must be passed here
	err = json.Unmarshal(buf[:dataLen], &msg)
	if err != nil {
		logger.Error("json.Unmarshal error: %v\n", err)
	}
	return
}

func (dispatcher Dispatcher) SendData(data []byte) (err error) {
	//First send the length of the data to server
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4],dataLen)

	//client sends the message length
	writeLen, err := dispatcher.Conn.Write(bytes[:1])
	if writeLen != 4 || err != nil {
		logger.Error("send data to server error: %v\n", err)
		return
	}

	//client sends the message
	writeLen, err = dispatcher.Conn.Write(data)
	if err != nil {
		logger.Error("send data length to server error: %v\n", err)
		return
	}
	return
}