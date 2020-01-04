package companylist

import (
	"encoding/json"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	pycompanylist "new_erp_agent_by_go/models/companylist"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/companylist"
)

type OpenApiCompanylistOperationController struct {
	controllers.CommonController
}
// @Title  添加公司
// @Summary 添加公司
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param name          raw  	string  	true  		`公司名称`
// @router /add_companylist_info [post]
func (this *OpenApiCompanylistOperationController) AddCompanylistInfo() {
	var data request.JccCompanylist
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &data); err != nil {
		this.Fail(err.Error())
	}

	//自身验证排重
	err := companylist.CheckName(data.Erp_company_info)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "存在相同数据")
	}

	//数据库查重
	err = companylist.CheckDatabase(data.Erp_company_info)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "数据库存在相同数据")
	}
	err = companylist.AddCompanylist(data.Erp_company_info)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "添加失败")
	}
	this.Success(data.Erp_company_info, "操作成功")
}
// @Title  查询公司
// @Summary 查询公司
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param page          raw  	string  	true  		`页数`
// @Param per_page          raw  	string  	true  		`每页几条`
// @router /query_companylist_info [post]
func (this *OpenApiCompanylistOperationController) QueryCompanylistInfo() {
	checkreq := new(request.Page)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &checkreq); err != nil {
		this.Fail(err.Error())
	}
	list, count, per_last, err := companylist.QueryCompanylist(checkreq)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "查询失败")
	}

	result := pycompanylist.T{}
	result.Data = list
	result.Last_page = per_last
	result.Total = count

	this.Success(result, "操作成功")
}
// @Title   修改公司
// @Summary 修改公司
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param id            raw  	string  	true  		`公司id`
// @Param name          raw  	string  	true  		`公司名称`
// @router /edit_companylist_info [post]
func (this *OpenApiCompanylistOperationController) EditCompanylistInfo() {
	checkreq := new(pycompanylist.JccJicanchuCompanylist)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &checkreq); err != nil {
		this.Fail(err.Error())
	}

	err := companylist.EditCompanylist(*checkreq)

	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "修改失败")
	}
	this.Success(nil, "操作成功")
}
// @Title   删除公司
// @Summary 删除公司
// @Param source        raw  	string  	true  		`请求来源erp传erp`
// @Param id            raw  	string  	true  		`公司名称`
// @router /delete_companylist_info [post]
func (this *OpenApiCompanylistOperationController) DeleteCompanylistInfo() {
	checkreq := new(request.JccCompanylistId)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &checkreq); err != nil {
		this.Fail(err.Error())
	}

	err := companylist.DeleteCompanylist(*checkreq)
	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "删除失败")
	}
	this.Success(nil, "删除成功")
}
