package common

const (
	LoginMessageType = "LoginMessage"
	RegisterMessageType = "RegisterMessage"
	LoginResponseMessageType = "LoginResponseMessageType"
	RegisterResponseMessageType = "RegisterResponseMessage"
	UserSendGroupMessageType = "UserSendGroupMessageType"
	SendGroupMessageToClientType = "SendGroupMessageToClientType"
	ShowAllOnlineUsersType = "ShowAllOnlineUsersType"
	PointToPointMessageType = "PointToPointMessageType"

	ServerError = 500

	//status code fo login
	LoginError = 403
	NotExit = 404
	LoginSucceed = 200

	//status code fo register
	HasExited = 403
	RegisterSucceed = 200
	PasswordNotMatch = 402
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMessage struct {
	UserName string
	Password string
}

type ResponseMessage struct {
	Type string
	Code int 
	Error string
	Data string
}

type RegisterMessage struct {
	UserName string
	Password string
	PasswordConfirm string
}

type UserSendGroupMessage struct {
	GroupID int //tatget group id, 0 => all users
	UserName string //current user name
	Content string //message content
}

type SendGroupMessageToClient struct {
	GroupID int //group id, 0 => all users
	UserName string //current user
	Content string //message
}

type PointToPointMessage struct {
	SourceUserID int
	SourceUserName string
	TatgetUserID int
	TatgetUserName string
	Content string
}

//online user info
type UserInfo struct {
	ID int
	UserName string
}