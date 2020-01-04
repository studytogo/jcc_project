package task

import (
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/error_message"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

//启动监控器
func startMonitor() {
	c := cron.New()
	spec := "*/" + beego.AppConfig.String("monitor.output_time") + " * * * * ?"
	if err := c.AddFunc(spec, func() {
		outPutStatus()
	}); err != nil {
		helper.Log.Error(error_message.AddMonitorErr, err)
	}
	c.Start()
	helper.Log.Info(" monitor start ")
	helper.Log.Monitor()
	select {}
}

func outPutStatus() {
	helper.Log.Monitor()
}
