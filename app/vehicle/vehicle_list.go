package vehicle

import (
	"bytes"
	"github.com/jinzhu/gorm"
	frameAdmin "github.com/qor/admin"
	"github.com/qor/qor"
	frameUtils "github.com/qor/qor/utils"
	"html/template"
	"my_qor_test/model/gta"
	"my_qor_test/utils/log_utils"
)

type voVehicle struct {
	gta.Vehicle
	VoMainImage string
}

func setupVehiclePage(Admin *frameAdmin.Admin) {
	vehiclePageResource := Admin.AddResource(&voVehicle{}, &frameAdmin.Config{Name: "Vehicle List", Menu: []string{"Vehicle"}})
	setVehiclePage(Admin, vehiclePageResource)
}

func setVehiclePage(admin *frameAdmin.Admin, res *frameAdmin.Resource) {

	res.IndexAttrs("VoMainImage", "CategoryName", "Brand", "Name")

	res.UseTheme("grid")
	//res.UseTheme("vehicle_list")

	//admin.GetRouter().Post("/gta/vehicle_list_get_scope_quantity", getManifestScopeQuantity)

	res.Scope(&frameAdmin.Scope{
		Name: "All",
		Handler: func(dbs *gorm.DB, context *qor.Context) *gorm.DB {
			return context.GetDB()
		},
		Default: true,
	})
	res.Scope(&frameAdmin.Scope{
		Name:  "LikeIt",
		Label: "Like",
		Handler: func(dbs *gorm.DB, context *qor.Context) *gorm.DB {
			return dbs.Where("like_it = ?", true)
		},
	})
	res.Scope(&frameAdmin.Scope{
		Name: "Wishlist",
		Handler: func(dbs *gorm.DB, context *qor.Context) *gorm.DB {
			return dbs.Where("wishlist = ?", true)
		},
	})
	res.Scope(&frameAdmin.Scope{
		Name: "AlreadyHas",
		Handler: func(dbs *gorm.DB, context *qor.Context) *gorm.DB {
			return dbs.Where("already_has = ?", true)
		},
	})

	//	region Filter
	res.Filter(&frameAdmin.Filter{
		Name: "CategoryName",
		Type: "select_one",
		Config: &frameAdmin.SelectOneConfig{
			Collection: func(record interface{}, context *frameAdmin.Context) [][]string {
				var (
					tx       = context.GetDB()
					vehicles []gta.Vehicle
					options  [][]string
				)
				if err := tx.Select("DISTINCT category_name").Find(&vehicles).Error; err != nil {
					log_utils.PrintErrorStackTrace(err)
				} else {
					for _, v := range vehicles {
						options = append(options, []string{frameUtils.ToString(v.CategoryName), v.CategoryName})
					}
				}
				return options
			},
			AllowBlank: true,
		},
	})
	res.Filter(&frameAdmin.Filter{
		Name: "Brand",
		Type: "select_one",
		Config: &frameAdmin.SelectOneConfig{
			Collection: func(record interface{}, context *frameAdmin.Context) [][]string {
				var (
					tx       = context.GetDB()
					vehicles []gta.Vehicle
					options  [][]string
				)
				if err := tx.Select("DISTINCT brand").Find(&vehicles).Error; err != nil {
					log_utils.PrintErrorStackTrace(err)
				} else {
					for _, v := range vehicles {
						options = append(options, []string{frameUtils.ToString(v.Brand), v.Brand})
					}
				}
				return options
			},
			AllowBlank: true,
		},
	})
	res.Filter(&frameAdmin.Filter{
		Name: "Name",
		Type: "select_one",
		Config: &frameAdmin.SelectOneConfig{
			Collection: func(record interface{}, context *frameAdmin.Context) [][]string {
				var (
					tx       = context.GetDB()
					vehicles []gta.Vehicle
					options  [][]string
				)
				if err := tx.Select("DISTINCT name").Find(&vehicles).Error; err != nil {
					log_utils.PrintErrorStackTrace(err)
				} else {
					for _, v := range vehicles {
						options = append(options, []string{frameUtils.ToString(v.Name), v.Name})
					}
				}
				return options
			},
			AllowBlank: true,
		},
	})
	//	endregion

	//	region Meta
	res.Meta(&frameAdmin.Meta{Name: "VoMainImage", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*voVehicle); ok {
			result := bytes.NewBufferString("")
			tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
			err := tmpl.Execute(result, map[string]string{"image": p.MainImageURL})
			if err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return template.HTML(result.String())
		}
		return ""
	}})
	//	endregion

	//	region Action
	res.Action(&frameAdmin.Action{
		Name: "Delete",
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			return false
		},
	})

	res.Action(&frameAdmin.Action{
		Name: "Like",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("like", true).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if !this.LikeIt {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	res.Action(&frameAdmin.Action{
		Name: "UnLike",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("like_it", false).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if this.LikeIt {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	res.Action(&frameAdmin.Action{
		Name: "Wishlist",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("wishlist", true).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if !this.Wishlist {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	res.Action(&frameAdmin.Action{
		Name: "RemoveWishlist",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("wishlist", false).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if this.Wishlist {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	res.Action(&frameAdmin.Action{
		Name: "AlreadyHas",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("already_has", true).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if !this.AlreadyHas {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	res.Action(&frameAdmin.Action{
		Name: "UnAlreadyHas",
		Handler: func(argument *frameAdmin.ActionArgument) error {
			var (
				this gta.Vehicle
				dbs  = argument.Context.GetDB()
			)
			if err := dbs.Select("id").Where("id = ?", argument.Context.ResourceID).Find(&this).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			if err := dbs.Model(&gta.Vehicle{}).Where("id = ?", this.ID).Update("already_has", false).Error; err != nil {
				return log_utils.PrintErrorStackTraceFriendly(err)
			}
			return nil
		},
		Visible: func(record interface{}, context *frameAdmin.Context) bool {
			var (
				this = record.(*voVehicle)
			)
			if this.AlreadyHas {
				return true
			}
			return false
		},
		Modes: []string{"menu_item"},
	})
	//	region

}
