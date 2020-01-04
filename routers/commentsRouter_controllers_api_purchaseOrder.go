package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/purchaseOrder:PurchaseOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/purchaseOrder:PurchaseOrderController"],
        beego.ControllerComments{
            Method: "QueryOrderInfo",
            Router: `/query_order_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
