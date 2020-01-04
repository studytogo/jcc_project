package printService

import (
	"new_erp_agent_by_go/models/purchaseOrder"

	"github.com/astaxie/beego/orm"
)

func CheckList(company_id string, orderid string, start_page int, end_page int, page int) ([]purchaseOrder.CheckPurchaseOrder, int, error) {
	db := orm.NewOrm()
	list, count, err := purchaseOrder.CheckList(company_id, orderid, start_page, end_page, page, db)
	if err != nil {
		return nil, 0, err
	}
	return list, count, err
}
