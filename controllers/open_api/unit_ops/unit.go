package unit_ops

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/service/open_api/unit_ops"
)

type OpenApiUnitOperationController struct {
	controllers.CommonController
}

// @Title  查询所有单位信息
// @Summary 查询所有单位信息
// @Description 查询所有单位信息
// @Param source     raw  	string  	true  		`请求来源值`
// @Success 201 {object} unit.JccUnit
// @Failure 403 body is empty
// @router /query_all_unit_info [post]
func (this *OpenApiUnitOperationController) QueryAllUnitInfo() {
	//checkResq := new(request.CheckUnit)
	//this.GetParamStruct(checkResq)
	//this.CheckValidation(checkResq)
	//param := new(unit.CheckUnit)
	//err := helper.ReflectiveStruct(param, checkResq)
	//if err != nil {
	//	this.CheckError(err, "")
	//}
	//获取商品之前单位id
	data, err := unit_ops.QueryAllUnitInfo()
	if err != nil {
		this.CheckError(errors.New("查询单位信息失败。。。"), "")
	}

	this.Success(data, "操作成功")
}
