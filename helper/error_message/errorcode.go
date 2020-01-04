package error_message

import "errors"

const (
	JsonUnmarshalErr     = "140000" // 解析json表单出错
	LostParamError       = "140001" // 缺少必要的参数
	CheckParamError      = "140002" // 校验参数出错
	INSERT_PRODUCT_ERROR = "150001" //商品spu插入失败
	Add_CATEGORY_ERROR   = "160001" //添加分类名称重复
	CATEGORY_LEVEL_ERROR = "160002" //添加分类等级错误
	CATEGORY_NO_ERROR    = "160003" //添加分类编号错误
	ORDER_SN_ERROR       = "170001" //订单编号不存在
	UPDATE_ORDER_ERROR   = "170002" //修改订单信息失败
	GOOD_ORDER_ERROR     = "170003" //修改订单信息失败
)

type NewError struct {
	Code    string
	Message string
	Err     error
}

func (newErr *NewError) Initialize(code string, message string, err error) {
	newErr.Code = code
	newErr.Message = message
	newErr.Err = err
}

func (e NewError) Error() string {
	if e.Err == nil {
		return "错误提示：" + e.Message
	}
	//return errors.New("错误提示：" + message + "----" + "错误原因：" + err.Error())
	return "错误提示：" + e.Message + "----" + "错误原因：" + e.Error()
}

//公共返回错误方法
func ErrMessage(message string, err error) error {
	if err == nil {
		return errors.New("错误提示：" + message)
	}
	return errors.New("错误提示：" + message + "----" + "错误原因：" + err.Error())
}
