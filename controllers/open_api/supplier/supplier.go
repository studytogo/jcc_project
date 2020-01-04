package supplier

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/service/open_api/supplier"
)

type OpenApiSupplierOperationController struct {
	controllers.CommonController
}

// @Title  查询线上供应商信息
// @Summary 查询线上供应商信息
// @Description 查询线上供应商信息
// @Param token         raw  	string  	true  		`token值`
// @Param source        raw  	string  	true  		`请求来源`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /query_online_supplier_info [post]
func (this *OpenApiSupplierOperationController) QueryOnlineSupplierInfo() {
	data, err := supplier.QueryAllOnlineBrandInfo()
	if err != nil {
		this.CheckError(errors.New("查询线上供应商信息失败。。。"), "")
	}

	this.Success(data, "操作成功")
}
