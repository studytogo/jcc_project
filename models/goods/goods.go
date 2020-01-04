package goods

import (
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/childGoods"
)

func QueryExistBySkuAndCompany(params childGoods.JccGoods) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Read(&params, "sku", "companyid", "is_del")
	return params.Id, err
}

func QuerygoodsById(id int64) (list childGoods.JccGoods, err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_goods where id = ?`
	err = o.Raw(sql, id).QueryRow(&list)
	return list, err
}

//修改商品属性
func UpdateGoodAttribute(erpId interface{}, condition string, db orm.Ormer, companyId string) error {

	sql := `update jcc_goods ` + condition + ` where erp_id = ? and companyid = ?`

	_, err := db.Raw(sql, erpId, companyId).Exec()

	return err
}
