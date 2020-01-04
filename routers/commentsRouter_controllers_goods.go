package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:ChildGoodsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:ChildGoodsController"],
        beego.ControllerComments{
            Method: "QueryChildGoods",
            Router: `/querychildgoods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:ChildGoodsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:ChildGoodsController"],
        beego.ControllerComments{
            Method: "SwitchParentOrChild",
            Router: `/switchparentorchild`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "AddChildGoods",
            Router: `/add_child_goods`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "GetGoodsUnitConversion",
            Router: `/goodsunitconversion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/goods:SyntheticSplitController"],
        beego.ControllerComments{
            Method: "Synthetic",
            Router: `/synthetic`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
