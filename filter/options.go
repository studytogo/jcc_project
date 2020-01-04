package filter

import (
	"github.com/astaxie/beego/context"
	"new_erp_agent_by_go/controllers"
)

//var url string
var OptionsFilter = func(ctx *context.Context) {

	//url = "https://cs.jicanchu.net"
	//logs.Info(url)
	if ctx.Request.Method == "OPTIONS" {
		optionsAuth(ctx)
	}

}

func optionsAuth(ctx *context.Context) {
	output := new(controllers.Output)
	ctx.Output.JSON(output.OptionsOutput("成功"), true, false)
}
