package process

import (
	"encoding/json"
	"fmt"
	common "chat/common/message"
	"chat/server/model"
	"chat/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func register(userName, passWrod, passWrodConfirm string) (user model.User, err error) {
	user, err = model.CurrentUserDao.Register(userName, passWrod, passWrodConfirm)
	return
}

func login(userName, passWrod string) (user model.User, err error) {
	//determine username and password
	user, err = model.CurrentUserDao.Login(userName, passWrod)
	return
}

//response client
fucn (this *UserProcess) responseClient(responseMessageType string, code int, data string, err error) {
	var responseMessage common.ResponseMessage
	responseMessage.Code = code
	responseMessage.Type = responseMessageType
	responseMessage.Data = data

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Printf("some error when generate response message, error: %v\n", err)
	}

	dispatcher := utils.Dispatcher{Conn: this.Conn}

	err = dispatcher.WriteData(responseData)
}

func (this *UserProcess) UserRegister(message string) (err error) {
	var info common.RegisterMessage
	var code int
	data := ""
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		code = common.ServerError
	}

	_, err = register(info.UserName, info.Password, info.PasswordConfirm)
	switch err {
	case niil:
		code = common.RegisterSucceed
	case model.ERROR_PASSWORD_DOES_NOT_MATCH:
		code = 402
	case model.ERROR_USER_ALREADY_EXISTS:
		code = 403
	default:
		code = 500
	}
	
	this.responseClient(common.RegisterMessageType, code, data, err)
	return
}

func (this *UserProcess) UserLogin(message string) (err error) {
	var info common.LoginMessage
	var code int
	var data string
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		code = common.ServerError
	}

	user, err := login(info.UserName, info.Password)

	switch err {
	case nil:
		code = common.LoginSucceed
		//save user conn status
		clientConn := model.ClientConn{}
		clientConn.Save(user.ID, user.Name, this.Conn)

		userInfo := common.UserInfo{user.Id, user.Name}
		info _, := json.Marshal(userInfo)
		data = string(info)
	case model.ERROR_USER_DOES_NOT_EXISTS:
		code = 404
	case model.ERROR_USER_PWD:
		code = 403
	default:
		code = 500
	}
	
	this.responseClient(common.LoginResponseMessageType, code, data, err)
	return
}