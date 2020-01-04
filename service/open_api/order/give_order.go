package order

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/util"
	"new_erp_agent_by_go/models/centerGoods"
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/goods"
	"new_erp_agent_by_go/models/jccBoss"
	"new_erp_agent_by_go/models/putOrder"
	"new_erp_agent_by_go/models/receiveOrder"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"time"
)

func UpdateGiveOrdersStatusById(params *request.GiveOrderStatus) error {
	//审批通过
	if params.Status == 2 {
		db := orm.NewOrm()
		db.Begin()
		//查询代理进货单是否存在
		//poExit := purchaseOrder.JccPurchaseOrder{Id:params.PurchaseOrderId}
		//err := purchaseOrder.CheckExistById(poExit)
		//if err != nil {
		//	db.Rollback()
		//	return errors.New("代理进货单不存在。。。")
		//}
		//查询代理进货单参数
		//purchaseOrder_list,err := purchaseOrder.QueryPurchaseOrderById(params.PurchaseOrderId)
		//if err != nil {
		//	db.Rollback()
		//	return errors.New("获取代理进货单信息失败。。。")
		//}

		RDH := util.CreateGiveOrder()
		//生成赠送单
		//根据ordersn查看获赠单是否已经添加
		count, err := receiveOrder.QueryGiveOrderExsitByOrdersn(params.Fromorder)
		if err != nil {
			db.Rollback()
			return errors.New("赠送单查询失败。。。")
		}
		if count > 0 {
			db.Rollback()
			return errors.New("赠送单已存在。。。")
		}
		boss_companyid, err := jccBoss.QueryCompanyIdById(params.BossId)
		if err != nil {
			db.Rollback()
			return errors.New("查询失败。。。")
		}
		var giveOrderParams = new(receiveOrder.JccReceiveOrder)
		giveOrderParams.BossId = params.BossId
		giveOrderParams.Ordersn = params.Fromorder
		giveOrderParams.Status = 1
		giveOrderParams.CompanyroomId = 0
		giveOrderParams.SupplierId = 1
		giveOrderParams.Type = 3
		giveOrderParams.Kind = "online"
		giveOrderParams.AffCompany = boss_companyid
		giveOrderParams.CreatedAt = time.Now().Unix()
		giveOrderParams.Savetime = time.Now().Unix()
		giveOrderid, err := receiveOrder.AddGiveOrder(db, giveOrderParams)

		if err != nil {
			db.Rollback()
			return errors.New("赠送单生成失败。。。")
		}
		//生成入库单
		var putOrderParams = new(putOrder.JccPutOrder)
		putOrderParams.RDh = RDH
		putOrderParams.Fromorder = params.Fromorder
		putOrderParams.Fromordertype = 4
		putOrderParams.RStatus = 6
		putOrderParams.OrderMoney = params.OrderMoney
		//putOrderParams.House1 = params.CompanyroomId
		putOrderParams.House = 0
		putOrderParams.Shop = strconv.Itoa(boss_companyid)
		putOrderParams.Remark = params.Remark
		putOrderParams.CreatedAt = time.Now().Unix()
		putOrderParams.Auditor = params.BossId
		putOrderParams.CustomerId = params.BossId
		putOrderParams.RPeople = 1
		putOrderParams.Operator, _ = strconv.Atoi(params.Approver)
		//查询是否已经添加
		poExsit := putOrder.JccPutOrder{Fromorder: putOrderParams.Fromorder}
		exist := putOrder.QueryPutOrderExitByOrderId(poExsit)
		//报错，就是不存在，需要添加入库单，赠送单
		if !exist {
			//入库单
			_, err := putOrder.AddPutOrder(db, putOrderParams)
			if err != nil {
				db.Rollback()
				return errors.New("入库单生成失败。。。")
			}
			//赠品：代理商必须有，如果未进货或者删除，则报错
			//加入判断中，否则会出现bug
			for id, num := range params.Gift {
				goodsid := id
				center_goods_list, err := centerGoods.QueryCenterGoodsById(goodsid)
				if err != nil {
					db.Rollback()
					return errors.New("查询商品中心失败。。。")
				}
				//fmt.Println(center_goods_list)
				//查询该公司是否存在该商品
				goods_Exist := childGoods.JccGoods{Sku: center_goods_list.Sku, Companyid: boss_companyid, IsDel: 0}
				gid, err := goods.QueryExistBySkuAndCompany(goods_Exist)
				if err != nil {
					db.Rollback()
					fmt.Println(err)
					return errors.New("该公司不存在该商品或已删除。。。")
				}

				//查询商品信息
				goods_list, err := goods.QuerygoodsById(gid)
				if err != nil {
					db.Rollback()
					return errors.New("查询商品失败。。。")
				}
				//构建赠送商品单
				var giveGoodsParams = new(receiveOrder.JccReceiveGoods)
				giveGoodsParams.GoodsId = int(gid)
				giveGoodsParams.OrderId = int(giveOrderid)
				giveGoodsParams.BuyingPrice = goods_list.BuyingPrice
				giveGoodsParams.Num = float64(num)

				giveGoodsParams.UnitId = goods_list.UnitId
				giveGoodsParams.Money = 0
				giveGoodsParams.CreatedAt = time.Now().Unix()
				giveGoodsParams.SupplierId = 1

				//存入数据库
				err = receiveOrder.AddGiveGoods(db, giveGoodsParams)
				if err != nil {
					db.Rollback()
					return errors.New("赠送商品单生成失败。。。")
				}
				//构建入库商品单
				var putItemParams = new(putOrder.JccPutItem)
				putItemParams.UnitId = goods_list.UnitId
				putItemParams.GoodsId = int(gid)
				putItemParams.Number = float64(num)
				putItemParams.Rdh = RDH
				putItemParams.Price = goods_list.BuyingPrice
				putItemParams.Supplier = strconv.Itoa(1)
				putItemParams.Money = 0

				//存入数据库
				err = putOrder.AddPutItem(db, putItemParams)
				if err != nil {
					db.Rollback()
					return errors.New("出库商品单生成失败。。。")
				}

				////增加或修改jcc_goods(防止事务套事务)
				//err = api.AddOnlineGoodsInfo(id,strconv.Itoa(boss_companyid))
				//if err != nil {
				//	db.Rollback()
				//	return errors.New("新增商品失败。。。")
				//}
			}
		}

		db.Commit()
		return nil
	}
	//出库
	if params.Status == 3 {
		db := orm.NewOrm()
		db.Begin()
		//更新赠送单状态
		err := receiveOrder.CompletePurchaseOrderByOrdersn(db, params.Fromorder)
		if err != nil {
			db.Rollback()
			return errors.New("赠送单状态修改失败。。。")
		}
		//更新入库单状态
		err = putOrder.UpdatePutOrder(db, params.Fromorder)
		if err != nil {
			db.Rollback()
			return errors.New("代理进货单状态修改失败。。。")
		}
		db.Commit()
		return nil
	}
	return nil
}
