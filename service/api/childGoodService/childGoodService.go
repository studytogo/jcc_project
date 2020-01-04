package childGoodService

import (
	"errors"
	"new_erp_agent_by_go/helper/util"
	"new_erp_agent_by_go/models/actual"
	"new_erp_agent_by_go/models/categoryIds"
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/models/supplier"
	"new_erp_agent_by_go/models/unit"
	"new_erp_agent_by_go/service/api/unitService"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

func QueryChildInfoByParentId(parentId, companyRoomId int) (*[]childGoods.JccGoodsDetailInfo, error) {
	//通过父商品id查询子商品id
	db := orm.NewOrm()
	childIds, err := childGoods.QueryChildIdByParentId(parentId, db)
	if err != nil {
		return nil, err
	}

	//查询父的单位id
	unitId, err := unit.QueryUnitIdByGoodsId(parentId)
	if err != nil {
		return nil, errors.New("父商品的id查询不到")
	}

	//通过id查询子商品信息
	var childGoodsList []childGoods.JccGoodsDetailInfo
	for _, v := range childIds {
		if v == 0 {
			continue
		}
		childInfo, err := childGoods.QueryChildGoodInfoById(v, companyRoomId, db)
		if err != nil {
			return nil, err
		}
		//增加综合价格信息，数组第一个值是零售价，之后是预设价
		lingshou := map[string]string{
			"key":   "0",
			"name":  "零售价",
			"value": strconv.FormatFloat(childInfo.RetailPrice, 'f', -1, 64),
		}
		childInfo.Redetermined_price_list = append(childInfo.Redetermined_price_list, lingshou)
		//增加预设价格的返回值
		if childInfo.PredeterminedPrices != "" {
			prices := strings.Split(childInfo.PredeterminedPrices, ",")
			for i := 1; i <= len(prices); i++ {
				predPrice := map[string]string{
					"key":   strconv.Itoa(i),
					"name":  "预设价格" + strconv.Itoa(i),
					"value": prices[i-1],
				}
				childInfo.Redetermined_price_list = append(childInfo.Redetermined_price_list, predPrice)
			}
		}
		//增加供应商的返回值(可能后期会修改)
		if childInfo.Kind == "online" {
			childInfo.SupperlierInfo = nil
		} else {
			supplierName, supplierId, err := supplier.QueryCategoryInfoByGoodId(v)
			if err != nil {
				return nil, errors.New("商品供应商查询失败。。。")
			}
			childInfo.SupperlierInfo = map[string]interface{}{
				"Id":   supplierId,
				"Name": supplierName,
			}
		}

		//循环查找分类信息
		categoryInfos, err := categoryIds.QueryAllCategoryInfoById(childInfo.CategoryId)
		if err != nil {
			return nil, errors.New("分类信息不存在。。。。")
		}

		childInfo.CategoryIds = categoryInfos

		childInfo.ParentUnitId = unitId
		//如果单位是国际单位需要除以相应换算单位
		num := unitService.QueryConversionNumByUnitName(childInfo.Unit)
		//库存在计算之后要保留2位小数
		if num != 0 {
			ac := float64(childInfo.ActualInt) / float64(num)
			childInfo.Actual = util.Decimal(ac, 2)
			childInfo.Allactual = util.Decimal(ac, 2)
		} else {
			//前端需要字段不同所以都给赋值
			childInfo.Actual = strconv.Itoa(childInfo.ActualInt)
			childInfo.Allactual = strconv.Itoa(childInfo.ActualInt)
		}

		//如果仓库id传过来是空字符串，我要传这个商品所在公司的所有库存
		if companyRoomId == 0 {
			allActual, _ := actual.QueryAllActualByCompanyId(childInfo.Companyid, v)
			if num != 0 {
				childInfo.Allactual = util.Decimal(float64(allActual)/float64(num), 2)
			} else {
				childInfo.Allactual = strconv.Itoa(allActual)
			}

		}

		//如果无库存显示无记录
		if childInfo.Allactual == "" {
			childInfo.Allactual = "无记录"
		}

		childGoodsList = append(childGoodsList, childInfo)
	}

	if err != nil {
		return nil, err
	}

	return &childGoodsList, err
}

func SwitchParentOrChild(param *request.SwitchGoods) (*[]childGoods.JccGoodsDetailInfo, error) {

	db := orm.NewOrm()

	ids, err := childGoods.QueryChildIdOrParentId(param, db)

	if err != nil {
		return nil, err
	}

	var resultList []childGoods.JccGoodsDetailInfo

	for _, v := range ids {
		result, err := childGoods.QueryChildOrParentGoodInfoById(v, param, db)
		if err != nil {
			return nil, err
		}

		//查询该商品是否有过库存
		_, err = actual.GoodsHavedActualByGoodId(v)
		//默认该商品有过库存
		result.IsHaveActual = 1
		if err != nil {
			if err.Error() == "<QuerySeter> no row found" {
				result.IsHaveActual = 0
			} else {
				return nil, errors.New("判断商品是否有过库存失败。。。")
			}
		}
		//增加供应商的返回值(可能后期会修改)
		if result.Kind == "online" {
			result.SupperlierInfo = nil
		} else {
			supplierName, supplierId, err := supplier.QueryCategoryInfoByGoodId(v)
			if err != nil {
				return nil, errors.New("商品供应商查询失败。。。")
			}
			result.SupperlierInfo = map[string]interface{}{
				"Id":   supplierId,
				"Name": supplierName,
			}
		}

		//循环查找分类信息
		categoryInfos, err := categoryIds.QueryAllCategoryInfoById(result.CategoryId)
		if err != nil {
			return nil, errors.New("分类信息不存在。。。。" + err.Error())
		}

		result.CategoryIds = categoryInfos

		//如果单位是国际单位需要除以相应换算单位
		num := unitService.QueryConversionNumByUnitName(result.Unit)
		//库存在计算之后要保留2位小数
		if num != 0 {
			ac := float64(result.ActualInt) / float64(num)
			result.Actual = util.Decimal(ac, 2)
		} else {
			result.Actual = strconv.Itoa(result.ActualInt)
		}

		resultList = append(resultList, result)
	}

	return &resultList, err

}
