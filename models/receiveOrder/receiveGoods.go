package receiveOrder

import "github.com/astaxie/beego/orm"



type JccReceiveGoods struct {
	Id          int     `orm:"column(id);auto"`
	GoodsId     int     `orm:"column(goods_id)" description:"商品id"`
	OrderId     int     `orm:"column(order_id)" description:"订单id"`
	SupplierId  int     `orm:"column(supplier_id)" description:"供应商id"`
	BuyingPrice float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"单价"`
	Num         float64 `orm:"column(num);null;digits(11);decimals(2)" description:"数量"`
	UnitId      int     `orm:"column(unit_id)" description:"单位id"`
	Money       float64 `orm:"column(money);null;digits(11);decimals(2)" description:"金额"`
	CreatedAt   int64     `orm:"column(created_at);null" description:"创建时间"`
	UpdatedAt   int     `orm:"column(updated_at);null" description:"修改时间"`
	IsDel       int8    `orm:"column(is_del);null" description:"是否删除"`
}

func init() {
	orm.RegisterModel(new(JccReceiveGoods))
}

func AddGiveGoods(db orm.Ormer,params *JccReceiveGoods) error {

	_, err := db.Insert(params)
	return err
}
