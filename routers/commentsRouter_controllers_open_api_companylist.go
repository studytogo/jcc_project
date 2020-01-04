package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"],
        beego.ControllerComments{
            Method: "AddCompanylistInfo",
            Router: `/add_companylist_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"],
        beego.ControllerComments{
            Method: "DeleteCompanylistInfo",
            Router: `/delete_companylist_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"],
        beego.ControllerComments{
            Method: "EditCompanylistInfo",
            Router: `/edit_companylist_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/companylist:OpenApiCompanylistOperationController"],
        beego.ControllerComments{
            Method: "QueryCompanylistInfo",
            Router: `/query_companylist_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
