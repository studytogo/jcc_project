package purchaseOrder

import "github.com/astaxie/beego/orm"

type JccPurchaseGoods struct {
	Id          int     `orm:"column(id);auto"`
	GoodsId     int     `orm:"column(goods_id)" description:"商品id"`
	OrderId     int     `orm:"column(order_id)" description:"订单id"`
	SupplierId  int     `orm:"column(supplier_id)" description:"供应商id"`
	BuyingPrice float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"单价"`
	Num         float64 `orm:"column(num);null;digits(11);decimals(2)" description:"数量"`
	UnitId      int     `orm:"column(unit_id)" description:"单位id"`
	Discount    string  `orm:"column(discount);size(50);null" description:"折扣"`
	Money       float64 `orm:"column(money);null;digits(11);decimals(2)" description:"金额"`
	CreatedAt   int     `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt   int     `orm:"column(updated_at);null" description:"修改时间"`
	IsDel       int8    `orm:"column(is_del);null" description:"是否删除"`
}

func init() {
	orm.RegisterModel(new(JccPurchaseGoods))
}

func QueryPurchaseGoods(order_id int64)(list []JccPurchaseGoods,err error){
	o := orm.NewOrm()
	sql := `select * from jcc_purchase_goods where order_id = ?`
	_,err = o.Raw(sql,order_id).QueryRows(&list)

	return list,err
}

func UpdatePurchaseGoods(o orm.Ormer,order_id int64,goods_id int64, goodss_id int) error {

	sql := `update jcc_purchase_goods set goods_id = ? where order_id = ? and goods_id = ?`
	_,err := o.Raw(sql,goods_id,order_id,goodss_id).Exec()
	return err
}