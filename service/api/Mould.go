package api

import (
	"errors"
	"github.com/astaxie/beego/orm"
	print2 "new_erp_agent_by_go/models/print"
)

func SelectMould(Id int) (info *print2.JccCommonMould, err error) {
	db := orm.NewOrm()
	_ = db.Begin()
	// 当有Id值时查一条
	res, err := print2.SelectMould(db, Id)
	if err != nil {
		return res, errors.New("查询打印机失败")
	}
	_ = db.Commit()
	return res, err
}

func SelectMouldList() (info *[]print2.JccCommonMould, err error) {
	db := orm.NewOrm()
	_ = db.Begin()
	// 查全部
	res, err := print2.SelectMouldList(db)
	if err != nil {
		return res, errors.New("查询失败")
	}
	_ = db.Commit()
	return res, err
}
