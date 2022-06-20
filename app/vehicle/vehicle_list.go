package vehicle

import (
	frameAdmin "github.com/qor/admin"
	"my_qor_test/model/gta"
)

func setupVehiclePage(Admin *frameAdmin.Admin) {
	vehiclePageResource := Admin.AddResource(&gta.Vehicle{}, &frameAdmin.Config{Name: "Vehicle List", Menu: []string{"Vehicle"}})
	setVehiclePage(Admin, vehiclePageResource)
}

func setVehiclePage(admin *frameAdmin.Admin, res *frameAdmin.Resource) {

}
