package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"],
        beego.ControllerComments{
            Method: "AddOrUpdateErpStock",
            Router: `/add_update_erp_stock`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"],
        beego.ControllerComments{
            Method: "ChangeErpStock",
            Router: `/change_erp_stock`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_stock:CenterStockController"],
        beego.ControllerComments{
            Method: "QueryErpStock",
            Router: `/query_erp_stock`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
