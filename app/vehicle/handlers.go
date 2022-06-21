package vehicle

import (
	"github.com/qor/admin"
	qorRender "github.com/qor/render"
	"my_qor_test/config/consts"
	"my_qor_test/model/gta"
	"my_qor_test/model/vo"
	"my_qor_test/utils/log_utils"
)

type Controller struct {
	View *qorRender.Render
}

func getManifestScopeQuantity(context *admin.Context) {
	var (
		dbs          = context.GetDB()
		functionName = context.Request.FormValue("functionName")
		taskStatus   = context.Request.FormValue("taskStatus")
		result       *vo.AjaxResult
	)
	defer func() {
		if _, err := context.Writer.Write(result.Marshal()); err != nil {
			log_utils.PrintErrorStackTrace(err)
		}
	}()

	switch functionName {
	case "All":
		taskStatus = consts.EMPTY
	case "Like":
		dbs = dbs.Where("like = ?", true)
		taskStatus = consts.NUM_0
	case "Wishlist":
		dbs = dbs.Where("wishlist = ?", true)
		taskStatus = consts.NUM_0
	case "AlreadyHas":
		dbs = dbs.Where("already_has = ?", true)
		taskStatus = consts.NUM_0
	}

	var count uint
	if taskStatus == consts.EMPTY {
		if err := dbs.Model(&gta.Vehicle{}).Count(&count).Error; err != nil {
			log_utils.PrintErrorStackTrace(err)
			result = vo.AjaxSystemError
			return
		}
	} else {
		if err := dbs.Model(&gta.Vehicle{}).Count(&count).Error; err != nil {
			log_utils.PrintErrorStackTrace(err)
			result = vo.AjaxSystemError
			return
		}
	}
	result = &vo.AjaxResult{
		Status: consts.SUCCESS,
		Result: map[string]uint{
			"Quantity": count,
		},
	}
}
