package erp_stock_record

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccErpStockRecord struct {
	Id          int       `orm:"column(id);auto" description:"ID" json:"id"  `
	CreatedAt   time.Time `orm:"column(created_at)" description:"创建时间" json:"-"  `
	UpdatedAt   time.Time `orm:"column(updated_at)" description:"修改时间" json:"-"  `
	OldStock    int       `orm:"column(old_stock)" description:"创建时间" json:"old_stock"  `
	NewStock    int       `orm:"column(new_stock)" description:"创建时间" json:"new_stock"  `
	Source      string    `orm:"column(source)" description:"创建时间" json:"source"  `
	ChangeStock int       `orm:"column(change_stock)" description:"创建时间" json:"change_stock"  `
	Companyid   int       `orm:"column(companyid)" description:"创建时间" json:"companyid"  `
	GoodsId     int       `orm:"column(goods_id)" description:"创建时间" json:"goods_id"  `
}

func init() {
	orm.RegisterModel(new(JccErpStockRecord)) //jcc_erp_stock_record
}

func AddErpStockRecord(param *JccErpStockRecord) (int64, error) {
	db := orm.NewOrm()
	return db.Insert(param)
}
