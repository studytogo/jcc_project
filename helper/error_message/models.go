package error_message

const (
	AddToModelError   = "130000" // 往数据库增加数据出错
	UpdateToModelErr  = "130001" // 往数据库更新数据出错
	DeleteToModelErr  = "130002" // 往数据库删除数据出错
	ReadToModelError  = "130003" // 往数据库查询数据出错
	OpModelError      = "130004" // 数据库操作出错
	DeleteToModelNone = "130005" // 数据库一条数据也没有被删除
	NotFindRecord     = "130006" // 查询不到对应的记录
	CheckAddModelErr  = "130007" // 先检查后增加
)
