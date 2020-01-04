package request

import (
	"new_erp_agent_by_go/models/actual"
	pycompanylist "new_erp_agent_by_go/models/companylist"
	"new_erp_agent_by_go/models/erp_goods"
	"new_erp_agent_by_go/models/user"
)

type EditGoodsCategory struct {
	CategoryId   int    `json:"category_id" valid:"Required" 分类id`
	CategoryName string `json:"category_name" valid:"Required" 分类名称`
}

type DeleteGoodsCategory struct {
	CategoryId int `json:"category_id" valid:"Required" 分类id`
}

type Page struct {
	Page     int `json:"page" valid:"Required" 当前页数`
	Per_page int `json:"per_page" valid:"Required" 每页条数`
}

type JccCompanylist struct {
	Erp_company_info []pycompanylist.JccJicanchuCompanylist `json:"erp_company_info" `
}

type JccCompanylistId struct {
	Id string `json:"id" `
}

type AddGoodsCategory struct {
	CategoryName string `json:"name"  valid:"Required" 分类名称`
	Pid          int    `json:"pid"   父类id`
	Level        int    `json:"level" valid:"Required" 分类等级`
	No           string `json:"no"    valid:"Required" 编码`
}

//分类返回请求结构体
type CategoryRespense struct {
	CategoryId int64 `json:"category_id" description:"分类id"`
}

type GoodsCategory struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	No    string `json:"no"`
	Level int    `json:"level"`
	Pid   string `json:"pid"`
}

type PurchaseOrderInfo struct {
	OrderSn    string           `json:"order_sn"  valid:"Required" 订单编号`
	TotalMoney float64          `json:"order_total_money"  订单总价`
	GoodsInfo  []GoodsOrderInfo `json:"goods_info"  订单商品信息`
}

type GoodsOrderInfo struct {
	GoodId         int     `json:"good_id" 订单编号`
	GoodTotalMoney float64 `json:"good_total_money" 商品总价`
}

type Address struct {
	Name string
	Pid  string
}

type AddErpGoods struct {
	ErpGoodsInfo []*erp_goods.JccCenterGoods `json:"erp_goods_info" valid:"Required" erp商品信息`
}

type UpdateErpGoods struct {
	ErpGoodsInfo []*erp_goods.JccCenterGoods `json:"erp_goods_info" valid:"Required" erp商品信息`
}

type DeleteErpGoods struct {
	GoodIds string `json:"good_ids" valid:"Required" 商品id`
}

type QueryErpGoods struct {
	GoodIds []int `json:"good_ids" valid:"Required" 商品id`
}

type AddOrUpdateErpStock struct {
	StockInfo []actual.JccCompanyStock `json:"stock_info" valid:"Required" erp库存信息`
	Source    string                   `json:"source" valid:"Required" 请求来源`
}

type PurchaseOrderStatus struct {
	PurchaseOrderId int64 `json:"purchase_order_id" valid:"Required" 代理进货订单id`
	Status          int   `json:"status" valid:"Required" 状态`
}

type PurchaseReturnOrderStatus struct {
	PurchaseReturnOrderId int64 `json:"purchase_return_order_id" valid:"Required" 代理退货订单id`
	Status                int   `json:"status" valid:"Required" 状态`
}

//type Gift struct {
//
//	Gift map[int]int `json:"gift_id" valid:"Required" 赠品id`
//}

type GiveOrderStatus struct {
	PurchaseOrderId int64  `json:"purchase_order_id"  代理进货订单id`
	Fromorder       string `json:"fromorder" valid:"Required" 赠送单号`
	Customerid      int    `json:"customerid" valid:"Required" 审批人`
	BossId          int    `json:"bossid" valid:"Required" 获赠人`
	//CompanyroomId int `json:"companyroomid" valid:"Required" 仓库id`
	Telphone   string         `json:"telphone" valid:"Required" 电话`
	Address    string         `json:"address" valid:"Required" 地址`
	AffCompany int            `json:"affCompany" valid:"Required" 公司id`
	OrderMoney float64        `json:"ordermony"  订单金额`
	Remark     string         `json:"remark"  备注`
	Approver   string         `json:"approver" valid:"Required" 操作人`
	Gift       map[string]int `json:"gift"  赠品id`
	Status     int            `json:"status" valid:"Required" 状态`
}

// 入库单
type PutOrder struct {
	Id            int64
	RDh           string
	Fromorder     string
	Fromordertype int64
	RPeople       int64
	RStatus       int64
	OrderMoney    string
	Shop          int64
	House         int64
	House1        int64
	ShippPeople   string
	Freconst      int64
	Remark        string
	CreatedAt     int64
	UpdatedAt     int64
	DeletedAt     int64
	IsDel         int64
	Auditor       int64
	Operator      int64
	PaymentType   int64
	CustomerId    int64
}

//公司开户
type OpenAccount struct {
	Name      string `json:"name" valid:"Required"`
	Address   string `json:"address" valid:"Required"`
	Tel       string `json:"tel" valid:"Required"`
	Remarks   string `json:"remarks"`
	Companyid int    `json:"companyid" valid:"Required"`
	Province  int    `json:"province"  `
	City      int    `json:"city" `
	District  int    `json:"district" `
	Password  string `json:"password" valid:"Required"`
	Uid       int    `json:"uid" valid:"Required"`
	IdCard    string `json:"idcard" valid:"Required"`
}

//删除公司开户
type DelOpenAccount struct {
	Id int `json:"id" valid:"Required"`
}

//修改公司开户
type UpdateOpenAccount struct {
	Id       int    `json:"id" valid:"Required"`
	Password string `json:"password" `
	Tel      string `json:"tel" `
	Address  string `json:"address" `
	Province int    `json:"province"  `
	City     int    `json:"city" `
	District int    `json:"district" `
}

type UserRoleLink struct {
	UserRoleLink []user.JccUserRoleLink `json:"data" valid:"Required"`
}
