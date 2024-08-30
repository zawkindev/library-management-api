package database

import (
	"library-management-api/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "user=postgres password=m dbname=library sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to the database: ", err)
	}

	// Auto-migrate the Book model
	DB.AutoMigrate(&model.Book{})
}
