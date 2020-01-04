package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/brand:OpenApiBrandOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/brand:OpenApiBrandOperationController"],
        beego.ControllerComments{
            Method: "OpenApiBrandOperationController",
            Router: `/query_online_brand_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
