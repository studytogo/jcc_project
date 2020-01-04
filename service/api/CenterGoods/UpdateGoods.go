package CenterGoods

import (
	"new_erp_agent_by_go/models/centerGoods"
	"new_erp_agent_by_go/models/request"
)

func QueryUpdateGoods(param *request.CenterGoods) error {
	return centerGoods.QueryUpdateGoods(param)
}
