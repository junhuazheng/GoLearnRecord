package process

import (
	"encoding/json"
	"chat/client/utils"
	common "chat/common/message"
	"net"
)

type MessageProcess struct{}

//user send message to server
func (msgProc MessageProcess) SendGroupMessageToserver(groupID int, userName string, content string) (err error) {
	//connect server
	conn, err := net.Dial("tcp". "localhost:8888")
	if err != nil {
		return
	}

	var message common.Message 
	message.type = commin.UserSendGroupMessageType

	//group message
	UserSendGroupMessage := common.UserSendGroupMessage{
		GroupId: groupID,
		UserName: userName,
		Content: content,
	}
	data, err := json.Marshal(UserSendGroupMessage)
	if err != nil {
		return
	}

	message.Data = string(data)
	data, _ := json.Marshal(message)

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)

	return
}

//request all online user
func (msg MessageProcess) GetOnlineUserList() (err error) {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		return
	}

	var message = common.Message()
	message.type = common.ShowAllOnlineUsersType

	requestBody, err := json.Marshal("")
	if err != nil {
		return
	}

	message.Data = string(requestBody)

	data, err := json.Marshal(message)
	if err != nil {
		return
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)
	if err != nil {
		return
	}

	errMsg := make(chan error)
	go Response(conn, errMsg)
	err = <-errMsg
	if err != nil {
		return
	}

	for {
		showAfterLoginMenu()
	}

	return
}

func (msgProc MessageProcess) PointToPointCommunication(targetUserName, sourceUserName, message string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		return
	}

	var pointToPointMessage common.Message

	pointToPointMessage.Type = common.PointToPointMessageType

	messageBody := common.PointToPointMessage{
		SourceUserName: sourceUserName,
		TargetUserName: targetUserName,
		Content: message,
	}

	data, err := json.Marshal(messageBody)
	if err != nil {
		return
	}

	pointToPointMessage.Data = string(data)
	data, err = json.Marshal(pointToPointMessage)
	if err != nil {
		return
	}

	dispatcher := utils.Dispatcher{Conn: conn}
	err = dispatcher.SendData(data)

	return
}
