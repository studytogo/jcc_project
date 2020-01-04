package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:GiveOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:GiveOrderController"],
        beego.ControllerComments{
            Method: "UpdateGiveOrderStatus",
            Router: `/give_order_status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"],
        beego.ControllerComments{
            Method: "UpdateOrderTotalMoney",
            Router: `/update_order_total_money`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"],
        beego.ControllerComments{
            Method: "UpdatePurchaseOrderStatus",
            Router: `/update_purchase_order_status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseOrderController"],
        beego.ControllerComments{
            Method: "UpdateReturnOrderTotalMoney",
            Router: `/update_return_total_money`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseReturnOrderController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/order:PurchaseReturnOrderController"],
        beego.ControllerComments{
            Method: "UpdatePurchaseReturnOrderStatus",
            Router: `/update_purchase_return_order_status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
