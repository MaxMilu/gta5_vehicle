package users

import (
	"fmt"
	"wyaoo/models"
)

type Address struct {
	models.LogicBaseModel
	//l10n.Locale

	UserID         uint
	ContactName    string `form:"contact-name"`
	Phone          string `form:"phone"`
	Email          string `form:"email"`
	Pcc            string `form:"pcc"`
	Province       string `form:"province"`
	City           string `form:"city"`
	Area           string `form:"area"`
	ZipCode        string `form:"zip_code"`
	DefaultAddress bool   `sql:"DEFAULT:0;"` // 是否是默认地址
	Address1       string `form:"address1"`
	Address2       string `form:"address2"`
	LanguageCode   string
}

func (Address) TableName() string {
	return "user_addresses"
}

func (address Address) Stringify() string {
	return fmt.Sprintf("%v, %v, %v", address.Address2, address.Address1, address.City)
}
