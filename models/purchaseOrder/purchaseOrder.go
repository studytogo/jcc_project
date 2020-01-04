package purchaseOrder

import (
	"errors"
	"new_erp_agent_by_go/models/request"
	"time"

	"github.com/astaxie/beego/orm"
)

type JccPurchaseOrder struct {
	Id            int64   `orm:"column(id);auto"`
	BossId        int     `orm:"column(boss_id)" description:"用户id"`
	Ordersn       string  `orm:"column(ordersn);size(255)" description:"订单编号"`
	BusinessType  int8    `orm:"column(business_type);null" description:"业务类型(0:默认为采购;-1为驳回;2:为异常)"`
	Address       string  `orm:"column(address);size(255);null" description:"收货地址"`
	Telphone      string  `orm:"column(telphone);size(30);null" description:"收货电话"`
	CompanyroomId int     `orm:"column(companyroom_id);null" description:"仓库id"`
	OrderMoney    float64 `orm:"column(order_money);null;digits(11);decimals(2)" description:"订单金额"`
	AffCompany    int     `orm:"column(aff_company);null" description:"公司id"`
	Status        int8    `orm:"column(status);null" description:"是否提交草稿箱"`
	Type          int8    `orm:"column(type);null" description:"审核入库('0':未审核;'1':待入库;'2':已完成)"`
	Approver      string  `orm:"column(approver);size(255)" description:"审批人"`
	Savetime      int     `orm:"column(savetime);null" description:"审核时间"`
	Remark        string  `orm:"column(remark);size(255);null" description:"备注"`
	CreatedAt     int     `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt     int     `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt     int     `orm:"column(deleted_at)" description:"删除时间"`
	IsDel         int8    `orm:"column(is_del);null" description:"是否删除"`
	Kind          string  `orm:"column(kind);size(20);null" description:"区分本地或代理（本地(offline)、代理(online)）默认的是本地(offline)"`
}

type CheckPurchaseOrder struct {
	Id             int    `orm:"column(id);auto"`
	Ordersn        string `orm:"column(ordersn);size(255)" description:"订单编号"`
	Printing_times int    `orm:"column(printing_times)" description:"打印次数"`
}

type T struct {
	Data      []CheckPurchaseOrder
	Total     int
	Last_page int
}

func init() {
	orm.RegisterModel(new(JccPurchaseOrder))
}

func AddJccPurchaseOrder(m *JccPurchaseOrder, o orm.Ormer) (id int64, err error) {
	id, err = o.Insert(m)
	return
}

func UpdateLog(OrderIDs string, o orm.Ormer) error {

	sql := `UPDATE jcc_purchase_order SET printing_times = printing_times + 1 WHERE id IN ( ` + OrderIDs + `)`

	_, err := o.Raw(sql).Exec()

	return err
}

func CheckList(company_id string, orderid string, start_page int, end_page int, page int, o orm.Ormer) ([]CheckPurchaseOrder, int, error) {

	var PurchaseList []CheckPurchaseOrder
	var count int
	sql := ``
	sql1 := ` `
	if orderid != "" {
		sql = `SELECT id,ordersn,printing_times FROM jcc_purchase_order WHERE business_type=0 and type in(2,3) and aff_company=` + company_id + ` and ordersn like '%` + orderid + `%' and is_del=0 ORDER BY created_at desc limit ?,?`

		sql1 = `SELECT count(*) FROM jcc_purchase_order WHERE business_type=0 and type in(2,3) and aff_company=` + company_id + ` and ordersn like '%` + orderid + `%' and is_del=0`
	} else {
		sql = `SELECT id,ordersn,printing_times FROM jcc_purchase_order WHERE business_type=0 and type in(2,3) and aff_company=` + company_id + ` and is_del=0 ORDER BY created_at desc limit ?,?`

		sql1 = `SELECT count(*) FROM jcc_purchase_order WHERE business_type=0 and type in(2,3) and aff_company=` + company_id + ` and is_del=0`
	}

	_, err := o.Raw(sql, start_page, end_page).QueryRows(&PurchaseList)
	err = o.Raw(sql1).QueryRow(&count)
	return PurchaseList, count, err
}

//通过jcc_purchase_goods查询goods_id的总数
func QueryGoodsIdTotalByOrderId(orderIds string) (total int, err error) {
	db := orm.NewOrm()

	sql := `SELECT count(*) from jcc_purchase_goods WHERE order_id in(` + orderIds + `)`

	err = db.Raw(sql).QueryRow(&total)

	if err != nil {
		return 0, errors.New("通过orderId查询goodId的总数量失败。。。")
	}

	return total, nil
}

//去重复操作并取查询中最新的一条数据
func QueryNewOrderId(orderIds string, page, pageSize, isPage int) (ids []int, err error) {
	db := orm.NewOrm()

	sql := `SELECT max(id) from jcc_purchase_goods WHERE order_id in(` + orderIds + `) GROUP BY goods_id `

	if isPage == 1 {
		sql += ` limit ?, ?`
		_, err = db.Raw(sql, (page-1)*pageSize, pageSize).QueryRows(&ids)
	} else {
		_, err = db.Raw(sql).QueryRows(&ids)
	}

	if err != nil {
		return nil, errors.New("查询orderIds失败。。。")
	}

	return ids, nil
}

//查询该订单的商品所在的仓库id
func QueryOrderCompanyIdByOrderId(orderId int) (companyRoomId int, err error) {
	db := orm.NewOrm()

	sql := `SELECT companyroom_id from jcc_purchase_order WHERE id = ?`

	err = db.Raw(sql, orderId).QueryRow(&companyRoomId)

	return companyRoomId, err
}

//修改订单总价
func UpdatePurchaseOrderTotleMoney(info *request.PurchaseOrderInfo, db orm.Ormer) (int64, error) {
	return db.QueryTable("jcc_purchase_order").Filter("ordersn", info.OrderSn).Update(orm.Params{
		"order_money": info.TotalMoney,
		"updated_at":  time.Now().Unix(),
	})
}

//通过订单编号查询订单id
func QueryOrderIdByOrderSn(orderSn string) (int, error) {
	db := orm.NewOrm()

	sql := `select id from jcc_purchase_order where ordersn = ? `

	var orderId int

	err := db.Raw(sql, orderSn).QueryRow(&orderId)

	return orderId, err
}

//修改订单商品总价
func UpdateGoodOrderTotalMoney(info request.GoodsOrderInfo, orderId int, db orm.Ormer) (int64, error) {

	sql := `update jcc_purchase_goods set money = ? , updated_at = ? where order_id = ? and goods_id = ?`

	num, err := db.Raw(sql, info.GoodTotalMoney, time.Now().Unix(), orderId, info.GoodId).Exec()

	u, _ := num.RowsAffected()

	return u, err
}

func CheckOrder(info request.GoodsOrderInfo, orderId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_purchase_goods where order_id = ? and goods_id = ?`

	var count int

	err := db.Raw(sql, orderId, info.GoodId).QueryRow(&count)

	return count, err
}

//修改退货订单总价
func UpdateReturnOrderTotleMoney(info *request.PurchaseOrderInfo, db orm.Ormer) (int64, error) {

	sql := `update jcc_purchase_return_order set order_money = ? and updated_at = ? where ordersn = ? `

	num, err := db.Raw(sql, info.TotalMoney, time.Now().Unix(), info.OrderSn).Exec()

	u, _ := num.RowsAffected()

	return u, err
}

//通过订单编号查询订单id
func QueryReturnOrderIdByOrderSn(orderSn string) (int, error) {
	db := orm.NewOrm()

	sql := `select id from jcc_purchase_return_order where ordersn = ? `

	var orderId int

	err := db.Raw(sql, orderSn).QueryRow(&orderId)

	return orderId, err
}

//修改订单商品总价
func UpdateReturnGoodOrderTotalMoney(info request.GoodsOrderInfo, orderId int, db orm.Ormer) (int64, error) {

	sql := `update jcc_purchase_return_goods set money = ? and updated_at = ? where order_id = ? and goods_id = ?`

	num, err := db.Raw(sql, info.GoodTotalMoney, time.Now().Unix(), orderId, info.GoodId).Exec()

	u, _ := num.RowsAffected()

	return u, err
}

func CheckReturnOrder(info request.GoodsOrderInfo, orderId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_purchase_return_goods where order_id = ? and goods_id = ?`

	var count int

	err := db.Raw(sql, orderId, info.GoodId).QueryRow(&count)

	return count, err
}

func RejectPurchaseOrderById(orderid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("jccPurchaseOrder").Filter("id", orderid).Update(orm.Params{
		"business_type": -1,
		"status":        0,
		"type":          0,
	})
	return err
}

func AgreePurchaseOrderById(orderid int64) error {
	db := orm.NewOrm()
	_, err := db.QueryTable("jccPurchaseOrder").Filter("id", orderid).Update(orm.Params{
		"type":   3,
		"status": 1,
	})
	return err
}

func QueryPurchaseOrderById(orderid int64) (list JccPurchaseOrder, err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_purchase_order where id=?`
	o.Raw(sql, orderid).QueryRow(&list)
	return list, err
}

func CompletePurchaseOrderById(db orm.Ormer, orderid int64) error {

	_, err := db.QueryTable("jccPurchaseOrder").Filter("id", orderid).Update(orm.Params{
		"type":       4,
		"updated_at": time.Now().Unix(),
	})
	return err
}

func CheckExistById(params JccPurchaseOrder) error {
	o := orm.NewOrm()
	err := o.Read(&params, "id")
	return err
}
