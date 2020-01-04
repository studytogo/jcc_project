package printService

import (
	"new_erp_agent_by_go/models/print"

	"github.com/astaxie/beego/orm"
)

func QueryList(company_id int) (*[]print.CheckUserMould, error) {
	db := orm.NewOrm()
	mould, err := print.QueryUserMould(company_id, db)

	if err != nil {
		return nil, err
	}
	return mould, err
}
