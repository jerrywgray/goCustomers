package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/jerrywgray/goCustomers/api/structs"
)

var db *gorm.DB

// Init Initializes all of our routes and gives us the reference to our db instance
func Init(_db *gorm.DB, config *structs.Config) *gin.Engine {
	db = _db

	gin.SetMode(config.GinMode)

	router := gin.Default()

	router.GET("/", root)

	return router
}
