package unitService

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/models"
	"new_erp_agent_by_go/models/unit"
)

func CheckUnit(param *unit.CheckUnit) error {
	//根据id查询单位名称
	unitName, err := unit.QueryUnitNameByUnitId(param.UnitId)
	if err != nil {
		return error_message.ErrMessage("QueryUnitNameByUnitId方法错误：", err)
	}
	//检查是否是国际单位
	count, _ := unit.ExistInternationalUnitByByName(unitName)
	if count > 0 {
		//检查该商品是否有库存
		actual, _ := unit.GoodStockByGoodId(param.GoodId)

		if actual > 0 {
			return errors.New("该商品有库存，不能更改成国际单位。。。")
		} else {
			return nil
		}
	}

	return nil
}

func QueryConversionNumByUnitName(unitName string) int {

	num, _ := unit.QueryConversionNumByUnitName(unitName)

	return num
}

//通过传入goodId 查询该商品的库存和转换大小
func QueryStockByGoodId(goodId int64, companyRoomId int64) (int, int, bool, error) {
	db := orm.NewOrm()
	var stock int64
	var unitConversion int
	var isHavedStock bool
	goodsUnitId, err := models.SelectGoodsUnitId(db, goodId)
	if err != nil {
		return 0, 0, false, err
	}

	unitName, err := unit.QueryUnitNameByUnitId(goodsUnitId)
	if err != nil {
		return 0, 0, false, err
	}

	stock, isHavedStock, err = models.SelectParentGoodsStock(db, goodId, companyRoomId)
	if err != nil {
		return 0, 0, false, err
	}

	unitConversion = QueryConversionNumByUnitName(unitName)

	return int(stock), unitConversion, isHavedStock, nil
}
