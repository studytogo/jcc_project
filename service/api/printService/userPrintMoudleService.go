package printService

import (
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/models/print"
	"time"
)

func UpdateOrAddUserPrintMoudle(param *print.JccUserMould) error {

	//判断是否存在记录
	exist, err := print.ExistUserMoudle(param.CompanyId)

	if err != nil {
		return error_message.ErrMessage("用户模板查询失败。。。", err)
	}

	if exist > 0 {
		//存在记录证明需要修改
		id, err := print.QueryUserMoudleId(param.CompanyId)
		if err != nil {
			return error_message.ErrMessage("用户模板查询失败。。。", err)
		}
		param.Id = id
		param.UpdatedAt = time.Now().Unix()
		db := orm.NewOrm()
		_, err = print.UpdateUserMoudle(param, db)
		if err != nil {
			return error_message.ErrMessage("修改用户模板失败。。。", err)
		}
	} else {
		//不存在记录需要增加记录
		param.CreatedAt = time.Now().Unix()
		_, err := print.AddUserPrintMoudle(param)
		if err != nil {
			return error_message.ErrMessage("添加用户模板失败。。。", err)
		}
	}

	return nil
}
