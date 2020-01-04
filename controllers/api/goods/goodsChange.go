package goods

import (
	"new_erp_agent_by_go/controllers"
)

type ChangeChild struct {
	controllers.CommonController
}

// @Title 				  查询返回修改项接口
// @Summary 			  查询返回修改项接口
// @Description 		  查询返回修改项接口
// @Param Token           raw  	string  	true  		`token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /goods_change [post]
func (this *ChangeChild) GetGoodsChange() {
	var data = map[string]string{
		"name":         "商品名称",
		"retail_price": "零售价",
		"unit":         "单位",
		"content":      "商品详情",
		"image":        "商品图片",
	}

	this.Success(data, "操作成功")

}
