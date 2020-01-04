package receiveOrder

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccReceiveOrder struct {
	Id            int     `orm:"column(id);auto"`
	BossId        int     `orm:"column(boss_id)" description:"用户id"`
	PurchaseId    string  `orm:"column(purchase_id);size(255);null" description:"采购订单id"`
	Ordersn       string  `orm:"column(ordersn);size(255)" description:"获赠订单编号"`
	BusinessType  int8    `orm:"column(business_type);null" description:"业务类型(0:默认获赠)"`
	AffCompany    int     `orm:"column(aff_company);null" description:"公司id"`
	CompanyroomId int     `orm:"column(companyroom_id)" description:"仓库id"`
	SupplierId    int     `orm:"column(supplier_id)" description:"供应商id"`
	Type          int8    `orm:"column(type);null" description:"订单状态"`
	Status        int8    `orm:"column(status);null" description:"是否提交"`
	OrderMoney    float64 `orm:"column(order_money);digits(11);decimals(2)" description:"订单金额"`
	Savetime      int64     `orm:"column(savetime);null" description:"保存时间"`
	Remark        string  `orm:"column(remark);size(255);null" description:"备注"`
	CreatedAt     int64     `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt     int     `orm:"column(updated_at);null" description:"修改时间"`
	DeletedAt     int     `orm:"column(deleted_at);null" description:"删除时间"`
	IsDel         int8    `orm:"column(is_del);null" description:"是否删除"`
	Kind          string  `orm:"column(kind);size(20);null" description:"区分本地或代理（本地(offline)、代理(online)）默认的是本地(offline)"`
}

func init() {
	orm.RegisterModel(new(JccReceiveOrder))
}

func AddGiveOrder(db orm.Ormer,params *JccReceiveOrder) (id int64,err error) {
	putOrderId, err := db.Insert(params)
	return putOrderId,err
}

func QueryGiveOrderExsitByOrdersn(ordersn string) (count int,err error) {
	o := orm.NewOrm()
	sql := `select count(*) from jcc_receive_order where ordersn = ?`
	o.Raw(sql,ordersn).QueryRow(&count)
	return count,err
}

func CompletePurchaseOrderByOrdersn(db orm.Ormer,ordersn string) error {

	_, err := db.QueryTable("JccReceiveOrder").Filter("ordersn", ordersn).Update(orm.Params{
		"status": 1,
		"type" : 1,
		"updated_at": time.Now().Unix(),
	})
	return err
}

