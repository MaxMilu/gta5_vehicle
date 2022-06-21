package vehicle

import (
	frameAdmin "github.com/qor/admin"
	"github.com/qor/qor-example/config/application"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	//controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("vehicle")}, "app/vehicle/views")}
	//funcmapmaker.AddFuncMapMaker(controller.View)

	admin := application.Admin
	admin.AddMenu(&frameAdmin.Menu{Name: "Vehicle", Priority: 1})
	setupVehiclePage(admin)
}
