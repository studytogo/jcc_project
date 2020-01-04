package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/unit_ops:OpenApiUnitOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/unit_ops:OpenApiUnitOperationController"],
        beego.ControllerComments{
            Method: "QueryAllUnitInfo",
            Router: `/query_all_unit_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
