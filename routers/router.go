package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"new_erp_agent_by_go/controllers/api/goods"
	"new_erp_agent_by_go/controllers/api/print"
	"new_erp_agent_by_go/controllers/api/purchaseOrder"
	"new_erp_agent_by_go/controllers/api/unit"
	"new_erp_agent_by_go/controllers/open_api/openAccount"
	"new_erp_agent_by_go/controllers/open_api/useRole"

	"new_erp_agent_by_go/controllers/open_api/companylist"

	"new_erp_agent_by_go/controllers/open_api/address"
	"new_erp_agent_by_go/controllers/open_api/brand"
	"new_erp_agent_by_go/controllers/open_api/center_good"
	"new_erp_agent_by_go/controllers/open_api/center_stock"

	"new_erp_agent_by_go/controllers/open_api/goodsCategory"
	"new_erp_agent_by_go/controllers/open_api/order"
	"new_erp_agent_by_go/controllers/open_api/supplier"
	"new_erp_agent_by_go/controllers/open_api/unit_ops"
	"new_erp_agent_by_go/filter"
	"new_erp_agent_by_go/helper"
)

func init() {

	ns := beego.NewNamespace("/api2",
		beego.NSNamespace("/goods",
			beego.NSInclude(
				&goods.SyntheticSplitController{},
				&goods.ChildGoodsController{},
				&goods.ErpGoodsOpsController{},
			),
		),
		beego.NSNamespace("/unit",
			beego.NSInclude(
				&unit.UnitOperationController{},
			),
		),
		beego.NSNamespace("/print",
			beego.NSInclude(
				&print.UserPrintMoudleController{},
				&print.GetPrintSettingsController{},
				&print.PrintLogController{},
				&print.QueryListController{},
				&print.UserPrintMoudleController{},
				&print.SelectPrintController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&purchaseOrder.PurchaseOrderController{},
			),
		),
		beego.NSNamespace("/goodsChange",
			beego.NSInclude(
				&goods.ChangeChild{},
			),
		),
		beego.NSNamespace("/getCenterGoods",
			beego.NSInclude(
				&goods.UpdateGoods{},
			),
		),
	)
	openApi := beego.NewNamespace("/openApi",
		beego.NSNamespace("/unit",
			beego.NSInclude(
				&unit_ops.OpenApiUnitOperationController{},
			),
		),
		beego.NSNamespace("/goods_category",
			beego.NSInclude(
				&goodsCategory.OpenApiGoodsCategoryOperationController{},
			),
		),

		beego.NSNamespace("/companylist",
			beego.NSInclude(
				&companylist.OpenApiCompanylistOperationController{},
			),
		),
		beego.NSNamespace("/online_supplier",
			beego.NSInclude(
				&supplier.OpenApiSupplierOperationController{},
			),
		),
		beego.NSNamespace("/online_brand",
			beego.NSInclude(
				&brand.OpenApiBrandOperationController{},
			),
		),
		beego.NSNamespace("/center_order",
			beego.NSInclude(
				&order.PurchaseOrderController{},
				&order.GiveOrderController{},
				&order.PurchaseReturnOrderController{},
			),
		),
		beego.NSNamespace("/center_goods",
			beego.NSInclude(
				&center_good.ErpGoodController{},
				&center_good.OnlineGoodsController{},
			),
		),
		beego.NSNamespace("/address",
			beego.NSInclude(
				&address.InfoController{},
			),
		),
		beego.NSNamespace("/erp_stock",
			beego.NSInclude(
				&center_stock.CenterStockController{},
			),
		),
		beego.NSNamespace("/openaccount",
			beego.NSInclude(
				&openAccount.OpenApiOpenAccountController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&useRole.UserController{},
			),
		),
	)
	beego.AddNamespace(ns, openApi)

	beego.InsertFilter("*", beego.BeforeRouter, filter.OptionsFilter)
	//将常用数据赋值提出来
	beego.InsertFilter("/api2/*", beego.BeforeRouter, filter.CommonlyParam)
	//可能暂时不会上权限验证
	openAuth, _ := beego.AppConfig.Bool("openAuth")
	if openAuth {
		helper.Log.Info("---------------------权限验证开启---------------------------")
		fmt.Println("---------------------权限验证开启---------------------------")
		beego.InsertFilter("/api2/*", beego.BeforeRouter, filter.CheckAuth)
	}

	//对外接口请求参数记录mysql
	beego.InsertFilter("/openApi/*", beego.BeforeRouter, filter.OpenApiRecord)
	// 特殊路由
	// beego.Router("/resource/*",&base.ResourceController{})
}
