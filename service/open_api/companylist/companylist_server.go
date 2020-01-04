package companylist

import (
	"errors"
	"math"
	pycompanylist "new_erp_agent_by_go/models/companylist"
	"new_erp_agent_by_go/models/request"
	"time"
)

func AddCompanylist(params []pycompanylist.JccJicanchuCompanylist) error {
	for k, _ := range params {
		params[k].Id = 0
		if params[k].Name == "" {
			return errors.New("公司名称不能为空")
		}
		params[k].CreatedAt = time.Now().Unix()
		params[k].IsDel = 0
	}

	err := pycompanylist.AddCompanylist(params)
	return err
}

func QueryCompanylist(params *request.Page) (list []pycompanylist.JccJicanchuCompanylist, count int, per_last int, err error) {

	list, err = pycompanylist.QueryCompanylist(params.Page, params.Per_page)

	if err != nil {
		return nil, 0, 0, err
	}
	count, err = pycompanylist.QueryCompanylistCount()
	if err != nil {
		return nil, 0, 0, err
	}
	per_last = int(math.Ceil(float64(count) / float64(params.Per_page)))

	return list, count, per_last, err
}

func EditCompanylist(params pycompanylist.JccJicanchuCompanylist) error {
	if params.Name == "" {
		return errors.New("公司名称不能为空")
	}

	params.UpdatedAt = time.Now().Unix()

	count, err := pycompanylist.QueryCompanylistOnly(params.Name, params.Id)
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("公司名称重复")
	}

	err = pycompanylist.EditCompanylist(params)

	return err
}

func DeleteCompanylist(params request.JccCompanylistId) error {
	if params.Id == "" {
		return errors.New("id不能为空")
	}
	err := pycompanylist.DeleteCompanylist(params.Id)
	return err
}

func CheckName(params []pycompanylist.JccJicanchuCompanylist) error {
	for k, v := range params {

		for _, vv := range params[k+1:] {

			if v.Name == vv.Name {
				return errors.New("存在相同数据")
			}
		}
	}
	return nil
}

func CheckDatabase(params []pycompanylist.JccJicanchuCompanylist) error {
	var companyname string
	for _, v := range params {
		companyname = companyname + ",'" + v.Name + "'"
	}
	companyname = companyname[1:]
	count, err := pycompanylist.QueryDatabase(companyname)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("数据存在相同数据")
	}
	return nil
}
