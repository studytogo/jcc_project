package brand

import (
	"new_erp_agent_by_go/models/brand"
)

//获取所有线上品牌信息
func QueryAllOnlineBrandInfo() ([]brand.JccBrand, error) {
	return brand.QueryOnlineBrandInfo()
}
