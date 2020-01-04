package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"],
        beego.ControllerComments{
            Method: "AddErpGoodsInfo",
            Router: `/add_erp_goods_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"],
        beego.ControllerComments{
            Method: "DeleteErpGoodsInfo",
            Router: `/delete_erp_goods_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"],
        beego.ControllerComments{
            Method: "QueryErpGoodsInfo",
            Router: `/query_erp_goods_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:ErpGoodController"],
        beego.ControllerComments{
            Method: "UpdateErpGoodsInfo",
            Router: `/update_erp_goods_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:OnlineGoodsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/center_good:OnlineGoodsController"],
        beego.ControllerComments{
            Method: "AddOnlineGoodsInfo",
            Router: `/add_online_goods_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
