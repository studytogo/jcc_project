package goodsCategory

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/goodsCategory"
)

type OpenApiGoodsCategoryOperationController struct {
	controllers.CommonController
}

// @Title  查询所有商品分类信息
// @Summary 查询所有商品分类信息
// @Description 查询所有商品分类信息
// @Param token         raw  	string  	true  		`token值`
// @Param source        raw  	string  	true  		`请求来源`
// @Param name          raw  	string  	false  		`搜索关键子名字`
// @Param no            raw  	string  	false  		`搜索关键子no`
// @Param pid           raw  	string  	false  		`搜索父id为这个的分类`
// @Success 201 {object} categoryIds.JccGoodsCategory
// @Failure 403 body is empty
// @router /query_all_goods_category_info [post]
func (this *OpenApiGoodsCategoryOperationController) QueryAllGoodsCategoryInfo() {
	checkResq := new(request.GoodsCategory)
	this.GetJsonStruct(checkResq)
	//param := new(categoryIds.JccGoodsCategory)
	//err := helper.ReflectiveStruct(param, checkResq)
	//if err != nil {
	//	this.CheckError(err, "")
	//}
	//获取所有商品分类
	data, err := goodsCategory.QueryAllGoodsCategoryInfo(checkResq)
	if err != nil {
		this.CheckError(errors.New("查询商品分类信息失败。。。"), "")
	}

	this.Success(data, "操作成功")
}


// @router /editgoodscategoryinfo [post]
func (this *OpenApiGoodsCategoryOperationController) EditGoodsCategoryInfo() {
	//接收检查数据
	checkResq := new(request.EditGoodsCategory)
	this.GetJsonStruct(checkResq)

	//发送service层
	err := goodsCategory.EditGoodsCategory(checkResq)
	if err != nil {
		this.CheckError(err, "修改商品分类失败")
	}

	this.Success(nil, "操作成功~！！！")

}

// @router /deletegoodscategory [post]
func (this *OpenApiGoodsCategoryOperationController) DeleteGoodsCategoryInfo() {
	//接收检查数据
	checkResq := new(request.DeleteGoodsCategory)
	this.GetJsonStruct(checkResq)

	//发送service层
	err := goodsCategory.DeleteGoodsCategory(checkResq)
	if err != nil {
		this.CheckError(err, "删除商品分类失败")
	}

	this.Success(nil, "操作成功~！！！")

}
// @Title  增加商品分类
// @Summary 增加商品分类
// @Description source字段，erp传erp,其它传其它的值
// @Param name            raw  	string  	true  		`分类名称`
// @Param pid             raw  	string  	    		`父类id`
// @Param level           raw  	string  	true  		`分类等级`
// @Param no              raw  	string  	true  		`分类编号`
// @Param source          raw  	string  	true  		`请求来源`
// @Success 201 {object} request.CategoryRespense
// @Failure 403 body is empty
// @router /add_goods_category [post]
func (this *OpenApiGoodsCategoryOperationController) AddGoodsCategory() {
	//使用json接受数据
	param := new(request.AddGoodsCategory)

	this.GetJsonStruct(param)

	id, err := goodsCategory.AddGoodCategory(param)

	if err != nil {
		this.CheckError(err, "")
	}

	response := request.CategoryRespense{
		CategoryId: id,
	}

	this.Success(response, "操作成功")
}
