package brand

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/service/open_api/brand"
)

type OpenApiBrandOperationController struct {
	controllers.CommonController
}

// @Title  查询线上品牌信息
// @Summary 查询线上品牌信息
// @Description 查询线上品牌信息
// @Param token         raw  	string  	true  		`token值`
// @Param source        raw  	string  	true  		`请求来源`
// @Success 201 {string}
// @Failure 403 body is empty
// @router /query_online_brand_info [post]
func (this *OpenApiBrandOperationController) OpenApiBrandOperationController() {
	data, err := brand.QueryAllOnlineBrandInfo()
	if err != nil {
		this.CheckError(errors.New("查询线上品牌信息失败。。。"), "")
	}

	this.Success(data, "操作成功")
}
