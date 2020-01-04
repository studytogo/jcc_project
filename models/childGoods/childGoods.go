package childGoods

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/categoryIds"
	"new_erp_agent_by_go/models/request"
	"strconv"
)

type JccGoods struct {
	Id                  int64   `orm:"column(id);auto" description:"ID"`
	Name                string  `orm:"column(name);size(150)" description:"名称"`
	Spu                 string  `orm:"column(spu);size(150)" description:"SPU"`
	Sku                 string  `orm:"column(sku);size(150)" description:"SKU"`
	Barcode             string  `orm:"column(barcode);size(150)" description:"条码"`
	Spec                string  `orm:"column(spec);size(150)" description:"规格"`
	BuyingPrice         float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"进货价"`
	RetailPrice         float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价"`
	PredeterminedPrices string  `orm:"column(predetermined_prices);size(255)" description:"预售价格集合 字符串,分隔"`
	InventoryUpperLimit string  `orm:"column(inventory_upper_limit);size(255);null" description:"库存上限"`
	InventoryLowerLimit string  `orm:"column(inventory_lower_limit)" description:"库存下限"`
	MnemonicWord        string  `orm:"column(mnemonic_word);size(255)" description:"助记词"`
	Remark              string  `orm:"column(remark);size(150)" description:"备注"`
	Image               string  `orm:"column(image);size(150)" description:"主图"`
	Content             string  `orm:"column(content)" description:"内容"`
	ProducingProvinceId int     `orm:"column(producing_province_id)" description:"产地省ID"`
	ProducingCityId     int     `orm:"column(producing_city_id)" description:"产地城市ID"`
	ProducingAreaId     int     `orm:"column(producing_area_id)" description:"产地区域ID"`
	ProducingAreaDetail string  `orm:"column(producing_area_detail);size(150)" description:"产地详情"`
	UnitId              int     `orm:"column(unit_id)" description:"单位ID"`
	BrandId             int     `orm:"column(brand_id)" description:"品牌ID"`
	CreatedAt           int64   `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt           int64   `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt           int     `orm:"column(deleted_at)" description:"删除时间"`
	IsDel               int8    `orm:"column(is_del)" description:"是否删除"`
	IsCancelProcurement int8    `orm:"column(is_cancel_procurement)" description:"是否取消采购"`
	Kind                string  `orm:"column(kind);size(20);null" description:"商品的类（本地(offline)、自营(online)）默认的是本地(offline)"`
	Companyid           int     `orm:"column(companyid)" description:"公司id"`
	Companyroomid       int     `orm:"column(companyroomid)" description:"公司仓库或门店id"`
	ErpId               int     `orm:"column(erp_id);null" description:"线上ERP2.0的商品ID"`
	ErpSku              string  `orm:"column(erp_sku);size(150)" description:"erp的sku"`
	Images              string  `orm:"column(images)" description:"多图"`
	IsSync              int     `orm:"column(is_sync);null" description:"是否同步过其它端"`
	IsParent            int8    `orm:"column(is_parent)" description:"是否父商品 （0不是 1是 默认1）"`
	ParentGoodsId       uint    `orm:"column(parent_goods_id);null" description:"子商品id"`
	IsHaveChild         int8    `orm:"column(is_have_child)" description:"是否存在子商品 （0不存在 1存在 默认0）"`
	GoodsUnitConversion string  `orm:"-" json:"goods_unit_conversion"`
	GoodsCategoryId     int64   `orm:"-"`
	SupplierId          int64   `orm:"-"`
	UnitName            string  `orm:"-"`
	CenterGoodsId       int     `orm:"-"`
}

type JccGoodsDetailInfo struct {
	Id                      int64   `orm:"column(id);auto" description:"ID"`
	Name                    string  `orm:"column(name);size(150)" description:"名称"`
	Spu                     string  `orm:"column(spu);size(150)" description:"SPU"`
	Sku                     string  `orm:"column(sku);size(150)" description:"SKU"`
	Barcode                 string  `orm:"column(barcode);size(150)" description:"条码"`
	Spec                    string  `orm:"column(spec);size(150)" description:"规格"`
	BuyingPrice             float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"进货价"`
	RetailPrice             float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价"`
	PredeterminedPrices     string  `orm:"column(predetermined_prices);size(255)" description:"预售价格集合 字符串,分隔"`
	InventoryUpperLimit     string  `orm:"column(inventory_upper_limit);size(255);null" description:"库存上限"`
	InventoryLowerLimit     string  `orm:"column(inventory_lower_limit)" description:"库存下限"`
	MnemonicWord            string  `orm:"column(mnemonic_word);size(255)" description:"助记词"`
	Remark                  string  `orm:"column(remark);size(150)" description:"备注"`
	Image                   string  `orm:"column(image);size(150)" description:"主图"`
	Content                 string  `orm:"column(content)" description:"内容"`
	ProducingProvinceId     int     `orm:"column(producing_province_id)" description:"产地省ID"`
	ProducingCityId         int     `orm:"column(producing_city_id)" description:"产地城市ID"`
	ProducingAreaId         int     `orm:"column(producing_area_id)" description:"产地区域ID"`
	ProducingAreaDetail     string  `orm:"column(producing_area_detail);size(150)" description:"产地详情"`
	UnitId                  int     `orm:"column(new_unit_id)" description:"单位ID"`
	BrandId                 int     `orm:"column(brand_id)" description:"品牌ID"`
	CreatedAt               int     `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt               int     `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt               int     `orm:"column(deleted_at)" description:"删除时间"`
	IsDel                   int8    `orm:"column(is_del)" description:"是否删除"`
	IsCancelProcurement     int8    `orm:"column(is_cancel_procurement)" description:"是否取消采购"`
	Kind                    string  `orm:"column(type);size(20);null" description:"商品的类（本地(offline)、自营(online)）默认的是本地(offline)"`
	Companyid               int     `orm:"column(companyid)" description:"公司id"`
	Companyroomid           int     `orm:"column(companyroomid)" description:"公司仓库或门店id"`
	ErpId                   int     `orm:"column(erp_id);null" description:"线上ERP2.0的商品ID"`
	ErpSku                  string  `orm:"column(erp_sku);size(150)" description:"erp的sku"`
	Images                  string  `orm:"column(images)" description:"多图"`
	IsSync                  int     `orm:"column(is_sync);null" description:"是否同步过其它端"`
	IsParent                int8    `orm:"column(is_parent)" description:"是否父商品 （0不是 1是 默认1）"`
	ParentGoodsId           uint    `orm:"column(parent_goods_id);null" description:"子商品id"`
	IsHaveChild             int8    `orm:"column(is_have_child)" description:"是否存在子商品 （0不存在 1存在 默认0）"`
	ActualInt               int     `orm:"column(actual)" description:"商品库存"`
	Actual                  string  `description:"商品库存显示字段"`
	Brand                   string  `orm:"column(brand_name)" description:"商品品牌"`
	L_id                    string  `orm:"column(company_name)" description:"仓库名称"`
	Unit                    string  `orm:"column(unit_name)" description:"单位名称"`
	Total_money             float64 `orm:"column(total)" description:"商品总额"`
	Num                     int     `orm:"column(num)" description:"单位转换数量"`
	CategoryId              int     `orm:"column(categoryId)" description:"分类id"`
	Redetermined_price_list []map[string]string
	ParentUnitId            int `description:"父商品id"`
	SupperlierInfo          map[string]interface{}
	CategoryIds             []categoryIds.JccGoodsCategory
	Croomid                 int    `orm:"column(croom_id)" description:"父商品id"`
	Allactual               string `description:"该商品所有库存"`
	IsHaveActual            int    `description:"该商品所有库存:1代表有过，0代表没有过"`
	IsUse                   int    `orm:"column(is_use)" description:"多图" json:"is_use"`
	StockId                 int    `orm:"column(stock_id)" description:"仓库id" json:"stock_id"`
}

type OrderGoodsInfo struct {
	Id            int64   `orm:"column(id);auto" description:"ID" json:"id"`
	Name          string  `orm:"column(name);size(150)" description:"名称" json:"name"`
	Barcode       string  `orm:"column(barcode);size(150)" description:"条码" json:"bar_code"`
	Spec          string  `orm:"column(spec);size(150)" description:"规格" json:"spec"`
	UnitId        int     `orm:"column(unit_id)" description:"单位ID" json:"unit_id"`
	RetailPrice   float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价" json:"retail_price"`
	GoodNum       float64 `orm:"column(num)" description:"订单中该商品的数量" json:"good_num"`
	IsParent      int     `orm:"column(is_parent)" description:"是否父商品 （0不是 1是 默认1）" json:"is_parent"`
	OrderId       int     `orm:"column(order_id)"  description:"订单id" json:"order_id"`
	IsHaveChild   int     `orm:"column(is_have_child)" description:"是否有孩子" json:"is_have_child"`
	UnitName      string  `orm:"column(-)" description:"单位名称" json:"unit_name"`
	CompanyRoomId int     `orm:"column(-)" description:"单位名称" json:"companyroom_id"`
}

func init() {
	orm.RegisterModel(new(JccGoods))
}

func QueryChildIdByParentId(parentId int, db orm.Ormer) ([]int, error) {

	var childId []int

	sql := `SELECT child_goods_id from jcc_goods_unit_conversion WHERE goods_id = ? `

	_, err := db.Raw(sql, parentId).QueryRows(&childId)

	return childId, err
}

//查询父或者子商品的所有信息，有些信息只查询了对应的id
func QueryChildOrParentGoodInfoById(id int, param *request.SwitchGoods, db orm.Ormer) (JccGoodsDetailInfo, error) {

	var childInfo JccGoodsDetailInfo

	sql := ``

	//在jcc_goods表中子商品的unit_id变化有问题，所以只能使用jcc_goods_unit_conversion表中的unit_id作为子商品id，但是父商品的unit_id是没问题的
	if param.IsParent == 0 {
		sql = `SELECT
				jg.*, js.actual,
				jb.name AS brand_name,
				jc.name AS company_name,
				ju.name AS unit_name,
				jguc.num as num,
                ju.is_use as is_use,
                jgci.goods_category_id as categoryId,
                jg.unit_id as new_unit_id,
				CASE jg.kind
			WHEN "offline" THEN
				"本地"
			ELSE
				"自营"
			END AS type,
			 jg.buying_price * js.actual AS total
			FROM
				(
					SELECT
						*
					FROM
						jcc_goods
					WHERE
						id = ?
				and is_del = 0
				) jg
			LEFT JOIN jcc_brand jb ON jg.brand_id = jb.id
			LEFT JOIN jcc_stock js ON jg.id = js.goods_id and js.l_id = ?
			LEFT JOIN jcc_companyroom jc ON js.l_id = jc.id
			LEFT JOIN jcc_unit ju ON jg.unit_id = ju.id
			LEFT JOIN jcc_goods_unit_conversion jguc ON jg.id = jguc.goods_id
            LEFT JOIN jcc_goods_category_index jgci ON jg.id = jgci.goods_id `

	} else {
		sql = `SELECT
				jg.*, js.actual,
				jb.name AS brand_name,
				jc.name AS company_name,
				ju.name AS unit_name,
                ju.is_use as is_use, 
				jguc.num as num,
                jgci.goods_category_id as categoryId,
                jguc.unit_id as new_unit_id,
				CASE jg.kind
			WHEN "offline" THEN
				"本地"
			ELSE
				"自营"
			END AS type,
			 jg.buying_price * js.actual AS total
			FROM
				(
					SELECT
						*
					FROM
						jcc_goods
					WHERE
						id = ?
				and is_del = 0
				) jg
			LEFT JOIN jcc_brand jb ON jg.brand_id = jb.id
			LEFT JOIN jcc_stock js ON jg.id = js.goods_id and js.l_id = ?
			LEFT JOIN jcc_companyroom jc ON js.l_id = jc.id
			LEFT JOIN jcc_unit ju ON jg.unit_id = ju.id
            LEFT JOIN jcc_goods_category_index jgci ON jg.id = jgci.goods_id
			LEFT JOIN jcc_goods_unit_conversion jguc ON jg.id = jguc.child_goods_id and jg.unit_id = ` + strconv.Itoa(param.UnitId)
	}

	err := db.Raw(sql, id, param.CompanyRoomId).QueryRow(&childInfo)

	return childInfo, err
}

//通过id查询商品的信息，部分信息只有对应id
func QueryChildGoodInfoById(childId int, companyId int, db orm.Ormer) (JccGoodsDetailInfo, error) {

	var childInfo JccGoodsDetailInfo

	//在jcc_goods表中子商品的unit_id变化有问题，所以只能使用jcc_goods_unit_conversion表中的unit_id作为子商品id
	sql := `SELECT
				jg.*, js.actual,
				jb.name AS brand_name,
				jc.name AS company_name,
				ju.name AS unit_name,
				jguc.num as num,
                ju.is_use as is_use,
                jgci.goods_category_id as categoryId,
                jg.unit_id as new_unit_id,
                js.l_id as croom_id,
				js.id   as stock_id,
				CASE jg.kind
			WHEN "offline" THEN
				"本地"
			ELSE
				"自营"
			END AS type,
			 jg.buying_price * js.actual AS total
			FROM
				(
					SELECT
						*
					FROM
						jcc_goods
					WHERE
						id = ?
				and is_del = 0
				) jg
			LEFT JOIN jcc_brand jb ON jg.brand_id = jb.id
			LEFT JOIN jcc_stock js ON jg.id = js.goods_id and js.l_id = ?
			LEFT JOIN jcc_companyroom jc ON js.l_id = jc.id
			LEFT JOIN jcc_goods_unit_conversion jguc ON jg.id = jguc.child_goods_id
			LEFT JOIN jcc_unit ju ON jguc.unit_id = ju.id
            LEFT JOIN jcc_goods_category_index jgci ON jg.id = jgci.goods_id `

	err := db.Raw(sql, childId, companyId).QueryRow(&childInfo)

	return childInfo, err
}

func QueryChildIdOrParentId(param *request.SwitchGoods, db orm.Ormer) ([]int, error) {

	var resultId []int

	sql := ``
	if param.IsParent == 1 {
		sql = `SELECT child_goods_id from jcc_goods_unit_conversion WHERE goods_id = ? and unit_id = ` + strconv.Itoa(param.UnitId)
	} else {
		sql = `SELECT goods_id from jcc_goods_unit_conversion WHERE child_goods_id = ?`
	}

	_, err := db.Raw(sql, param.Id).QueryRows(&resultId)

	return resultId, err
}

func ExistBarcode(barCode string, companyId int) (int, error) {

	db := orm.NewOrm()

	sql := `select count(1) from jcc_goods where barcode = ? and is_del = 0 and companyid = ?`

	var count int

	err := db.Raw(sql, barCode, companyId).QueryRow(&count)

	return count, err
}

func ExistHaveChild(parentId int) (int, error) {

	db := orm.NewOrm()

	sql := `select count(1) from jcc_goods where id = ? and is_have_child = 1`

	var count int

	err := db.Raw(sql, parentId).QueryRow(&count)

	return count, err
}

//通过goodsId查询商品信息,用in为条件
func QueryGoodsInfoByGoodIds(orderIds string) (info []OrderGoodsInfo, err error) {
	db := orm.NewOrm()

	sql := `SELECT
			jg.id,jg.name,jg.barcode,jg.spec,jpg.unit_id,jg.retail_price,jpg.num,jg.is_parent,jpg.order_id,jg.is_have_child
			FROM
			(SELECT goods_id, num, unit_id, order_id FROM jcc_purchase_goods WHERE id in(` + orderIds + `)) jpg
			INNER JOIN
			jcc_goods jg ON jpg.goods_id = jg.id`

	_, err = db.Raw(sql).QueryRows(&info)

	if err != nil {
		return nil, errors.New("通过in条件查询商品信息失败。。。。")
	}

	return info, err
}
