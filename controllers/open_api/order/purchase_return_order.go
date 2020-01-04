package order

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/order"
)

type PurchaseReturnOrderController struct {
	controllers.CommonController
}

// @router /update_purchase_return_order_status [post]
func (this *PurchaseReturnOrderController) UpdatePurchaseReturnOrderStatus() {
	request := new(request.PurchaseReturnOrderStatus)
	this.GetJsonStruct(request)
	err := order.UpdatePurchaseReturnOrdersStatusById(request)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil,"操作成功")
}