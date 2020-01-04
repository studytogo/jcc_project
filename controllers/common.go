package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/models/record"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type CommonController struct {
	beego.Controller
	token string
}

//成功的json
func (controller *CommonController) Success(data interface{}, msg string) {
	controller.Data["json"] = new(Output).SuccessOutput(data, msg).SetToken(controller.token)
	controller.responseRecord(data, msg)
	controller.ServeJSON()
	controller.StopRun()
}

//失败的json
func (controller *CommonController) Fail(msg string) {
	controller.Data["json"] = new(Output).ErrorOutput(msg).SetToken(controller.token)
	controller.responseRecord(nil, msg)
	controller.ServeJSON()
	controller.StopRun()
}

// validation验证，error直接返回
func (controller *CommonController) CheckValidation(requestData interface{}) {
	valid := validation.Validation{}
	res, err := valid.Valid(requestData)
	if err != nil {
		helper.Log.ControllerError(controller.Ctx.Input.URL(), "CheckValidation error result:", err)
		controller.Fail(err.Error())
	}
	// 包含错误信息
	if !res {
		errors := ""
		for _, err := range valid.Errors {
			helper.Log.ControllerError(controller.Ctx.Input.URL(), "CheckValidation error result:", err)
			errors += fmt.Sprintf("%s:%s;", err.Key, err.Message)
		}
		controller.Fail(errors)
	}
}

// 通过结构体获取数据
func (controller *CommonController) GetParamStruct(paramsStruct interface{}) {
	if err := controller.ParseForm(paramsStruct); err != nil {
		controller.Fail(err.Error())
	}
	//表单验证
	controller.CheckValidation(paramsStruct)
}

// 错误信息验证
func (controller *CommonController) CheckError(err error, msg string) {
	if err != nil {
		logs.Error(err)
		if msg == "" {
			msg = err.Error()
		}
		controller.Fail(msg)
	}
}

// 验证码图片地址有效
// func (controller *CommonController) CheckPictureUrl(pictureUrl string)  {
// 	reqManager := http_manager.NewRequestManager(false,false)
// 	_ ,err := reqManager.Get(pictureUrl, nil)
// 	controller.CheckError(err,"图片地址无效")
// }

// string to int64
func (controller *CommonController) String2Int64(str string) int64 {
	if str == "" {
		return 0
	}
	val, err := strconv.ParseInt(str, 10, 64)
	controller.CheckError(err, "strconv.ParseInt is error")
	return val
}

// string to uint8
func (controller *CommonController) String2Uint8(str string) uint8 {
	if str == "" {
		return 0
	}
	val, err := strconv.Atoi(str)
	controller.CheckError(err, "strconv.Atoi is error")
	return uint8(val)
}

// string to uint
func (controller *CommonController) String2Uint(str string) uint {
	if str == "" {
		return 0
	}
	val, err := strconv.Atoi(str)
	controller.CheckError(err, "strconv.Atoi is error")
	return uint(val)
}

// 方便获取页码
func (controller *CommonController) GetPageNo() uint {
	pageNo, err := controller.GetInt64("page_no")
	controller.CheckError(err, "")
	if pageNo == 0 {
		pageNo = 1
	}
	return uint(pageNo)
}

//动态设置token
func (controller *CommonController) SetToken(token string) {
	controller.token = token
}

// 通过注解的valid来校验，传入是结构体的地址
func (controller *CommonController) CheckParams(params interface{}) {
	valid := validation.Validation{}
	ok, err := valid.Valid(params)
	controller.CheckError(err, error_message.CheckParamError)
	if !ok {
		err := errors.New("verification error，maybe lack of necessary parameters ")
		controller.CheckError(err, error_message.LostParamError)
	}
}

// 把校验的内容删掉
func (controller *CommonController) ClearAuth(clearMap *map[string]interface{}) {
	delete(*clearMap, "CenterAppId")
	delete(*clearMap, "CenterAppName")
	delete(*clearMap, "CenterTimeStamp")
}

//专为修改单位验证接口使用的返回错误
func (controller *CommonController) UnitCheckError(data interface{}, err error) {
	controller.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  err.Error(),
		"data": data,
	}
	controller.ServeJSON()
	controller.StopRun()
}

//json接收数据转换的公用方法
func (controller *CommonController) GetJsonStruct(paramsStruct interface{}) {
	if err := json.Unmarshal(controller.Ctx.Input.RequestBody, &paramsStruct); err != nil {
		controller.Fail(err.Error())
	}
	//表单验证
	controller.CheckValidation(paramsStruct)
}

//记录返回值
func (controller *CommonController) responseRecord(data interface{}, msg string) {
	//将返回的json存入数据库-----
	requestId, _ := controller.Ctx.Input.GetData("requestId").(int)
	if requestId == 0 {
		return
	}
	fff := fmt.Sprintf("%+v", data)
	if data == nil {
		fff = msg
	} else {
		if value, ok := data.(string); ok {
			if value == "" {
				fff = msg
			}
		}
	}
	err := record.UpdateContext(requestId, fff)
	if err != nil {
		helper.Log.Error("插入Context_Param数据库失败", err)
	}
	//---------------------------
}
