package unit

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/models/unit"
	"new_erp_agent_by_go/service/api/unitService"
)

type UnitOperationController struct {
	controllers.CommonController
}

// @Title  验证单位是否是国际换算单位
// @Summary 验证单位是否是国际换算单位
// @Description 验证单位是否是国际换算单位
// @Param unit_id          raw  	string  	true  		`更改后的单位Id`
// @Param good_id          raw  	string  	true  		`商品id`
// @Param token            raw  	string  	true  		`token值`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /checkunit [post]
func (this *UnitOperationController) CheckUnit() {
	checkResq := new(request.CheckUnit)
	this.GetParamStruct(checkResq)
	this.CheckValidation(checkResq)
	param := new(unit.CheckUnit)
	err := helper.ReflectiveStruct(param, checkResq)
	if err != nil {
		this.CheckError(err, "")
	}
	//获取商品之前单位id
	unitId, err := unit.QueryUnitIdByGoodsId(param.GoodId)

	if err != nil {
		this.CheckError(errors.New("查询商品的单位id不存在。。。"), "")
	}

	//将老单位传给前端
	response := map[string]interface{}{
		"old_unit_id": unitId,
	}

	//验证可以更改成国际单位
	err = unitService.CheckUnit(param)

	if err != nil {
		this.UnitCheckError(response, err)
	}

	//如果之前单位是国际单位，库存不为0也不让换
	param.UnitId = unitId
	err = unitService.CheckUnit(param)

	if err != nil {
		this.UnitCheckError(response, err)
	}

	this.Success(response, "操作成功")
}
