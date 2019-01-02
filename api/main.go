package main

import (
	"os"
	"path/filepath"

	// Load postgres bindings for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"

	"github.com/jerrywgray/goCustomers/api/structs"

	log "github.com/sirupsen/logrus"
)

func main() {
	a := Context{}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":" + os.Getenv("APP_PORT")
	if port == ":" {
		port = ":8080"
	}

	var ginMode string

	if os.Getenv("GIN_MODE") == "debug" {
		ginMode = "debug"
	} else {
		ginMode = "release"
	}

	config := &structs.Config{
		DBConfig: &structs.DBConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
		DataFilename: filepath.Join("./database/data", os.Getenv("DATA_FILENAME")),
		Env:          os.Getenv("APP_ENV"),
		Port:         port,
		GinMode:      ginMode,
		GormMode:     os.Getenv("GORM_MODE"),
		Initialize:   os.Getenv("INITIALIZE") == "1",
	}

	a.Init(config)

	a.Run(config)
}
