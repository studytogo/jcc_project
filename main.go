package main

import (
	"github.com/astaxie/beego/orm"
	_ "new_erp_agent_by_go/models"
	_ "new_erp_agent_by_go/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {

		is_dev, _ := beego.AppConfig.Bool("is_dev")
		if is_dev {
			orm.Debug = true
		}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
