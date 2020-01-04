package erp_goods

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccCenterGoods struct {
	Id                    int     `orm:"column(id);auto" description:"ID" json:"id"  `
	Name                  string  `orm:"column(name);size(150)" description:"名称" json:"name"  valid:"Required"`
	Spu                   string  `orm:"column(spu);size(150)" description:"SPU" json:"spu"  `
	Sku                   string  `orm:"column(sku);size(150)" description:"SKU" json:"sku"  `
	Barcode               string  `orm:"column(barcode);size(150)" description:"条码" json:"barcode"  valid:"Required" `
	Spec                  string  `orm:"column(spec);size(150)" description:"规格" json:"spec" `
	BuyingPrice           float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"进货价" json:"agent_price"  `
	RetailPrice           float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价" json:"retail_price"  `
	InventoryUpperLimit   string  `orm:"column(inventory_upper_limit);size(255);null" description:"库存上限" json:"inventory_upper_limit" `
	InventoryLowerLimit   string  `orm:"column(inventory_lower_limit)" description:"库存下限" json:"inventory_lower_limit"  `
	MnemonicWord          string  `orm:"column(mnemonic_word);size(255)" description:"助记词" json:"mnemonic_word"  `
	Remark                string  `orm:"column(remark);size(150)" description:"备注" json:"remark"  `
	Image                 string  `orm:"column(image);size(150)" description:"主图" json:"image"  `
	Images                string  `orm:"column(images)" description:"多图" json:"images"  `
	Content               string  `orm:"column(content)" description:"内容" json:"content"  `
	ProducingProvinceId   int     `orm:"column(producing_province_id)" description:"产地省ID" json:"producing_province_id"  valid:"Required"`
	ProducingProvinceName string  `orm:"column(producing_province_name)" description:"产地省名称" json:"producing_province_name"  `
	ProducingCityId       int     `orm:"column(producing_city_id)" description:"产地城市ID" json:"producing_city_id"  valid:"Required"`
	ProducingCityName     string  `orm:"column(producing_city_name)" description:"产地城市名称" json:"producing_city_name"  `
	ProducingAreaId       int     `orm:"column(producing_area_id)" description:"产地区域ID" json:"producing_area_id"  valid:"Required"`
	ProducingAreaName     string  `orm:"column(producing_area_name)" description:"产地区域ID" json:"producing_area_name"  `
	ProducingAreaDetail   string  `orm:"column(producing_area_detail);size(150)" description:"产地详情" json:"producing_area_detail"  `
	UnitId                int     `orm:"column(unit_id)" description:"单位ID" json:"unit_id"  valid:"Required"`
	UnitName              string  `orm:"column(unit_name)" description:"单位名称" json:"unit_name"  `
	BrandId               int     `orm:"column(brand_id)" description:"品牌ID" json:"brand_id"  valid:"Required"`
	BrandName             string  `orm:"column(brand_name)" description:"品牌名称" json:"brand_name" `
	CreatedAt             int64   `orm:"column(created_at)" description:"创建时间" json:"-"  `
	UpdatedAt             int64   `orm:"column(updated_at)" description:"修改时间" json:"-"  `
	DeletedAt             int64   `orm:"column(deleted_at)" description:"删除时间" json:"-"  `
	IsDel                 int8    `orm:"column(is_del)" description:"是否删除" json:"-"  `
	A8Code                string  `orm:"column(a8_code)" description:"a8编码" json:"a8_code"  `
	IsCancelProcurement   int8    `orm:"column(is_cancel_procurement)" description:"是否取消采购" json:"is_cancel_procurement"  `
	CategoryId            int     `orm:"column(category_id)" description:"商品分类" json:"goods_category_ids"  `
}

func init() {
	orm.RegisterModel(new(JccCenterGoods))
}

//批量增加商品(取消)
func InsertMultiGoods(data []*JccCenterGoods) (int64, error) {
	db := orm.NewOrm()
	return db.InsertMulti(1, data)
}

//修改商品信息
func UpdateErpGood(data *JccCenterGoods, db orm.Ormer) (int64, error) {
	return db.Update(data, "name", "spu", "sku", "barcode", "spec",
		"buying_price", "retail_price", "inventory_upper_limit", "inventory_lower_limit",
		"mnemonic_word", "remark", "image", "images", "content", "producing_province_id", "producing_province_name", "producing_city_id",
		"producing_city_name", "producing_area_id", "producing_area_name", "producing_area_detail", "unit_id", "unit_name",
		"brand_id", "brand_name", "a8_code", "updated_at", "is_cancel_procurement", "category_id")
}

//删除商品
func DeleteErpGood(id int, db orm.Ormer) (int64, error) {
	return db.QueryTable("jcc_center_goods").Filter("id", id).Update(orm.Params{
		"is_del":     1,
		"deleted_at": time.Now().Unix(),
	})
}

//查询商品信息
func QueryErpGoodInfo(ids []int) ([]JccCenterGoods, error) {
	db := orm.NewOrm()

	var goodsInfo []JccCenterGoods
	_, err := db.QueryTable("jcc_center_goods").Filter("id__in", ids).Filter("is_del", 0).All(&goodsInfo)

	return goodsInfo, err
}

//查询条码是否存在
func ExistBarcode(barcode string, goodId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_center_goods where barcode = ? and id != ? and is_del = 0`

	var count int

	err := db.Raw(sql, barcode, goodId).QueryRow(&count)

	return count, err

}

//商品名称是否存在
func ExistGoodName(name string, goodId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_center_goods where name = ? and id != ? and is_del = 0`

	var count int
	err := db.Raw(sql, name, goodId).QueryRow(&count)

	return count, err
}

//判断是否是省
func CheckProvince(provinceId int) (int, string, error) {
	db := orm.NewOrm()

	sql := `select deep,name from jcc_address where id = ? `

	var deep int
	var name string

	err := db.Raw(sql, provinceId).QueryRow(&deep, &name)

	return deep, name, err
}

//判断地区是否合法
func CheckArea(areaId int) (int, string, error) {
	db := orm.NewOrm()

	sql := `select pid,name from jcc_address where id = ?`

	var areaPId int
	var areaName string

	err := db.Raw(sql, areaId).QueryRow(&areaPId, &areaName)

	return areaPId, areaName, err
}

func CheckUnitId(unitId int) (int, string, error) {
	db := orm.NewOrm()

	sql := `select count(*), name from jcc_unit where id = ? and is_del = 0 group by name`

	var count int
	var name string

	err := db.Raw(sql, unitId).QueryRow(&count, &name)

	return count, name, err
}

func CheckBrandId(brandId int) (int, string, error) {
	db := orm.NewOrm()

	sql := `select count(*), name from jcc_brand where id = ? and is_del = 0 group by name`

	var count int
	var brandName string

	err := db.Raw(sql, brandId).QueryRow(&count, &brandName)

	return count, brandName, err
}

//查询所有的商品名称
func QueryAllGoodName() ([]*SyncInfo, error) {
	db := orm.NewOrm()

	sql := `select id, name as arrtibute from jcc_center_goods where is_del = 0`

	var allName []*SyncInfo
	_, err := db.Raw(sql).QueryRows(&allName)

	return allName, err
}

type SyncInfo struct {
	Id        int    `orm:"column(id);auto" description:"ID" json:"id"  `
	Arrtibute string `orm:"column(arrtibute);auto" description:"ID" json:"arrtibute"  `
}

//查询所有商品条码
func QueryAllBarcode() ([]*SyncInfo, error) {
	db := orm.NewOrm()

	sql := `select id,barcode as arrtibute from jcc_center_goods where is_del = 0`

	var allBarcode []*SyncInfo
	_, err := db.Raw(sql).QueryRows(&allBarcode)

	return allBarcode, err
}

func ExistGoodsId(goodsId int) bool {
	db := orm.NewOrm()

	return db.QueryTable("jcc_center_goods").Filter("id", goodsId).Exist()
}

func ExistCompanyId(companyId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_jicanchu_companylist where id = ? `

	var count int
	err := db.Raw(sql, companyId).QueryRow(&count)

	return count, err
}

func QueryOnlineGoodsInfo(id string) (list JccCenterGoods, err error) {
	db := orm.NewOrm()

	sql := `select * from jcc_center_goods where id = ?`

	err = db.Raw(sql, id).QueryRow(&list)

	return list, err
}

func QueryOnlineGoodsCount(id string) (count int, err error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_center_goods where id = ?`

	err = db.Raw(sql, id).QueryRow(&count)

	return count, err
}

type ErpGoodsInfo struct {
	Id                    int     `orm:"column(id);auto" description:"ID" json:"id"  `
	Name                  string  `orm:"column(name);size(150)" description:"名称" json:"name"  valid:"Required"`
	Spu                   string  `orm:"column(spu);size(150)" description:"SPU" json:"spu"  `
	Sku                   string  `orm:"column(sku);size(150)" description:"SKU" json:"sku"  `
	Barcode               string  `orm:"column(barcode);size(150)" description:"条码" json:"barcode"  valid:"Required"`
	Spec                  string  `orm:"column(spec);size(150)" description:"规格" json:"spec" "`
	BuyingPrice           float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"进货价" json:"buying_price"  `
	RetailPrice           float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价" json:"retail_price"  `
	InventoryUpperLimit   string  `orm:"column(inventory_upper_limit);size(255);null" description:"库存上限" json:"inventory_upper_limit" `
	InventoryLowerLimit   string  `orm:"column(inventory_lower_limit)" description:"库存下限" json:"inventory_lower_limit"  `
	MnemonicWord          string  `orm:"column(mnemonic_word);size(255)" description:"助记词" json:"mnemonic_word"  `
	Remark                string  `orm:"column(remark);size(150)" description:"备注" json:"remark"  `
	Image                 string  `orm:"column(image);size(150)" description:"主图" json:"image"  `
	Images                string  `orm:"column(images)" description:"多图" json:"images"  `
	Content               string  `orm:"column(content)" description:"内容" json:"content"  `
	ProducingProvinceId   int     `orm:"column(producing_province_id)" description:"产地省ID" json:"producing_province_id"  valid:"Required"`
	ProducingProvinceName string  `orm:"column(producing_province_name)" description:"产地省名称" json:"producing_province_name"  `
	ProducingCityId       int     `orm:"column(producing_city_id)" description:"产地城市ID" json:"producing_city_id"  valid:"Required"`
	ProducingCityName     string  `orm:"column(producing_city_name)" description:"产地城市名称" json:"producing_city_name"  `
	ProducingAreaId       int     `orm:"column(producing_area_id)" description:"产地区域ID" json:"producing_area_id"  valid:"Required"`
	ProducingAreaName     string  `orm:"column(producing_area_name)" description:"产地区域ID" json:"producing_area_name"  `
	ProducingAreaDetail   string  `orm:"column(producing_area_detail);size(150)" description:"产地详情" json:"producing_area_detail"  `
	UnitId                int     `orm:"column(unit_id)" description:"单位ID" json:"unit_id"  valid:"Required"`
	UnitName              string  `orm:"column(unit_name)" description:"单位名称" json:"unit_name"  `
	BrandId               int     `orm:"column(brand_id)" description:"品牌ID" json:"brand_id"  valid:"Required"`
	BrandName             string  `orm:"column(brand_name)" description:"品牌名称" json:"brand_name" `
	CreatedAt             int64   `orm:"column(created_at)" description:"创建时间" json:"-"  `
	UpdatedAt             int64   `orm:"column(updated_at)" description:"修改时间" json:"-"  `
	DeletedAt             int64   `orm:"column(deleted_at)" description:"删除时间" json:"-"  `
	IsDel                 int8    `orm:"column(is_del)" description:"是否删除" json:"-"  `
	A8Code                int8    `orm:"column(a8_code)" description:"a8编码" json:"a8_code"  `
	IsCancelProcurement   int8    `orm:"column(is_cancel_procurement)" description:"是否取消采购" json:"is_cancel_procurement"  `
	CategoryId            int     `orm:"column(category_id)" description:"商品分类" json:"goods_category_ids"  `
	Option                int     `orm:"column(option)" description:"可定库存" json:"option"  `
	Uuid                  int     `orm:"column(uuid)" description:"前端判断唯一字段" json:"uuid"  `
}

//通过公司id查询上级公司id
func QueryErpCompanyIdByCompanyId(companyId int) (int, error) {
	db := orm.NewOrm()

	sql := `select companyid from jcc_companylist where id = ?`

	var erpCompanyId int
	err := db.Raw(sql, companyId).QueryRow(&erpCompanyId)

	return erpCompanyId, err
}

//通过erpId查询erp商品信息及库存
func QueryErpGoodsByErpCompanyId(condition string, page, pageSize int) ([]ErpGoodsInfo, error) {
	db := orm.NewOrm()

	sql := `SELECT 
			jcs.option, jcs.id as uuid, jcg.*
			FROM
			jcc_company_stock jcs
			INNER JOIN
			jcc_center_goods jcg on jcs.goods_id = jcg.id ` + condition + ` and is_del = 0 order by jcg.created_at limit ?, ? `

	var result []ErpGoodsInfo
	_, err := db.Raw(sql, (page-1)*pageSize, pageSize).QueryRows(&result)

	return result, err
}

//通过erpId查询erp商品总数
func QueryErpGoodsCount(condition string) (int, error) {
	db := orm.NewOrm()

	sql := `SELECT 
			COUNT(*)
			FROM
			jcc_company_stock jcs
			INNER JOIN
			jcc_center_goods jcg on jcs.goods_id = jcg.id ` + condition + ` and is_del = 0 `

	var count int
	err := db.Raw(sql).QueryRow(&count)

	return count, err
}

//判断公司端条码是否与加盟商的重复
func ExistBarCodeInAgent(barcode string) (int, error) {
	db := orm.NewOrm()
	sql := `select count(*) from jcc_center_goods where barcode = ?`

	var count int
	err := db.Raw(sql, barcode).QueryRow(&count)

	return count, err
}
