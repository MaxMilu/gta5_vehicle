package users

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/qor/media"
	"github.com/qor/media/oss"
	qorUtils "github.com/qor/qor/utils"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
	"time"
	"wyaoo/config/consts"
	"wyaoo/models"
	"wyaoo/utils/json_utils"
)

type User struct {
	models.LogicBaseModel
	FranchiseeID           uint `sql:"NOT NULL"`
	BelongToWarehouseCode  string
	OfficeID               uint
	Email                  string `form:"email"`
	Password               string
	APIToken               string
	Name                   string `form:"name"`
	Code                   string // 用户代码
	UserLabel              string // 用户标签
	ThirdPartyCode         string // 第三方用户代码(用户备案关联)
	Gender                 string
	Role                   string `gorm:"default:'ECV'"`
	Birthday               *time.Time
	Balance                float32
	DefaultBillingAddress  uint `form:"default-billing-address"`
	DefaultShippingAddress uint `form:"default-shipping-address"`
	Addresses              []Address
	Avatar                 AvatarImageStorage
	Referrer               uint
	Wechat                 string
	Alipay                 string
	BankAccount            string
	BankUsername           string
	BankName               string
	StoreIndex             uint
	// Confirm
	ConfirmToken string
	Confirmed    bool

	Emails UserEmails // 自定义用户邮箱

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time `gorm:"-"`

	// Accepts
	AcceptPrivate bool `form:"accept-private"`
	AcceptLicense bool `form:"accept-license"`
	AcceptNews    bool `form:"accept-news"`

	UserGroups                []GroupUser
	UserLevel                 uint            // 会员等级
	ConsumptionAmount         decimal.Decimal // 消费金额
	FingerprintCode           string          // 指纹识别码
	UserCategory              uint            // 用户分类（0：用户，1：正式员工，2：临时员工）
	DeductionPosition         uint            `gorm:"not null:default:1"` // 扣费时间点类型（0:下单扣费；1:Dispatch确认发货后扣费）
	AutoNotifyStorageShipping bool            // 自动通知仓储运单发货
}

func (User) TableName() string {
	return consts.TABLE_NAME_USER_USERS
}

func (user User) Stringify() string {
	return fmt.Sprintf("%v", user.Email)
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) Label() string {
	if strings.TrimSpace(user.Name) == consts.EMPTY {
		return user.Email
	} else {
		return fmt.Sprintf("%v[%v]", user.Email, user.Name)
	}
}

func (user User) AvailableLocales(req *http.Request) []string {
	var (
		languages               []FranchiseeDomainLanguage
		locales                 []string
		currentFranchiseeDomain = GetCurrentFranchiseeDomain(req)
		db                      = qorUtils.GetDBFromRequest(req)
	)
	db.Find(&languages, "domain=?", currentFranchiseeDomain.Domain)
	for _, language := range languages {
		countryLanguageCode := language.CountryCode + " " + language.LanguageRegionCode
		locales = append(locales, countryLanguageCode)
	}
	return locales
}

type AvatarImageStorage struct{ oss.OSS }

func (AvatarImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 50, Height: 50},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}

type UserConfig struct {
	models.LogicBaseModel
	UserID            string
	RequestLimitCount float64
}

func GetUser(db *gorm.DB, ID uint) User {
	var user User
	db.Where("id=?", ID).Find(&user)
	return user
}

type UserEmails []UserEmail

func (UserEmails *UserEmails) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, UserEmails)
	case string:
		if v != "" {
			return UserEmails.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (UserEmails UserEmails) Value() (driver.Value, error) {
	if len(UserEmails) == 0 {
		return nil, nil
	}
	if data, err := json_utils.Encode(UserEmails); err == nil {
		return qorUtils.ReplaceQuoted(string(data)), err
	} else {
		return nil, err
	}
}

type UserEmail struct {
	Email string
}
