package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:GetPrintSettingsController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:GetPrintSettingsController"],
        beego.ControllerComments{
            Method: "CheckList",
            Router: `/checklist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:PrintLogController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:PrintLogController"],
        beego.ControllerComments{
            Method: "PrintLog",
            Router: `/printlog`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:QueryListController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:QueryListController"],
        beego.ControllerComments{
            Method: "QueryListController",
            Router: `/querylist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:SelectPrintController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:SelectPrintController"],
        beego.ControllerComments{
            Method: "SelectMouldOneOrList",
            Router: `/select_mould_one_or_list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:UserPrintMoudleController"] = append(beego.GlobalControllerRouter["new_erp_agent_by_go/controllers/print:UserPrintMoudleController"],
        beego.ControllerComments{
            Method: "UpdateOrAddUserPrintMoudle",
            Router: `/update_user_print`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
