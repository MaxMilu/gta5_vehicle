package users

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wyaoo/config/consts"
	"wyaoo/models"
	"wyaoo/models/bases"
)

type FranchiseeTemplate struct {
	models.LogicBaseModel
	FranchiseeID        uint           `gorm:"not null;default:0"`
	SystemFunctionCode  string         `gorm:"not null"` // 模板类型
	TemplateID          uint           `gorm:"not null"` // 模板ID
	Template            bases.Template // 模板表
	TemplateDisplayName string         `gorm:"not null"` // 模板名称
	UserID              uint           // 用户ID(给在线打单默认模板用)
	WarehouseID         uint           // 仓库ID
	ServiceID           uint           // 服务ID
}

func (FranchiseeTemplate) TableName() string {
	return "user_franchisee_templates"
}

func GetFranchiseeTemplate(dbs *gorm.DB, functionName string, warehouseId uint, serviceId uint) ([]*FranchiseeTemplate, error) {
	var userTemplateList []*FranchiseeTemplate
	if err := dbs.Where("system_function_code=? AND warehouse_id=? AND service_id=?", functionName, warehouseId, serviceId).Find(&userTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var userServiceTemplateList []*FranchiseeTemplate
	if err := dbs.Where("system_function_code=? AND warehouse_id=? AND service_id=?", functionName, warehouseId, 0).Find(&userServiceTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var userWarehouseTemplateList []*FranchiseeTemplate
	if err := dbs.Where("system_function_code=? AND warehouse_id=? AND service_id=?", functionName, 0, 0).Find(&userWarehouseTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var serviceTemplateList []*FranchiseeTemplate
	if err := dbs.Set(consts.WYAOO_OMIT_CREATE_USER, true).Where("system_function_code=? AND user_id=? AND warehouse_id=? AND service_id=?", functionName, 0, warehouseId, serviceId).Find(&serviceTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var warehouseTemplateList []*FranchiseeTemplate
	if err := dbs.Set(consts.WYAOO_OMIT_CREATE_USER, true).Where("system_function_code=? AND user_id=? AND warehouse_id=? AND service_id=?", functionName, 0, warehouseId, 0).Find(&warehouseTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var franchiseeTemplateList []*FranchiseeTemplate
	if err := dbs.Set(consts.WYAOO_OMIT_CREATE_USER, true).Where("system_function_code=? AND user_id=? AND warehouse_id=? AND service_id=?", functionName, 0, 0, 0).Find(&franchiseeTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	userTemplateList = append(userTemplateList, userServiceTemplateList...)
	userTemplateList = append(userTemplateList, userWarehouseTemplateList...)
	userTemplateList = append(userTemplateList, serviceTemplateList...)
	userTemplateList = append(userTemplateList, warehouseTemplateList...)
	userTemplateList = append(userTemplateList, franchiseeTemplateList...)
	return userTemplateList, nil
}

func GetUserFranchiseeTemplate(dbs *gorm.DB, functionName string) ([]*FranchiseeTemplate, error) {
	var userTemplateList []*FranchiseeTemplate
	if err := dbs.Where("system_function_code=?", functionName).Find(&userTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var franchiseeTemplateList []*FranchiseeTemplate
	if err := dbs.Set(consts.WYAOO_OMIT_CREATE_USER, true).Where("system_function_code=? AND user_id=?", functionName, 0).Find(&franchiseeTemplateList).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	userTemplateList = append(userTemplateList, franchiseeTemplateList...)
	return userTemplateList, nil
}
