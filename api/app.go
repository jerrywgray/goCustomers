package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	// Load postgres bindings for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jerrywgray/goCustomers/api/database"
	"github.com/jerrywgray/goCustomers/api/routes"
	"github.com/jerrywgray/goCustomers/api/structs"
)

// Context holds references to our global instances
type Context struct {
	Router *gin.Engine
	DB     *gorm.DB
}

// Init initializes our router and DB
func (a *Context) Init(config *structs.Config) {
	var err error

	a.DB, err = database.Init(config)
	if err != nil {
		log.Fatal(err, "couldn't connect to db")
	}

	a.Router = routes.Init(a.DB, config)
}

// Run starts the router and defers closing the database
func (a *Context) Run(config *structs.Config) {
	defer a.DB.Close()
	a.Router.Run(config.Port)
}
