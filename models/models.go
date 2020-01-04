package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库配置
func init() {
	var (
		sqlConn    = beego.AppConfig.String("sql.conn")
		maxIdle, _ = beego.AppConfig.Int("sql.max_idle_conns")
		maxConn, _ = beego.AppConfig.Int("sql.max_open_conns")
	)
	if err := orm.RegisterDataBase("default", "mysql", sqlConn, maxIdle, maxConn); err != nil {
		panic("error connecting to the database! " + err.Error())
	}
}
