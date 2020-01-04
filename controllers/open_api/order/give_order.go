package order

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/order"
)

type GiveOrderController struct {
	controllers.CommonController
}

// @router /give_order_status [post]
func (this *GiveOrderController) UpdateGiveOrderStatus() {
	request := new(request.GiveOrderStatus)
	this.GetJsonStruct(request)

	err := order.UpdateGiveOrdersStatusById(request)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil,"操作成功")
}