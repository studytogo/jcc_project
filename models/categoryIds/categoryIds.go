package categoryIds

import (
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/request"
)

type JccGoodsCategory struct {
	Id        int    `orm:"column(id);auto" description:"ID"`
	Name      string `orm:"column(name);size(20)" description:"名称"`
	No        string `orm:"column(no);size(9)" description:"编码"`
	Level     int8   `orm:"column(level)" description:"等级"`
	Pid       int    `orm:"column(pid)" description:"父类ID"`
	CreatedAt int64  `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt int64  `orm:"column(updated_at)" description:"修改时间"`
	DeletedAt int64  `orm:"column(deleted_at)" description:"删除时间"`
	IsDel     int8   `orm:"column(is_del)" description:"是否删除"`
	//Kind      string             `orm:"column(kind);size(20);null" description:"商品的分类（本地(offline)、自营(online)）默认的是本地(offline)"`
	//Companyid int                `orm:"column(companyid)" description:"公司id"`
	//ErpId     int                `orm:"column(erp_id)" description:"erp的商品分类id"`
	Category []JccGoodsCategory `orm:"-"`
}

func init() {
	orm.RegisterModel(new(JccGoodsCategory))
}

func QueryAllCategoryInfoById(categoryId int) ([]JccGoodsCategory, error) {

	db := orm.NewOrm()

	var categoryInfo []JccGoodsCategory

	sql := `SELECT
			  jgc3.*
			FROM
			  (SELECT * FROM jcc_goods_category WHERE id = ?) jgc1
			INNER JOIN
			   jcc_goods_category jgc2 ON jgc1.pid = jgc2.id 
			INNER JOIN
			   jcc_goods_category jgc3 ON jgc2.pid = jgc3.id 
			UNION ALL
			SELECT
			  jgc2.*
			FROM
			  (SELECT * FROM jcc_goods_category WHERE id = ?) jgc1
			INNER JOIN
			  jcc_goods_category jgc2 ON jgc1.pid = jgc2.id 
			UNION ALL
			SELECT
				*
			FROM
				jcc_goods_category 
			WHERE
				id = ? `

	_, err := db.Raw(sql, categoryId, categoryId, categoryId).QueryRows(&categoryInfo)

	return categoryInfo, err
}

//查询所有商品分类信息
func QueryAllGoodsCategoryInfo(param *request.GoodsCategory) ([]JccGoodsCategory, error) {
	db := orm.NewOrm()

	// 有搜索条件 no
	noSql := ``
	if param.No != "" {
		noSql = `AND no = ` + "'" + param.No + "'"
	}
	// 有搜索条件 名字
	nameSql := ``
	if param.Name != "" {
		nameSql = `AND name = ` + "'" + param.Name + "'"
	}

	// 有搜索条件 父id
	pidSql := ``
	if param.Pid != "" {
		pidSql = `AND pid = ` + param.Pid
	}
	var goodsCategoryInfo []JccGoodsCategory

	sql := `SELECT * FROM jcc_goods_category WHERE is_del = 0 ` + noSql + nameSql + pidSql

	_, err := db.Raw(sql).QueryRows(&goodsCategoryInfo)

	return goodsCategoryInfo, err
}

//查询商品分类名字是否唯一
func QueryGoodsCategoryNameOnly(params *request.EditGoodsCategory) (sum int, err error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_goods_category where name=?`

	err = db.Raw(sql, params.CategoryName).QueryRow(&sum)

	return sum, err
}

//修改商品分类名称
func EditGoodsCategory(params *request.EditGoodsCategory) (err error) {
	db := orm.NewOrm()

	sql := `update jcc_goods_category set name = ? where id = ?`

	_, err = db.Raw(sql, params.CategoryName, params.CategoryId).Exec()

	return err
}

//查询商品种类level
func QueryGoodsCategoryLevel(params *request.DeleteGoodsCategory) (level int, err error) {
	db := orm.NewOrm()

	sql := `select level from jcc_goods_category where id = ?`

	err = db.Raw(sql, params.CategoryId).QueryRow(&level)

	return level, err
}

//查询商品种类是否存在子分类
func QueryGoodsCategoryChildren(params *request.DeleteGoodsCategory) (sum int, err error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_goods_category where pid = ? and is_del = 0`

	err = db.Raw(sql, params.CategoryId).QueryRow(&sum)

	return sum, err
}

//查询商品种类是否存在商品
func QueryGoodsCategoryindexChildren(params *request.DeleteGoodsCategory) (sum int, err error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_goods_category_index where goods_category_id = ?`

	err = db.Raw(sql, params.CategoryId).QueryRow(&sum)

	return sum, err
}

//删除商品种类
func DeleteGoodsCategory(params *request.DeleteGoodsCategory) (err error) {
	db := orm.NewOrm()

	sql := `delete from jcc_goods_category where id = ?`

	_, err = db.Raw(sql, params.CategoryId).Exec()

	return err
}

//查询未删除状态的分类名称是否存在
func QueryExistCategoryName(categoryName string) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_goods_category where name = ? and is_del = 0`

	var count int

	err := db.Raw(sql, categoryName).QueryRow(&count)

	return count, err
}

//查询父类的等级
func QueryParentCategoryInfo(pid int) (int, string, error) {
	db := orm.NewOrm()

	sql := `select level,no from jcc_goods_category where id = ?`

	var level int
	var no string

	err := db.Raw(sql, pid).QueryRow(&level, &no)

	return level, no, err
}

//查询相同父类的分类个数
func CountSameParentCategory(pid int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_goods_category where pid = ?`

	var level int

	err := db.Raw(sql, pid).QueryRow(&level)

	return level, err
}

//添加分类
func AddCategory(category JccGoodsCategory) (int64, error) {
	db := orm.NewOrm()
	return db.Insert(&category)

}

//判断是否存在三级分类
func ExistCategory(categoryId int) bool {
	db := orm.NewOrm()
	return db.QueryTable("jcc_goods_category").Filter("level", 3).Filter("is_del", 0).Filter("id", categoryId).Exist()
}

//通过id查询
func QueryCategoryById(categoryId int) (list JccGoodsCategory,err error) {
	db := orm.NewOrm()
	sql := `select * from jcc_goods_category where id = ?`
	db.Raw(sql,categoryId).QueryRow(&list)
	return list,err
}
