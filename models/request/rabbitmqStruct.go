package request

type RabbitmqGoodsCategory struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	Pid      int64  `json:"pid"`
	IsDel    int64  `json:"is_del"`
	IsUpdate bool   `json:"is_update"`
}

type RabbitmqBoss struct {
	Address     string `json:"address"`
	Boss        string `json:"boss"`
	Companyid   int64  `json:"companyid"`
	IdCard      string `json:"id_card"`
	Uid         int    `json:"uid"`
	IsDel       int64  `json:"is_del"`
	Password    string `json:"password"`
	Pricing     string `json:"pricing"`
	Telephone   string `json:"telephone"`
	CreatedAt   int64  `json:"created_at"`
	DeletedAt   int64  `json:"deleted_at"`
	UpdatedAt   int64  `json:"updated_at"`
	ErpId       int64  `json:"id"`
	Province    int64  `json:"province"`
	City        int64  `json:"city"`
	District    int64  `json:"district"`
	SaleAddress string `json:"sale_address"`
}
