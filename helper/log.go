package helper

import (
	"fmt"
	"new_erp_agent_by_go/helper/monitor"

	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type logger struct {
	*logs.BeeLogger
}

var Log = logger{
	logs.NewLogger(),
}

//日志初始化
func init() {

	//创建日志目录
	if _, err := os.Stat("logs"); err != nil {
		_ = os.Mkdir("logs", os.ModePerm)
	}

	var level = 7

	var (
		maxLine    = beego.AppConfig.String("log.max_line")
		maxsize, _ = beego.AppConfig.Int("log.max_size")
		maxDays    = beego.AppConfig.String("log.max_days")
	)
	conf := fmt.Sprintf(`{"filename":"logs/logs.log","level":%v,"separate": ["info", "error"],"maxlines":%v,"maxsize":%v,"daily":true,"maxdays":%v}`, level,
		maxLine, 1<<20*maxsize, maxDays)
	if err := Log.SetLogger(logs.AdapterMultiFile, conf); err != nil {
		panic("parsing log conf error! " + err.Error())
	}
	//是否异步输出日志
	Log.Async(1e3)

	Log.EnableFuncCallDepth(true)
}

// 重写错误log
func (l *logger) Error(errCode string, err error) {
	l.BeeLogger.Error("MEG: " + errCode + " INFO: " + err.Error())
}

// 数据库相关的错误
func (l *logger) ModelsError(tableName string, errCode string, err error) {
	l.BeeLogger.Error("TABLE: " + tableName + " MEG: " + errCode + " INFO: " + err.Error())
}

//controller层报错
func (l *logger) ControllerError(url, message string, err error) {
	l.BeeLogger.Error("URL: " + url + "  INFO: " + message + "   MES:" + err.Error())
}

// 输出系统信息
func (l *logger) Monitor() {
	l.BeeLogger.Info(fmt.Sprintf("%+v", monitor.StandardOutput()))
}

//只输出信息,没有状态码
func (l *logger) ErrorString(message string) {
	l.BeeLogger.Error(message)
}

//记录请求参数信息
func (l *logger) RecordParam(param string) {
	l.BeeLogger.Info(param)
}
