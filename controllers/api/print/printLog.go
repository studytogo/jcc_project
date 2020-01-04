package print

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/printService"
	"strings"
)

type PrintLogController struct {
	controllers.CommonController
}

// @router /printlog [post]
func (this *PrintLogController) PrintLog() {

	var Order_IDs = new(request.OrderIDs)
	this.GetParamStruct(Order_IDs)
	orderids_list := strings.Split(Order_IDs.OrderIDs, ",")

	orderids_str := strings.Join(orderids_list, ",")
	err := printService.PrintLog(orderids_str)
	// fmt.Println(err)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功~！！！")
}
