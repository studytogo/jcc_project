package center_stock

import (
	"encoding/json"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/actual"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/erp_stock_server"
)

type CenterStockController struct {
	controllers.CommonController
}

// @Title  批量添加公司端商品
// @Summary 批量添加公司端商品
// @Description 一下参数是数组方式传送
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param company_id    raw  	string  	true  		`公司id`
// @Param goods_id      raw  	string  	true  		`对应商品id`
// @Param option        raw  	string  	true  		`可订库存`
// @Success 201 {object} []erp_goods.JccCenterGoods
// @Failure 403 body is empty
// @router /add_update_erp_stock [post]
func (this *CenterStockController) AddOrUpdateErpStock() {
	param := new(request.AddOrUpdateErpStock)
	this.GetJsonStruct(param)

	err := erp_stock_server.AddOrUpdateErpStock(param.StockInfo, param.Source)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success("", "操作成功")
}

// @Title  批量添加公司端商品库存
// @Summary 批量添加公司端商品库存
// @Description 批量添加公司端商品库存
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param company_id    raw  	string  	true  		`公司id`
// @Param goods_id      raw  	string  	true  		`对应商品id`
// @Param num           raw  	string  	true  		`变化库存`
// @Success 201 {object} []erp_goods.JccCenterGoods
// @Failure 403 body is empty
// @router /change_erp_stock [post]
func (this *CenterStockController) ChangeErpStock() {
	param := new(request.AddOrUpdateErpStock)
	this.GetJsonStruct(param)

	err := erp_stock_server.ChangeErpStock(param.StockInfo, param.Source)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success("", "操作成功")
}

// @router /query_erp_stock [post]
func (this *CenterStockController) QueryErpStock() {
	checkreq := new(request.Page)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &checkreq); err != nil {
		this.Fail(err.Error())
	}
	list, count, last_page, err := erp_stock_server.QueryErpStock(checkreq)
	if err != nil {
		this.CheckError(err, "")
	}

	result := actual.T{}
	result.Data = list
	result.Total = count
	result.Last_page = last_page

	this.Success(result, "操作成功")
}
