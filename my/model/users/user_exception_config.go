package users

import "wyaoo/models"

type UserExceptionConfig struct {
	models.LogicBaseModel
	UserID      uint   // 用户ID
	ServiceCode string // 服务Code
	Type        string // 异常类型Dictionaries.type=8
	Days        uint   // 异常天数
}

func (UserExceptionConfig) TableName() string {
	return "user_exception_configs"
}
