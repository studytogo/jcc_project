package outOrder

import "github.com/astaxie/beego/orm"

type JccOutItem struct {
	Id                  int     `orm:"column(id);auto" description:"ID"`
	GoodsId             int     `orm:"column(goods_id)" description:"商品sku"`
	Number              float64 `orm:"column(number);null;digits(11);decimals(2)" description:"数量"`
	Rdh                 string  `orm:"column(rdh);size(30)" description:"关联出库单号"`
	Price               float64 `orm:"column(price);null" description:"商品单价"`
	Supplier            string  `orm:"column(supplier);size(22);null" description:"供应商"`
	Money               float64 `orm:"column(money);null;digits(10);decimals(2)" description:"实收金额"`
	Discount            string  `orm:"column(discount);size(255);null" description:"折扣率"`
	UnitId              int     `orm:"column(unit_id);null" description:"单位"`
	PredeterminedPrices string  `orm:"column(predetermined_prices);size(255);null" description:"价格体系"`
	UnitIdNum           float64 `orm:"column(unit_id_num);null;digits(11);decimals(2)" description:"原来单位数量"`
	ConversionMiniNum   string  `orm:"column(conversion_mini_num);size(255);null" description:"最小单位换算关系"`
}

func init() {
	orm.RegisterModel(new(JccOutItem))
}

func AddOutItem(db orm.Ormer,params *JccOutItem) error {

	_, err := db.Insert(params)
	return err
}
