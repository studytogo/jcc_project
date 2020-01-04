package printService

import (
	"new_erp_agent_by_go/models/purchaseOrder"

	"github.com/astaxie/beego/orm"
)

func PrintLog(OrderIDs string) error {
	db := orm.NewOrm()
	err := purchaseOrder.UpdateLog(OrderIDs, db)
	return err
}
