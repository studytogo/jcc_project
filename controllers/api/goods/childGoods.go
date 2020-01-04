package goods

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/childGoodService"
	"strconv"
)

type ChildGoodsController struct {
	controllers.CommonController
}

// @Title  查询子商品信息
// @Summary 查询子商品信息
// @Description 查询子商品信息
// @Param Id               raw  	string  	true  		`父商品的id`
// @Param CompanyRoomId    raw  	string  	  		    `仓库id`
// @Param token            raw  	string  	true  		`token值`
// @Success 201 {object} []childGoods.JccGoodsDetailInfo
// @Failure 403 body is empty
// @router /querychildgoods [post]
func (this *ChildGoodsController) QueryChildGoods() {
	fatherId := new(request.ChildInfo)
	this.GetParamStruct(fatherId)
	this.CheckValidation(fatherId)
	IdInt, _ := strconv.Atoi(fatherId.Id)
	companyRoomId, _ := strconv.Atoi(fatherId.CompanyRoomId)
	childInfo, err := childGoodService.QueryChildInfoByParentId(IdInt, companyRoomId)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "")
	}

	this.Success(childInfo, "操作成功~！！！")
}

// @Title  切换子父商品信息
// @Summary 切换子父商品信息
// @Description 切换子父商品信息
// @Param Id              raw  	string  	true  		`商品的id`
// @Param IsParent        raw  	string  	true  		`1是父商品,0不是`
// @Param UnitId          raw  	string  	true  		`商品单位的id`
// @Param CompanyRoomId   raw  	string  	  		    `仓库id`
// @Param token           raw  	string  	true  	    `token值`
// @Success 201 {object} []childGoods.JccGoods
// @Failure 403 body is empty
// @router /switchparentorchild [post]
func (this *ChildGoodsController) SwitchParentOrChild() {
	var requestParam = new(request.SwitchGoodsString)
	this.GetParamStruct(requestParam)
	this.CheckValidation(requestParam)
	var param = new(request.SwitchGoods)
	err := helper.ReflectiveStruct(param, requestParam)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "数据转换失败", err)
		this.CheckError(err, "")
	}
	resultList, err := childGoodService.SwitchParentOrChild(param)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "")
	}

	this.Success(resultList, "操作成功~！！！")
}
