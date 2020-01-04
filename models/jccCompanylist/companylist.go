package jccCompanylist

import "github.com/astaxie/beego/orm"

type JccCompanylist struct {
	Id          int    `orm:"column(id);auto"`
	Name        string `orm:"column(name);size(255)" description:"公司名称"`
	Companycode string `orm:"column(companycode);size(100)" description:"公司标识"`
	Sort        int    `orm:"column(sort)" description:"排序"`
	Level       int8   `orm:"column(level)" description:"等级"`
	Pid         int    `orm:"column(pid)" description:"上级id"`
	IsDel       int8   `orm:"column(is_del)" description:"是否删除默认0不删除"`
	People      string `orm:"column(people);size(100);null" description:"联系人"`
	Address     string `orm:"column(address);size(255);null" description:"公司地址"`
	Tel         string `orm:"column(tel);size(50);null" description:"公司电话"`
	Fax         string `orm:"column(fax);size(50);null" description:"传真地址"`
	Remarks     string `orm:"column(remarks);null" description:"备注"`
	Province    string `orm:"column(province);size(10);null" description:"省"`
	City        string `orm:"column(city);size(10);null" description:"市"`
	District    string `orm:"column(district);size(10);null" description:"区"`
	CreatedAt   int64    `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt   int64    `orm:"column(updated_at);null" description:"最后修改时间"`
	DeletedAt   int64    `orm:"column(deleted_at);null" description:"删除时间"`
	Regionid    int    `orm:"column(regionid)" description:"区域id"`
	Companyid   int    `orm:"column(companyid)" description:"所属公司id"`
}

func init() {
	orm.RegisterModel(new(JccCompanylist))
}

func CheckExistByName(params string) (count int,err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_companylist where name = ? and is_del = 0`
	err = o.Raw(sql,params).QueryRow(&count)
	return count,err
}

func AddCompany(o orm.Ormer,params *JccCompanylist) (id int64,err error) {
	id,err = o.Insert(params)
	return id,err
}

func DelCompany(o orm.Ormer,params *JccCompanylist) error {
	_,err := o.Update(params,"is_del","deleted_at")
	return err
}

func QueryCompanylistIdById(id int) (list JccCompanylist,err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_companylist where id = ? `
	err = o.Raw(sql,id).QueryRow(&list)
	return list,err
}

func UpdateCompany(o orm.Ormer,params *JccCompanylist) error {
	_,err := o.Update(params)
	return err
}
