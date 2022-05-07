package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"net/http"
)

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})

	// Initialize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

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
