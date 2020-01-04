package print

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type JccCommonMould struct {
	Id          int    `orm:"column(id)" description:"ID" json:"id"`
	FieldOrder  string `orm:"column(field_order)" description:"字段顺序" json:"field_order"`
	MouldWeight int    `orm:"column(mould_weight)" description:"模板宽度" json:"mould_weight"`
	MouldHeight int    `orm:"column(mould_height)" description:"模板高度" json:"mould_height"`
	PrintList   int    `orm:"column(print_list)" description:"打印条数" json:"print_list"`
	CreatedAt   int    `orm:"column(created_at)" description:"创建时间" json:"created_at"`
	UpdatedAt   int    `orm:"column(updated_at)" description:"修改时间" json:"updated_at"`
	IsDefault   int    `orm:"column(is_default)" description:"是否是默认模板" json:"is_default"`
	IsDel       int    `orm:"column(is_del)" description:"是否删除" json:"is_del"`
}

func init() {
	orm.RegisterModel(new(JccCommonMould))
}

// 查询一条
func SelectMould(ormer orm.Ormer, Id int) (info *JccCommonMould, err error) {
	var mould = new(JccCommonMould)
	err = ormer.QueryTable("jcc_common_mould").Filter("id", Id).Filter("is_del", 0).One(mould)
	return mould, err
}

// 查询列表
func SelectMouldList(ormer orm.Ormer) (info *[]JccCommonMould, err error) {
	var mould = new([]JccCommonMould)
	fmt.Println("+++")
	_, err = ormer.QueryTable("jcc_common_mould").Filter("is_del", 0).All(mould)
	return mould, err
}
