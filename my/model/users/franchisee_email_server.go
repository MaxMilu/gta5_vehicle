package users

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wyaoo/config/consts"
	"wyaoo/models"
)

type EmailServer struct {
	models.LogicBaseModel
	FranchiseeID uint   `gorm:"not null;default:0"` // 加盟ID
	Host         string `gorm:"size:50;not null"`   // 邮件服务器主机
	Port         uint   `gorm:"size:50;not null"`   // 邮件服务器端口
	Username     string `gorm:"size:50;not null"`   // 邮件服务器登录用户
	Password     string `gorm:"size:50;not null"`   // 邮件服务器登录密码
}

func (EmailServer) TableName() string {
	return "user_franchisee_email_servers"
}

func (server EmailServer) Label() string {
	return fmt.Sprintf("%v@%v:%v", server.Username, server.Host, server.Port)
}

func GetEmailServer(db *gorm.DB, partnerID uint, franchiseeID uint) (*EmailServer, error) {
	db = db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Set(consts.WYAOO_OMIT_FRANCHISEE_ID, true)
	var emailServer EmailServer
	if err := db.Where("partner_id=? and franchisee_id=?", partnerID, franchiseeID).FirstOrInit(&emailServer).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	if emailServer.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=?", partnerID, partnerID).FirstOrInit(&emailServer).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailServer.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=?", 0, 0).FirstOrInit(&emailServer).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailServer.ID == 0 {
		return nil, errors.New("email server doesn't exist")
	}
	return &emailServer, nil
}
