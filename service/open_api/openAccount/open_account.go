package openAccount

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/rs/xid"
	"new_erp_agent_by_go/helper/rabbitmq"
	"new_erp_agent_by_go/models/jccBoss"
	"new_erp_agent_by_go/models/jccCompanylist"
	"new_erp_agent_by_go/models/request"
	"strconv"
	"time"
)

func OpenAccount(params *request.OpenAccount) (id int64,err error) {
	fmt.Println(params.Name)
	//检查参数
	count,err := jccBoss.CheckExistByBoss(params.Name)
	if err != nil {
		return 0,errors.New("查询用户名失败")
	}
	if count != 0 {
		return 0,errors.New("存在相同用户名")
	}
	count,err = jccBoss.CheckExistByIDCard(params.IdCard)
	if err != nil {
		return 0,errors.New("查询IDCard名失败")
	}
	if count != 0 {
		return 0,errors.New("存在相同IDCard")
	}
	count,err = jccCompanylist.CheckExistByName(params.Name+"公司")
	if err != nil {
		return 0,errors.New("查询公司名失败")
	}
	if count != 0 {
		return 0,errors.New("存在相同公司名")
	}
	//参数
	uuid := xid.New()
	companylist := new(jccCompanylist.JccCompanylist)
	companylist.Name = params.Name + "公司"
	companylist.Companycode = "C" + uuid.String()
	companylist.People = params.Name
	companylist.Tel = params.Tel
	companylist.Address = params.Address
	companylist.CreatedAt = time.Now().Unix()
	companylist.Companyid = params.Companyid
	companylist.Province = strconv.Itoa(params.Province)
	companylist.City = strconv.Itoa(params.City)
	companylist.District = strconv.Itoa(params.District)


	//添加company
	db := orm.NewOrm()
	db.Begin()
	company_id,err := jccCompanylist.AddCompany(db,companylist)
	if err != nil {
		db.Rollback()
		return 0,errors.New("添加公司失败")
	}
	//添加boss
	boss := new(jccBoss.JccBoss)
	boss.Boss = params.Name
	boss.Uid = params.Uid
	boss.IdCard = params.IdCard
	boss.Password = params.Password
	boss.Pricing = "1,2,3,4,5,6,7,8,9"
	boss.CreatedAt = time.Now().Unix()
	boss.AffCompany = company_id
	boss.Telephone = params.Tel
	boss.Address = params.Address
	boss.PrimaryAccount = 1

	boss_id,err := jccBoss.AddBoss(db,boss)
	if err != nil {
		db.Rollback()
		return 0,errors.New("添加主账号失败")
	}
	db.Commit()

	//添加mq信息
	message := new(request.RabbitmqBoss)
	message.Address = params.Address
	message.Boss = params.Name
	message.Companyid = int64(params.Companyid)
	message.IdCard = params.IdCard
	message.Uid = params.Uid
	message.Password = params.Password
	message.Pricing = "1,2,3,4,5,6,7,8,9"
	message.Telephone = params.Tel
	message.CreatedAt = time.Now().Unix()
	message.Province = int64(params.Province)
	message.City = int64(params.City)
	message.District = int64(params.District)

	go rabbitmq.SendRabbitMqMessage("agent.syncCompany", "sync", message)

	return boss_id,nil
}

func DelOpenAccount(id int) error {
	//根据bossid查询公司id
	company_id,err := jccBoss.QueryCompanyIdById(id)
	if err != nil {
		return errors.New("查询公司id失败")
	}

	boss := new(jccBoss.JccBoss)
	boss.Id = id
	boss.IsDel = 1
	boss.DeletedAt = time.Now().Unix()

	company := new(jccCompanylist.JccCompanylist)
	company.Id = company_id
	company.IsDel = 1
	company.DeletedAt = time.Now().Unix()

	db := orm.NewOrm()
	db.Begin()
	err = jccBoss.DelBoss(db,boss)
	if err != nil {
		db.Rollback()
		return errors.New("主账号删除失败")
	}

	err = jccCompanylist.DelCompany(db,company)
	err = jccBoss.DelBoss(db,boss)
	if err != nil {
		db.Rollback()
		return errors.New("公司删除失败")
	}

	db.Commit()

	//添加mq信息
	list, err := jccBoss.QueryBossAlllistIdById(id)
	if err != nil {
		return errors.New("查询分类失败")
	}
	//根据bossid查询公司id
	company_id,err = jccBoss.QueryCompanyIdById(id)
	if err != nil {
		return errors.New("查询公司id失败")
	}
	company_list,err := jccCompanylist.QueryCompanylistIdById(company_id)
	if err != nil {
		return errors.New("查询公司信息失败")
	}

	Province,_ := strconv.Atoi(company_list.Province)
	City,_ := strconv.Atoi(company_list.City)
	District,_ := strconv.Atoi(company_list.District)

	message := new(request.RabbitmqBoss)
	message.Address = list.Address
	message.Boss = list.Boss
	message.Companyid = list.AffCompany
	message.IdCard = list.IdCard
	message.Uid = list.Uid
	message.Password = list.Password
	message.Pricing = list.Pricing
	message.Telephone = list.Telephone
	message.CreatedAt = list.CreatedAt
	message.UpdatedAt = list.UpdatedAt
	message.DeletedAt = list.DeletedAt
	message.IsDel = int64(list.IsDel)
	message.Province = int64(Province)
	message.City = int64(City)
	message.District = int64(District)

	go rabbitmq.SendRabbitMqMessage("agent.syncCompany", "sync", message)

	return nil
}

func UpdateOpenAccount(params *request.UpdateOpenAccount) error {
	//查询boss
	boss_list,err := jccBoss.QueryCompanylistIdById(params.Id)
	if err != nil {
		return errors.New("查询用户信息失败")
	}
	//根据bossid查询公司id
	company_id,err := jccBoss.QueryCompanyIdById(params.Id)
	if err != nil {
		return errors.New("查询公司id失败")
	}
	company_list,err := jccCompanylist.QueryCompanylistIdById(company_id)
	if err != nil {
		return errors.New("查询公司信息失败")
	}
	//构造数据
	if params.Address != "" {
		boss_list.Address = params.Address
		company_list.Address = params.Address
	}
	if params.Password != "" {
		boss_list.Password = params.Password
	}
	if params.Tel != "" {
		boss_list.Telephone = params.Tel
		company_list.Address = params.Address
	}
	if params.Province != 0 {
		company_list.Province = strconv.Itoa(params.Province)
	}
	if params.City != 0 {
		company_list.City = strconv.Itoa(params.City)
	}
	if params.District != 0 {
		company_list.District = strconv.Itoa(params.District)
	}

	boss_list.UpdatedAt = time.Now().Unix()
	company_list.UpdatedAt = time.Now().Unix()

	//存数据
	db := orm.NewOrm()
	db.Begin()
	err = jccBoss.UpdateBoss(db,&boss_list)
	if err != nil {
		db.Rollback()
		return errors.New("修改用户信息失败")
	}
	err = jccCompanylist.UpdateCompany(db,&company_list)
	if err != nil {
		db.Rollback()
		return errors.New("修改公司信息失败")
	}
	db.Commit()

	//添加mq信息
	list, err := jccBoss.QueryBossAlllistIdById(params.Id)
	if err != nil {
		return errors.New("查询分类失败")
	}
	company_list,err = jccCompanylist.QueryCompanylistIdById(company_id)
	if err != nil {
		return errors.New("查询公司信息失败")
	}

	Province,_ := strconv.Atoi(company_list.Province)
	City,_ := strconv.Atoi(company_list.City)
	District,_ := strconv.Atoi(company_list.District)

	message := new(request.RabbitmqBoss)
	message.Address = list.Address
	message.Boss = list.Boss
	message.Companyid = list.AffCompany
	message.IdCard = list.IdCard
	message.Uid = list.Uid
	message.Password = list.Password
	message.Pricing = list.Pricing
	message.Telephone = list.Telephone
	message.CreatedAt = list.CreatedAt
	message.UpdatedAt = list.UpdatedAt
	message.DeletedAt = list.DeletedAt
	message.IsDel = int64(list.IsDel)
	message.Province = int64(Province)
	message.City = int64(City)
	message.District = int64(District)

	go rabbitmq.SendRabbitMqMessage("agent.syncCompany", "sync", message)

	return nil
}
