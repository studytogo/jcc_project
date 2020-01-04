package center_good

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/erp_goods_server"
)

type ErpGoodController struct {
	controllers.CommonController
}

// @Title  批量添加公司端商品
// @Summary 批量添加公司端商品
// @Description 一下参数是数组方式传送
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param name          raw  	string  	true  		`名称`
// @Param spu           raw  	string  	false  		`SPU`
// @Param sku           raw  	string  	false  		`SKU`
// @Param barcode       raw  	string  	true  		`条码`
// @Param spec          raw  	string  	false  		`规格`
// @Param buying_price  raw  	float64  	false  		`进货价`
// @Param retail_price  raw  	float64  	false  		`零售价`
// @Param inventory_upper_limit  raw  	string  	false  		`库存上限`
// @Param inventory_lower_limit  raw  	string  	false  		`库存下限`
// @Param mnemonic_word          raw  	string  	false  		`助记词`
// @Param remark        raw  	string  	false  		`备注`
// @Param image         raw  	string  	false  		`主图`
// @Param images        raw  	string  	false  		`多图`
// @Param content       raw  	string  	false  		`内容`
// @Param producing_province_id  raw  	int  	   false  		`产地省ID`
// @Param producing_city_id      raw  	int  	   false  		`产地城市ID`
// @Param producing_area_id      raw  	int  	   false  		`产地区域ID`
// @Param producing_area_detail  raw  	string     false  		`产地详情`
// @Param unit_id       raw  	int  	    false  		`单位ID`
// @Param brand_id      raw  	int  	    false  		`品牌ID`
// @Param a8_code       raw  	string  	false  		`a8编码`
// @Param is_cancel_procurement  raw  	int  	  false  		`是否取消采购`
// @Success 201 {object} []erp_goods.JccCenterGoods
// @Failure 403 body is empty
// @router /add_erp_goods_info [post]
func (this *ErpGoodController) AddErpGoodsInfo() {
	erpGoods := new(request.AddErpGoods)
	this.GetJsonStruct(erpGoods)

	err := erp_goods_server.AddErpGoods(erpGoods.ErpGoodsInfo)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(erpGoods, "操作成功")
}

// @Title  批量修改公司端商品
// @Summary 批量修改公司端商品
// @Description 一下参数是数组方式传送
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param id            raw  	id  	    true  		`商品Id`
// @Param name          raw  	string  	true  		`名称`
// @Param spu           raw  	string  	false  		`SPU`
// @Param sku           raw  	string  	false  		`SKU`
// @Param barcode       raw  	string  	true  		`条码`
// @Param spec          raw  	string  	false  		`规格`
// @Param buying_price  raw  	float64  	false  		`进货价`
// @Param retail_price  raw  	float64  	false  		`零售价`
// @Param inventory_upper_limit  raw  	string  	false  		`库存上限`
// @Param inventory_lower_limit  raw  	string  	false  		`库存下限`
// @Param mnemonic_word          raw  	string  	false  		`助记词`
// @Param remark        raw  	string  	false  		`备注`
// @Param image         raw  	string  	false  		`主图`
// @Param images        raw  	string  	false  		`多图`
// @Param content       raw  	string  	false  		`内容`
// @Param producing_province_id  raw  	int  	   false  		`产地省ID`
// @Param producing_city_id      raw  	int  	   false  		`产地城市ID`
// @Param producing_area_id      raw  	int  	   false  		`产地区域ID`
// @Param producing_area_detail  raw  	string     false  		`产地详情`
// @Param unit_id       raw  	int  	    false  		`单位ID`
// @Param brand_id      raw  	int  	    false  		`品牌ID`
// @Param a8_code       raw  	string  	false  		`a8编码`
// @Param is_cancel_procurement  raw  	int  	  false  		`是否取消采购`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /update_erp_goods_info [post]
func (this *ErpGoodController) UpdateErpGoodsInfo() {
	erpGoods := new(request.UpdateErpGoods)
	this.GetJsonStruct(erpGoods)

	err := erp_goods_server.UpdateErpGood(erpGoods.ErpGoodsInfo)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功")
}

// @Title  批量删除公司端商品
// @Summary 批量删除公司端商品
// @Description good_ids用“,”分割
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param good_ids      raw  	string  	true  		`商品id`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /delete_erp_goods_info [post]
func (this *ErpGoodController) DeleteErpGoodsInfo() {
	erpGoods := new(request.DeleteErpGoods)
	this.GetJsonStruct(erpGoods)

	err := erp_goods_server.DeleteErpGood(erpGoods.GoodIds)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil, "操作成功")
}

// @Title  批量查询公司端商品
// @Summary 批量查询公司端商品
// @Description good_ids，int数组
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param good_ids      raw  	[]int  	    true  		`商品id`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /query_erp_goods_info [post]
func (this *ErpGoodController) QueryErpGoodsInfo() {
	erpGoods := new(request.QueryErpGoods)
	this.GetJsonStruct(erpGoods)

	result, err := erp_goods_server.QueryErpGoodInfo(erpGoods.GoodIds)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(result, "操作成功")
}
