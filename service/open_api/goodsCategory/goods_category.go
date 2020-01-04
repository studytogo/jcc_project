package goodsCategory

import (
	"errors"
	"fmt"
	"new_erp_agent_by_go/helper"
	"new_erp_agent_by_go/helper/error_message"
	"new_erp_agent_by_go/helper/rabbitmq"
	"new_erp_agent_by_go/models/categoryIds"
	"new_erp_agent_by_go/models/request"
	"time"
)

//获取所有商品分类信息
func QueryAllGoodsCategoryInfo(param *request.GoodsCategory) ([]categoryIds.JccGoodsCategory, error) {
	return categoryIds.QueryAllGoodsCategoryInfo(param)
}

//添加商品分类
func  AddGoodCategory(param *request.AddGoodsCategory) (int64, error) {
	//查询商品名称是否有重复
	sameName, err := categoryIds.QueryExistCategoryName(param.CategoryName)

	if err != nil || sameName != 0 {
		helper.Log.Error(error_message.Add_CATEGORY_ERROR, error_message.ErrMessage("添加分类名称重复", err))
		return 0, error_message.ErrMessage("添加分类名称重复", err)
	}

	//分类等级验证

	if param.Pid == 0 && param.Level != 1 {
		helper.Log.Error(error_message.CATEGORY_LEVEL_ERROR, error_message.ErrMessage("分类等级不应该是1", nil))
		return 0, error_message.ErrMessage("分类等级不应该是1", nil)
	}

	if param.Level == 1 && param.Pid != 0 {
		helper.Log.Error(error_message.CATEGORY_LEVEL_ERROR, error_message.ErrMessage("分类等级不应该是1", nil))
		return 0, error_message.ErrMessage("分类等级不应该是1", nil)
	}

	if param.Level > 3 {
		helper.Log.Error(error_message.CATEGORY_LEVEL_ERROR, error_message.ErrMessage("分类等级参数错误", nil))
		return 0, error_message.ErrMessage("分类等级参数错误", nil)
	}

	//查询父分类的等级和编号
	level, no, err := categoryIds.QueryParentCategoryInfo(param.Pid)

	if err != nil && err.Error() != "<QuerySeter> no row found" {
		helper.Log.Error(error_message.CATEGORY_LEVEL_ERROR, error_message.ErrMessage("分类等级查询错误", err))
		return 0, error_message.ErrMessage("分类等级查询错误", err)
	}
	if param.Level != 1 {
		if param.Level != level+1 {
			helper.Log.Error(error_message.CATEGORY_LEVEL_ERROR, error_message.ErrMessage("分类等级参数错误", nil))
			return 0, error_message.ErrMessage("分类等级参数错误", nil)
		}
	}

	//编号验证
	//如果是1级分类是没有父分类编号的
	sameParent, err := categoryIds.CountSameParentCategory(param.Pid)
	if err != nil {
		helper.Log.Error(error_message.CATEGORY_NO_ERROR, error_message.ErrMessage("查询编号错误", err))
		return 0, error_message.ErrMessage("查询编号错误", err)
	}

	checkNo := ""

	if param.Level == 1 {
		checkNo = fmt.Sprintf("%03d", (sameParent + 1))
	} else {
		checkNo = no + fmt.Sprintf("%03d", (sameParent+1))
	}

	if param.No != checkNo {
		helper.Log.Error(error_message.CATEGORY_NO_ERROR, error_message.ErrMessage("分类编号错误", nil))
		return 0, error_message.ErrMessage("分类编号错误", nil)
	}

	//增加操作
	category := categoryIds.JccGoodsCategory{
		Name:      param.CategoryName,
		No:        param.No,
		Level:     int8(param.Level),
		Pid:       param.Pid,
		CreatedAt: time.Now().Unix(),
	}
	id, err := categoryIds.AddCategory(category)
	if err != nil {
		return 0, error_message.ErrMessage("添加商品分类失败", err)
	}
	//添加mq信息
	message := new(request.RabbitmqGoodsCategory)
	message.Id = id
	message.Name = param.CategoryName
	message.Level = int64(param.Level)
	message.Pid = int64(param.Pid)
	message.IsDel = 0
	message.IsUpdate = false
	go rabbitmq.SendRabbitMqMessage("agent.scOrder", "category", message)

	return id, nil
}

//修改分类名称
func EditGoodsCategory(params *request.EditGoodsCategory) error {

	//查询房老师所传种类修改名称是否存在相同
	count, err := categoryIds.QueryGoodsCategoryNameOnly(params)
	if err != nil {
		helper.Log.Error("", err)
		return err
	}
	if count != 0 {
		return errors.New("该商品种类名称已存在")
	}
	//修改数据库
	err = categoryIds.EditGoodsCategory(params)
	if err != nil {
		helper.Log.Error("", err)
		return err
	}

	//添加mq信息
	list, err := categoryIds.QueryCategoryById(params.CategoryId)
	if err != nil {
		return errors.New("查询分类失败")
	}
	message := new(request.RabbitmqGoodsCategory)
	message.Id = int64(list.Id)
	message.Name = list.Name
	message.Level = int64(list.Level)
	message.Pid = int64(list.Pid)
	message.IsDel = 0
	message.IsUpdate = true

	go rabbitmq.SendRabbitMqMessage("agent.scOrder", "category", message)

	return nil
}

//删除分类名称
func DeleteGoodsCategory(params *request.DeleteGoodsCategory) error {
	//检查商品分类level
	level, err := categoryIds.QueryGoodsCategoryLevel(params)
	if err != nil {
		helper.Log.Error("", err)
		return err
	}
	//如果不是三级分类，不能有子分类
	if level != 3 {
		count, err := categoryIds.QueryGoodsCategoryChildren(params)

		if err != nil {
			helper.Log.Error("", err)
			return err
		}

		if count != 0 {
			return errors.New("该商品种类存在子分类")
		}
		//如果是三级分类，不能有商品
	} else {
		count, err := categoryIds.QueryGoodsCategoryindexChildren(params)

		if err != nil {
			helper.Log.Error("", err)
			return err
		}

		if count != 0 {
			return errors.New("该商品种类存在商品")
		}
	}

	//修改数据库
	err = categoryIds.DeleteGoodsCategory(params)

	if err != nil {
		helper.Log.Error("", err)
		return err
	}

	//添加mq信息
	list, err := categoryIds.QueryCategoryById(params.CategoryId)
	if err != nil {
		return errors.New("查询分类失败")
	}
	message := new(request.RabbitmqGoodsCategory)
	message.Id = int64(list.Id)
	message.Name = list.Name
	message.Level = int64(list.Level)
	message.Pid = int64(list.Pid)
	message.IsDel = 1
	message.IsUpdate = true

	go rabbitmq.SendRabbitMqMessage("agent.scOrder", "category", message)

	return nil
}
