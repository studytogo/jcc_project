package purchaseOrder

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/purchaseOrderService"
)

type PurchaseOrderController struct {
	controllers.CommonController
}

// @Title  查询订单中的商品信息
// @Summary 查询订单中的商品信息
// @Description 查询订单中的商品信息
// @Param order_ids       raw  	string  	true  		`订单Id,用逗号分隔`
// @Param page        	  raw  	string  	true  		`当前页`
// @Param page_size       raw  	string  	true  		`每页数据量`
// @Param is_page         raw  	string  	true  		`是否分页 0不分页；1分页`
// @Param token           raw  	string  	true  	    `token值`
// @Success 0 {object} []childGoods.OrderGoodsInfo
// @Failure 1 body is empty
// @router /query_order_info [post]
func (this *PurchaseOrderController) QueryOrderInfo() {
	req := new(request.OrderGoodsInfoReq)
	this.GetParamStruct(req)
	info, total, err := purchaseOrderService.QueryOrderGoodsInfoByOrderId(req)
	if err != nil {
		this.CheckError(err, "")
	}
	result := new(request.ResponsePage)
	result.Data = info
	result.Total = total

	this.Success(result, "操作成功！！！！")

}
