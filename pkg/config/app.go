package config

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	db, err := gorm.Open(sqlite.Open("bookStore.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect to the database! \n", err)
		os.Exit(2)
	}

	fmt.Println("Connected to the database successfully")
	// db.Logger = logger.Default.LogMode(logger.Info)
	fmt.Println("Running Migrations")

	Database = DbInstance{Db: db}
}

func GetDB() *gorm.DB {
	return Database.Db
}
