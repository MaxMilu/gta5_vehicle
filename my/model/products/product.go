package products

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/qor/l10n"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/media/oss"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	qorUtils "github.com/qor/qor/utils"
	qorSeo "github.com/qor/seo"
	"github.com/qor/slug"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"github.com/shopspring/decimal"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"wyaoo/config/consts"
	"wyaoo/config/db"
	"wyaoo/models"
	"wyaoo/models/sites"
	"wyaoo/models/users"
	"wyaoo/utils/json_utils"
)

type Product struct {
	models.LogicBaseModel
	l10n.Locale
	sorting.SortingDESC
	Type                         string // 商品类型（0：自营商品，9：代购商品）
	Name                         string
	NameWithSlug                 slug.Slug `l10n:"sync"`
	Featured                     bool
	Code                         string          `l10n:"sync"`
	BrandID                      uint            `l10n:"sync"`
	Brand                        ProductBrand    `l10n:"sync"`
	MaterialID                   uint            `l10n:"sync"`
	Material                     ProductMaterial `l10n:"sync"`
	ProductCategories            []ProductCategory
	Collections                  []Collection `gorm:"many2many:product_product_collections;"`
	MadeCountry                  string       `l10n:"sync"`
	Gender                       string       `l10n:"sync"`
	MainImage                    media_library.MediaBox
	Price                        decimal.Decimal // 商品价值
	Quantity                     uint
	BriefDescription             string // 商品简介
	DressSizeID                  uint
	DressSize                    DressSize
	Description                  string                     `sql:"type:text"` // 商品说明
	ProductVariations            []ProductVariation         // 商品选项列表
	ProductVariationsSorter      sorting.SortableCollection `sql:"type:text"` // 商品选项排序
	ColorVariations              []ColorVariation
	ColorVariationsSorter        sorting.SortableCollection
	ProductProperties            ProductProperties `sql:"type:text"`
	Seo                          qorSeo.Setting
	SupplierID                   uint   `l10n:"sync"`
	Currency                     string `sql:"size:3"`
	Specification                string
	PlugSize                     string                 //电器的标准 请填写Chinese Plug 或 US Plug 或 UK Plug 或 EU Plug
	OtherImages                  media_library.MediaBox // 其它商品图片
	Freight                      decimal.Decimal        // 派送费
	FreightCurrency              string                 `sql:"size:3"` // 派送地币种
	OverseasPurchasingFeePercent decimal.Decimal        // 代购费扣费百分比
	OverseasPurchasingFeeFix     decimal.Decimal        // 代购费扣费固定值
	PartnerMarkupPercent         decimal.Decimal        // 平台加价百分比
	PartnerMarkupFix             decimal.Decimal        // 平台加价固定值
	PartnerMarkupPrice           decimal.Decimal        // 平台加价价格
	SellerAgentAllow             bool                   // 是否允许代销
	SellerAgentMarkupPercentMin  decimal.Decimal        // 允许代销商加价额度百分比下限
	SellerAgentMarkupPercentMax  decimal.Decimal        // 允许代销商加价额度百分比上限
	SellerAgentMarkupFix         decimal.Decimal        // 允许代销商加价额度固定值
	SellerAgentMarkupPrice       decimal.Decimal        // 允许代销商加价额度
	PartnerAgentAllow            bool                   // 是否允许其他合作商代销
	OriginalPartnerID            uint                   // 商品归属合作商ID
	PartnerAgentMarkupPercent    decimal.Decimal        // 代理合作商加价百分比
	PartnerAgentMarkupFix        decimal.Decimal        // 代理合作商加价固定值
	PartnerAgentMarkupPrice      decimal.Decimal        // 代理合作商加价价格
	FinalPrice                   decimal.Decimal        // 最终售价
	PercentOff                   decimal.Decimal        `gorm:"not null;default:0"` // 折扣百分比
	PriceOff                     decimal.Decimal        `gorm:"not null;default:0"` // 折扣固定金额
	OffValue                     decimal.Decimal        `gorm:"not null;default:0"` // 折扣值
	MinFinalPrice                decimal.Decimal        // 最小售价价值
	MaxFinalPrice                decimal.Decimal        // 最大售价价值
	RetailPrice                  decimal.Decimal        // 参考零售价
	AvailableDate                string                 // 发售日期
	SalesVolume                  decimal.Decimal        // 销量
	Evaluations                  []Evaluation           // 评论
	EvaluationCount              uint                   // 评论次数
	HSCode                       uint
	TaxCode                      uint
	BarCode                      uint
	Weight                       decimal.Decimal // 重量
	WeightUnit                   string          // 重量单位
	ThirdPartyProductUrl         string          // 第三方商品URL
	Favourite                    bool            `gorm:"-"` // 第三方商品URL
	publish2.Version
	publish2.Schedule
	publish2.Visible
}

func (Product) TableName() string {
	return "product_products"
}

type ProductVariation struct {
	models.LogicBaseModel
	l10n.Locale
	ProductID               uint            `gorm:"not null;default:0;"` // 商品ID
	Product                 Product         // 商品
	Option1ID               uint            `gorm:"not null;default:0;"`                            // 商品选项1ID
	Color                   Color           `gorm:"foreignkey:Option1ID;association_foreignkey:ID"` // 商品选项1
	Option2ID               uint            `gorm:"not null;default:0;"`                            // 商品选项2ID
	Size                    Size            `gorm:"foreignkey:Option2ID;association_foreignkey:ID"` // 商品选项2
	Option3ID               uint            `gorm:"not null;default:0;"`                            // 商品选项3ID
	Option3                 Option3         // 商品选项3
	Option4ID               uint            `gorm:"not null;default:0;"` // 商品选项4ID
	Option4                 Option4         // 商品选项4
	Option5ID               uint            `gorm:"not null;default:0;"` // 商品选项5ID
	Option5                 Option5         // 商品选项5
	Option6ID               uint            `gorm:"not null;default:0;"` // 商品选项6ID
	Option6                 Option6         // 商品选项6
	MaterialID              uint            `gorm:"not null;default:0"` // 材质ID
	Material                ProductMaterial // 材质
	Code                    string          // 用户自己输入
	SKU                     string          `gorm:"not null;unique_index:uix_product_variations_2" l10n:"sync"` // SKU
	ReceiptName             string          // 发票显示名称
	Featured                bool            // 是否推荐商品
	PartnerMarkupPrice      decimal.Decimal `gorm:"not null;default:0"` // 平台加价价格 TODO 没用了等待被删除
	SellerAgentMarkupPrice  decimal.Decimal `gorm:"not null;default:0"` // 允许代销商加价额度 TODO 没用了等待被删除
	PartnerAgentMarkupPrice decimal.Decimal `gorm:"not null;default:0"` // 代理合作商加价价格 TODO 没用了等待被删除
	Price                   decimal.Decimal `gorm:"not null;default:0"` // 商品原价
	OverseasPurchasingFee   decimal.Decimal // 代购费
	//PercentOff              decimal.Decimal        `gorm:"not null;default:0"` // 折扣百分比
	//PriceOff                decimal.Decimal        `gorm:"not null;default:0"` // 折扣固定金额
	//OffValue           decimal.Decimal        `gorm:"not null;default:0"` // 折扣值
	SellingPrice       decimal.Decimal        `gorm:"not null;default:0"` // 商品最终售价
	MinSellingQuantity uint                   `gorm:"not null;default:0"` // 最低起售数量
	MaxSellingQuantity uint                   `gorm:"not null;default:0"` // 最高销售数量
	AvailableQuantity  uint                   `gorm:"not null;default:0"` // 商品库存
	Images             media_library.MediaBox // 图片
	publish2.SharedVersion
}

func (productVariation ProductVariation) MainImageURL() string {
	if len(productVariation.Images.Files) > 0 {
		return productVariation.Images.URL()
	}
	return "/images/default_product.png"
}

func (ProductVariation) TableName() string {
	return "product_variations"
}

type ProductSelectMany struct {
	ProductSelectMany []string
}

func (product Product) GetSEO() *qorSeo.SEO {
	return sites.SEOCollection.GetSEO("Product Page")
}

func (product Product) RenderSEO(req *http.Request) template.HTML {
	var tagValues map[string]string
	replace := func(str string) string {
		re := regexp.MustCompile("{{([a-zA-Z0-9]*)}}")
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			str = strings.Replace(str, match[0], tagValues[match[1]], 1)
		}
		return str
	}

	product.Seo.Title = replace(product.Seo.Title)
	product.Seo.Description = replace(product.Seo.Description)
	product.Seo.Keywords = replace(product.Seo.Keywords)
	product.Seo.Type = replace(product.Seo.Type)
	product.Seo.OpenGraphURL = replace(product.Seo.OpenGraphURL)
	product.Seo.OpenGraphImageURL = replace(product.Seo.OpenGraphImageURL)
	product.Seo.OpenGraphType = replace(product.Seo.OpenGraphType)
	for idx, metadata := range product.Seo.OpenGraphMetadata {
		product.Seo.OpenGraphMetadata[idx] = qorSeo.OpenGraphMetadata{
			Property: replace(metadata.Property),
			Content:  replace(metadata.Content),
		}
	}
	return product.Seo.FormattedHTML(&qor.Context{Request: req})
}

func (product Product) DefaultPath() string {
	defaultPath := fmt.Sprintf("/products/%s", product.Code)
	return defaultPath
}

//4.28 加 更新商品信息
func UpdateProductFromFranchiseebyPartnerAgentAllow(db *gorm.DB, currentUser *users.User, partnerAgentAllow bool) {

	var productArray []Product
	db.Set(consts.WYAOO_OMIT_PARTNER_ID, true).Where("partner_id=?", currentUser.PartnerID).Find(&productArray)

	for _, v := range productArray {
		db.Model(v).Update("PartnerAgentAllow", partnerAgentAllow)
	}

}

func UpdateProductFromFranchiseebyPartnerMarkupPercent(db *gorm.DB, currentUser *users.User, franchiseeInformation *users.FranchiseeInformation, a bool, b bool, c bool, d bool) {
	franchiseeInformationMap := make(map[string]interface{})

	var productArray []Product
	db.Where("partner_id=?", currentUser.PartnerID).Find(&productArray)

	for _, v := range productArray {
		if a {
			franchiseeInformationMap["PartnerMarkupPercent"] = franchiseeInformation.PartnerMarkupPercent
			franchiseeInformationMap["PartnerMarkupPrice"] = franchiseeInformation.PartnerMarkupPercent.Mul(v.Price)
			v.PartnerMarkupPercent = franchiseeInformation.PartnerMarkupPercent
			//v.PartnerMarkupPrice = franchiseeInformation.PartnerMarkupPercent.Mul(v.Price)
		}
		if b {
			franchiseeInformationMap["SellerAgentMarkupPercentMin"] = franchiseeInformation.SellerAgentMarkupPercentMin
			v.SellerAgentMarkupPercentMin = franchiseeInformation.SellerAgentMarkupPercentMin

		}
		if c {
			franchiseeInformationMap["SellerAgentMarkupPercentMax"] = franchiseeInformation.SellerAgentMarkupPercentMax
			franchiseeInformationMap["SellerAgentMarkupPrice"] = franchiseeInformation.SellerAgentMarkupPercentMax.Mul(v.Price)
			v.SellerAgentMarkupPercentMax = franchiseeInformation.SellerAgentMarkupPercentMax
			//v.SellerAgentMarkupPrice = franchiseeInformation.SellerAgentMarkupPercentMax.Mul(v.Price)
		}
		if d {
			if franchiseeInformation.PartnerAgentAllow {
				franchiseeInformationMap["PartnerAgentAllow"] = true
				//v.PartnerAgentAllow = true
			} else {
				franchiseeInformationMap["PartnerAgentAllow"] = false
				//v.PartnerAgentAllow = false
			}
		}
		//franchiseeInformationMap["FinalPrice"] = v.Price.Add(v.PartnerMarkupPrice).Add(v.SellerAgentMarkupPrice)
		//db.Model(v).Update(v)
		db.Model(&Product{}).Where("id = ?", v.ID).UpdateColumns(franchiseeInformationMap)
	}
}

func (product Product) MainImageURL(styles ...string) string {
	style := "main"
	if len(styles) > 0 {
		style = styles[0]
	}

	if len(product.MainImage.Files) > 0 {
		return product.MainImage.URL(style)
	}

	for _, cv := range product.ProductVariations {
		return cv.MainImageURL()
	}

	return "/images/default_product.png"
}

func (product Product) CurrencySign() string {
	signMap := make(map[string]string)
	signMap["CNY"] = "￥"
	signMap["USD"] = "$"
	signMap["EUR"] = "€"
	signMap["KRW"] = "₩"
	return signMap[product.Currency]
}

func (product Product) Validate(db *gorm.DB) {
	if strings.TrimSpace(product.Name) == "" {
		db.AddError(validations.NewError(product, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(product.Code) == "" {
		db.AddError(validations.NewError(product, "Code", "Code can not be empty"))
	}
}

type ProductImage struct {
	models.LogicBaseModel
	FranchiseeID uint `gorm:"not null;default:0"` // 供应商ID
	Title        string
	Color        Color
	ColorID      uint
	Category     Category
	CategoryID   uint
	SelectedType string
	File         media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
}

func (ProductImage) TableName() string {
	return "product_images"
}

func (productImage ProductImage) Validate(db *gorm.DB) {
	if strings.TrimSpace(productImage.Title) == "" {
		db.AddError(validations.NewError(productImage, "Title", "Title can not be empty"))
	}
}

func (productImage *ProductImage) SetSelectedType(typ string) {
	productImage.SelectedType = typ
}

func (productImage *ProductImage) GetSelectedType() string {
	return productImage.SelectedType
}

func (productImage *ProductImage) ScanMediaOptions(mediaOption media_library.MediaOption) error {
	if bytes, err := json_utils.Encode(mediaOption); err == nil {
		return productImage.File.Scan(bytes)
	} else {
		return err
	}
}

func (productImage *ProductImage) GetMediaOption() (mediaOption media_library.MediaOption) {
	mediaOption.Video = productImage.File.Video
	mediaOption.FileName = productImage.File.FileName
	mediaOption.URL = productImage.File.URL()
	mediaOption.OriginalURL = productImage.File.URL("original")
	mediaOption.CropOptions = productImage.File.CropOptions
	mediaOption.Sizes = productImage.File.GetSizes()
	mediaOption.Description = productImage.File.Description
	return
}

type ProductProperties []ProductProperty

type ProductProperty struct {
	Name  string
	Value string
}

func (productProperties *ProductProperties) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json_utils.Decode(v, productProperties)
	case string:
		if v != "" {
			return productProperties.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (productProperties ProductProperties) Value() (driver.Value, error) {
	if len(productProperties) == 0 {
		return nil, nil
	}
	if data, err := json_utils.Encode(productProperties); err == nil {
		return qorUtils.ReplaceQuoted(string(data)), err
	} else {
		return nil, err
	}
}

type ColorVariation struct {
	models.LogicBaseModel
	Product        Product
	ProductID      uint
	ColorID        uint
	Color          Color
	ColorCode      string
	Images         media_library.MediaBox
	SizeVariations []SizeVariation
	publish2.SharedVersion
}

func (ColorVariation) TableName() string {
	return "product_color_variations"
}

// ViewPath view path of color variation
func (colorVariation ColorVariation) ViewPath() string {
	defaultPath := ""
	var product Product
	if !db.DB.First(&product, "id = ?", colorVariation.ProductID).RecordNotFound() {
		defaultPath = fmt.Sprintf("/products/%s_%s", product.Code, colorVariation.ColorCode)
	}
	return defaultPath
}

// 获取数量
func (colorVariation ColorVariation) SizeVariationCount() uint {
	var sumCount uint
	for _, obj := range colorVariation.SizeVariations {
		sumCount += obj.AvailableQuantity
	}
	return sumCount
}

func (colorVariation ColorVariation) StrID() string {
	return qorUtils.ToString(colorVariation.ID)
}

type ColorVariationImage struct {
	gorm.Model
	ColorVariationID uint
	Image            ColorVariationImageStorage `sql:"type:varchar(4096)"`
}

func (ColorVariationImage) TableName() string {
	return "product_color_variation_images"
}

type ColorVariationImageStorage struct{ oss.OSS }

func (colorVariation ColorVariation) MainImageURL() string {
	if len(colorVariation.Images.Files) > 0 {
		return colorVariation.Images.URL()
	}
	return "/images/default_product.png"
}

func (ColorVariationImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 320, Height: 320},
		"middle": {Width: 640, Height: 640},
		"big":    {Width: 1280, Height: 1280},
	}
}

type SizeVariation struct {
	models.LogicBaseModel
	l10n.Locale
	ColorVariationID        uint
	ColorVariation          ColorVariation
	SizeID                  uint
	Size                    Size
	Sku                     string          // 唯一码
	AvailableQuantity       uint            // 库存数量
	Price                   decimal.Decimal // 商品价格
	PartnerMarkupPrice      decimal.Decimal // 平台加价价格
	SellerAgentMarkupPrice  decimal.Decimal // 允许代销商加价额度
	PartnerAgentMarkupPrice decimal.Decimal // 代理合作商加价价格
	SalePrice               decimal.Decimal // 商品销售价格
	SalesVolume             decimal.Decimal // 销量
	publish2.SharedVersion
}

func (SizeVariation) TableName() string {
	return "product_size_variations"
}

func SizeVariations() []SizeVariation {
	sizeVariations := make([]SizeVariation, 0)
	if err := db.DB.Preload("ColorVariation.Color").Preload("ColorVariation.Product").Preload("Size").Find(&sizeVariations).Error; err != nil {
		log.Fatalf("query sizeVariations (%v) failure, got err %v", sizeVariations, err)
		return sizeVariations
	}
	return sizeVariations
}

func (sizeVariation SizeVariation) Stringify() string {
	if colorVariation := sizeVariation.ColorVariation; colorVariation.ID != 0 {
		product := colorVariation.Product
		return fmt.Sprintf("%s (%s-%s-%s)", product.Name, product.Code, colorVariation.Color.Code, sizeVariation.Size.Code)
	}
	return fmt.Sprint(sizeVariation.ID)
}

type ProductCategory struct {
	models.PhysicsBaseModel
	FranchiseeID     uint
	Type             uint
	ProductPartnerID uint
	ProductID        uint
	Product          Product
	CategoryID       uint
	Category         Category
}

func (ProductCategory) TableName() string {
	return "product_product_categories"
}
