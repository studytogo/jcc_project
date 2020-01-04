package print

import (
	"fmt"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	print2 "new_erp_agent_by_go/models/print"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api"
)

type SelectPrintController struct {
	controllers.CommonController
}

// @Title 				  条码模板查看接口
// @Summary 			  条码模板查看接口
// @Description 		  条码模板查看接口
// @Param id  raw  	string  	false  		`需要查看的木板id 不传为查全部`
// @Param Token           raw  	string  	true  		`token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /select_mould_one_or_list [post]
func (this *SelectPrintController) SelectMouldOneOrList() {
	var mouldString = new(request.JccCommonMould)
	this.GetParamStruct(mouldString)
	var mould = new(print2.JccCommonMould)
	err := helper.ReflectiveStruct(mould, mouldString)
	if err != nil {
		helper.Log.Info("")
	}
	fmt.Println("mould", mould)
	fmt.Println("id", mould.Id)

	// 查列表
	info, err := api.SelectMouldList()

	// 查一条
	if mould.Id > 0 {
		info1, err := api.SelectMould(mould.Id)
		if err != nil {
			helper.Log.Info("")
		}
		this.Success(info1, "成功！")
	}

	if err != nil {
		helper.Log.Info("")
	}

	this.Success(info, "成功！")

}
