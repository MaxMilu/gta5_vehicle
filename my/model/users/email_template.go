package users

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/qor/l10n"
	"github.com/qor/qor/utils"
	"wyaoo/config/consts"
	"wyaoo/models"
	"wyaoo/utils/json_utils"
)

type EmailTemplate struct {
	models.LogicBaseModel
	FranchiseeID    uint `gorm:"not null;default:0" l10n:"sync"` // 加盟ID
	Franchisee      FranchiseeInformation
	UserID          uint `gorm:"not null;default:0" l10n:"sync"` // 用户ID
	User            User
	Type            string    `gorm:"size:20;not null" l10n:"sync"` // 邮件类型（参考字典表Type14）
	Subject         string    `gorm:"size:100;not null"`            // 邮件标题
	ContentTemplate string    `gorm:"type:text;not null"`           // 模板内容模板
	DefaultFrom     string    // 默认发件人
	DefaultTo       DefaultTo // 默认收件人
	DefaultCC       DefaultCC // 默认抄送
	l10n.Locale
}

func (EmailTemplate) TableName() string {
	return "user_email_templates"
}

type EmailTemplateEmail struct {
	Email string
}

type DefaultTo []EmailTemplateEmail

func (defaultTo *DefaultTo) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, defaultTo)
	case string:
		if v != "" {
			return defaultTo.Scan([]byte(v))
		}
	case []string:
		for _, str := range v {
			if err := defaultTo.Scan(str); err != nil {
				return err
			}
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (defaultTo DefaultTo) Value() (driver.Value, error) {
	results, err := json_utils.Encode(defaultTo)
	return utils.ReplaceQuoted(string(results)), err
}

type DefaultCC []EmailTemplateEmail

func (defaultCC *DefaultCC) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, defaultCC)
	case string:
		if v != "" {
			return defaultCC.Scan([]byte(v))
		}
	case []string:
		for _, str := range v {
			if err := defaultCC.Scan(str); err != nil {
				return err
			}
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (defaultCC DefaultCC) Value() (driver.Value, error) {
	results, err := json_utils.Encode(defaultCC)
	return utils.ReplaceQuoted(string(results)), err
}

func GetPartnerEmailTemplate(db *gorm.DB, partnerID uint, emailType string) (*EmailTemplate, error) {
	db = db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Set(consts.WYAOO_OMIT_FRANCHISEE_ID, true).Set(consts.WYAOO_OMIT_CREATE_USER, true)
	var emailTemplate EmailTemplate
	if emailTemplate.ID == 0 {
		if err := db.Preload("EmailServer").Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Preload("EmailServer").Where("partner_id=? and franchisee_id=? and user_id=? and type=?", 0, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &emailTemplate, nil
}

func GetFranchiseeEmailTemplate(db *gorm.DB, partnerID uint, franchiseeID uint, emailType string) (*EmailTemplate, error) {
	db = db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Set(consts.WYAOO_OMIT_FRANCHISEE_ID, true).Set(consts.WYAOO_OMIT_CREATE_USER, true)
	var emailTemplate EmailTemplate
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, franchiseeID, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", 0, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &emailTemplate, nil
}

func GetUserEmailTemplate(db *gorm.DB, partnerID uint, franchiseeID uint, userID uint, emailType string) (*EmailTemplate, error) {
	db = db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Set(consts.WYAOO_OMIT_FRANCHISEE_ID, true).Set(consts.WYAOO_OMIT_CREATE_USER, true)
	var emailTemplate EmailTemplate
	if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, franchiseeID, userID, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, franchiseeID, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", 0, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &emailTemplate, nil
}

func GetEmailTemplate(db *gorm.DB, partnerID uint, franchiseeID uint, userID uint, emailType string) (*EmailTemplate, error) {
	db = db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Set(consts.WYAOO_OMIT_FRANCHISEE_ID, true).Set(consts.WYAOO_OMIT_CREATE_USER, true).Set(consts.L10N_MODE, consts.ANY)
	var emailTemplate EmailTemplate
	if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, franchiseeID, userID, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, franchiseeID, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	if emailTemplate.ID == 0 {
		if err := db.Where("partner_id=? and franchisee_id=? and user_id=? and type=?", partnerID, 0, 0, emailType).FirstOrInit(&emailTemplate).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return &emailTemplate, nil
}
