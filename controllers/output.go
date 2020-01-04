package controllers

type Output struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Custom interface{} `json:"custom"`
	Token  string      `json:"token"`
}

const (
	SuccessCode = 1
	ErrorCode   = 0
	NoPermCode  = 0
	OptionsCode = 1
)

const (
	SuccessMsg = "操作成功"
	ErrorMsg   = "系统异常"
	NoPermMsg  = "没有权限"
)

var DefaultErrorMsgList = map[string]string{
	"<QuerySeter> no row found": "未知数据",
}

func (output *Output) SuccessOutput(data interface{}, msg string) *Output {
	output.Code = SuccessCode
	if msg == "" {
		output.Msg = SuccessMsg
	} else {
		output.Msg = msg
	}
	output.Data = data
	return output
}

func (output *Output) ErrorOutput(msg string) *Output {
	output.Code = ErrorCode
	if msg == "" {
		output.Msg = ErrorMsg
	} else {
		output.Msg = output.SetDefaultErrorMsg(msg)
	}
	return output
}

func (output *Output) NoPermOutput(msg string) *Output {
	output.Code = NoPermCode
	if msg == "" {
		output.Msg = NoPermMsg
	} else {
		output.Msg = output.SetDefaultErrorMsg(msg)
	}
	return output
}

func (output *Output) SetToken(token string) *Output {
	output.Token = token
	return output
}

// 获取error默认提示信息
func (output *Output) SetDefaultErrorMsg(errMsg string) string {
	if DefaultErrorMsgList[errMsg] != "" {
		return DefaultErrorMsgList[errMsg]
	}
	return errMsg
}

func (output *Output) OptionsOutput(msg string) *Output {
	output.Code = OptionsCode
	if msg == "" {
		output.Msg = NoPermMsg
	} else {
		output.Msg = output.SetDefaultErrorMsg(msg)
	}
	return output
}
