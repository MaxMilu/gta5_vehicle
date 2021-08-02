package users

import "wyaoo/models"

type Group struct {
	models.LogicBaseModel
	FranchiseeID uint        `gorm:"not null;default:0"` // 加盟ID
	Name         string      `gorm:"not null"`           // 用户组名称
	OwnerID      uint        `gorm:"not null"`           // 主账号用户ID
	Type         string      `gorm:"not null"`           // 组类型 (0.系统生成 1.手动创建)
	Users        []GroupUser // 用户组成员
}

func (Group) TableName() string {
	return "user_groups"
}

type GroupUser struct {
	models.LogicBaseModel
	FranchiseeID uint `gorm:"not null;default:0"`                           // 加盟ID
	GroupID      uint `gorm:"not null;unique_index:uix_user_group_users_1"` // 用户组ID
	Group        Group
	UserID       uint `gorm:"not null;unique_index:uix_user_group_users_1"` // 用户ID
	User         User
	Owner        bool `gorm:"not null"`           // 是否主账号
	PayType      uint `gorm:"not null"`           // 支付方式（1：个人支付，2：主账号支付）
	Confirm      bool `gorm:"not null;default:0"` // 确认状态
}

func (GroupUser) TableName() string {
	return "user_group_users"
}
