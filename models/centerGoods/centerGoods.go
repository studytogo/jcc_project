package centerGoods

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"strings"
	"time"
)

type JccCenterGoods struct {
	Id                    int64   `orm:"column(id);auto" description:"ID"`
	Name                  string  `orm:"column(name);size(150)" description:"名称"`
	Spu                   string  `orm:"column(spu);size(150)" description:"SPU"`
	Sku                   string  `orm:"column(sku);size(150)" description:"SKU"`
	Barcode               string  `orm:"column(barcode);size(150)" description:"条码"`
	Spec                  string  `orm:"column(spec);size(150)" description:"规格"`
	BuyingPrice           float64 `orm:"column(buying_price);digits(11);decimals(2)" description:"进货价"`
	RetailPrice           float64 `orm:"column(retail_price);digits(11);decimals(2)" description:"零售价"`
	InventoryUpperLimit   string  `orm:"column(inventory_upper_limit);size(255);null" description:"库存上限"`
	InventoryLowerLimit   string  `orm:"column(inventory_lower_limit)" description:"库存下限"`
	MnemonicWord          string  `orm:"column(mnemonic_word);size(255)" description:"助记词"`
	Remark                string  `orm:"column(remark);size(150)" description:"备注"`
	Image                 string  `orm:"column(image);size(150)" description:"主图"`
	Images                string  `orm:"column(images)" description:"多图"`
	Content               string  `orm:"column(content)" description:"内容"`
	ProducingProvinceId   int     `orm:"column(producing_province_id)" description:"产地省ID"`
	ProducingProvinceName string  `orm:"column(producing_province_name)" description:"产地省名称"`
	ProducingCityId       int     `orm:"column(producing_city_id)" description:"产地城市ID"`
	ProducingCityName     string  `orm:"column(producing_city_name)" description:"产地城市名称"`
	ProducingAreaId       int     `orm:"column(producing_area_id)" description:"产地区域ID"`
	ProducingAreaName     string  `orm:"column(producing_area_name)" description:"产地区域名称"`
	ProducingAreaDetail   string  `orm:"column(producing_area_detail);size(150)" description:"产地详情"`
	UnitId                int     `orm:"column(unit_id)" description:"单位ID"`
	UnitName              string  `orm:"column(unit_name)" description:"单位名称"`
	BrandId               int     `orm:"column(brand_id)" description:"品牌ID"`
	BrandName             string  `orm:"column(brand_name)" description:"品牌名称"`
	A8Code                string  `orm:"column(a8_code)" description:"A8编码"`
	CreatedAt             int64   `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt             int     `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt             int     `orm:"column(deleted_at)" description:"删除时间"`
	IsDel                 int8    `orm:"column(is_del)" description:"是否删除"`
}

//更新商品
func QueryUpdateGoods(param *request.CenterGoods) error {
	db := orm.NewOrm()

	var centerGoods []JccCenterGoods
	//var goods childGoods.JccGoods

	// 通过商品表传过来商品中心的商品id  去商品中心表找
	centerGoodsSql := `WHERE id in (` + param.CenterId + ")"
	sql := `SELECT * FROM jcc_center_goods ` + centerGoodsSql
	_, err := db.Raw(sql).QueryRows(&centerGoods)

	fmt.Println("中心表的商品信息", centerGoods)

	centerGoodsIds := strings.Split(param.CenterId, ",")

	// 循环的是接收到的要修改的商品的centerGoodsId
	for _, centerGoodsId := range centerGoodsIds {
		centerGoodsIdInt64, _ := strconv.ParseInt(centerGoodsId, 10, 64)

		// 循环的是 中心表里商品的信息
		for _, v := range centerGoods {

			// 如果中心表的id能对上商品表的center_goods_id 则修改
			if v.Id == centerGoodsIdInt64 {

				//有更新字段 名字
				nameSql := ``
				if param.Name != "" {
					nameSql = `, name = ` + "'" + v.Name + "'"
				}

				//有更新字段 零售价
				priceSql := ``
				if param.RetailPrice != "" {
					// 需要先把float64转成string
					s2 := strconv.FormatFloat(v.RetailPrice, 'E', -1, 64)
					priceSql = `, retail_price = ` + "'" + s2 + "'"
				}

				//有更新字段 商品详情
				contentSql := ``
				if param.Content != "" {
					contentSql = `, content = ` + "'" + v.Content + "'"
				}

				//有更新字段 主图
				imageSql := ``
				if param.Image != "" {
					imageSql = `, image = ` + "'" + v.Image + "'"
				}

				//有更新字段 单位
				unitSql := ``
				if param.UnitId != "" {
					unitIdString := strconv.Itoa(v.UnitId)
					unitSql = `, unit_id = ` + "'" + unitIdString + "'"
				}

				// 用centerGoodsId作为查询条件更新商品信息
				goodsUpdateSql := "UPDATE jcc_goods SET updated_at = ?" + nameSql + priceSql + contentSql + imageSql + unitSql + " WHERE center_goods_id = " + centerGoodsId
				res, err := db.Raw(goodsUpdateSql, time.Now().Unix()).Exec()
				if err == nil {
					num, _ := res.RowsAffected()
					fmt.Println("mysql row affected nums: ", num)
				} else {
					return errors.New("更新商品信息失败")
				}
			}
		}

	}

	return err
}

func QueryCenterGoodsById(id string) (list JccCenterGoods, err error) {
	o := orm.NewOrm()
	sql := `select * from jcc_center_goods where id = ?`
	err = o.Raw(sql, id).QueryRow(&list)
	return list, err
}

//查询需要同步的商品信息
func QuerySyncGoodAttribute(erpId interface{}) (JccCenterGoods, error) {
	db := orm.NewOrm()

	sql := `select * from jcc_center_goods where id = ?`

	var result JccCenterGoods

	err := db.Raw(sql, erpId).QueryRow(&result)

	return result, err
}
