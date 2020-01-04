package order

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/util"
	"new_erp_agent_by_go/models/purchaseOrder"
	"new_erp_agent_by_go/models/putOrder"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api"
	"strconv"
	"time"
)

func UpdateOrderTotalMoney(info *request.PurchaseOrderInfo) error {
	//查询订单编号对应的订单id
	orderId, err := purchaseOrder.QueryOrderIdByOrderSn(info.OrderSn)
	if err != nil {
		return errors.New("订单编号不存在。。。" + err.Error())
	}

	//修改订单总金额
	db := orm.NewOrm()
	db.Begin()
	_, err = purchaseOrder.UpdatePurchaseOrderTotleMoney(info, db)
	if err != nil {
		db.Rollback()
		return errors.New("修改订单失败,请检查订单编号或者重复修改了。。。" + err.Error())
	}
	//修改商品订单信息
	for _, v := range info.GoodsInfo {
		//校验数据
		count, err := purchaseOrder.CheckOrder(v, orderId)
		if err != nil {
			db.Rollback()
			return errors.New("参数有错误，请检查商品id和订单编号。。。" + err.Error())
		}

		//修改条件不存在
		if count == 0 {
			db.Rollback()
			return errors.New("参数有错误，请检查商品id和订单编号。。。")
		}

		_, err = purchaseOrder.UpdateGoodOrderTotalMoney(v, orderId, db)
		if err != nil {
			db.Rollback()
			return errors.New("修改商品订单失败，请检查商品id和订单编号。。。" + err.Error())
		}
	}
	db.Commit()
	return nil
}

func UpdateReturnOrderTotalMoney(info *request.PurchaseOrderInfo) error {
	//查询订单编号对应的订单id
	orderId, err := purchaseOrder.QueryReturnOrderIdByOrderSn(info.OrderSn)
	if err != nil {
		return errors.New("订单编号不存在。。。" + err.Error())

	}
	//修改订单总金额
	db := orm.NewOrm()
	db.Begin()
	_, err = purchaseOrder.UpdateReturnOrderTotleMoney(info, db)
	if err != nil {
		db.Rollback()
		return errors.New("修改订单失败,请检查订单编号或者重复修改了。。。" + err.Error())
	}
	//修改商品订单信息
	for _, v := range info.GoodsInfo {
		//校验数据
		count, err := purchaseOrder.CheckReturnOrder(v, orderId)
		if err != nil {
			db.Rollback()
			return errors.New("参数有错误，请检查商品id和订单编号。。。" + err.Error())
		}

		//修改条件不存在
		if count == 0 {
			db.Rollback()
			return errors.New("参数有错误，请检查商品id和订单编号。。。")
		}

		_, err = purchaseOrder.UpdateReturnGoodOrderTotalMoney(v, orderId, db)
		if err != nil {
			db.Rollback()
			return errors.New("修改商品订单失败，请检查商品id和订单编号或者重复修改了。。。" + err.Error())
		}
	}
	db.Commit()
	return nil
}

func UpdatePurchaseOrdersStatusById(params *request.PurchaseOrderStatus) error {

	//审批驳回
	if params.Status == -1 {
		err := purchaseOrder.RejectPurchaseOrderById(params.PurchaseOrderId)
		if err != nil {
			return errors.New("订单驳回失败。。。")
		}
	}
	//审批通过
	if params.Status == 2 {

		//查询代理进货通过id
		list, err := purchaseOrder.QueryPurchaseOrderById(params.PurchaseOrderId)
		if err != nil {

			return errors.New("订单查询通过失败。。。")
		}

		//改商品id

		//生成入库单
		var putOrderParams = new(putOrder.JccPutOrder)
		RDH := util.CreatePurchaseOrder()
		putOrderParams.RDh = RDH
		putOrderParams.Fromorder = list.Ordersn
		putOrderParams.Fromordertype = 1
		putOrderParams.RStatus = 6
		putOrderParams.OrderMoney = list.OrderMoney
		putOrderParams.House1 = list.CompanyroomId
		putOrderParams.House = 0
		putOrderParams.Shop = strconv.Itoa(list.AffCompany)
		putOrderParams.Remark = list.Remark
		putOrderParams.CreatedAt = time.Now().Unix()
		putOrderParams.Auditor = list.BossId
		putOrderParams.CustomerId = 1
		//1是集餐厨
		putOrderParams.RPeople = 1
		putOrderParams.Operator = list.BossId

		//查询代理进货商品表
		goods_list, err := purchaseOrder.QueryPurchaseGoods(params.PurchaseOrderId)

		if err != nil {

			return errors.New("查询代理进货商品失败。。。")
		}

		//查询是否已经添加
		poExit := putOrder.JccPutOrder{Fromorder: putOrderParams.Fromorder}
		exist := putOrder.QueryPutOrderExitByOrderId(poExit)

		db := orm.NewOrm()
		db.Begin()

		//更新代理进货状态
		err = purchaseOrder.AgreePurchaseOrderById(params.PurchaseOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("订单审批通过失败。。。")
		}
		//报错，就是不存在，需要添加入库单
		if !exist {
			_, err := putOrder.AddPutOrder(db, putOrderParams)
			if err != nil {
				db.Rollback()
				return errors.New("出库单生成失败。。。")
			}

			//	return nil
			//入库商品单
			for _, purchasegoods := range goods_list {
				//增加或修改jcc_goods

				iddss, err := api.AddOnlineOneGoodsInfo(strconv.Itoa(purchasegoods.GoodsId), strconv.Itoa(list.BossId))

				if err != nil {
					db.Rollback()
					return errors.New("新增商品失败。。。" + err.Error())
				}
				//修改代理进货

				err = purchaseOrder.UpdatePurchaseGoods(db, params.PurchaseOrderId, iddss, purchasegoods.GoodsId)
				if err != nil {
					db.Rollback()
					return errors.New("修改代理进货失败。。。" + err.Error())
				}
				//构造入库商品信息
				putItem := new(putOrder.JccPutItem)
				putItem.UnitId = purchasegoods.UnitId
				putItem.GoodsId = int(iddss)
				putItem.Number = purchasegoods.Num
				putItem.Rdh = RDH
				putItem.Price = purchasegoods.BuyingPrice
				putItem.Supplier = strconv.Itoa(1)
				putItem.Money = purchasegoods.Money
				putItem.Discount = purchasegoods.Discount

				//存入数据库
				err = putOrder.AddPutItem(db, putItem)
				if err != nil {
					db.Rollback()
					return errors.New("出库商品单生成失败。。。")
				}
			}
		}

		db.Commit()
		return nil
	}
	//出库
	if params.Status == 3 {
		db := orm.NewOrm()
		db.Begin()
		//更新代理进货单状态
		err := purchaseOrder.CompletePurchaseOrderById(db, params.PurchaseOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("代理进货单出库失败失败。。。")
		}
		//查询代理进货单
		list, err := purchaseOrder.QueryPurchaseOrderById(params.PurchaseOrderId)
		if err != nil {
			db.Rollback()
			return errors.New("代理进货单查询失败。。。")
		}
		//更新入库单状态
		err = putOrder.UpdatePutOrder(db, list.Ordersn)
		if err != nil {
			db.Rollback()
			return errors.New("代理商品进货单查询失败。。。")
		}
		db.Commit()
		return nil
	}
	return nil
}
