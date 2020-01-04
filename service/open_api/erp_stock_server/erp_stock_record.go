package erp_stock_server

import "new_erp_agent_by_go/models/erp_stock_record"

func ErpStockRecord(record *erp_stock_record.JccErpStockRecord) (int64, error) {
	return erp_stock_record.AddErpStockRecord(record)
}
