package vehicle

import (
	"bytes"
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

	res.Meta(&frameAdmin.Meta{Name: "VoMainImage", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*voVehicle); ok {
			result := bytes.NewBufferString("")
			tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
			tmpl.Execute(result, map[string]string{"image": p.MainImageURL})
			return template.HTML(result.String())
		}
		return ""
	}})

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

}
