package erp_goods_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/redis"
	"new_erp_agent_by_go/models/categoryIds"
	"new_erp_agent_by_go/models/erp_goods"
	"strconv"
	"strings"
	"sync"
	"time"
)

//防止并发操作进行限制
var erpGoodLock sync.Mutex

//批量增加商品信息
func AddErpGoods(param []*erp_goods.JccCenterGoods) (err error) {
	_, _, allBarcodes, allGoodsName, err := QueryGoodNameAndBarcode(param)
	defer func() {
		if err != nil {
			//如果报错将redis的数据还原
			err2 := redis.SetJson("names", allGoodsName)
			err2 = redis.SetJson("barcodes", allBarcodes)

			if err2 != nil {
				err = err2
			}
		}
	}()

	//判断商品名称和条码是否重复
	for _, goodsInfo := range param {
		for _, goodName := range allGoodsName {
			if goodsInfo.Name == goodName.Arrtibute {
				return errors.New(strconv.Itoa(goodsInfo.Id) + "商品名称存在。。。")
			}
		}

		for _, goodBarcode := range allBarcodes {
			if goodsInfo.Barcode == goodBarcode.Arrtibute {
				return errors.New(strconv.Itoa(goodsInfo.Id) + "商品条码存在。。。")
			}
		}
	}
	//erpGoodLock.Lock()
	//defer erpGoodLock.Unlock()
	//barcode, goodName, err := QueryGoodNameAndBarcode()
	//if err != nil {
	//	return err
	//}
	for k, _ := range param {
		param[k].Id = 0
		err = checkErpGoodsInfo(param[k])
		if err != nil {
			return err
		}
		param[k].CreatedAt = time.Now().Unix()
	}

	_, err = erp_goods.InsertMultiGoods(param)
	if err != nil {
		return errors.New("添加商品失败，请重新添加。。。")
	}
	//插入成功，将正确的信息增加到redis中
	var correctGoodsName []*erp_goods.SyncInfo
	correctGoodsName = append(correctGoodsName, allGoodsName...)
	var correctBarcodes []*erp_goods.SyncInfo
	correctBarcodes = append(correctBarcodes, allBarcodes...)
	for k, _ := range param {
		one := erp_goods.SyncInfo{}
		one.Id = param[k].Id
		one.Arrtibute = param[k].Name
		correctGoodsName = append(correctGoodsName, &one)
		one.Arrtibute = param[k].Barcode
		correctBarcodes = append(correctBarcodes, &one)
	}
	err = redis.SetJson("barcodes", correctBarcodes)
	if err != nil {
		return errors.New("更新reids商品条码失败")
	}

	err = redis.SetJson("names", correctGoodsName)
	if err != nil {
		return errors.New("更新reids商品名称失败")
	}

	return nil
}

//批量修改商品信息
func UpdateErpGood(param []*erp_goods.JccCenterGoods) (err error) {
	_, _, allBarcodes, allGoodsName, err := QueryGoodNameAndBarcode(param)
	defer func() {
		//修改后的缓存
		err2 := redis.SetJson("names", allGoodsName)
		err2 = redis.SetJson("barcodes", allBarcodes)

		if err2 != nil {
			err = err2
		}
	}()

	//判断商品名称和条码是否重复
	for _, goodsInfo := range param {
		for k, goodName := range allGoodsName {
			if goodsInfo.Name == goodName.Arrtibute && goodsInfo.Id != goodName.Id {
				return errors.New(strconv.Itoa(goodsInfo.Id) + "商品名称存在。。。")
			}

			if goodsInfo.Id == goodName.Id {
				allGoodsName[k].Arrtibute = goodsInfo.Name
			}
		}

		for k, goodBarcode := range allBarcodes {
			if goodsInfo.Barcode == goodBarcode.Arrtibute && goodsInfo.Id != goodBarcode.Id {
				return errors.New(strconv.Itoa(goodsInfo.Id) + "商品条码存在。。。")
			}

			if goodsInfo.Id == goodBarcode.Id {
				allGoodsName[k].Arrtibute = goodsInfo.Barcode
			}
		}
	}

	db := orm.NewOrm()
	db.Begin()
	for k, _ := range param {

		//err := CheckSameByUpdate(param[k])
		//
		//if err != nil {
		//	db.Rollback()
		//	return err
		//}

		err = checkErpGoodsInfo(param[k])

		if err != nil {
			db.Rollback()
			return err
		}

		param[k].UpdatedAt = time.Now().Unix()

		_, err = erp_goods.UpdateErpGood(param[k], db)
		if err != nil {
			db.Rollback()
			return errors.New(param[k].Name + "修改商品失败。。。")
		}
	}

	db.Commit()
	return nil
}

//批量删除商品
func DeleteErpGood(goodIds string) error {
	goodId := strings.Split(goodIds, ",")

	db := orm.NewOrm()
	for _, k := range goodId {
		id, _ := strconv.Atoi(k)
		if id == 0 {
			return errors.New("商品信息不能为0。。。")
		}
		existGoodId := erp_goods.ExistGoodsId(id)
		if !existGoodId {
			return errors.New(k + "商品信息不存在。。。")
		}
		_, err := erp_goods.DeleteErpGood(id, db)
		if err != nil {
			db.Rollback()
			return errors.New(k + "删除失败。。。")
		}
	}
	db.Commit()
	return nil

}

//批量查询商品信息
func QueryErpGoodInfo(ids []int) ([]erp_goods.JccCenterGoods, error) {
	if len(ids) == 0 {
		return nil, errors.New("商品id不能为空。。。")
	}

	return erp_goods.QueryErpGoodInfo(ids)
}

//检验公司端的商品参数
func checkErpGoodsInfo(goodInfo *erp_goods.JccCenterGoods) error {
	////商品名称不能空，也不能重复
	//if goodInfo.Name == "" {
	//	return errors.New(goodInfo.Name + "商品名称不能为空。。。")
	//}
	//
	//nameExist, err := erp_goods.ExistGoodName(goodInfo.Name, goodInfo.Id)
	//
	//if err != nil {
	//	return errors.New(goodInfo.Name + "商品名称错误。。。")
	//}
	//
	//if nameExist != 0 {
	//	return errors.New(goodInfo.Name + "商品名称重复。。。")
	//}
	//
	////条码不能重复
	////条码不能为空
	//if goodInfo.Barcode == "" {
	//	return errors.New("条码不能为空。。。")
	//}
	//
	//exist, err := erp_goods.ExistBarcode(goodInfo.Barcode, goodInfo.Id)
	//
	//if err != nil {
	//	return errors.New(goodInfo.Name + "条码错误。。。")
	//}
	//
	//if exist != 0 {
	//	return errors.New(goodInfo.Name + "条码重复。。。")
	//}

	//判断分类是否是3级分类
	//categorys, _ := redis.GetJson("Category")
	//if categorys == nil {
	//	//redis没有查询数据库，并做缓存
	//
	//}
	categoryExist := categoryIds.ExistCategory(goodInfo.CategoryId)

	if !categoryExist {
		return errors.New(goodInfo.Name + "分类参数不对。。。")
	}

	//判断规格是否正确
	right := checkSpec(goodInfo.Spec)

	if !right {
		return errors.New(goodInfo.Name + "规格格式错误。。。")
	}

	//判断image,images是否正确
	pass := checkImages(goodInfo.Image, goodInfo.Images)

	if !pass {
		return errors.New(goodInfo.Name + "图片参数有误，请重新请求。。。")
	}

	//检查是否为省
	if goodInfo.ProducingProvinceId != 0 {
		deep, name, err := erp_goods.CheckProvince(goodInfo.ProducingProvinceId)
		if err != nil {
			return errors.New(goodInfo.Name + "省不存在。。。")
		}

		if deep != 1 {
			return errors.New(goodInfo.Name + "省参数错误。。。")
		}

		goodInfo.ProducingProvinceName = name
	}

	//检查市是否正确
	if goodInfo.ProducingCityId != 0 {
		cityPid, cityName, err := erp_goods.CheckArea(goodInfo.ProducingCityId)

		if err != nil {
			return errors.New(goodInfo.Name + "城市错误。。。")
		}

		if cityPid != goodInfo.ProducingProvinceId {
			return errors.New(goodInfo.Name + "城市和省不匹配。。。")
		}

		goodInfo.ProducingCityName = cityName
	}

	//检查区域是否正确
	if goodInfo.ProducingAreaId != 0 {
		areaPid, areaName, err := erp_goods.CheckArea(goodInfo.ProducingAreaId)

		if err != nil {
			return errors.New(goodInfo.Name + "区域错误。。。")
		}

		if areaPid != goodInfo.ProducingCityId {
			return errors.New(goodInfo.Name + "区域和城市不匹配...")
		}

		goodInfo.ProducingAreaName = areaName
	}

	//判断单位是否存在
	if goodInfo.UnitId == 0 {
		return errors.New("goodInfo.Name" + "单位参数不能为0。。。")
	}
	unit, unitName, err := erp_goods.CheckUnitId(goodInfo.UnitId)

	if err != nil {
		return errors.New(goodInfo.Name + "单位参数错误。。。")
	}

	if unit == 0 {
		return errors.New(goodInfo.Name + "单位不存在....")
	}

	goodInfo.UnitName = unitName

	if goodInfo.BrandId == 0 {
		return errors.New("goodInfo.Name" + "品牌参数不能为0。。。")
	}
	goodInfo.BrandId = 1
	goodInfo.BrandName = "集餐厨"

	return nil
}

//检查规格是否符合要求
func checkSpec(spec string) bool {
	if spec == "" {
		return true
	}
	if strings.Index(spec, ";") == -1 {
		return false
	}
	specs := strings.Split(spec, ";")
	for _, v := range specs {
		if v != "" {
			if strings.Index(v, ":") == -1 {
				return false
			}
		}
	}

	return true
}

//检查图片是否符合要求  "images的第一张图应该和image保持一致"
func checkImages(image string, images string) bool {
	if (image == "" && images != "") || (image != "" && images == "") {
		return false
	}

	imageSlience := strings.Split(images, ",")
	if image != imageSlience[0] {
		return false
	}

	return true
}

//func QueryGoodNameAndBarcode() ([]string, []string, error) {
//	barcodes, err := erp_goods.QueryAllBarcode()
//	if err != nil {
//		return nil, nil, errors.New("查询商品条码失败。。。")
//	}
//
//	allName, err := erp_goods.QueryAllGoodName()
//	if err != nil {
//		return nil, nil, errors.New("查询商品名称失败。。。")
//	}
//
//	return barcodes, allName, nil
//}

func QueryGoodNameAndBarcode(goodInfo []*erp_goods.JccCenterGoods) ([]*erp_goods.SyncInfo, []*erp_goods.SyncInfo, []*erp_goods.SyncInfo, []*erp_goods.SyncInfo, error) {
	for {
		//分布式锁
		sync, err := redis.GetString("syncLock")
		if err != nil || sync == "0" || sync == "" {
			break
		}
	}
	redis.SetOperation("syncLock", "1", 2)
	//条码验证
	barcodesByte, _ := redis.GetJson("barcodes")
	fmt.Println("------------------", barcodesByte == nil)

	var allBarcodes []*erp_goods.SyncInfo
	if barcodesByte == nil {
		var err error
		allBarcodes, err = erp_goods.QueryAllBarcode()
		if err != nil {
			return nil, nil, nil, nil, errors.New("mysql查询商品条码失败。。。")
		}
		err = redis.SetJson("barcodes", allBarcodes)
		if err != nil {
			return nil, nil, nil, nil, errors.New("redis存储商品条码失败。。。")
		}
	} else {
		err := json.Unmarshal(barcodesByte, &allBarcodes)
		if err != nil {
			return nil, nil, nil, nil, errors.New("redis查询商品条码失败。。。")
		}
	}
	//商品名称验证
	goodsNameByte, _ := redis.GetJson("names")
	fmt.Println("++++++++++++++++++", goodsNameByte == nil)

	var allGoodsName []*erp_goods.SyncInfo
	if goodsNameByte == nil {
		var err error
		allGoodsName, err = erp_goods.QueryAllGoodName()
		if err != nil {
			return nil, nil, nil, nil, errors.New("mysql查询商品名称失败。。。")
		}
		err = redis.SetJson("barcodes", allGoodsName)
		if err != nil {
			return nil, nil, nil, nil, errors.New("redis存储商品名称失败。。。")
		}
	} else {
		err := json.Unmarshal(goodsNameByte, &allGoodsName)
		if err != nil {
			return nil, nil, nil, nil, errors.New("redis查询商品名称失败。。。")
		}
	}
	//先将数据存到redis中,此时的操做中数据还没有id,目的是防止并发时，其它请求插入相同数据
	//添加操作
	var newAllBarcodes []*erp_goods.SyncInfo
	var newAllGoodsName []*erp_goods.SyncInfo
	newAllBarcodes = append(newAllBarcodes, allBarcodes...)
	newAllGoodsName = append(newAllGoodsName, allGoodsName...)

	for _, v := range goodInfo {
		if v.Name == "" {
			return nil, nil, nil, nil, errors.New(strconv.Itoa(v.Id) + "商品名称不能为空")
		}
		one := &erp_goods.SyncInfo{}
		one.Id = v.Id
		one.Arrtibute = v.Name
		newAllGoodsName = append(newAllGoodsName, one)
		if v.Barcode == "" {
			return nil, nil, nil, nil, errors.New(strconv.Itoa(v.Id) + "商品条码不能为空")
		}
		one.Arrtibute = v.Barcode
		fmt.Println("--------------------", one)
		newAllBarcodes = append(newAllBarcodes, one)
	}
	err := redis.SetJson("barcodes", newAllBarcodes)
	if err != nil {
		return nil, nil, nil, nil, errors.New("更新reids商品条码失败")
	}

	err = redis.SetJson("names", newAllGoodsName)
	if err != nil {
		return nil, nil, nil, nil, errors.New("更新reids商品名称失败")
	}

	redis.SetOperation("syncLock", "0", 2)

	return newAllBarcodes, newAllGoodsName, allBarcodes, allGoodsName, nil

}

func CheckSameByUpdate(goodInfo *erp_goods.JccCenterGoods) error {
	//商品名称不能空，也不能重复
	if goodInfo.Name == "" {
		return errors.New(goodInfo.Name + "商品名称不能为空。。。")
	}

	nameExist, err := erp_goods.ExistGoodName(goodInfo.Name, goodInfo.Id)

	if err != nil {
		return errors.New(goodInfo.Name + "商品名称错误。。。")
	}

	if nameExist != 0 {
		return errors.New(goodInfo.Name + "商品名称重复。。。")
	}

	//条码不能重复
	//条码不能为空
	if goodInfo.Barcode == "" {
		return errors.New("条码不能为空。。。")
	}

	exist, err := erp_goods.ExistBarcode(goodInfo.Barcode, goodInfo.Id)

	if err != nil {
		return errors.New(goodInfo.Name + "条码错误。。。")
	}

	if exist != 0 {
		return errors.New(goodInfo.Name + "条码重复。。。")
	}

	return nil

}
