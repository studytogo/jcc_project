package actual

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

func GoodsHavedActualByGoodId(goodId int) (int, error) {

	db := orm.NewOrm()

	sql := "SELECT `actual` FROM jcc_stock WHERE goods_id = ?"

	var goodActual int
	err := db.Raw(sql, goodId).QueryRow(&goodActual)

	return goodActual, err
}

func QueryAllActualByCompanyId(companyId, goodId int) (int, error) {
	db := orm.NewOrm()

	sql := `SELECT sum(jsk.actual) as all_actual
			FROM
			(SELECT id FROM jcc_companyroom WHERE comanyid = ?) jcm
			LEFT JOIN jcc_stock jsk ON jsk.l_id = jcm.id 
            WHERE jsk.goods_id = ? `

	var allActual int

	err := db.Raw(sql, companyId, goodId).QueryRow(&allActual)

	return allActual, err
}

type JccCompanyStock struct {
	Id         int64  `orm:"column(id);auto" description:"ID" json:"id" `
	GoodsId    int    `orm:"column(goods_id)" description:"对应商品id" json:"goods_id"`
	Option     int    `orm:"column(option)" description:"可订库存" json:"option"`
	UpdatedAt  int64  `orm:"column(updated_at)" description:"修改时间" json:"-"`
	CompanyId  int    `orm:"column(companyid)" description:"公司id" json:"company_id"`
	CreatedAt  int64  `orm:"column(created_at)" description:"创建时间" json:"-"`
	IsAddOrSub string `orm:"-" description:"加减状态" json:"is_add_or_sub"`
}

type T struct {
	Data      []JccCompanyStock
	Total     int
	Last_page int
}

func init() {
	orm.RegisterModel(new(JccCompanyStock))
}

//判断库存是否存在
func ExistErpStock(goodId, companyId int) (int64, error) {
	db := orm.NewOrm()

	return db.QueryTable("jcc_company_stock").Filter("goods_id", goodId).Filter("companyid", companyId).Count()
}

func AddErpStock(data JccCompanyStock, db orm.Ormer) (int64, error) {
	return db.Insert(&data)
}

func UpdateErpStock(data JccCompanyStock, db orm.Ormer) (int64, error) {
	return db.Update(&data, "option", "updated_at")
}

func QueryErpStock(Page int, Per_page int) (list []JccCompanyStock, err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_company_stock limit ?,?`
	_, err = o.Raw(sql, (Page-1)*Per_page, Per_page).QueryRows(&list)
	return list, err
}

func QuertErpStockCount() (count int, err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_company_stock`
	err = o.Raw(sql).QueryRow(&count)
	return count, err
}

//加减公司端商品库存
func ChangeErpStock(data JccCompanyStock, db orm.Ormer) error {
	sql := ``

	switch data.IsAddOrSub {
	case "add":
		sql = "update jcc_company_stock set `option` = `option` + ? where companyid = ? and goods_id = ?"
	case "sub":
		sql = "update jcc_company_stock set `option` = `option` - ? where companyid = ? and goods_id = ?"
	default:
		return errors.New("请输入IsAddOrSub参数")
	}

	_, err := db.Raw(sql, data.Option, data.CompanyId, data.GoodsId).Exec()

	return err
}

//查询单条商品库存
func QueryOneErpStock(goodId, companyId int) (int, error) {
	db := orm.NewOrm()

	sql := "select `option` from jcc_company_stock where companyid = ? and goods_id = ?"

	var stock int
	err := db.Raw(sql, companyId, goodId).QueryRow(&stock)

	return stock, err
}

//插入单挑商品库存
func InsertOneErpStock(param JccCompanyStock, db orm.Ormer) (int64, error) {
	return db.Insert(&param)
}
