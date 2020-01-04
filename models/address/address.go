package address

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"new_erp_agent_by_go/models/request"
)

type JccAddress struct {
	Id   int
	Name string
	Pid  int
	Sort int
	Deep int
}

//查询地区信息
func QueryAddressInfo(param *request.Address) ([]JccAddress, error) {
	db := orm.NewOrm()

	// 有搜索条件 名字
	noSql := ``
	if param.Name != "" {
		noSql = `WHERE name = ` + "'" + param.Name + "'"
	}
	fmt.Println("============noSql", noSql)
	// 有搜索条件 pid
	nameSql := ``
	if param.Pid != "" {
		nameSql = `WHERE pid = ` + "'" + param.Pid + "'"
	}
	fmt.Println("============nameSql", nameSql)

	var addressInfo []JccAddress

	sql := `SELECT * FROM jcc_address ` + noSql + nameSql

	_, err := db.Raw(sql).QueryRows(&addressInfo)
	fmt.Println("============addressInfo", addressInfo)

	return addressInfo, err
}
