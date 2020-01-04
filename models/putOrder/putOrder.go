package putOrder

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccPutOrder struct {
	Id               int     `orm:"column(id);auto" description:"ID"`
	RDh              string  `orm:"column(r_dh);size(30)" description:"订单号"`
	Fromorder        string  `orm:"column(fromorder);size(30)" description:"来源单据"`
	Fromordertype    int     `orm:"column(fromordertype)" description:"来源单据类型"`
	RPeople          int     `orm:"column(r_people)" description:"来源单据制单人"`
	RStatus          int16   `orm:"column(r_status);null" description:"订单状态"`
	OrderMoney       float64 `orm:"column(order_money);null" description:"商品总价"`
	Shop             string  `orm:"column(shop);size(22);null" description:"公司"`
	House            int     `orm:"column(house);null" description:"出库仓库ID"`
	House1           int     `orm:"column(house1);null" description:"入库仓库"`
	ShippPeople      string  `orm:"column(shipp_people);size(30);null" description:"收货人"`
	Freconst         int     `orm:"column(freconst);null" description:"运费"`
	Remark           string  `orm:"column(remark);null" description:"备注"`
	CreatedAt        int64   `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt        int     `orm:"column(updated_at);null" description:"修改时间"`
	DeletedAt        int     `orm:"column(deleted_at);null" description:"删除时间"`
	IsDel            int     `orm:"column(is_del);null" description:"是否删除"`
	Auditor          int     `orm:"column(auditor);null" description:"来源单据审核人"`
	Operator         int     `orm:"column(operator);null" description:"操作人"`
	PaymentType      int8    `orm:"column(payment_type);null" description:"应付款状态"`
	CustomerId       int     `orm:"column(customer_id);null" description:"客户id"`
	FcgCheckPayables int8    `orm:"column(fcg_check_payables);null" description:"针对采购设计（内部应付款判断）"`
}

func init() {
	orm.RegisterModel(new(JccPutOrder))
}

func QueryPutOrderExitByOrderId(OrderId JccPutOrder) bool {
	o := orm.NewOrm()
	//err := o.Read(&OrderId, "fromorder")
	return o.QueryTable("jcc_put_order").Filter("fromorder", OrderId.Fromorder).Exist()
}

func AddPutOrder(db orm.Ormer, params *JccPutOrder) (id int64, err error) {

	putOrderId, err := db.Insert(params)
	return putOrderId, err
}

//func QueryPutOrderExitByOrdersn(Ordersn string) (id int64,err error) {
//	o := orm.NewOrm()
//	sql := `select id from jcc_put_order where fromorder = ?`
//	o.Raw(sql,Ordersn).QueryRow(&id)
//	return id,err
//}

func UpdatePutOrder(db orm.Ormer, fromorder string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("jccPutOrder").Filter("fromorder", fromorder).Update(orm.Params{
		"r_status":   0,
		"updated_at": time.Now().Unix(),
	})
	return err
}
