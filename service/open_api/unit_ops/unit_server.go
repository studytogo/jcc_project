package unit_ops

import "new_erp_agent_by_go/models/unit"

//获取所有单位信息
func QueryAllUnitInfo() ([]unit.JccUnit, error) {
	return unit.QueryAllUnitInfo()
}
