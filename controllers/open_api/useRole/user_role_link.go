package useRole

import (
	"errors"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/user_server"
)

type UserController struct {
	controllers.CommonController
}

// @Title 				  增加用户权限表
// @Summary 			  增加用户权限表
// @Description 		  增加用户权限表
// @Param source           raw  	string  	true  		`数据来源`
// @Param data             raw  	[]user.JccUserRoleLink  	true  		`数据`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /add_user_role_link [post]
func (this *UserController) AddUserRoleLink() {

	param := new(request.UserRoleLink)

	this.GetJsonStruct(param)

	if param == nil {
		this.CheckError(errors.New("参数不能为空"), "")
	}

	err := user_server.AddUserRoleLink(param)

	if err != nil {
		this.CheckError(err, "")
	}

	this.Success("", "成功！！！")
}
