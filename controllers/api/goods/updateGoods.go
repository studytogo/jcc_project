package goods

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/CenterGoods"
)

type UpdateGoods struct {
	controllers.CommonController
}

// @Title 				  更新商品
// @Summary 			  更新商品
// @Description 		  更新商品
// @Param center_id       raw  	string  	true  		`要更新的商品的center_goods_id字段，以逗号分隔的形式 eg: 1,2,3`
// @Param name        	  raw  	string  	false  		`商品名称`
// @Param retail_price    raw  	string  	false  		`零售价`
// @Param unit_id    	  raw  	string  	false  		`单位`
// @Param content         raw  	string  	false  		`商品详情`
// @Param image        	  raw  	string  	false  		`商品图片`
// @Param Token           raw  	string  	true  		`token`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /getCenterGoods [post]
func (this *UpdateGoods) UpdateGoods() {
	checkResq := new(request.CenterGoods)
	this.GetJsonStruct(checkResq)

	//获取地区
	err := CenterGoods.QueryUpdateGoods(checkResq)
	if err != nil {
		this.CheckError(errors.New("更新商品失败"), "")
	}

	this.Success("", "操作成功")

}
