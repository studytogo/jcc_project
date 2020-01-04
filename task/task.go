package task

import "github.com/astaxie/beego"

func init() {
	if openMonitor, _ := beego.AppConfig.Bool("monitor.start"); openMonitor {
		go startMonitor()
	}
}
