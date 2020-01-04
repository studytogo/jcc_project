package purchaseReturnOrder

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccPurchaseReturnOrder struct {
	Id            int     `orm:"column(id);auto"`
	BossId        int     `orm:"column(boss_id)" description:"用户id"`
	Ordersn       string  `orm:"column(ordersn);size(255)" description:"订单编号"`
	BusinessType  int8    `orm:"column(business_type);null" description:"业务类型(0:默认为采购退货;-1为驳回;2:为异常)"`
	CompanyroomId int     `orm:"column(companyroom_id)" description:"仓库id"`
	AffCompany    int     `orm:"column(aff_company);null" description:"公司id"`
	OrderMoney    float64 `orm:"column(order_money);digits(11);decimals(2)" description:"订单金额"`
	SupplierId    int     `orm:"column(supplier_id)" description:"供应商id"`
	Type          int8    `orm:"column(type);null" description:"0待审核;1待出库;2已完成"`
	Status        int8    `orm:"column(status);null" description:"是否提交草稿箱"`
	Approver      string  `orm:"column(approver);size(255);null" description:"审批人"`
	Savetime      int     `orm:"column(savetime);null" description:"审核时间"`
	Remark        string  `orm:"column(remark);size(255);null" description:"备注"`
	CreatedAt     int     `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt     int     `orm:"column(updated_at);null" description:"修改时间"`
	DeletedAt     int     `orm:"column(deleted_at);null" description:"删除时间"`
	IsDel         int8    `orm:"column(is_del);null" description:"是否删除"`
	Kind          string  `orm:"column(kind);size(20);null" description:"区分本地或代理（本地(offline)、代理(online)）默认的是本地(offline)"`
}

func init() {
	orm.RegisterModel(new(JccPurchaseReturnOrder))
}

func RejectPurchaseReturnOrderById(orderid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("jccPurchaseReturnOrder").Filter("id", orderid).Update(orm.Params{
		"business_type": -1,
		"status":        0,
		"type":          0,
		"updated_at":    time.Now().Unix(),
	})
	return err
}

func AgreePurchaseReturnOrderById(db orm.Ormer,orderid int64) error {

	_, err := db.QueryTable("jccPurchaseReturnOrder").Filter("id", orderid).Update(orm.Params{
		"type":   3,
		"updated_at": time.Now().Unix(),
	})
	return err
}

func QueryPurchaseReturnOrderById(orderid int64) (list JccPurchaseReturnOrder,err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_purchase_return_order where id=?`
	o.Raw(sql,orderid).QueryRow(&list)
	return list,err
}
