package error_message

const (
	AuthError     = "110001" // 权鉴失败
	SignLenError  = "110002" // 传入的sign长度不正确
	AppidError    = "110003" // 传入Appid出错
	FormatError   = "110004" // 权鉴校验格式出错
	SignCompError = "110005" // sign比较出错
	AuthTimeout   = "110006" // 超过设置权鉴超时
	AuthTimeErr   = "110007" // 权鉴时间配置错误
)
