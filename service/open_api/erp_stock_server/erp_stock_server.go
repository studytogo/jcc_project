package erp_stock_server

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"math"
	"new_erp_agent_by_go/models/actual"
	"new_erp_agent_by_go/models/erp_goods"
	"new_erp_agent_by_go/models/erp_stock_record"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"sync"
	"time"
)

var erpStockLock sync.Mutex

func AddOrUpdateErpStock(param []actual.JccCompanyStock, source string) error {
	erpStockLock.Lock()
	defer erpStockLock.Unlock()
	db := orm.NewOrm()
	db.Begin()

	for k, _ := range param {
		companyId := strconv.Itoa(param[k].CompanyId)
		//判断商品id是否存在
		existGoodId := erp_goods.ExistGoodsId(param[k].GoodsId)
		if !existGoodId {
			db.Rollback()
			return errors.New(strconv.Itoa(param[k].GoodsId) + "商品不存在。。。")
		}
		//判断公司id是否存在
		exist, err := erp_goods.ExistCompanyId(param[k].CompanyId)
		if err != nil {
			db.Rollback()
			return errors.New(companyId + "公司不存在。。。")
		}

		if exist == 0 {
			db.Rollback()
			return errors.New(companyId + "公司不存在。。。")
		}
		//判断是否存在
		count, err := actual.ExistErpStock(param[k].GoodsId, param[k].CompanyId)

		if err != nil {
			db.Rollback()
			return errors.New(companyId + "数据不存在。。。")
		}

		//不存在就进行添加操作，存在就进行修改操作
		if count != 0 {
			//存在，修改
			param[k].UpdatedAt = time.Now().Unix()
			_, err := actual.UpdateErpStock(param[k], db)

			if err != nil {
				db.Rollback()
				return errors.New(companyId + "修改失败。。。")
			}
			//对库存增加记录
			stock, err := actual.QueryOneErpStock(param[k].GoodsId, param[k].CompanyId)

			//报错暂时不记录信息
			if err == nil {
				erpStockRecord := new(erp_stock_record.JccErpStockRecord)
				erpStockRecord.GoodsId = param[k].GoodsId
				erpStockRecord.Companyid = param[k].CompanyId
				erpStockRecord.CreatedAt = time.Now()
				erpStockRecord.ChangeStock = param[k].Option - stock
				erpStockRecord.OldStock = stock
				erpStockRecord.NewStock = param[k].Option
				erpStockRecord.Source = source
				ErpStockRecord(erpStockRecord)
			}

		} else {
			param[k].Id = 0
			//不存在，增加
			param[k].CreatedAt = time.Now().Unix()
			_, err := actual.AddErpStock(param[k], db)
			if err != nil {
				db.Rollback()
				return errors.New(companyId + "添加失败。。。")
			}
			//对库存增加记录
			erpStockRecord := new(erp_stock_record.JccErpStockRecord)
			erpStockRecord.GoodsId = param[k].GoodsId
			erpStockRecord.Companyid = param[k].CompanyId
			erpStockRecord.CreatedAt = time.Now()
			erpStockRecord.ChangeStock = param[k].Option
			erpStockRecord.OldStock = 0
			erpStockRecord.NewStock = param[k].Option
			erpStockRecord.Source = source
			ErpStockRecord(erpStockRecord)
		}
	}

	db.Commit()
	return nil
}

func QueryErpStock(params *request.Page) (list []actual.JccCompanyStock, count int, last_page int, err error) {
	list, err = actual.QueryErpStock(params.Page, params.Per_page)
	if err != nil {
		return nil, 0, 0, errors.New("查询失败")
	}
	count, err = actual.QuertErpStockCount()
	if err != nil {
		return nil, 0, 0, errors.New("查询失败")
	}
	last_page = int(math.Ceil(float64(count) / float64(params.Per_page)))
	return list, count, last_page, err
}

func ChangeErpStock(param []actual.JccCompanyStock, source string) error {

	//循环修改数据
	db := orm.NewOrm()
	db.Begin()
	for k, _ := range param {
		companyId := strconv.Itoa(param[k].CompanyId)
		//判断商品id是否存在
		existGoodId := erp_goods.ExistGoodsId(param[k].GoodsId)
		if !existGoodId {
			return errors.New(strconv.Itoa(param[k].GoodsId) + "商品不存在。。。")
		}
		//判断公司id是否存在
		exist, err := erp_goods.ExistCompanyId(param[k].CompanyId)
		if err != nil {
			return errors.New(companyId + "公司不存在。。。" + err.Error())
		}

		if exist == 0 {
			return errors.New(companyId + "公司不存在。。。")
		}
		//判断是否存在
		count, err := actual.ExistErpStock(param[k].GoodsId, param[k].CompanyId)

		if err != nil {
			return errors.New(companyId + "数据不存在。。。" + err.Error())
		}

		if count == 0 {
			_, err := actual.InsertOneErpStock(param[k], db)
			if err != nil {
				db.Rollback()
				return errors.New(companyId + "增加数据失败。。。")
			}
			//增加库存操作
			erpStockRecord := new(erp_stock_record.JccErpStockRecord)
			erpStockRecord.GoodsId = param[k].GoodsId
			erpStockRecord.Companyid = param[k].CompanyId
			erpStockRecord.CreatedAt = time.Now()
			erpStockRecord.ChangeStock = param[k].Option
			erpStockRecord.OldStock = 0
			erpStockRecord.NewStock = param[k].Option
			erpStockRecord.Source = source
			ErpStockRecord(erpStockRecord)
		} else {
			err := actual.ChangeErpStock(param[k], db)
			if err != nil {
				db.Rollback()
				return errors.New(strconv.Itoa(param[k].CompanyId) + "添加失败。。。" + err.Error())
			}

			//对库存做记录
			stock, err := actual.QueryOneErpStock(param[k].GoodsId, param[k].CompanyId)
			if err == nil {
				erpStockRecord := new(erp_stock_record.JccErpStockRecord)
				erpStockRecord.GoodsId = param[k].GoodsId
				erpStockRecord.Companyid = param[k].CompanyId
				erpStockRecord.CreatedAt = time.Now()
				erpStockRecord.ChangeStock = param[k].Option - stock
				erpStockRecord.OldStock = stock
				erpStockRecord.NewStock = param[k].Option
				erpStockRecord.Source = source
				ErpStockRecord(erpStockRecord)
			}
		}
	}
	db.Commit()

	return nil
}
