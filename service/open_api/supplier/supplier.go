package supplier

import (
	"new_erp_agent_by_go/models/supplier"
)

//获取所有线上供应商信息
func QueryAllOnlineBrandInfo() ([]supplier.JccSupplier, error) {
	return supplier.QueryAllOnlineSupplierInfo()
}
