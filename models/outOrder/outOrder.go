package outOrder

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type JccOutOrder struct {
	Id                  int       `orm:"column(id);auto" description:"ID"`
	RDh                 string    `orm:"column(r_dh);size(30)" description:"订单号"`
	Fromorder           string    `orm:"column(fromorder);size(30)" description:"来源单据"`
	Fromordertype       int       `orm:"column(fromordertype)" description:"来源单据类型"`
	RPeople             int       `orm:"column(r_people)" description:"制单人"`
	RStatus             int16     `orm:"column(r_status)" description:"订单状态"`
	OrderMoney          float64   `orm:"column(order_money);null" description:"订单金额"`
	Shop                string    `orm:"column(shop);size(22);null" description:"门店ID"`
	House               int     `orm:"column(house);null" description:"出库仓库ID"`
	House1              int     `orm:"column(house1);null" description:"入库仓库"`
	ShippPeople         string    `orm:"column(shipp_people);size(30);null" description:"收货人"`
	Freconst            int       `orm:"column(freconst);null" description:"运费"`
	Remark              string    `orm:"column(remark);null" description:"备注"`
	CreatedAt           int64       `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt           int       `orm:"column(updated_at);null" description:"修改时间"`
	DeletedAt           int       `orm:"column(deleted_at);null" description:"删除时间"`
	IsDel               int       `orm:"column(is_del);null" description:"是否删除"`
	Auditor             int       `orm:"column(auditor);null" description:"审核人"`
	Operator            int       `orm:"column(operator);null" description:"操作人"`
	ShippAddress        string    `orm:"column(shipp_address);null" description:"收货地址"`
	ShippPhone          string    `orm:"column(shipp_phone);size(15);null" description:"收货电话"`
	PayMethod           string    `orm:"column(pay_method);size(100);null" description:"付款方式"`
	DeliveryMethod      int8      `orm:"column(delivery_method);null" description:"送货方式"`
	DeliveryDate        time.Time `orm:"column(delivery_date);type(date);null" description:"期望送货日期"`
	DeliveryPeople      string    `orm:"column(delivery_people);size(30);null" description:"送货人"`
	PaymentType         int8      `orm:"column(payment_type)" description:"应收款状态"`
	CustomerId          int       `orm:"column(customer_id);null" description:"客户id"`
	SupplierId          int       `orm:"column(supplier_id);null" description:"供应商id"`
	FcgCheckReceivables int8      `orm:"column(fcg_check_receivables);null" description:"针对采购设计（内部应收款判断）"`
	IsSelf              int8      `orm:"column(is_self);null" description:"是否自提"`
}

func init() {
	orm.RegisterModel(new(JccOutOrder))
}

func QueryOutOrderExitByOrderId(OrderId JccOutOrder) error {
	o := orm.NewOrm()
	err := o.Read(&OrderId, "fromorder")
	return err
}

func AddOutOrder(db orm.Ormer,params *JccOutOrder) (id int64,err error) {

	putOrderId, err := db.Insert(params)
	return putOrderId,err
}
