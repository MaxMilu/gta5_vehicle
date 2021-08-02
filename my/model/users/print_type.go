package users

import "wyaoo/models"

type FranchiseePrintType struct {
	models.LogicBaseModel
	FranchiseeID uint   `gorm:"not null;default:0"` // 加盟ID
	Code         string `gorm:"not null"`           // 打印类型
	Name         string `gorm:"not null"`           // 打印类型名称
}

func (FranchiseePrintType) TableName() string {
	return "user_franchisee_print_types"
}
