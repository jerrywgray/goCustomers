package database

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/jerrywgray/goCustomers/api/schemas"
	"github.com/jerrywgray/goCustomers/api/structs"
	"github.com/jinzhu/gorm"

	// Load postgres bindings for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/palantir/stacktrace"
)

var db *gorm.DB

// Init returns the postgres db
func Init(config *structs.Config) (*gorm.DB, error) {
	options := config.DBConfig

	if options.User == "" || options.Password == "" || options.Database == "" || options.Host == "" || options.Port == "" {
		return nil, stacktrace.NewError("Need all db options in environment")
	}

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		options.Host,
		options.Port,
		options.User,
		options.Database,
		options.Password)

	var err error

	db, err = gorm.Open("postgres", connectionStr)
	if err != nil {
		return nil, stacktrace.Propagate(err, "couldn't connect to db")
	}

	if config.GormMode == "debug" {
		db.LogMode(true)
	}

	if config.Initialize {
		db.DropTableIfExists(&schemas.Customers{})

		db.AutoMigrate(&schemas.Customers{})

		err = loadInitialData(config)
		if err != nil {
			return nil, stacktrace.Propagate(err, "db not initialized")
		}
	}

	return db, nil
}

func loadInitialData(config *structs.Config) error {
	csvPath := config.DataFilename
	csvFile, _ := os.Open(csvPath)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	lineCount := 1
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return stacktrace.Propagate(err, "failed to read line in CSV")
		}
		if line[0] == "" {
			continue
		}

		customers := &schemas.Customers{}

		db.Create(customers)

		lineCount++
	}

	return nil
}
