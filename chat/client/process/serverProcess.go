package process

import (
	"encoding/json"
	"errors"
	"chat/client/logger"
	"chat/client/model"
	"chat/client/utils"
	common "chat/common/message"
	"net"
)

func dealLoginResponse(responseMsg common.ResponseMessage) (err error) {
	switch responseMsg.Code {
	case 200:
		//parses current user infomation
		var userInfo common.UserInfo 
		err = json.Unmarshal([]byte(responseMsg.Data), &userInfo)
		if err != nil {
			return
		}

		//initialize the CurrentUser
		user := model.User{}
		err = user.InitCurrentUser(userInfo.ID, userInfo.UserName)
		if err != nil {
			return
		}
		logger.Success("Login succeed!\n")
		logger.Notice("Current user, id: %d, name: %v\n", model.CurrentUser.UserID, model.CurrentUser.UserName)
	case 500:
		err = errors.New("Server error!")
	case 404:
		err = errors.New("User does not exist!")
	case 403:
		err = errors.New("Passwor invalid!")
	default:
		err = errors.New("Some error!")
	}
	return
}

func dealRegisterResponse(responseMsg common.ResponseMessage) (err, error) {
	switch responseMsg.Code {
	case 200:
		logger.Success("Register succeed!\n")
	case 500:
		err = errors.New("Server error!")
	case 403:
		err = errors.New("User already exists!")
	case 402:
		err = errors.New("Password invalid!")
	default:
		err = errors.New("Some error!")
	}
	return
}

func dealGroupMessage(ResponseMsg common.ResponseMessage) (err error) {
	var groupMessage common.SendGroupMessageToClient
	err = json.Unmarshal([]byte(responseMsg.Data), &groupMessage)
	if err != nil {
		return
	}
	logger.Info("%v send you:", groupMessage.UserName)
	logger.Notice("\t%v\n", groupMessage.Content)
	return
}

func showAllOnlineUsersList(responseMeg common.ResponseMessage) (err error) {
	if responseMsg.Code != 200 {
		err = errors.New("server Error!")
		return
	}

	var userList []common.UserInfo
	err = json.Unmarshal([]byte(responseMeg.Data), &userList)
	if err != nil {
		return
	}

	logger.Success("Online user list(%v users)\n", len(userList))
	logger.Notice("\t\tID\t\tname\n")
	for _, info := range userList {
		logger.Success("\t\t%v\t\t%v\n", info.ID, info.UserName)
	}

	return
}

func showPonitToPointMessage(responseMsg common.ResponseMessage) (err error) {
	if responseMsg.Code != 200 {
		err = errors.New(responseMsg.Error)
		return
	}

	var pointToPointMessage common.PointToPointMessage
	err = json.Unmarshal([]byte(responseMsg.Data), &pointToPointMessage)
	if err != nil {
		return
	}

	logger.Info("\r\n\r\n%v say: ", pointToPointMessage.SourceUserName)
	logger.Notice("\t%v\n", pointToPointMessage.Content)
	return
}

//handle the server's return
func Response(conn net.Conn, errMsg chan error) (err error) {
	var responseMsg common.ResponseMessage
	dispatcher := utils.Dispatcher{Conn: conn}

	for {
		responseMsg, err = dispatcher.ReadDate()
		if err != nil {
			logger.Error("Waiting response error: %v\n", err)
			return
		}

		//According to the message type returnde by the server,the corresponding processing is carried out
		switch responseMsg.Type {
		case common.LoginResponseMessageType:
			err = dealLoginResponse(responseMsg)
			errMsg <- err
		case common.RegisterResponseMessageType:
			err = dealGroupMessage(responseMsg)
			errMsg <- err
		case common.SendGroupMessageToClientType:
			err = dealGroupMessage(responseMsg)
			if err != nil {
				logger.Error("%v\n", err)
			}
		case common.ShowAllOnlineUsersType:
			err = showAllOnlineUsersList(responseMsg)
			errMsg <- err
		case common.PointToPointMessageType:
			err = showPonitToPointMessage(responseMsg)
			errMsg <- err
		default:
			logger.Error("Unkonwn message type!")
		}
	}
}
