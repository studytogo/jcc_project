package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"],
        beego.ControllerComments{
            Method: "AddOpenAccountInfo",
            Router: `/add_open_account_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"],
        beego.ControllerComments{
            Method: "DeleteOpenAccountInfo",
            Router: `/delete_open_account_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/openAccount:OpenApiOpenAccountController"],
        beego.ControllerComments{
            Method: "UpdateOpenAccountInfo",
            Router: `/update_open_account_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
