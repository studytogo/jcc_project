package record

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(JccContextParam))
}

type JccContextParam struct {
	Id            int
	CreatedAt     string
	RequsetParam  string
	ResponseParam string
	Url           string
	AppName       string
}

func (cp *JccContextParam) AddContext() (id int64, err error) {
	o := orm.NewOrm()
	return o.Insert(cp)
}

func UpdateContext(id int, response string) (err error) {
	o := orm.NewOrm()
	sql := `update jcc_context_param set response_param = ? where id = ?`
	_, err = o.Raw(sql, response, id).Exec()
	return err
}
