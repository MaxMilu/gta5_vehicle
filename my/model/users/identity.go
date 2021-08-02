package users

import (
	"github.com/qor/auth/auth_identity"
	"wyaoo/models"
)

type AuthIdentity struct {
	models.LogicBaseModel
	FranchiseeID uint `gorm:"not null;default:0"` // 加盟商ID
	auth_identity.Basic
	auth_identity.SignLogs
}

func (AuthIdentity) TableName() string {
	return "user_auth_identities"
}
