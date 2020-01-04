package brand

import (
	"github.com/astaxie/beego/orm"
)

type JccBrand struct {
	Id        int
	Name      string
	CreatedAt int
	UpdatedAt int
	DeletedAt int
	IsDel     int
	Kind      string
	Companyid int
	ErpId     int
}

func QueryOnlineBrandInfo() ([]JccBrand, error) {
	db := orm.NewOrm()
	var brand []JccBrand

	sql := `SELECT * FROM jcc_brand WHERE is_del = 0 AND kind = "online"`

	_, err := db.Raw(sql).QueryRows(&brand)

	return brand, err
}
