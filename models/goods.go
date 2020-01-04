package models

import (
	"errors"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/erp_goods"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type JccGoodsUnitConversion struct {
	Id           int64  `orm:"column(id);auto" description:"ID"`
	GoodsId      int64  `orm:"column(goods_id)" description:"父商品id" json:"goods_id"`
	ChildGoodsId int64  `orm:"column(child_goods_id)" description:"子商品id" json:"child_goods_id"`
	CreatedAt    int64  `orm:"column(created_at)" description:"创建时间"`
	UnitId       int64  `orm:"column(unit_id)" description:"单位id"`
	Num          int64  `orm:"column(num)" description:"换算数量" json:"num"`
	Mode         string `orm:"-" json:"mode"`
	LId          int64  `orm:"-" description:"所在仓库" json:"l_id"`
	UnitName     string `orm:"-"`
	GoodsName    string `orm:"-"`
}

type JccStock struct {
	Id        int64 `orm:"column(id);auto" description:"ID"`
	GoodsId   int64 `orm:"column(goods_id)" description:"商品id" json:"goods_id"`
	Actual    int64 `orm:"column(actual)" description:"实际库存" json:"actual"`
	Option    int64 `orm:"column(option)" description:"可定库存" json:"option"`
	LId       int64 `orm:"column(l_id)" description:"所在仓库"`
	UpdatedAt int64 `orm:"column(updated_at)" description:"修改时间"`
	Companyid int   `orm:"column(companyid)" description:"公司id" json:"companyid"`
}

type JccUnit struct {
	Id        int64  `orm:"column(id);auto" description:"ID"`
	Name      string `orm:"column(name)" description:"单位id" json:"unit_id"`
	IsDel     int64  `orm:"column(is_del)" description:"是否删除" json:"is_del"`
	UpdatedAt int64  `orm:"column(updated_at)" description:"修改时间"`
	//Companyid int    `orm:"column(companyid)" description:"公司id" json:"companyid"`
}

type JccBrand struct {
	Id        int64  `orm:"column(id);auto" description:"ID"`
	Name      string `orm:"column(name)" description:"品牌id" json:"brand_id"`
	IsDel     int64  `orm:"column(is_del)" description:"是否删除" json:"is_del"`
	UpdatedAt int64  `orm:"column(updated_at)" description:"修改时间"`
	Companyid int    `orm:"column(companyid)" description:"公司id" json:"companyid"`
}

type JccGoodsCategoryIndex struct {
	Id              int64
	GoodsId         int64
	GoodsCategoryId int64
	CreatedAt       int64
}

type JccGoodsSupplierIndex struct {
	Id         int64
	GoodsId    int64
	SupplierId int64
	CreatedAt  int64
}

func init() {
	orm.RegisterModel(new(JccGoodsUnitConversion), new(JccStock), new(JccUnit), new(JccGoodsCategoryIndex), new(JccGoodsSupplierIndex))
}

// 通过childGoodsId查找unitConversion
func ChildSelect(ormer orm.Ormer, childGoodsId int64) (goodsUnitConversionInfo *JccGoodsUnitConversion, err error) {
	goodsUnitConversionInfo = new(JccGoodsUnitConversion)
	err = ormer.QueryTable("jcc_goods_unit_conversion").Filter("child_goods_id", childGoodsId).One(goodsUnitConversionInfo)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品查询失败")
	}

	return goodsUnitConversionInfo, err
}

// 通过GoodsId查找unitConversion
func ParentSelect(ormer orm.Ormer, GoodsId int64) (goodsUnitConversionInfo *JccGoodsUnitConversion, err error) {
	goodsUnitConversionInfo = new(JccGoodsUnitConversion)
	err = ormer.QueryTable("jcc_goods_unit_conversion").Filter("goods_id", GoodsId).One(goodsUnitConversionInfo)
	if err != nil {
		helper.Log.Error("父商品查询失败", err)
		return nil, err
	}

	return goodsUnitConversionInfo, err
}

// 拆分查子库存
func SelectGoodsStock(ormer orm.Ormer, childGoodsId int64, newStrock int64, lId int64) (strockId int64, goodsStock int64, err error) {
	var strock = new(JccStock)
	err = ormer.QueryTable("jcc_stock").Filter("goods_id", childGoodsId).Filter("l_id", lId).One(strock)
	if err != nil {
		strock.GoodsId = childGoodsId
		strock.Actual = newStrock
		strock.LId = lId
		_, err = ormer.Insert(strock)
		if err != nil {
			return 0, 0, errors.New("子商品库存添加失败")
		}
	} else {
		strock.Actual = strock.Actual + newStrock
		_, err = ormer.QueryTable("jcc_stock").Filter("goods_id", childGoodsId).Filter("l_id", lId).Update(orm.Params{
			"Actual":    strock.Actual,
			"UpdatedAt": time.Now().Unix(),
		})

	}
	return strock.Id, strock.Actual, err
}

// 合成查子库存
func SelectChildGoodsStock(ormer orm.Ormer, childGoodsId int64, newStrock int64, lId int64) (strockId int64, goodsStock int64, err error) {
	var strock = new(JccStock)
	err = ormer.QueryTable("jcc_stock").Filter("goods_id", childGoodsId).Filter("l_id", lId).One(strock)

	if err != nil {
		return 0, 0, errors.New("子商品没有库存")
	} else {
		strock.Actual = strock.Actual - newStrock
		if strock.Actual < 0 {
			return 0, 0, errors.New("子商品库存不足")
		}
		_, err = ormer.QueryTable("jcc_stock").Filter("goods_id", childGoodsId).Filter("l_id", lId).Update(orm.Params{
			"Actual":    strock.Actual,
			"UpdatedAt": time.Now().Unix(),
		})

	}
	return strock.Id, strock.Actual, err
}

// 查父库存
func SelectParentGoodsStock(ormer orm.Ormer, goodsId int64, LId int64) (goodsStock int64, isHavedStock bool, err error) {
	var strock = new(JccStock)
	isHavedStock = true
	err = ormer.QueryTable("jcc_stock").Filter("goods_id", goodsId).Filter("l_id", LId).One(strock)
	if err != nil {
		if err.Error() != "<QuerySeter> no row found" {
			return 0, false, errors.New("库存查询失败")
		} else {
			isHavedStock = false
		}
	}

	return strock.Actual, isHavedStock, nil
}

// 对库存进行修改
func UpdateGoodsStock(ormer orm.Ormer, goodsId int64, newChildStock int64, lId int64) error {
	_, err := ormer.QueryTable("jcc_stock").Filter("goods_id", goodsId).Filter("l_id", lId).Update(orm.Params{
		"Actual":    newChildStock,
		"UpdatedAt": time.Now().Unix(),
	})
	return err
}

// 添加子商品
func AddChildGoods(ormer orm.Ormer, goods *childGoods.JccGoods) (goodsBack *childGoods.JccGoods, ChildId int64, err error) {
	goods.IsParent = 0
	goods.IsHaveChild = 0
	ChildId, err = ormer.Insert(goods)
	if err != nil {
		helper.Log.Error("添加失败", err)
	}
	return goods, ChildId, err

}

// 添加商品单位换算
func AddGoodsUnitConversion(ormer orm.Ormer, unitId int64, Num int64, ParentGoodsId int64, ChildGoodsId int64) error {
	conversion := new(JccGoodsUnitConversion)
	conversion.GoodsId = ParentGoodsId
	conversion.UnitId = unitId
	conversion.Num = Num
	conversion.CreatedAt = time.Now().Unix()
	conversion.ChildGoodsId = ChildGoodsId
	_, err := ormer.Insert(conversion)
	if err != nil {
		helper.Log.Error("添加商品单位换算失败", err)
	}
	return err
}

// 修改父商品的parent_id字段
func UpdateParentGoods(ormer orm.Ormer, goodsId int64) error {
	_, err := ormer.QueryTable("jcc_goods").Filter("id", goodsId).Update(orm.Params{
		"IsHaveChild": 1,
		"UpdatedAt":   time.Now().Unix(),
	})
	return err
}

// 通过child_goods_id  找父商品的单位
func ChildSelectConversion(ormer orm.Ormer, childGoodsId int64) (goodsUnitConversionInfo *JccGoodsUnitConversion, err error) {
	goodsUnitConversionInfo = new(JccGoodsUnitConversion)
	err = ormer.QueryTable("jcc_goods_unit_conversion").Filter("child_goods_id", childGoodsId).One(goodsUnitConversionInfo)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品查询失败")
	}

	// 传1  通过child_goods_id  找父商品的单位
	parentGoodsInfo := new(childGoods.JccGoods)
	err = ormer.QueryTable("jcc_goods").Filter("id", goodsUnitConversionInfo.GoodsId).One(parentGoodsInfo)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品单位查询失败")
	}

	goodsUnit := new(JccUnit)
	err = ormer.QueryTable("jcc_unit").Filter("id", parentGoodsInfo.UnitId).One(goodsUnit)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品单位查询失败")
	}

	goodsUnitConversionInfo.UnitName = goodsUnit.Name
	goodsUnitConversionInfo.GoodsName = parentGoodsInfo.Name

	return goodsUnitConversionInfo, err
}

// 通过GoodsId查找unitConversion
func ParentSelectConversion(ormer orm.Ormer, GoodsId int64) (goodsUnitConversionInfo *JccGoodsUnitConversion, err error) {
	goodsUnitConversionInfo = new(JccGoodsUnitConversion)
	err = ormer.QueryTable("jcc_goods_unit_conversion").Filter("goods_id", GoodsId).One(goodsUnitConversionInfo)
	if err != nil {
		helper.Log.Error("父商品查询失败", err)
		return nil, err
	}

	// 传0  通过goods_id  找子商品的单位
	childGoodsInfo := new(childGoods.JccGoods)
	err = ormer.QueryTable("jcc_goods").Filter("id", goodsUnitConversionInfo.ChildGoodsId).One(childGoodsInfo)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品单位查询失败")
	}

	goodsUnit := new(JccUnit)
	err = ormer.QueryTable("jcc_unit").Filter("id", childGoodsInfo.UnitId).One(goodsUnit)
	if err != nil {
		return goodsUnitConversionInfo, errors.New("子商品单位查询失败")
	}

	goodsUnitConversionInfo.UnitName = goodsUnit.Name
	goodsUnitConversionInfo.GoodsName = childGoodsInfo.Name

	return goodsUnitConversionInfo, err
}

func AddGoodsCategory(goodsCategory *JccGoodsCategoryIndex, ormer orm.Ormer, GoodsId int64, GoodsCategoryId int64) error {
	goodsCategory.GoodsId = GoodsId
	goodsCategory.GoodsCategoryId = GoodsCategoryId
	goodsCategory.CreatedAt = time.Now().Unix()
	_, err := ormer.Insert(goodsCategory)
	if err != nil {
		return errors.New("子商品商品分类关系添加失败")
	}
	return err
}

func AddGoodsSupplier(goodsSupplier *JccGoodsSupplierIndex, ormer orm.Ormer, GoodsId int64, GoodsSupplierId int64) error {
	goodsSupplier.GoodsId = GoodsId
	goodsSupplier.SupplierId = GoodsSupplierId
	goodsSupplier.CreatedAt = time.Now().Unix()
	_, err := ormer.Insert(goodsSupplier)
	if err != nil {
		return errors.New("子商品供应商关系添加失败")
	}
	return err
}

// 通过商品id查询单位id
func SelectGoodsUnitId(ormer orm.Ormer, goodsId int64) (unitId int, err error) {
	var goods = new(childGoods.JccGoods)
	err = ormer.QueryTable("jcc_goods").Filter("id", goodsId).One(goods)
	if err != nil {
		return goods.UnitId, errors.New("商品单位查询失败")
	}
	return goods.UnitId, err
}

func SelectUnitName(ormer orm.Ormer, unitId int) (unitName string, err error) {
	var unit = new(JccUnit)
	err = ormer.QueryTable("jcc_unit").Filter("id", unitId).One(unit)
	if err != nil {
		return unit.Name, errors.New("添加子商品单位不存在" + strconv.Itoa(unitId))
	}
	return unit.Name, err
}

//增加库存信息
func (stock *JccStock) AddStock(db orm.Ormer) error {
	_, err := db.Insert(stock)

	return err
}

func QuerySKUExit(sku string, companyid string) (count int, err error) {
	db := orm.NewOrm()
	sql := `select count(*) from jcc_goods where sku = ? and companyid = ?`

	err = db.Raw(sql, sku, companyid).QueryRow(&count)

	return count, err
}

func AddOnlineGoods(db orm.Ormer, list erp_goods.JccCenterGoods, companyid int, erp_id string) (id int64, err error) {

	jcc_goods := new(childGoods.JccGoods)
	jcc_goods.Name = list.Name
	jcc_goods.Spu = list.Spu
	jcc_goods.Sku = list.Sku
	jcc_goods.Barcode = list.Barcode
	jcc_goods.Spec = list.Spec
	jcc_goods.BuyingPrice = list.BuyingPrice
	jcc_goods.RetailPrice = list.RetailPrice
	jcc_goods.InventoryUpperLimit = list.InventoryUpperLimit
	jcc_goods.InventoryLowerLimit = list.InventoryLowerLimit
	jcc_goods.MnemonicWord = list.MnemonicWord
	jcc_goods.Remark = list.Remark
	jcc_goods.Image = list.Image
	jcc_goods.Content = list.Content
	jcc_goods.ProducingProvinceId = list.ProducingProvinceId
	jcc_goods.ProducingCityId = list.ProducingCityId
	jcc_goods.ProducingAreaId = list.ProducingAreaId
	jcc_goods.ProducingAreaDetail = list.ProducingAreaDetail
	jcc_goods.UnitId = list.UnitId
	jcc_goods.BrandId = 1
	jcc_goods.CreatedAt = time.Now().Unix()
	jcc_goods.UpdatedAt = time.Now().Unix()
	jcc_goods.IsCancelProcurement = list.IsCancelProcurement
	jcc_goods.Companyid = companyid
	jcc_goods.Kind = "online"
	jcc_goods.Images = list.Images
	jcc_goods.IsParent = 1
	erp_id_temp, _ := strconv.Atoi(erp_id)
	jcc_goods.ErpId = erp_id_temp

	id, err = db.Insert(jcc_goods)

	return id, err
}

func AddOnlineGoodsCategory(db orm.Ormer, GoodsId int64, GoodsCategoryId int64) error {

	goodsCategory := new(JccGoodsCategoryIndex)
	goodsCategory.GoodsId = GoodsId
	goodsCategory.GoodsCategoryId = GoodsCategoryId
	goodsCategory.CreatedAt = time.Now().Unix()
	_, err := db.Insert(goodsCategory)
	if err != nil {
		return errors.New("商品分类关系添加失败")
	}
	return err
}

func QuerySKUId(sku string, companyid int) (id int64, err error) {
	db := orm.NewOrm()
	sql := `select id from jcc_goods where sku = ? and companyid = ?`

	err = db.Raw(sql, sku, companyid).QueryRow(&id)

	return id, err
}

func QueryCategoryId(goodsid int64) (id int64, err error) {
	db := orm.NewOrm()
	sql := `select id from jcc_goods_category_index where goods_id = ?`

	err = db.Raw(sql, goodsid).QueryRow(&id)

	return id, err
}

func   UpdateOnlineGoods(db orm.Ormer, list erp_goods.JccCenterGoods, id int64) error {
	jcc_goods := new(childGoods.JccGoods)

	jcc_goods.Id = id
	jcc_goods.Name = list.Name
	jcc_goods.Spu = list.Spu
	jcc_goods.Spec = list.Spec
	jcc_goods.BuyingPrice = list.BuyingPrice
	jcc_goods.RetailPrice = list.RetailPrice
	jcc_goods.InventoryUpperLimit = list.InventoryUpperLimit
	jcc_goods.InventoryLowerLimit = list.InventoryLowerLimit
	jcc_goods.MnemonicWord = list.MnemonicWord
	jcc_goods.Remark = list.Remark
	jcc_goods.Image = list.Image
	jcc_goods.Content = list.Content
	jcc_goods.UnitId = list.UnitId
	jcc_goods.Barcode = list.Barcode
	jcc_goods.ProducingAreaId = list.ProducingAreaId
	jcc_goods.ProducingCityId = list.ProducingCityId
	jcc_goods.ProducingProvinceId = list.ProducingProvinceId
	jcc_goods.UpdatedAt = time.Now().Unix()
	jcc_goods.IsCancelProcurement = list.IsCancelProcurement
	jcc_goods.ProducingAreaDetail = list.ProducingAreaDetail

	jcc_goods.Images = list.Images

	_, err := db.Update(jcc_goods, "BuyingPrice", "Barcode", "ProducingAreaId", "ProducingCityId", "ProducingProvinceId", "UnitId", "UpdatedAt", "IsCancelProcurement","ProducingAreaDetail")
	return err
}

func UpdateOnlineGoodsCategory(db orm.Ormer, id int64, GoodsCategoryId int) error {

	goodsCategory := new(JccGoodsCategoryIndex)
	goodsCategory.Id = id
	goodsCategory.GoodsCategoryId = int64(GoodsCategoryId)
	//goodsCategory.CreatedAt = time.Now().Unix()
	_, err := db.Update(goodsCategory, "GoodsCategoryId")
	if err != nil {
		return errors.New("商品分类关系修改失败")
	}
	return err
}
