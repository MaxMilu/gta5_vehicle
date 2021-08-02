package users

import (
	"database/sql/driver"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
	"net/http"
	"wyaoo/config/consts"
	"wyaoo/models"
	"wyaoo/utils/json_utils"
)

type FranchiseeDomain struct {
	models.LogicBaseModel
	FranchiseeID        uint   `gorm:"not null;default:0"`
	Domain              string `gorm:"not null"`
	ServiceEmail        string `gorm:"not null"`
	LayoutType          string
	LoginType           string
	HomeType            string
	HomeTemplate        string
	HomeLogo            oss.OSS
	AdminLogo           oss.OSS
	FaviconLogo         oss.OSS
	Region              string `gorm:"not null;default'US'"`
	Language            string `gorm:"not null;default'en-US'"`
	DefaultRegisterRole string `gorm:"not null;default'ECV'"`
	Type                uint   //1.合作商 2.供货商 3.代销商 4.转运商
	//WeChatAppID         string `gorm:"size:50"` // 小程序AppID
	//WeChatAppSecret     string `gorm:"size:50"` // 小程序App Secret
}

func (FranchiseeDomain) TableName() string {
	return "user_franchisee_domains"
}

func (franchiseeDomain *FranchiseeDomain) Validate(db *gorm.DB) {
	if franchiseeDomain.Domain == "" {
		db.AddError(validations.NewError(franchiseeDomain, "Domain", "Please input domain"))
	}
	if franchiseeDomain.HomeType == "" {
		db.AddError(validations.NewError(franchiseeDomain, "HomeType", "Please select home type"))
	}
}

func GetCurrentFranchiseeDomain(req *http.Request) (currentFranchiseeDomain *FranchiseeDomain) {
	if v := req.Context().Value(consts.REQUEST_FRANCHISEE_DOMAIN); v != nil {
		currentFranchiseeDomain = v.(*FranchiseeDomain)
	} else {
		currentFranchiseeDomain = &FranchiseeDomain{}
	}
	return
}

func GetAllFranchiseeDomains(db *gorm.DB) (franchiseeDomains []FranchiseeDomain) {
	db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Find(&franchiseeDomains)
	return
}

func GetFranchiseeDomain(db *gorm.DB, domain string) (franchiseeDomains FranchiseeDomain) {
	db.Where("domain=?", domain).First(&franchiseeDomains)
	return
}

type FranchiseeDomainCountry struct {
	models.LogicBaseModel
	FranchiseeID uint                       `gorm:"not null;default:0"`
	Domain       string                     `sql:"not null" gorm:"unique_index:uix_user_franchisee_domain_countries_1"`
	CountryCode  string                     `sql:"size:2;not null" gorm:"unique_index:uix_user_franchisee_domain_countries_1"`
	Languages    []FranchiseeDomainLanguage `gorm:"foreignkey:CountryCode;association_foreignkey:CountryCode"`
}

func (FranchiseeDomainCountry) TableName() string {
	return "user_franchisee_domain_countries"
}

func GetFranchiseeDomainCountries(db *gorm.DB, domain string) (franchiseeCountries []FranchiseeDomainCountry) {
	db.Model(&FranchiseeDomainCountry{}).Where("domain=?", domain).Order("country_code asc").Find(&franchiseeCountries)
	return
}

func (FranchiseeDomainLanguage) TableName() string {
	return "user_franchisee_domain_languages"
}

type FranchiseeDomainLanguage struct {
	models.LogicBaseModel
	FranchiseeID       uint   `gorm:"not null;default:0"`
	Domain             string `gorm:"not null;unique_index:uix_user_franchisee_domain_languages_1"`
	CountryCode        string `gorm:"size:2;not null;unique_index:uix_user_franchisee_domain_languages_1"`
	LanguageRegionCode string `gorm:"size:5;not null;unique_index:uix_user_franchisee_domain_languages_1"`
	Currency           string `gorm:"size:3;not null"`
	WeightUnit         string `gorm:"size:2;not null"`
	LengthUnit         string `gorm:"size:2;not null"`
	Default            bool   `gorm:"not null;default:0"`
}

func GetFranchiseeDomainLanguages(db *gorm.DB, domain string, currentCountryCode string) (franchiseeLanguages []FranchiseeDomainLanguage) {
	db.Model(&FranchiseeDomainLanguage{}).Where("domain=? and country_code=?", domain, currentCountryCode).Order("language_region_code").Find(&franchiseeLanguages)
	return
}

type FranchiseeDomainWidget struct {
	models.LogicBaseModel
	FranchiseeID   uint           `gorm:"not null;default:0"`
	Domain         string         `gorm:"not null;unique_index:uix_user_franchisee_domain_widgets_1"`
	WidgetSettings WidgetSettings `gorm:"type:text"`
}

func (FranchiseeDomainWidget) TableName() string {
	return "user_franchisee_domain_widgets"
}

type DomainWidgetSetting struct {
	Location string
	Widgets  []DomainWidget
}

type DomainWidget struct {
	Index      int
	WidgetName string
}

type WidgetSettings []DomainWidgetSetting

func (setting *WidgetSettings) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, setting)
	case string:
		if v != "" {
			return setting.Scan([]byte(v))
		}
	case []string:
		for _, str := range v {
			if err := setting.Scan(str); err != nil {
				return err
			}
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (setting WidgetSettings) Value() (driver.Value, error) {
	results, err := json_utils.Encode(setting)
	return utils.ReplaceQuoted(string(results)), err
}
