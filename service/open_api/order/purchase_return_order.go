package order

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/util"
	"new_erp_agent_by_go/models/outOrder"
	"new_erp_agent_by_go/models/purchaseReturnOrder"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"time"
)

func UpdatePurchaseReturnOrdersStatusById(params *request.PurchaseReturnOrderStatus) error {
	if params.Status == -1 {
		err := purchaseReturnOrder.RejectPurchaseReturnOrderById(params.PurchaseReturnOrderId)
		if err != nil {
			return errors.New("订单驳回失败。。。")
		}
	}
	if params.Status == 2 {
		db := orm.NewOrm()
		db.Begin()

		//更新代理退货状态
		err := purchaseReturnOrder.AgreePurchaseReturnOrderById(db,params.PurchaseReturnOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("订单审批通过失败。。。")
		}

		//查询代理退货通过id
		list,err := purchaseReturnOrder.QueryPurchaseReturnOrderById(params.PurchaseReturnOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("订单查询通过失败。。。")
		}

		//生成出库单
		var outOrderParams = new(outOrder.JccOutOrder)
		RDH := util.CreatePurchaseReturnOrder()
		outOrderParams.RDh = RDH
		outOrderParams.Fromorder = list.Ordersn
		outOrderParams.Fromordertype = 1
		outOrderParams.RPeople = list.BossId
		outOrderParams.RStatus = 0
		outOrderParams.OrderMoney = list.OrderMoney
		outOrderParams.Shop = strconv.Itoa(list.AffCompany)
		outOrderParams.House = list.CompanyroomId
		outOrderParams.ShippPeople = "集餐厨"
		outOrderParams.Remark = list.Remark
		outOrderParams.CreatedAt = time.Now().Unix()
		outOrderParams.Auditor,_ = strconv.Atoi(list.Approver)
		outOrderParams.Operator = 0
		outOrderParams.ShippPeople = "集餐厨"
		outOrderParams.ShippAddress = "集餐厨"
		outOrderParams.ShippPhone = "0"
		outOrderParams.PayMethod = "0"
		outOrderParams.DeliveryMethod = 0
		outOrderParams.DeliveryPeople = "0"
		outOrderParams.PaymentType = 0
		outOrderParams.CustomerId = 0
		outOrderParams.SupplierId = 0

		//查询代理退货商品表
		goods_list,err := purchaseReturnOrder.QueryPurchaseReturnGoods(params.PurchaseReturnOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("查询代理退货商品失败。。。")
		}

		//查询是否已经添加
		poExit := outOrder.JccOutOrder{Fromorder: outOrderParams.Fromorder}
		err = outOrder.QueryOutOrderExitByOrderId(poExit)
		//报错，就是不存在，需要添加出库单
		if err != nil {
			_,err := outOrder.AddOutOrder(db,outOrderParams)
			if err != nil {
				db.Rollback()
				return errors.New("出库单生成失败。。。")
			}
			//出库商品单
			for _,purchasereturngoods := range goods_list{
				//构造入库商品信息
				outItem := new(outOrder.JccOutItem)
				outItem.UnitId = purchasereturngoods.UnitId
				outItem.GoodsId = purchasereturngoods.GoodsId
				outItem.Number = purchasereturngoods.Num
				outItem.Rdh = RDH
				outItem.Price = purchasereturngoods.BuyingPrice
				outItem.Supplier = strconv.Itoa(1)
				outItem.Money = purchasereturngoods.Money
				outItem.Discount = purchasereturngoods.Discount

				//存入数据库
				err = outOrder.AddOutItem(db,outItem)
				if err != nil {
					db.Rollback()
					return errors.New("出库商品单生成失败。。。")
				}
			}
		}

		db.Commit()
		return nil
	}
	return nil
}
