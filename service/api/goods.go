package api

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/models"
	"new_erp_agent_by_go/models/childGoods"
	"new_erp_agent_by_go/models/erp_goods"
	"new_erp_agent_by_go/models/jccBoss"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/api/unitService"
	"strconv"
	"strings"
	"time"
)

// 通过child_goods_id去单位关系表里查  子商品合成父商品
func SelectConversionChild(info *models.JccGoodsUnitConversion) (newChildStock, newParentstock, isHavedStockType int64, err error) {
	// 通过child_goods_id去单位关系表里查
	o := orm.NewOrm()
	goodsUnitConversionInfo, err := models.ChildSelect(o, info.ChildGoodsId)
	if err != nil {
		return 0, 0, 0, err
	}

	//查询父商品的库存
	parentStock, parentUnitConversion, isHavedStock, err := unitService.QueryStockByGoodId(goodsUnitConversionInfo.GoodsId, info.LId)

	if err != nil {
		return 0, 0, 0, err
	}
	//查询子商品库存
	childStock, childUnitConversion, _, err := unitService.QueryStockByGoodId(goodsUnitConversionInfo.ChildGoodsId, info.LId)

	if err != nil {
		return 0, 0, 0, err
	}

	//根据国际单位修改库存
	if childUnitConversion != 0 {
		childStock = childStock - (int(info.Num) * childUnitConversion * int(goodsUnitConversionInfo.Num))
		newChildStock = int64(childStock / childUnitConversion)
	} else {
		childStock = childStock - (int(info.Num) * int(goodsUnitConversionInfo.Num))
		newChildStock = int64(childStock)
	}

	if childStock < 0 {
		return 0, 0, 0, errors.New("子商品库存不足。。。")
	}

	if parentUnitConversion != 0 {
		parentStock = parentStock + int(info.Num)*parentUnitConversion
		newParentstock = int64(parentStock / parentUnitConversion)
	} else {
		parentStock = parentStock + int(info.Num)
		newParentstock = int64(parentStock)
	}

	db := orm.NewOrm()
	db.Begin()
	//修改父商品的库存
	if isHavedStock {
		err = models.UpdateGoodsStock(db, goodsUnitConversionInfo.GoodsId, int64(parentStock), info.LId)
	} else {
		jccStock := new(models.JccStock)
		jccStock.GoodsId = goodsUnitConversionInfo.GoodsId
		jccStock.UpdatedAt = time.Now().Unix()
		jccStock.Actual = int64(parentStock)
		jccStock.LId = info.LId
		err = jccStock.AddStock(db)
	}

	if err != nil {
		db.Rollback()
		return 0, 0, 0, err
	}

	//修改子商品的库存
	err = models.UpdateGoodsStock(db, goodsUnitConversionInfo.ChildGoodsId, int64(childStock), info.LId)

	if err != nil {
		db.Rollback()
		return 0, 0, 0, err
	}

	db.Commit()

	if isHavedStock {
		isHavedStockType = 1
	} else {
		isHavedStockType = 0
	}

	return newChildStock, newParentstock, isHavedStockType, nil

}

// 父商品拆分成子商品
func SelectConversionParent(info *models.JccGoodsUnitConversion) (newChildStock int64, newParentstock, isHavedStockType int64, err error) {
	o := orm.NewOrm()
	// 通过goods_id去单位关系表里查
	goodsUnitConversionInfo, err := models.ParentSelect(o, info.GoodsId)

	if err != nil {
		return 0, 0, 0, err
	}

	//查询父商品的库存
	parentStock, parentUnitConversion, _, err := unitService.QueryStockByGoodId(goodsUnitConversionInfo.GoodsId, info.LId)

	if err != nil {
		return 0, 0, 0, err
	}
	//查询子商品库存
	childStock, childUnitConversion, isHavedStock, err := unitService.QueryStockByGoodId(goodsUnitConversionInfo.ChildGoodsId, info.LId)

	if err != nil {
		return 0, 0, 0, err
	}

	// 父商品拆分成子商品

	//根据国际单位修改库存
	if parentUnitConversion != 0 {
		parentStock = parentStock - int(info.Num)*parentUnitConversion
		newParentstock = int64(parentStock / parentUnitConversion)
	} else {
		parentStock = parentStock - int(info.Num)
		newParentstock = int64(parentStock)
	}

	if parentStock < 0 {
		return 0, 0, 0, errors.New("父商品库存不足。。。")
	}

	if childUnitConversion != 0 {
		childStock = childStock + (int(info.Num) * childUnitConversion * int(goodsUnitConversionInfo.Num))
		newChildStock = int64(childStock / childUnitConversion)
	} else {
		childStock = childStock + (int(info.Num) * int(goodsUnitConversionInfo.Num))
		newChildStock = int64(childStock)
	}

	db := orm.NewOrm()
	db.Begin()
	//修改父商品的库存
	err = models.UpdateGoodsStock(db, goodsUnitConversionInfo.GoodsId, int64(parentStock), info.LId)

	if err != nil {
		db.Rollback()
		return 0, 0, 0, err
	}

	//修改子商品的库存
	if isHavedStock {
		err = models.UpdateGoodsStock(db, goodsUnitConversionInfo.ChildGoodsId, int64(childStock), info.LId)
	} else {
		jccStock := new(models.JccStock)
		jccStock.GoodsId = goodsUnitConversionInfo.ChildGoodsId
		jccStock.LId = info.LId
		jccStock.Actual = int64(childStock)
		jccStock.UpdatedAt = time.Now().Unix()
		jccStock.AddStock(db)
	}

	if err != nil {
		db.Rollback()
		return 0, 0, 0, err
	}

	db.Commit()

	if isHavedStock {
		isHavedStockType = 1
	} else {
		isHavedStockType = 0
	}

	return newChildStock, newParentstock, isHavedStockType, nil
}

// 添加子商品
func AddChildGoods(goods *childGoods.JccGoods) (goodsBack *childGoods.JccGoods, ChildId int64, err error) {

	if goods.RetailPrice < 0 {
		return nil, 0, error_message.ErrMessage("零售价不能为负", nil)
	}

	//商品条码不能为空,不能重复
	if goods.Barcode == "" {
		return nil, 0, error_message.ErrMessage("条码不能为空", nil)
	}
	count, err := childGoods.ExistBarcode(goods.Barcode, goods.Companyid)

	if count > 0 || err != nil {
		return nil, 0, error_message.ErrMessage("条码重复", err)
	}
	if goods.RetailPrice == 0.00 {
		return nil, 0, error_message.ErrMessage("零售价不能为空", nil)
	}

	specs := []request.SpecInfo{}
	err = json.Unmarshal([]byte(goods.Spec), &specs)

	//不要前端传过来的值，重新赋值规格
	goods.Spec = ""
	for _, v := range specs {
		goods.Spec += v.Label + ":" + v.Value + ";"
	}

	db := orm.NewOrm()
	db.Begin()

	now := time.Now().Format("20060102150405")
	sku := "C" + now
	goods.Sku = sku

	//切片的第一个值是单位id
	strs := strings.Split(goods.GoodsUnitConversion, "|")
	if strs[0] == "" {
		return nil, 0, errors.New("单位转换参数有误。。。。")
	}

	//如果父有子商品则不让添加子商品
	if len(strs) < 3 {
		return nil, 0, errors.New("没有传父id。。。。")
	}

	parentId, _ := strconv.Atoi(strs[2])
	existHaveChild, err := childGoods.ExistHaveChild(parentId)

	if existHaveChild > 0 || err != nil {
		return nil, 0, error_message.ErrMessage("已经存在子商品。。。。", err)
	}

	unitName, err := models.SelectUnitName(db, goods.UnitId)
	if err != nil {
		db.Rollback()
		return nil, 0, err
	}

	goods.UnitId, _ = strconv.Atoi(strs[0])
	goods.CreatedAt = time.Now().Unix()

	goodsBack, ChildId, err = models.AddChildGoods(db, goods)
	if err != nil {
		helper.Log.Error("添加失败", err)
		db.Rollback()
		return goodsBack, ChildId, err
	}
	goodsBack.UnitName = unitName

	err = db.Commit()

	//增加分类关联表信息
	var goodsCategory = new(models.JccGoodsCategoryIndex)
	err = models.AddGoodsCategory(goodsCategory, db, goodsBack.Id, goods.GoodsCategoryId)
	if err != nil {
		db.Rollback()
		return goodsBack, ChildId, err
	}

	//增加供应商关联表信息
	var goodsSupplier = new(models.JccGoodsSupplierIndex)
	err = models.AddGoodsSupplier(goodsSupplier, db, goodsBack.Id, goods.SupplierId)
	if err != nil {
		db.Rollback()
		return goodsBack, ChildId, err
	}

	err = AddGoodsUnitConversion(goods, ChildId)
	if err != nil {
		db.Rollback()
		return goodsBack, ChildId, err
	}

	return goodsBack, ChildId, err
}

// 添加商品单位换算
func AddGoodsUnitConversion(goods *childGoods.JccGoods, ChildId int64) error {
	db := orm.NewOrm()
	db.Begin()
	s := strings.Split(goods.GoodsUnitConversion, "|")
	unitId, _ := strconv.ParseInt(s[0], 10, 64)
	Num, _ := strconv.ParseInt(s[1], 10, 64)
	ParentGoodsId, _ := strconv.ParseInt(s[2], 10, 64)
	err := models.AddGoodsUnitConversion(db, unitId, Num, ParentGoodsId, ChildId)
	if err != nil {
		db.Rollback()
		return err
	}

	err = models.UpdateParentGoods(db, ParentGoodsId)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return err

}

/*为前端运算返回*/
func SelectGoodsUnitConversion(info *models.JccGoodsUnitConversion) (goodsUnitConversionInfo *models.JccGoodsUnitConversion, err error) {
	// 通过child_goods_id去单位关系表里查
	db := orm.NewOrm()
	goodsUnitConversionInfo, err = models.ChildSelectConversion(db, info.ChildGoodsId)
	if err != nil {
		helper.Log.Error("", err)
	}
	return goodsUnitConversionInfo, err
}

/*为前端运算返回*/
func SelectParentGoodsUnitConversion(info *models.JccGoodsUnitConversion) (goodsUnitConversionInfo *models.JccGoodsUnitConversion, err error) {
	// 通过goods_id去单位关系表里查
	db := orm.NewOrm()
	goodsUnitConversionInfo, err = models.ParentSelectConversion(db, info.GoodsId)
	if err != nil {
		helper.Log.Error("", err)
	}
	return goodsUnitConversionInfo, err
}

//生成或更新线上商品信息接口
func AddOnlineGoodsInfo(goodsid string, companyid string) error {
	goods_id := strings.Split(goodsid, ",")

	boss, _ := strconv.Atoi(companyid)

	boss_companyid, err := jccBoss.QueryCompanyIdById(boss)

	if err != nil {
		return errors.New("查询失败。。。")
	}

	db := orm.NewOrm()
	db.Begin()

	for _, id := range goods_id {
		//查询
		count, err := erp_goods.QueryOnlineGoodsCount(id)
		if err != nil {
			db.Rollback()
			return errors.New(id + "商品查询失败。。。")
		}
		if count == 0 {
			db.Rollback()
			return errors.New(id + "商品不存在。。。")
		}

		list, err := erp_goods.QueryOnlineGoodsInfo(id)

		//验证jcc_goods表是否含有
		count, err = models.QuerySKUExit(list.Sku, strconv.Itoa(boss_companyid))
		if err != nil {
			db.Rollback()
			return errors.New(id + "商品SKU查询失败。。。")
		}
		if err != nil {
			db.Rollback()
			return errors.New("公司id错误。。。")
		}
		//新增

		if count == 0 {
			id64, err := models.AddOnlineGoods(db, list, boss_companyid, id)
			if err != nil {
				db.Rollback()
				return errors.New(id + "商品插入失败。。。")
			}
			err = models.AddOnlineGoodsCategory(db, id64, int64(list.CategoryId))
			if err != nil {
				db.Rollback()
				return errors.New("商品分类插入失败。。。")
			}

		}
		//修改
		if count == 1 {
			idss, err := models.QuerySKUId(list.Sku, boss_companyid)
			if err != nil {
				db.Rollback()
				return errors.New(id + "商品查询失败。。。")
			}
			err = models.UpdateOnlineGoods(db, list, idss)
			if err != nil {
				db.Rollback()
				return errors.New(id + "商品修改失败。。。")
			}
			//查询商品id对应的中间表id
			catrgory_id, err := models.QueryCategoryId(idss)
			//根据商品id修改单位
			//fmt.Println(catrgory_id)
			err = models.UpdateOnlineGoodsCategory(db, catrgory_id, list.CategoryId)
			if err != nil {
				db.Rollback()
				return errors.New(id + "商品分类修改失败。。。")
			}
		}
	}

	db.Commit()
	return nil
}

//单个更新线上商品信息接口
func AddOnlineOneGoodsInfo(goodsid string, companyid string) (idsss int64, err error) {

	boss, _ := strconv.Atoi(companyid)

	boss_companyid, err := jccBoss.QueryCompanyIdById(boss)

	if err != nil {
		return 0, errors.New("查询失败。。。")
	}

	db := orm.NewOrm()
	db.Begin()

	//查询
	count, err := erp_goods.QueryOnlineGoodsCount(goodsid)
	if err != nil {
		db.Rollback()
		return 0, errors.New(goodsid + "商品查询失败。。。")
	}
	if count == 0 {
		db.Rollback()
		return 0, errors.New(goodsid + "商品不存在。。。")
	}

	list, err := erp_goods.QueryOnlineGoodsInfo(goodsid)

	//验证jcc_goods表是否含有
	count, err = models.QuerySKUExit(list.Sku, strconv.Itoa(boss_companyid))
	if err != nil {
		db.Rollback()
		return 0, errors.New(goodsid + "商品SKU查询失败。。。")
	}
	if err != nil {
		db.Rollback()
		return 0, errors.New("公司id错误。。。")
	}
	//新增
	var id64 int64 = 0

	if count == 0 {
		id64, err = models.AddOnlineGoods(db, list, boss_companyid, goodsid)
		if err != nil {
			db.Rollback()
			return 0, errors.New(goodsid + "商品插入失败。。。")
		}
		err = models.AddOnlineGoodsCategory(db, id64, int64(list.CategoryId))
		if err != nil {
			db.Rollback()
			return 0, errors.New("商品分类插入失败。。。")
		}

	}
	//修改
	if count != 0 {
		id64, err = models.QuerySKUId(list.Sku, boss_companyid)
		if err != nil {
			db.Rollback()
			return 0, errors.New(goodsid + "商品查询失败。。。")
		}
		err = models.UpdateOnlineGoods(db, list, id64)
		if err != nil {
			db.Rollback()
			return 0, errors.New(goodsid + "商品修改失败。。。")
		}
		//查询商品id对应的中间表id
		catrgory_id, err := models.QueryCategoryId(id64)
		//根据商品id修改单位
		//fmt.Println(catrgory_id)
		err = models.UpdateOnlineGoodsCategory(db, catrgory_id, list.CategoryId)
		if err != nil {
			db.Rollback()
			return 0, errors.New(goodsid + "商品分类修改失败。。。")
		}
	}

	db.Commit()
	return id64, err
}
