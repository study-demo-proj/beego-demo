package models

type CommonInfo struct {
	Username string
	RemoteIP string
}

//响应Body
type ResponseBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//响应的状态码1
const (
	CODE_SUCCESS         = 0
	CODE_SYS_ERROR       = 1
	CODE_PARAMETER_ERROR = 2
	CODE_NO_AUTH_ERROR   = 3
	CODE_NO_LOGIN_ERROR  = 9
	CODE_BUSINESS_ERROR  = 4
)

//响应信息
const (
	MSG_SUCCESS         = "success"
	MSG_SYS_ERROR       = "system error"
	MSG_PARAMETER_ERROR = "invalid request"
	MSG_NO_AUTH__ERROR  = "do not have permission"
	MSG_NO_LOGIN_ERROR  = "no login"
)

