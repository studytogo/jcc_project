package print

import "github.com/astaxie/beego/orm"

type JccUserMould struct {
	Id             int   `orm:"column(id);auto" description:"ID"`
	MouldId        int   `orm:"column(mould_id)" description:"模板id" json:"mould_id"`
	TopMargin      int   `orm:"column(top_margin)" description:"上边距" json:"top_margin"`
	CreatedAt      int64 `orm:"column(created_at)" description:"创建时间" json:"created_at"`
	LeftMargin     int   `orm:"column(left_margin)" description:"左边距" json:"left_margin"`
	IsDisplayPrice int   `orm:"column(is_display_price)" description:"是否显示零售价" json:"is_display_price"`
	CompanyId      int   `orm:"column(company_id)" description:"公司id" json:"num" json:"company_id"`
	UpdatedAt      int64 `orm:"column(updated_at)" description:"修改时间" json:"updated_at"`
	IsDefault      int   `orm:"column(is_default)" description:"修改时间" json:"is_default"`
}

type CheckUserMould struct {
	MouldId        int    `orm:"column(mould_id)" description:"模板id" json:"mould_id"`
	FieldOrder     string `orm:"column(field_order)" description:"字段顺序" json:"field_order"`
	Mould_weight   int    `orm:"column(mould_weight)" description:"模板宽度" json:"mould_weight"`
	Mould_height   int    `orm:"column(mould_height)" description:"模板宽度" json:"mould_height"`
	PrintList      int    `orm:"column(print_list)" description:"打印调试" json:"print_list"`
	TopMargin      int    `orm:"column(top_margin)" description:"上边距" json:"top_margin"`
	LeftMargin     int    `orm:"column(left_margin)" description:"左边距" json:"left_margin"`
	IsDisplayPrice int    `orm:"column(is_display_price)" description:"是否显示零售价" json:"is_display_price"`
}

func init() {
	orm.RegisterModel(new(JccUserMould))
}

//插入用户打印模板
func AddUserPrintMoudle(userMould *JccUserMould) (int64, error) {
	db := orm.NewOrm()
	return db.Insert(userMould)
}

//是否存在模板
func ExistUserMoudle(companyId int) (int, error) {
	db := orm.NewOrm()

	sql := `select count(*) from jcc_user_mould where company_id = ?`

	var exist int

	err := db.Raw(sql, companyId).QueryRow(&exist)

	return exist, err
}

func UpdateUserMoudle(userMould *JccUserMould, db orm.Ormer) (int64, error) {
	return db.Update(userMould, "mould_id", "top_margin", "left_margin", "is_display_price", "updated_at")
}

func QueryUserMoudleId(companyId int) (int, error) {
	db := orm.NewOrm()

	sql := `select id from jcc_user_mould where company_id = ?`

	var moudleId int

	err := db.Raw(sql, companyId).QueryRow(&moudleId)

	return moudleId, err
}

func QueryUserMould(companyId int, db orm.Ormer) (*[]CheckUserMould, error) {

	var UserMould = new([]CheckUserMould)

	sql := `SELECT mould_id,field_order,print_list,top_margin,left_margin,is_display_price,mould_weight,mould_height FROM jcc_user_mould u INNER JOIN jcc_common_mould c ON u.mould_id = c.id WHERE u.company_id = ? `

	_, err := db.Raw(sql, companyId).QueryRows(UserMould)

	return UserMould, err
}
