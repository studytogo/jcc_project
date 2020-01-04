package user_server

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/models/user"
)

func AddUserRoleLink(param *request.UserRoleLink) error {
	//先删除用户的对应权限
	db := orm.NewOrm()
	db.Begin()

	err := user.DeleteUserRoleLink(db, param.UserRoleLink[0].Uid)
	if err != nil {
		db.Rollback()
		return errors.New("添加用户权限失败。。。")
	}
	//批量增加对应的权限
	_, err = user.AddUserRoleLink(db, param.UserRoleLink)
	if err != nil {
		db.Rollback()
		return errors.New("添加用户权限失败。。。")
	}

	db.Commit()
	return nil
}
