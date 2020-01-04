package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/useRole:UserController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/useRole:UserController"],
        beego.ControllerComments{
            Method: "AddUserRoleLink",
            Router: `/add_user_role_link`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
