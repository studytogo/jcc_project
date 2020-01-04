package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/supplier:OpenApiSupplierOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/supplier:OpenApiSupplierOperationController"],
        beego.ControllerComments{
            Method: "QueryOnlineSupplierInfo",
            Router: `/query_online_supplier_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
