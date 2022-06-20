package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/qor-example/config/application"
	"my_qor_test/app/vehicle"
	"my_qor_test/config/db"
	"net/http"
)

func main() {

	// Initialize
	var (
		Router      = chi.NewRouter()
		Admin       = admin.New(&admin.AdminConfig{DB: db.DB})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	Application.Use(vehicle.New(&vehicle.Config{}))

	// initialize an HTTP request multiplexer
	mux := http.NewServeMux()
	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
	})
	r.Any("/admin/*resources", gin.WrapH(mux))
	r.Run("0.0.0.0:" + "9000")
}
