package purchaseOrderService

import (
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/purchaseOrder"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/models/unit"
	"strconv"
	"strings"
)

//查询订单的商品信息
func QueryOrderGoodsInfoByOrderId(param *request.OrderGoodsInfoReq) (info []childGoods.OrderGoodsInfo, total int, err error) {
	//通过订单id查询商品id
	pageInt, err := strconv.Atoi(param.Page)
	if err != nil {
		return nil, 0, err
	}

	pageSizeInt, err := strconv.Atoi(param.PageSize)
	if err != nil {
		return nil, 0, err
	}

	isPage, _ := strconv.Atoi(param.IsPage)
	//先筛选重复的消息，取最新的一条
	orderIds, err := purchaseOrder.QueryNewOrderId(param.OrderIds, pageInt, pageSizeInt, isPage)
	if err != nil {
		return nil, 0, err
	}
	var orderIdStr []string
	for _, v := range orderIds {
		orderIdStr = append(orderIdStr, strconv.Itoa(v))
	}
	args := strings.Join(orderIdStr, ",")

	//查询数据的总数量
	total, err = purchaseOrder.QueryGoodsIdTotalByOrderId(param.OrderIds)
	if err != nil {
		return nil, 0, err
	}

	info, err = childGoods.QueryGoodsInfoByGoodIds(args)

	if err != nil {
		return nil, 0, err
	}

	for k, _ := range info {
		info[k].UnitName, _ = unit.QueryUnitNameByUnitId(info[k].UnitId)
		info[k].CompanyRoomId, _ = purchaseOrder.QueryOrderCompanyIdByOrderId(info[k].OrderId)
	}
	//通过单位id查询单位名称
	return info, total, err
}
