package user

import "github.com/astaxie/beego/orm"

type JccUserRoleLink struct {
	Id        int  `orm:"column(id);auto" json:"id"`
	BossId    int  `orm:"column(boss_id)" description:"操作人" json:"boss_id"`
	Uid       int  `orm:"column(uid)" description:"使用人id" json:"uid"`
	RoleId    int  `orm:"column(role_id)" description:"权限id" json:"role_id"`
	GroupId   int  `orm:"column(group_id);null" description:"员组表id" json:"group_id"`
	IsDel     int8 `orm:"column(is_del);null" description:"是否删除" json:"is_del"`
	IsStart   int8 `orm:"column(is_start);null" description:"是否启用" json:"is_start"`
	CreatedAt int  `orm:"column(created_at);null" description:"创建时间" json:"created_at"`
	UpdatedAt int  `orm:"column(updated_at);null" description:"修改时间" json:"updated_at"`
	DeletedAt int  `orm:"column(deleted_at);null" description:"删除时间" json:"deleted_at"`
}

func init() {
	orm.RegisterModel(new(JccUserRoleLink))
}

//批量删除用户权限表信息
func DeleteUserRoleLink(db orm.Ormer, uid int) error {
	sql := `DELETE FROM jcc_user_role_link WHERE uid = ? `

	_, err := db.Raw(sql, uid).Exec()

	return err
}

//批量增加用户权限表信息
func AddUserRoleLink(db orm.Ormer, userRoleLink []JccUserRoleLink) (int64, error) {
	return db.InsertMulti(1, &userRoleLink)
}
