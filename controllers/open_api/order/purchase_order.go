package order

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/order"
)

type PurchaseOrderController struct {
	controllers.CommonController
}

// @Title  修改订单金额
// @Summary 修改订单金额
// @Description 修改订单金额
// @Param source        				raw  	string  	true  		`erp端默认是erp`
// @Param order_sn      				raw  	string  	true  		`订单编号`
// @Param goods_info       				raw  	[]request.GoodsOrderInfo  	    true  		`商品信息`
// @Param order_total_money             raw  	float64  	false  		`订单总价`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /update_order_total_money [post]
func (this *PurchaseOrderController) UpdateOrderTotalMoney() {
	request := new(request.PurchaseOrderInfo)
	this.GetJsonStruct(request)

	err := order.UpdateOrderTotalMoney(request)

	if err != nil {
		helper.Log.ErrorString(this.Ctx.Input.URL() + err.Error())
		this.CheckError(err, "")
	}

	this.Success("", "操作成功")
}

// @Title  修改退款订单金额
// @Summary 修改退款订单金额
// @Description 修改退款订单金额
// @Param source        				raw  	string  	true  		`erp端默认是erp`
// @Param order_sn      				raw  	string  	true  		`订单编号`
// @Param goods_info       				raw  	[]request.GoodsOrderInfo  	    true  		`商品信息`
// @Param order_total_money             raw  	float64  	false  		`订单总价`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /update_return_total_money [post]
func (this *PurchaseOrderController) UpdateReturnOrderTotalMoney() {
	request := new(request.PurchaseOrderInfo)
	this.GetJsonStruct(request)

	err := order.UpdateReturnOrderTotalMoney(request)

	newErr, ok := err.(*error_message.NewError)

	if ok {
		helper.Log.ErrorString(this.Ctx.Input.URL() + newErr.Error())
		this.CheckError(errors.New(newErr.Error()), "")
	}

	this.Success("", "操作成功")
}

// @router /update_purchase_order_status [post]
func (this *PurchaseOrderController) UpdatePurchaseOrderStatus() {
	request := new(request.PurchaseOrderStatus)
	this.GetJsonStruct(request)

	err := order.UpdatePurchaseOrdersStatusById(request)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功")
}
