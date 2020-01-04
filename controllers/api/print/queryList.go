package print

import (
	"math"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/purchaseOrder"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/printService"
	"strconv"
)

type QueryListController struct {
	controllers.CommonController
}

// @router /querylist [post]
func (this *QueryListController) QueryListController() {

	companyId := this.Ctx.Request.Form.Get("Companyid")

	var Order_IDs = new(request.OrderIDs)
	this.GetParamStruct(Order_IDs)

	var orderid = Order_IDs.OrderIDs
	//分页
	var page, _ = strconv.Atoi(Order_IDs.Page)
	var per_page, _ = strconv.Atoi(Order_IDs.Per_page)
	var start_page = (page - 1) * per_page
	var end_page = per_page

	list, count, err := printService.CheckList(companyId, orderid, start_page, end_page, page)

	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "")
	}
	var last_page = math.Ceil(float64(count) / float64(per_page))
	result := purchaseOrder.T{}
	result.Data = list
	result.Total = count
	result.Last_page = int(last_page)
	this.Success(result, "操作成功~！！！")
}
