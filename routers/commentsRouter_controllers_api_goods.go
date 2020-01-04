package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChangeChild"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChangeChild"],
        beego.ControllerComments{
            Method: "GetGoodsChange",
            Router: `/goods_change`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChildGoodsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChildGoodsController"],
        beego.ControllerComments{
            Method: "QueryChildGoods",
            Router: `/querychildgoods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChildGoodsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ChildGoodsController"],
        beego.ControllerComments{
            Method: "SwitchParentOrChild",
            Router: `/switchparentorchild`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ErpGoodsOpsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ErpGoodsOpsController"],
        beego.ControllerComments{
            Method: "QueryErpGoodsInfoForPc",
            Router: `/query_erp_goods_info_pc`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ErpGoodsOpsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:ErpGoodsOpsController"],
        beego.ControllerComments{
            Method: "SyncErpGoodsInfoForPc",
            Router: `/sync_erp_goods_info_pc`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "AddChildGoods",
            Router: `/add_child_goods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "GetGoodsUnitConversion",
            Router: `/goodsunitconversion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "Synthetic",
            Router: `/synthetic`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:UpdateGoods"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/api/goods:UpdateGoods"],
        beego.ControllerComments{
            Method: "UpdateGoods",
            Router: `/getCenterGoods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
