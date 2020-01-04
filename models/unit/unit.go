package unit

import (
	"github.com/astaxie/beego/orm"
)

type JccUnit struct {
	Id        int    `orm:"column(id);auto" description:"ID"`
	Name      string `orm:"column(name);size(20)" description:"名称"`
	CreatedAt int    `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt int    `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt int    `orm:"column(deleted_at)" description:"删除时间"`
	IsDel     int8   `orm:"column(is_del)" description:"是否删除"`
	Kind      string `orm:"column(kind);size(20);null" description:"商品的分类（本地(offline)、自营(online)）默认的是本地(offline)"`
	Companyid int    `orm:"column(companyid)" description:"公司id"`
	ErpId     int    `orm:"column(erp_id)" description:"erp的单位id"`
}

func QueryUnitIdByGoodsId(goodId int) (int, error) {
	db := orm.NewOrm()

	sql := `select unit_id from jcc_goods where id = ? and is_del = 0`

	var unitId int

	err := db.Raw(sql, goodId).QueryRow(&unitId)

	return unitId, err
}

type CheckUnit struct {
	UnitId int
	GoodId int
}

func QueryUnitNameByUnitId(unitId int) (string, error) {
	db := orm.NewOrm()

	sql := `select name from jcc_unit where id = ? and is_del = 0`

	var unitName string

	err := db.Raw(sql, unitId).QueryRow(&unitName)

	return unitName, err
}

func ExistInternationalUnitByByName(unitName string) (int, error) {
	db := orm.NewOrm()

	sql := `select count(1) from jcc_unit_conversion where unit_name = ? and  is_del = 0`

	var count int

	err := db.Raw(sql, unitName).QueryRow(&count)

	return count, err
}

func GoodStockByGoodId(goodId int) (int, error) {
	db := orm.NewOrm()

	sql := "select sum(`actual`) from jcc_stock where goods_id = ?"

	var actual int

	err := db.Raw(sql, goodId).QueryRow(&actual)

	return actual, err

}

func QueryConversionNumByUnitName(unitName string) (int, error) {
	db := orm.NewOrm()

	sql := `select conversion_num from jcc_unit_conversion where unit_name = ? and  is_del = 0`

	var conversionNum int

	err := db.Raw(sql, unitName).QueryRow(&conversionNum)

	return conversionNum, err
}

//查询所有单位信息
func QueryAllUnitInfo() ([]JccUnit, error) {
	db := orm.NewOrm()

	sql := `select * from jcc_unit`

	var unitInfo []JccUnit

	_, err := db.Raw(sql).QueryRows(&unitInfo)

	return unitInfo, err
}
