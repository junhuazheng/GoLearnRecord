package process

import (
	"encoding/json"
	"fmt"
	common "chat/common/message"
	"chat/server/model"
	"cht/server/utils"
	"net"
)

type OnlineInfoProcess struct {
	Conn net.Conn
}

type UserInfo = common.UserInfo

func (this OnlineInfoProcess) showAllOnlineUserList() (err error) {
	var onlineUserList []UserInfo
	var code int
	for _, connInfo := range model.ClientConnsMap {
		user, err := model.CurrentUserDao.GetUserByUserName(connInfo.UserName)
		if err != nil {
			continue
		}
		UserInfo := UserInfo{user.ID, user.Name}
		onlineUserList = append(onlineUserList, userInfo)
	}

	data, err := json.Marshal(onlineUserList)
	if err != nil {
		code = common.ServerError
	} else {
		code = 200
	}

	err = responseClient(this.Conn, code, string(data), fmt.Sprintf("%v", err))
	if err != nil {
		fmt.Printf("point to point communicate, response client error: %v\n", err)
	}
	return
}

func responseClient(conn net.Conn, code int, data string, errMsg string) (err error) {
	responseMessage := common.ResponseMessage{
		Code: code,
		Type: common.ShowAllOnlineUsersType,
		Data: data,
		Error: errMsg,
	}

	responseData, err := json.Marshal(responseMessage)
	if err != nil {
		fmt.Printf("some error when generate response message, error: %v\n", err)
	}
	
	dispatcher := utils.WriteData(responseData)
	if err != nil {
		fmt.Printf("some error: %v\n", err)
	}
	return
}