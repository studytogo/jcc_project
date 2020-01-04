package goods

import (
	"fmt"
	"github.com/astaxie/beego"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/http_query"
	"new_erp_agent_by_go/models"
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api"
)

type SyntheticSplitController struct {
	controllers.CommonController
}

// @Title 				  子商品合成父商品 / 父商品拆分成子商品
// @Summary 			  子商品合成父商品 / 父商品拆分成子商品
// @Description 		  子商品合成父商品 / 父商品拆分成子商品
// @Param child_goods_id     raw  	string  	true  		`子商品的id`
// @Param goods_id           raw  	string  	true  		`父商品的id`
// @Param l_id               raw  	string  	true  		`仓库id`
// @Param num       	     raw  	string  	true  		`合成的父商品的数量`
// @Param mode               raw  	string  	true  		`拆分或合成的标识:synthesis => 子商品合成父商品; split => 父商品拆分成子商品`
// @Param Token              raw  	string  	true  		`token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /synthetic [post]
func (this *SyntheticSplitController) Synthetic() {
	var goodsUnitConversionString = new(request.JccGoodsUnitConversion)

	this.GetParamStruct(goodsUnitConversionString)

	var goodsUnitConversion = new(models.JccGoodsUnitConversion)
	err := helper.ReflectiveStruct(goodsUnitConversion, goodsUnitConversionString)

	if err != nil {
		helper.Log.Info("", err)
		this.CheckError(err, "")
	}
	// 子商品合成父商品
	var response = new(request.Response)
	if goodsUnitConversion.Mode == "synthesis" {
		childStrock, parentStrock, isHavedStock, err := api.SelectConversionChild(goodsUnitConversion)
		if err != nil {
			helper.Log.Error("", err)
			this.CheckError(err, "")
		}
		response.ChildStock = childStrock
		response.ParentStock = parentStrock
		response.IsHaveStock = isHavedStock
		this.Success(response, "操作成功!")
	}

	// 父商品拆分成子商品
	if goodsUnitConversion.Mode == "split" {
		childStrock, parentStrock, isHavedStock, err := api.SelectConversionParent(goodsUnitConversion)
		if err != nil {
			helper.Log.Error("", err)
			this.CheckError(err, "")
		}
		response.ChildStock = childStrock
		response.ParentStock = parentStrock
		response.IsHaveStock = isHavedStock
		this.Success(response, "操作成功!")
	}
}

// @Title 		添加子商品
// @Summary 	添加子商品
// @Description 添加子商品  	IsParent IsHaveChild   这两个字段不用传
// @Param id        				raw  	string  	    false  		`ID`
// @Param name        				raw  	string  	    true  		`名称`
// @Param spu        				raw  	string  	    false  		`SPU`
// @Param sku        				raw  	string  	    false  		`SKU`
// @Param barcode        			raw  	string  	    true  		`条码`
// @Param spec        				raw  	string  	    true  		`规格`
// @Param buying_price        		raw  	string  	    true  		`进货价`
// @Param retail_price        		raw  	string  	    true  		`零售价`
// @Param predetermined_prices      raw  	string  	    true  		`预售价格集合 字符串,分隔`
// @Param inventory_upper_limit     raw  	string  	    true  		`库存上限`
// @Param inventory_lower_limit     raw  	string  	    true  		`库存下限`
// @Param mnemonic_word        		raw  	string  	    true  		`助记词`
// @Param remark        			raw  	string  	    true  		`备注`
// @Param image        				raw  	string  	    true  		`主图`
// @Param content        			raw  	string  	    true  		`内容`
// @Param producing_province_id     raw  	string  	    true  		`产地省ID`
// @Param producing_city_id        	raw  	string  	    true  		`产地城市ID`
// @Param producing_area_id        	raw  	string  	    true  		`产地区域ID`
// @Param producing_area_detail     raw  	string  	    true  		`产地详情`
// @Param unit_id        			raw  	string  	    true  		`单位ID`
// @Param brand_id        			raw  	string  	    true  		`品牌ID`
// @Param created_at        		raw  	string  	    true  		`创建时间`
// @Param updated_at        		raw  	string  	    true  		`修改时间`
// @Param deleted_at        		raw  	string  	    true  		`删除时间`
// @Param is_del        			raw  	string  	    true  		`是否删除`
// @Param is_cancel_procurement     raw  	string  	    true  		`是否取消采购`
// @Param kind        				raw  	string  	    true  		`商品的类（本地(offline)、自营(online)）默认的是本地(offline)`
// @Param companyid        			raw  	string  	    true  		`公司id`
// @Param companyroomid        		raw  	string  	    true  		`公司仓库或门店id`
// @Param erp_id        			raw  	string  	    true  		`线上ERP2.0的商品ID`
// @Param erp_sku        			raw  	string  	    true  		`erp的sku`
// @Param images        			raw  	string  	    true  		`多图`
// @Param is_sync        			raw  	string  	    true  		`是否同步过其它端`
// @Param is_parent        			raw  	string  	    false  		`是否父商品 （0不是 1是 默认1）`
// @Param parent_goods_id        	raw  	string  	    false  		`父商品id`
// @Param is_have_child        		raw  	string  	    false  		`是否存在子商品 （0不存在 1存在 默认0）`
// @Param goods_unit_conversion   	raw  	string  	    true  		`商品单位换算 单位id|数量|父商品id`
// @Param token        				raw  	string  	    true  		`Token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /add_child_goods [post]
func (this *SyntheticSplitController) AddChildGoods() {
	var goodsString = new(request.JccGoods)
	this.GetParamStruct(goodsString)

	var goods = new(childGoods.JccGoods)
	err := helper.ReflectiveStruct(goods, goodsString)

	goodsBack, _, err := api.AddChildGoods(goods)

	if err != nil {
		this.CheckError(err, "")
	}

	isOpen, _ := beego.AppConfig.Bool("agent_sync")
	if isOpen {
		go func() {
			//城市电商添加商品
			//是否开启城市电商
			//判断是否开启连接挪到父类中进行
			//处理批量删除参数
			var params = make(map[string]string)
			params["title"] = goodsBack.Name                           //商品名称
			params["type"] = "1"                                       //商品类型 1 实体物品 2 虚拟物品 3 虚拟物品(卡密)
			params["merch_id"] = fmt.Sprint(goods.Companyid)           //店铺id
			params["unit"] = goodsBack.UnitName                        //单位
			params["thumb"] = goodsBack.Image                          //商品图
			params["market_price"] = fmt.Sprint(goodsBack.RetailPrice) //现价
			params["total"] = "0"                                      //库存
			params["sku_code"] = goodsBack.Sku                         //sku码
			params["come_where"] = "0"                                 //0自营1直营
			params["thumb_url"] = goodsBack.Images                     //相册图片路径json(最多6张)
			params["content"] = goodsBack.Content
			httpQuery := http_query.NewInit("agent")
			httpQuery.SetRunMode(beego.AppConfig.String("runmode"))
			httpQuery.SetDingWebHookUrl(beego.AppConfig.String("ding_web_hook_url")) //详情
			agentUrl := beego.AppConfig.String("agent_url")

			// 请求城市电商1.0
			httpQuery = http_query.NewInit("agent")
			httpQuery.AgentPost(agentUrl+"goods/add_goods", params)
			//if res["code"].(string) != "100000" {
			//	// 错误记录
			//	// this.CheckError(err, "同步添加城市电商子商品失败")
			//	helper.Log.ErrorString("同步添加城市电商子商品失败" + res["msg"].(string))
			//}
		}()
	}

	this.Success(goodsBack, "成功！")
}

// @Title 				  获取父子单位关系
// @Summary 			  获取父子单位关系
// @Description 		  获取父子单位关系
// @Param child_goods_id  raw  	string  	true  		`子商品的id`
// @Param goods_id        raw  	string  	true  		`父商品的id`
// @Param mode            raw  	string  	true  		`0  通过goods_id  找子商品的单位；1 通过child_goods_id 找父商品的单位`
// @Param Token           raw  	string  	true  		`token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /goodsunitconversion [post]
func (this *SyntheticSplitController) GetGoodsUnitConversion() {
	var goodsUnitConversionString = new(request.JccGoodsUnitConversion)
	this.GetParamStruct(goodsUnitConversionString)

	var goodsUnitConversion = new(models.JccGoodsUnitConversion)
	err := helper.ReflectiveStruct(goodsUnitConversion, goodsUnitConversionString)
	if err != nil {
		helper.Log.Info("")
	}

	if goodsUnitConversion.Mode == "0" {
		info, err := api.SelectParentGoodsUnitConversion(goodsUnitConversion)
		if err != nil {
			helper.Log.Info("")
		}
		this.Success(info, "成功！")
	}

	if goodsUnitConversion.Mode == "1" {
		info, err := api.SelectGoodsUnitConversion(goodsUnitConversion)
		if err != nil {
			helper.Log.Info("")
		}
		this.Success(info, "成功！")
	}

}
