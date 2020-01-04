package pycompanylist

import (
	"github.com/astaxie/beego/orm"
)

type JccJicanchuCompanylist struct {
	Id        int64  `orm:"column(id);auto" description:"ID"`
	Name      string `orm:"column(name);size(255)" description:"公司名称" json:"name" valid:"Required"`
	CreatedAt int64  `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt int64  `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt int64  `orm:"column(deleted_at)" description:"删除时间"`
	IsDel     int8   `orm:"column(is_del)" description:"是否删除"`
}

type T struct {
	Data      []JccJicanchuCompanylist
	Total     int
	Last_page int
}

func init() {
	orm.RegisterModel(new(JccJicanchuCompanylist))
}

func AddCompanylist(parmas []JccJicanchuCompanylist) error {
	o := orm.NewOrm()
	_, err := o.InsertMulti(1, &parmas)

	return err
}

func QueryCompanylist(Page int, Per_page int) (list []JccJicanchuCompanylist, err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_jicanchu_companylist limit ?,?`
	_, err = o.Raw(sql, (Page-1)*Per_page, Per_page).QueryRows(&list)

	return list, err
}

func QueryCompanylistCount() (count int, err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_jicanchu_companylist`
	err = o.Raw(sql).QueryRow(&count)

	return count, err
}
func EditCompanylist(params JccJicanchuCompanylist) error {
	o := orm.NewOrm()
	_, err := o.Update(&params, "name", "updated_at")
	return err
}

func DeleteCompanylist(params string) error {
	o := orm.NewOrm()

	sql := `update jcc_jicanchu_companylist set is_del=1,deleted_at= unix_timestamp(now()) where id in (` + params + `)`
	_, err := o.Raw(sql).Exec()
	return err
}

func QueryCompanylistOnly(name string, id int64) (count int, err error) {
	o := orm.NewOrm()

	sql := `select count(*) from jcc_jicanchu_companylist where name = ? and is_del = 0 and id <> ?`
	err = o.Raw(sql, name, id).QueryRow(&count)
	return count, err
}

func QueryDatabase(params string) (count int, err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_jicanchu_companylist where is_del = 0 and name in (` + params + `)`
	err = o.Raw(sql).QueryRow(&count)
	return count, err
}
