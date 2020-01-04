# JCC-AGENT

JCC-AGENT

### 使用bee工具生成orm结构体
- bee api AGENT -tables="" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/jcc_erp_agent"

### 使用github.com/fatih/gomodifytags 
##### 快速添加json标签
- gomodifytags -file ./models/goucenter/user.go -struct User -add-tags json 

##### 快速去除orm标签
- gomodifytags -file ./models/goucenter/user.go -struct User -remove-tags orm 

##### 综合使用，把orm标签转换成form内容
- gomodifytags -file controllers/shop/param_structs.go -struct addBrandParams -remove-tags orm,description -add-tags form

### 使用tools下ormStruct2Swagger.go将orm结构体内容生成对应swagger param 文档
- 修改代码 convertStruct:= xxx结构体
- 运行main函数
