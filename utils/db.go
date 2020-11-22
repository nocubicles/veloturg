package utils

import (
	"fmt"
	"os"

	"github.com/nocubicles/veloturg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DbConnection returns the connection to use the db
func DbConnection() (db *gorm.DB) {
	dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to open db connection")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Ad{},
		&models.Session{},
	)

	if err != nil {
		panic(err)
	}

	return db
}
