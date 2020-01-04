package address

import (
	"errors"
	"fmt"
	"new_erp_agent_by_go/controllers"
	"new_erp_agent_by_go/models/request"
	"new_erp_agent_by_go/service/open_api/address"
)

type InfoController struct {
	controllers.CommonController
}

// @Title 				  查询区域
// @Summary 			  查询区域
// @Description 		  查询区域
// @Param pid  raw  	  raw  	string		false  		`父id`
// @Param name        	  raw  	string  	false  		`名称`
// @Param Token           raw  	string  	true  		`token`
// @Param source          raw  	string  	true  		`source`
// @Success 201 {map}
// @Failure 403 body is empty
// @router /address [post]
func (this *InfoController) GetAddress() {
	checkResq := new(request.Address)
	this.GetJsonStruct(checkResq)

	//获取地区
	data, err := address.QueryAddressInfo(checkResq)
	if err != nil {
		this.CheckError(errors.New("查询商品分类信息失败。。。"), "")
	}

	fmt.Println("++", data)
	this.Success(data, "操作成功")

}
