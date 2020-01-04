package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"],
        beego.ControllerComments{
            Method: "AddGoodsCategory",
            Router: `/add_goods_category`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"],
        beego.ControllerComments{
            Method: "DeleteGoodsCategoryInfo",
            Router: `/deletegoodscategory`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"],
        beego.ControllerComments{
            Method: "EditGoodsCategoryInfo",
            Router: `/editgoodscategoryinfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/goodsCategory:OpenApiGoodsCategoryOperationController"],
        beego.ControllerComments{
            Method: "QueryAllGoodsCategoryInfo",
            Router: `/query_all_goods_category_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
