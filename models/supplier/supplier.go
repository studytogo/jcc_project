package supplier

import (
	"github.com/astaxie/beego/orm"
)

type JccSupplier struct {
	Id        int
	Name      string
	Address   string
	Phone     string `orm:"column(phone)"`
	Linkman   string
	CreatedAt int
	UpdatedAt int
	DeletedAt int
	IsDel     int
	Kind      string `orm:"column(kind)"`
}

func QueryCategoryInfoByGoodId(goodId int) (string, int, error) {

	db := orm.NewOrm()

	var supplierId int
	var supplierName string

	sql := `SELECT
				js.id,
				js.name
			FROM
				jcc_goods_supplier_index jgsi
			INNER JOIN jcc_supplier js ON jgsi.supplier_id = js.id
			WHERE
				jgsi.goods_id = ?`

	err := db.Raw(sql, goodId).QueryRow(&supplierId, &supplierName)

	return supplierName, supplierId, err
}

func QueryAllOnlineSupplierInfo() ([]JccSupplier, error) {
	db := orm.NewOrm()
	var supplier []JccSupplier

	sql := `SELECT * FROM jcc_supplier WHERE is_del = 0 AND kind = "online"`

	_, err := db.Raw(sql).QueryRows(&supplier)

	return supplier, err
}
