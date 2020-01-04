package address

import (
	"new_erp_agent_by_go/models/address"
	"new_erp_agent_by_go/models/request"
)

//获取地区信息
func QueryAddressInfo(param *request.Address) ([]address.JccAddress, error) {
	return address.QueryAddressInfo(param)
}
