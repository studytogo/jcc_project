package request

//请求子商品信息
type ChildInfo struct {
	Id            string `valid:"Required" 父商品id`
	CompanyRoomId string `仓库Id`
}

type SwitchGoods struct {
	Id            int `商品id`
	IsParent      int `1是父商品 0不是`
	UnitId        int `单位Id`
	CompanyRoomId int `仓库Id`
}

type SwitchGoodsString struct {
	Id            string `valid:"Required" 商品id`
	IsParent      string `valid:"Required" 1是父商品 0不是`
	UnitId        string `valid:"Required" 单位Id`
	CompanyRoomId string `仓库Id`
}

type JccGoodsUnitConversion struct {
	Id           string ``
	GoodsId      string `form:"goods_id"`
	ChildGoodsId string `form:"child_goods_id"`
	CreatedAt    string ``
	UnitId       string `form:"unit_id"`
	Num          string `form:"num"`
	Mode         string `form:"mode"`
	LId          string `form:"l_id"`
}

type JccStock struct {
	Id        string ``
	GoodsId   string ``
	Actual    string ``
	Option    string ``
	LId       string ``
	UpdatedAt string ``
	Companyid string ``
}

type JccGoods struct {
	Name                string `form:"name"`
	Spu                 string `form:"spu"`
	Sku                 string `form:"sku"`
	Barcode             string `form:"barcode"`
	Spec                string `form:"spec"`
	BuyingPrice         string `form:"buying_price"`
	RetailPrice         string `form:"retail_price"`
	PredeterminedPrices string `form:"predetermined_prices"`
	InventoryUpperLimit string `form:"inventory_upper_limit"`
	InventoryLowerLimit string `form:"inventory_lower_limit"`
	MnemonicWord        string `form:"mnemonic_word"`
	Remark              string `form:"remark"`
	Image               string `form:"image"`
	Content             string `form:"content"`
	ProducingProvinceId string `form:"producing_province_id"`
	ProducingCityId     string `form:"producing_city_id"`
	ProducingAreaId     string `form:"producing_area_id"`
	ProducingAreaDetail string `form:"producing_area_detail"`
	UnitId              string `form:"unit_id"`
	BrandId             string `form:"brand_id"`
	CreatedAt           string `form:"created_at"`
	UpdatedAt           string `form:"updated_at"`
	DeletedAt           string `form:"deleted_at"`
	IsDel               string `form:"is_del"`
	IsCancelProcurement string `form:"is_cancel_procurement"`
	Kind                string `form:"kind"`
	Companyid           string `form:"Companyid"`
	Companyroomid       string `form:"companyroomid"`
	ErpId               string `form:"erp_id"`
	ErpSku              string `form:"erp_sku"`
	Images              string `form:"images"`
	IsSync              string `form:"is_sync"`
	IsParent            string `form:"is_parent"`
	ParentGoodsId       string `form:"parent_goods_id"`
	IsHaveChild         string `form:"is_have_child"`
	GoodsUnitConversion string `form:"goods_unit_conversion"`
	Uid                 string
	GoodsCategoryId     string `form:"goods_category_ids"`
	SupplierId          string `form:"supplier_ids"`
}

type JccUnit struct {
	Name      string `form:"name"`
	IsDel     string `form:"is_del"`
	UpdatedAt string `form:"updated_at"`
	Companyid string `form:"companyid"`
}

type JccBrand struct {
	Name      string `form:"name"`
	IsDel     string `form:"is_del"`
	UpdatedAt string `form:"updated_at"`
	Companyid string `form:"companyid"`
}

type Response struct {
	ChildStock  int64
	ParentStock int64
	IsHaveStock int64
}

type SpecInfo struct {
	Label string
	Value string
}

//修改单位或者新增商品时单位验证是否时国际单位
type CheckUnit struct {
	UnitId string `form:"unit_id" valid:"Required"`
	GoodId string `form:"good_id" valid:"Required"`
}

type UserPrintMoudle struct {
	MouldId        string `form:"mould_id"`
	TopMargin      string `form:"top_margin"`
	LeftMargin     string `form:"left_margin"`
	IsDisplayPrice string `form:"is_display_price"`
	CompanyId      string `form:"Companyid"`
}

type OrderIDs struct {
	OrderIDs string `form:"OrderIDs"`
	Page     string `form:"page"`
	Per_page string `form:"per_page"`
}
type JccCommonMould struct {
	Id          string `form:"id"`
	FieldOrder  string `form:"field_order"`
	MouldWeight string `form:"mould_weight"`
	MouldHeight string `form:"mould_height"`
	CreatedAt   string `form:"created_at"`
	UpdatedAt   string `form:"updated_at"`
	IsDefault   string `form:"is_default"`
	IsDel       string `form:"is_del"`
}

type OrderGoodsInfoReq struct {
	OrderIds string `form:"order_ids" valid:"Required" 订单id`
	PageParam
}

type PageParam struct {
	Page     string `form:"page" valid:"Required" 页数`
	PageSize string `form:"page_size" valid:"Required" 每页数量`
	IsPage   string `form:"is_page" 是否分页`
}

type ResponsePage struct {
	Data  interface{}
	Total int
}

type CenterGoods struct {
	CenterId    string `json:"center_id" valid:"Required"`
	Name        string `json:"name"`
	RetailPrice string `json:"retail_price"`
	UnitId      string `json:"unit_id"`
	Content     string `json:"content"`
	Image       string `json:"image"`
}

type AddOnlineGoodsInfo struct {
	GoodsIds  string `json:"goods_ids" valid:"Required" 商品id`
	CompanyId string `json:"company_id" 公司端公司id`
}

//代理商进货
type AgentJinHuo struct {
	CompanyId    string `form:"goods_companylist_id" 公司端公司id`
	GoodArribute string `form:"goods_search_word"  公司端商品姓名`
	PageParam
}

//同步公司端商品到加盟商
type SyncErpToAgent struct {
	ErpIds         string `form:"erp_ids" valid:"Required" 公司端商品id`
	GoodsAttribute string `form:"goods_attribute" valid:"Required" 公司端商品属性`
	Companyid      string `form:"Companyid"`
}
