package goodService

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/centerGoods"
	"new_erp_agent_by_go/models/erp_goods"
	"new_erp_agent_by_go/models/goods"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"strings"
)

//查询erp商品信息
func QueryErpGoodsInfo(param *request.AgentJinHuo) ([]erp_goods.ErpGoodsInfo, int, error) {
	//查询加盟商进货公司id
	condition := ``
	if param.CompanyId != "" {
		companyId, _ := strconv.Atoi(param.CompanyId)
		erpCompanyId, err := erp_goods.QueryErpCompanyIdByCompanyId(companyId)
		if err != nil {
			return nil, 0, errors.New("erp公司不存在。。。")
		}
		condition += ` where jcs.companyid = ` + strconv.Itoa(erpCompanyId)
	}

	if param.GoodArribute != "" {
		if condition != "" {
			condition += ` and (jcg.name LIKE '%` + param.GoodArribute + `%' or jcg.brand_name LIKE '%` + param.GoodArribute + `%' or jcg.barcode = '` + param.GoodArribute + `')`
		} else {
			condition += ` where jcg.name LIKE '%` + param.GoodArribute + `%' or jcg.brand_name LIKE '%` + param.GoodArribute + `%' or jcg.barcode = '` + param.GoodArribute + `'`
		}
	}

	pageInt, _ := strconv.Atoi(param.Page)
	pageSizeInt, _ := strconv.Atoi(param.PageSize)

	erpGoodsInfos, err := erp_goods.QueryErpGoodsByErpCompanyId(condition, pageInt, pageSizeInt)

	if err != nil {
		return nil, 0, errors.New("商品信息不存在。。。")
	}

	total, err := erp_goods.QueryErpGoodsCount(condition)

	if err != nil {
		return nil, 0, errors.New("查询总数错误。。。")
	}

	return erpGoodsInfos, total, nil

}

//同步加盟商商品信息
func SyncErpGoods(param *request.SyncErpToAgent) error {
	subs := strings.Split(param.GoodsAttribute, ",")

	//验证商品属性参数
	for _, v := range subs {
		if v == "" {
			continue
		}
		if v != "name" && v != "buying_price" && v != "unit_id" && v != "content" && v != "image" {
			return errors.New("参数错误。。。")
		}

	}

	erpIds := strings.Split(param.ErpIds, ",")
	db := orm.NewOrm()
	db.Begin()
	for _, erpid := range erpIds {
		if erpid == "" {
			continue
		}

		attribute, err := centerGoods.QuerySyncGoodAttribute(erpid)
		if err != nil {
			return errors.New("同步商品失败。。。。原因" + err.Error())
		}

		condition := `set `
		for k, one := range subs {
			if k == 0 {
				switch one {
				case "name":
					condition += " name = '" + attribute.Name + "'"
				case "buying_price":
					condition += " retail_price = " + strconv.FormatFloat(attribute.RetailPrice, 'f', 2, 64)
				case "unit_id":
					condition += " unit_id = " + strconv.Itoa(attribute.UnitId)
				case "content":
					condition += " content = '" + attribute.Content + "'"
				case "image":
					condition += " image = '" + attribute.Image + "' and images = '" + attribute.Images + "'"
				}
			} else {
				switch one {
				case "name":
					condition += " , name = '" + attribute.Name + "'"
				case "buying_price":
					condition += " , retail_price = " + strconv.FormatFloat(attribute.RetailPrice, 'f', 2, 64)
				case "unit_id":
					condition += " , unit_id = " + strconv.Itoa(attribute.UnitId)
				case "content":
					condition += " , content = '" + attribute.Content + "'"
				case "image":
					condition += " , image = '" + attribute.Image + "' , images = '" + attribute.Images + "'"
				}
			}
		}
		//修改加盟商商品
		err = goods.UpdateGoodAttribute(erpid, condition, db, param.Companyid)
		if err != nil {
			db.Rollback()
			return errors.New("修改数据失败。。。" + err.Error())
		}
	}

	db.Commit()
	return nil
}
