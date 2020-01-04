package openAccount

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/openAccount"
)

type OpenApiOpenAccountController struct {
	controllers.CommonController
}

// @router /add_open_account_info [post]
func (this *OpenApiOpenAccountController) AddOpenAccountInfo(){
	req := new(request.OpenAccount)
	this.GetJsonStruct(req)
	id,err := openAccount.OpenAccount(req)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(id,"操作成功")
}

// @router /delete_open_account_info [post]
func (this *OpenApiOpenAccountController) DeleteOpenAccountInfo(){
	req := new(request.DelOpenAccount)
	this.GetJsonStruct(req)
	err :=openAccount.DelOpenAccount(req.Id)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil,"操作成功")
}

// @router /update_open_account_info [post]
func (this *OpenApiOpenAccountController) UpdateOpenAccountInfo(){
	req := new(request.UpdateOpenAccount)
	this.GetJsonStruct(req)
	err := openAccount.UpdateOpenAccount(req)
	if err != nil {
		this.CheckError(err, "")
	}

	this.Success(nil,"操作成功")
}
