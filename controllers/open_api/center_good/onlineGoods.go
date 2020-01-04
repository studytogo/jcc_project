package center_good

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api"
)

type OnlineGoodsController struct {
	controllers.CommonController
}

// @router /add_online_goods_info [post]
func (this *OnlineGoodsController) AddOnlineGoodsInfo() {
	erpGoods := new(request.AddOnlineGoodsInfo)
	this.GetJsonStruct(erpGoods)

	//
	err := api.AddOnlineGoodsInfo(erpGoods.GoodsIds,erpGoods.CompanyId)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil,"操作成功")
}