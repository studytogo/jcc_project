package models

import (
	"github.com/astaxie/beego/orm"
)

type AppParam struct {
	AppKey   string `orm:"column(app_key);auto"`
	AppValue string `orm:"column(app_value);auto"`
}

func QueryAppParam(appKey string) (*AppParam, error) {
	o := orm.NewOrm()
	var app *AppParam
	sql := `select app_key, app_value from applications where app_key = ?`

	err := o.Raw(sql, appKey).QueryRow(&app)

	return app, err

}
