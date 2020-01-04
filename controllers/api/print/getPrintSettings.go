package print

import (
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/service/api/printService"
	"strconv"
)

type GetPrintSettingsController struct {
	controllers.CommonController
}

// @router /checklist [post]
func (this *GetPrintSettingsController) CheckList() {

	companyId := this.Ctx.Request.Form.Get("Companyid")
	company_id, err := strconv.Atoi(companyId)

	mould, err := printService.QueryList(company_id)

	if err != nil {
		helper.Log.ControllerError(this.Ctx.Input.URL(), "------------", err)
		this.CheckError(err, "")
	}

	this.Success(mould, "操作成功~！！！")

}
