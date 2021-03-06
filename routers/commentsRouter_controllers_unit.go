package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/unit:UnitOperationController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/unit:UnitOperationController"],
        beego.ControllerComments{
            Method: "CheckUnit",
            Router: `/checkunit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
