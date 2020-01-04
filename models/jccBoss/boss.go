package jccBoss

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type JccBoss struct {
	Id             int    `orm:"column(id);auto"`
	Uid            int    `orm:"column(uid);null" description:"上级id"`
	Boss           string `orm:"column(boss);size(20)" description:"用户名称"`
	IdCard         string `orm:"column(id_card);size(20);null" description:"显示用ID"`
	Password       string `orm:"column(password);size(32)" description:"密码"`
	Loginip        string `orm:"column(loginip);size(15);null" description:"登陆IP"`
	Pricing        string `orm:"column(pricing);size(100);null" description:"价格体系"`
	Rank           int16  `orm:"column(rank);null" description:"权限数字"`
	IsDel          int    `orm:"column(is_del);null" description:"是否删除"`
	CreatedAt      int64    `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt      int64    `orm:"column(updated_at);null" description:"修改时间"`
	DeletedAt      int64    `orm:"column(deleted_at);null" description:"删除时间"`
	AffCompany     int64    `orm:"column(aff_company);null" description:"所属公司"`
	SectionId      string `orm:"column(section_id);size(50);null" description:"所属部门"`
	Telephone      string `orm:"column(telephone);size(50);null" description:"联系电话"`
	Address        string `orm:"column(address);size(255);null" description:"联系地址"`
	Operator       string `orm:"column(operator);size(50);null" description:"操作员"`
	Openid         string `orm:"column(openid);size(30)" description:"微信Openid"`
	ErpId          int    `orm:"column(erp_id)" description:"erp端的代理商id"`
	GroupId        int    `orm:"column(group_id)" description:"操作员组ID"`
	PrimaryAccount int8   `orm:"column(primary_account)" description:"是否是主账号"`
}

func init() {
	orm.RegisterModel(new(JccBoss))
}

func CheckExistByBoss(params string) (count int,err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_boss where boss = ? and is_del = 0`
	err = o.Raw(sql,params).QueryRow(&count)
	return count,err
}

func CheckExistByIDCard(params string) (count int,err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_boss where id_card = ? and is_del = 0`
	err = o.Raw(sql,params).QueryRow(&count)
	return count,err
}

func AddBoss(o orm.Ormer,params *JccBoss) (id int64,err error) {
	id,err = o.Insert(params)
	return id,err
}

func QueryCompanyIdById(id int) (companyid int,err error) {
	o := orm.NewOrm()
	sql := `select aff_company from jcc_boss where id = ? and primary_account = 1`
	err = o.Raw(sql,id).QueryRow(&companyid)
	fmt.Println(err)
	return companyid,err
}

func DelBoss(o orm.Ormer,params *JccBoss) error {
	_,err := o.Update(params,"is_del","deleted_at")
	return err
}

func QueryCompanylistIdById(id int) (list JccBoss,err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_boss where id = ? and is_del = 0 and primary_account = 1`
	err = o.Raw(sql,id).QueryRow(&list)
	return list,err
}

func UpdateBoss(o orm.Ormer,params *JccBoss) error {
	_,err := o.Update(params)
	return err
}

func QueryBossAlllistIdById(id int) (list JccBoss,err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_boss where id = ? `
	err = o.Raw(sql,id).QueryRow(&list)
	return list,err
}