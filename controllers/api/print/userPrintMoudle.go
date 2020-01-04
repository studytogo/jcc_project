package print

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/print"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/printService"
)

type UserPrintMoudleController struct {
	controllers.CommonController
}

// @Title  查询用户的打印模板信息
// @Summary 查询用户的打印模板信息
// @Description 查询用户的打印模板信息
// @Param mould_id               raw  	string  	true  		`模板id`
// @Param top_margin    		 raw  	string  	true  		`上边距`
// @Param left_margin            raw  	string  	true  		`左边距`
// @Param is_display_price       raw  	string  	true  		`是否显示价格`
// @Param token            		 raw  	string  	true  		`token值`
// @Success 0 {}
// @Failure 1 body is empty
// @router /update_user_print [post]
func (this *UserPrintMoudleController) UpdateOrAddUserPrintMoudle() {
	request := new(request.UserPrintMoudle)
	this.GetParamStruct(request)
	userMoudle := new(print.JccUserMould)
	err := helper.ReflectiveStruct(userMoudle, request)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "数据转换失败", err)
		this.CheckError(err, "")
	}

	err = printService.UpdateOrAddUserPrintMoudle(userMoudle)

	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "失败", err)
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功~！！！")

}
