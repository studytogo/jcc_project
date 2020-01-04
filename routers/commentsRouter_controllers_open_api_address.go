package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/address:InfoController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/open_api/address:InfoController"],
        beego.ControllerComments{
            Method: "GetAddress",
            Router: `/address`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
