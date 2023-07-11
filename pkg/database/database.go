package database

import (
	"fmt"
	"go-rest-api-sample/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {

	connString := "host=localhost user=dbadmin password=dbadmin dbname=gorestdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Disable GORM's default logger
	})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		panic("Failed to migrate Product table: " + err.Error())
	}

	fmt.Println("Database connection established!")

	return db
}
