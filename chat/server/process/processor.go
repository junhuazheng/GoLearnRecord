package process

import (
	"fmt"
	common "chat/common/message"
	"chat/server/model"
	"chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//Process the message
//use the appropriate processing according to the message type
func (this *Process) messageProcess(message common.Message) (err error) {
	switch message.Type {
	case common.LoginMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserLogin(message.Data)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
		}
	case common.RegisterMessageType:
		up := UserProcess{Conn: this.Conn}
		err = up.UserRegister(message.Data)
		if err != nil {
			fmt.Printf("some error when register: %v\n", err)
		}
	case common.UserSendGroupMessageType:
		fmt.Println("user send group message!")
		gmp := GroupMessageProcess{}
		gmp.sendToGroupUsers(message.Data)
	case common.ShowAllOnlineUsersType:
		olP := OnlineInfoProcess{this.Conn}
		err = olP.ShowAllOnlineUsersList()
		if err != nil {
			fmt.Prinltn("get all online user list error: %v\n", err)
		}
	case common.PointToPointMessageType:
		fmt.Prinltn("point to point comminite!")
		pop := PointToPointMessageProcess{}
		err = pop.sendMessageToTargetUser(message.Data)
		var code int
		if err != nil {
			code = 400
		} else {
			code = 100
		}

		err := pop.responseClient(this.Conn, code, "", err)
		if err != nil {
			fmt.Printf("some err wehn popmessage: %v\n", err)
		}
	default:
		fmt.Println("other type\n")
	}
	return
}

//handle communication with users
func (this *Process) MainProcess() {

	//loop reads the infomation from the client
	for {
		dispatcher := utils.Dispatcher{Conn: this.Conn}
		message, err := dispatcher.ReadData()
		if err != nil {
			if err == io.EOF {
				cc := model.ClientConn{}
				cc.Del(this.Conn)
				fmt.Println("client closed!\n")
				break
			}
			fmt.Printf("get login message error: %v\n", err)
		}

		//process infomation from the client
		//different processing methods are used depending on the type of message
		err = this.messageProcess(message)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
			break
		}
	}
}