package goods

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/goodService"
)

type ErpGoodsOpsController struct {
	controllers.CommonController
}

// @Title  查询公司端商品信息
// @Summary 查询公司端商品信息
// @Description 查询公司端商品信息
// @Param token        raw  	string  	true  		`token信息`
// @Success 201 {object} []erp_goods.ErpGoodsInfo
// @Failure 403 body is empty
// @router /query_erp_goods_info_pc [post]
func (this *ErpGoodsOpsController) QueryErpGoodsInfoForPc() {
	erpGoods := new(request.AgentJinHuo)
	this.GetParamStruct(erpGoods)

	data, totol, err := goodService.QueryErpGoodsInfo(erpGoods)

	if err != nil {
		this.CheckError(err, "")
	}

	result := map[string]interface{}{
		"data":  data,
		"total": totol,
	}

	this.Success(result, "操作成功")
}

// @Title  同步公司端商品信息到加盟商
// @Summary 同步公司端商品信息到加盟商
// @Description 同步公司端商品信息到加盟商
// @Param token        		raw  	string  	true  		`token信息`
// @Param erp_id       		raw  	string  	true  		`加盟商商品id`
// @Param good_attribute    raw  	string  	true  		`同步加盟商商品属性`
// @Success 201 {object} []erp_goods.ErpGoodsInfo
// @Failure 403 body is empty
// @router /sync_erp_goods_info_pc [post]
func (this *ErpGoodsOpsController) SyncErpGoodsInfoForPc() {
	erpGoods := new(request.SyncErpToAgent)
	this.GetParamStruct(erpGoods)

	err := goodService.SyncErpGoods(erpGoods)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功")
}
