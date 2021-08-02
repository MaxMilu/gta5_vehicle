package users

import (
	"database/sql/driver"
	"github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/qor/media/oss"
	qorUtils "github.com/qor/qor/utils"
	"github.com/qor/validations"
	"github.com/shopspring/decimal"
	"time"
	"wyaoo/config"
	"wyaoo/models"
	"wyaoo/utils/json_utils"
	"wyaoo/utils/log_utils"
)

type FranchiseeInformation struct {
	models.LogicBaseModel
	Applier                     uint    `sql:NOT NULL`
	Type                        uint8   `sql:"NOT NULL"` //1.合作商 2.供货商 3.代销商 4.转运商
	Name                        string  `sql:"NOT NULL" gorm:"unique_index"`
	ContactPerson               string  `sql:"NOT NULL"`
	ContactPhone                string  `sql:"NOT NULL"`
	IdentificationType          uint8   `sql:"NOT NULL"`
	IdCardNo                    string  `sql:"NOT NULL"`
	PassportNo                  string  `sql:"NOT NULL"`
	PassportFile                oss.OSS `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}/{{filename_with_hash}}"`
	Status                      string  `sql:"size:1;NOT NULL;DEFAULT:0"`
	CustomsCode                 string  `sql:"NOT NULL"` // 企业海关备案号
	CustomsName                 string  `sql:"NOT NULL"` // 企业海关备案名
	PartnerMarkupPercent        decimal.Decimal
	SellerAgentMarkupPercentMin decimal.Decimal
	SellerAgentMarkupPercentMax decimal.Decimal
	PartnerAgentAllow           bool
	PartnerTransportAllow       bool
	TopCategoriesAllow          TopCategoriesAllows `sql:"type:text"`
	HomeLogo                    oss.OSS             // 后台Logo
	AdminLogo                   oss.OSS             // 痈台Logo
	CommissionMinimum           decimal.Decimal     `sql:"not null;default:100"`
	CommissionTerm              uint                `sql:"not null;default:7"`
	DefaultFranchisee           uint                // 默认转运商
}

func (FranchiseeInformation) TableName() string {
	return "user_franchisees"
}

type TopCategoriesAllows []TopCategoriesAllow
type TopCategoriesAllow struct {
	CategoryID uint
}

func (topCategoriesAllow *TopCategoriesAllows) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, topCategoriesAllow)
	case string:
		if v != "" {
			return topCategoriesAllow.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (topCategoriesAllow TopCategoriesAllows) Value() (driver.Value, error) {
	if len(topCategoriesAllow) == 0 {
		return nil, nil
	}
	if data, err := json_utils.Encode(topCategoriesAllow); err == nil {
		return qorUtils.ReplaceQuoted(string(data)), err
	} else {
		return nil, err
	}
}

//func (franchisee *FranchiseeInformation) Validate(db *gorm.DB) {
//	if franchisee.Name == "" {
//
//		db.AddError(validations.NewError(franchisee, "Name", "Please input name"))
//	} else if !checks.CheckSpecialChars(franchisee.Name) {
//		db.AddError(validations.NewError(franchisee, "Name", "Name can not contains special characters"))
//	}
//	if franchisee.ContactPerson == "" {
//		db.AddError(validations.NewError(franchisee, "ContactPerson", "Please input contact person"))
//	} else if !checks.CheckSpecialChars(franchisee.ContactPerson) {
//		db.AddError(validations.NewError(franchisee, "ContactPerson", "Contact person can not contains special characters"))
//	}
//	if franchisee.ContactPhone == "" {
//		db.AddError(validations.NewError(franchisee, "ContactPhone", "Please input contact phone"))
//	} else if !checks.CheckSpecialChars(franchisee.ContactPhone) {
//		db.AddError(validations.NewError(franchisee, "ContactPhone", "Contact phone can not contains special characters"))
//	}
//	//如果是销售商
//	if franchisee.Type != 3 {
//		if franchisee.IdentificationType == 0 {
//			db.AddError(validations.NewError(franchisee, "IdentificationType", "Please select identification type"))
//		} else if franchisee.IdentificationType == 1 {
//			if franchisee.IdCardNo == "" {
//				db.AddError(validations.NewError(franchisee, "IdCardNo", "Please input id card number"))
//			} else if !checks.CheckIdCardNo(franchisee.IdCardNo) {
//				db.AddError(validations.NewError(franchisee, "IdCardNo", "Please input correct id card number"))
//			}
//		} else if franchisee.IdentificationType == 2 {
//			if franchisee.PassportNo == "" {
//				db.AddError(validations.NewError(franchisee, "PassportNo", "Please input passport number"))
//			}
//			if franchisee.PassportFile.FileName == "" {
//				local, _ := db.Get(consts.L10N_LOCALE)
//				errorInfo := i18n.I18n.T(local.(string), "user_franchisee_franchiseeCheck_passport_file_select", "Please select passport file")
//				db.AddError(validations.NewError(franchisee, "PassportFile", string(errorInfo)))
//			}
//		}
//	}
//	if len(db.GetErrors()) == 0 {
//		var count uint
//		db.Model(&User{}).Where("franchisee_id=?", franchisee.Applier).Count(&count)
//		if count > 0 {
//			var temp FranchiseeInformation
//			db.Model(&FranchiseeInformation{}).Where("applier=? and id!=?", franchisee.Applier, franchisee.ID).Find(&temp)
//			db.AddError(validations.NewError(franchisee, "Applier", "Duplicate apply"))
//
//			return
//		}
//		db.Model(&FranchiseeInformation{}).Where("applier=? and id!=?", franchisee.Applier, franchisee.ID).Count(&count)
//		if count > 0 {
//			db.AddError(validations.NewError(franchisee, "Applier", "Duplicate apply"))
//			return
//		}
//		db.Model(&FranchiseeInformation{}).Where("name=? and id!=?", franchisee.Name, franchisee.ID).Count(&count)
//		if count > 0 {
//			db.AddError(validations.NewError(franchisee, "Name", "Name already exists"))
//			return
//		}
//	}
//}

func GetFranchiseeInformation(db *gorm.DB, id uint) (franchiseeInformation FranchiseeInformation) {
	db.Model(&FranchiseeInformation{}).Where("id=?", id).Find(&franchiseeInformation)
	return
}

type FranchiseeAddress struct {
	models.LogicGeneralBaseModel
	FranchiseeID uint   `sql:"NOT NULL;default:0"`
	CountryCode  string `sql:"size:2;NOT NULL"`
	Provinces    string `sql:"NOT NULL"`
	City         string `sql:"NOT NULL"`
	County       string
	ZipCode      string `sql:"NOT NULL"`
	Address      string `sql:"NOT NULL"`
	Phone        string `sql:"NOT NULL"`
}

func (FranchiseeAddress) TableName() string {
	return "user_franchisee_addresses"
}

func (franchiseeAddress *FranchiseeAddress) Validate(db *gorm.DB) {
	if franchiseeAddress.CountryCode == "" {
		db.AddError(validations.NewError(franchiseeAddress, "CountryCode", "Please select country code"))
	}
	if franchiseeAddress.Provinces == "" {
		db.AddError(validations.NewError(franchiseeAddress, "Provinces", "Please input provinces"))
	}
	if franchiseeAddress.City == "" {
		db.AddError(validations.NewError(franchiseeAddress, "City", "Please input city"))
	}
	if franchiseeAddress.County == "" {
		db.AddError(validations.NewError(franchiseeAddress, "County", "Please input county"))
	}
	if franchiseeAddress.ZipCode == "" {
		db.AddError(validations.NewError(franchiseeAddress, "ZipCode", "Please input zip code"))
	}
	if franchiseeAddress.Address == "" {
		db.AddError(validations.NewError(franchiseeAddress, "Address", "Please input address"))
	}
	if franchiseeAddress.Phone == "" {
		db.AddError(validations.NewError(franchiseeAddress, "Phone", "Please input phone"))
	}
	if len(db.GetErrors()) == 0 {
		var count uint
		db.Model(&FranchiseeAddress{}).Where("country_code=? and id!=?", franchiseeAddress.CountryCode, franchiseeAddress.ID).Count(&count)
		if count > 0 {
			db.AddError(validations.NewError(franchiseeAddress, "Phone", "Address already exists"))
		}
	}
}

type FranchiseePayAccount struct {
	models.LogicBaseModel

	ShowName            string           // 前台显示名
	Name                string           // 后台自用名
	NotifyUrlDomains    NotifyUrlDomains // 后台自用名
	Type                string           // 类型:base_dictionaries type=3对应的Detail
	AccountId           string           // 账号或者AppID
	PublicKey           string           // 公钥
	PrivateKey          string           // 私钥
	AliPublicKey        string           // 支付宝公钥
	SecrecyKey          string           // 密钥(针对无公钥私钥系统)
	Currency            string           // 支持的币种用,分割
	SellerShowName      string           // 卖家显示名称
	IsProduction        bool             // 是否是正式环境
	PayAgentStatus      bool             // 是否使用代理支付
	IsCollectIdCard     bool             // 是否收集身份证
	PaymentCompanyName  string           // 支付企业名称
	PaymentCompanyCode  string           // 支付企业代码
	RemitName           string           // 汇款名
	RemitAccount        string           // 汇款账号
	RemitBankName       string           // 汇款银行名称
	RemitBankBranchName string           // 汇款银行分行名称
	RemitPhone          string           // 汇款联系电话
	LanguageCode        string           // 本地化
}

func (FranchiseePayAccount) TableName() string {
	return "user_franchisee_pay_accounts"
}

type NotifyUrlDomains []NotifyUrlDomain

type NotifyUrlDomain struct {
	DBWorkID  uint16 // 0:默认,1:中国,3:美国
	NotifyUrl string // 异步地址
}

func (notifyUrlDomains *NotifyUrlDomains) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, notifyUrlDomains)
	case string:
		if v != "" {
			return notifyUrlDomains.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (notifyUrlDomains NotifyUrlDomains) Value() (driver.Value, error) {
	if len(notifyUrlDomains) == 0 {
		return nil, nil
	}
	if data, err := json_utils.Encode(notifyUrlDomains); err == nil {
		return qorUtils.ReplaceQuoted(string(data)), err
	} else {
		return nil, err
	}
}

func (franchiseePayAccount FranchiseePayAccount) GetNotifyUrlDomain() string {
	if franchiseePayAccount.ID == 0 || franchiseePayAccount.NotifyUrlDomains == nil || len(franchiseePayAccount.NotifyUrlDomains) <= 0 {
		return ""
	}

	notifyEntity := linq.From(franchiseePayAccount.NotifyUrlDomains).Where(func(c interface{}) bool {
		return c.(NotifyUrlDomain).DBWorkID == config.Config.DB.MachineID
	}).Single().(NotifyUrlDomain)

	if notifyEntity.DBWorkID == 0 && notifyEntity.NotifyUrl == "" {
		notifyEntity = linq.From(franchiseePayAccount.NotifyUrlDomains).Where(func(c interface{}) bool {
			return c.(NotifyUrlDomain).DBWorkID == 0
		}).Single().(NotifyUrlDomain)
	}

	return notifyEntity.NotifyUrl
}

func GetAllPartnerFranchisees(db *gorm.DB, partnerId uint) []FranchiseeInformation {
	var franchisees []FranchiseeInformation
	if err := db.Where("partner_id = ?", partnerId).Find(&franchisees).Error; err != nil {
		log_utils.PrintErrorStackTrace(err)
	}
	return franchisees
}

func GetPartnerFranchisees(db *gorm.DB, partnerId uint) []FranchiseeInformation {
	var (
		allFranchisees = GetAllPartnerFranchisees(db, partnerId)
		franchisees    []FranchiseeInformation
	)
	linq.From(allFranchisees).Where(func(record interface{}) bool {
		var franchisee = record.(FranchiseeInformation)
		if franchisee.ID == partnerId {
			return false
		} else {
			return true
		}
	}).ToSlice(&franchisees)
	return franchisees
}

func GetFranchisee(db *gorm.DB, partnerId uint) *FranchiseeInformation {
	var franchisee FranchiseeInformation
	db.Where("id = ?", partnerId).Find(&franchisee)
	return &franchisee
}

func GetUserFranchiseePayAccountByPartnerId(db *gorm.DB, partnerId uint, payType string, languageCode string) *FranchiseePayAccount {
	var franchiseeAccount FranchiseePayAccount
	db.Where("partner_id = ? and type = ? and language_code = ?", partnerId, payType, languageCode).Find(&franchiseeAccount)
	return &franchiseeAccount
}

type PartnerShipper struct {
	models.LogicBaseModel
	ShipperID uint
}

type FranchiseeWeChatApplet struct {
	models.LogicBaseModel
	FranchiseeID  uint   `gorm:"not null;default:0"` // 加盟ID
	Name          string `gorm:"size:100"`           // 小程序名称
	AppID         string `gorm:"size:50"`            // 小程序AppID
	AppSecret     string `gorm:"size:50"`            // 小程序Secret
	WebViewDomain string // 小程序内嵌网站域名
	ExtJson       string `gorm:"type:text"`

	EarlyAccessVersion string `gorm:"size:20"` // 体验版本
	OnlineVersion      string `gorm:"size:20"` // 当前线上版本
	AuditID            uint   // 审核编号
	AuditStatus        uint   `gorm:"size:1"`    // 审核状态（1：已上传；2：待审核；3：审核通过，4，审核不通过，5：已发布）
	Reason             string `gorm:"size:1000"` // 审核不通过的原因

	AuthorizerAccessToken            string     // 接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
	AuthorizerAccessTokenRefreshTime *time.Time // 接口调用令牌创建时间
	AuthorizerAccessTokenExpiresIn   uint       // authorizer_access_token 的有效期（在授权的公众号/小程序具备API权限时，才有此返回值），单位：秒
	AuthorizerRefreshToken           string     // 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
	FuncInfo                         string     // 授权给开发者的权限集列表
}

func (FranchiseeWeChatApplet) TableName() string {
	return "user_franchisee_we_chat_applets"
}
