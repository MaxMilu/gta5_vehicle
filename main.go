package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"my_qor_test/config/db"
	"my_qor_test/model/gta"
	"net/http"
)

func main() {

	// Initialize
	Admin := admin.New(&admin.AdminConfig{DB: db.DB})

	Admin.AddResource(&gta.Vehicle{})

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
