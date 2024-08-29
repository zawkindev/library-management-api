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

	// Insert books into the database
	books := []model.Book{
		{Title: "The Fellowship of the Ring", Author: "J.R.R. Tolkien", Year: 1954, Genre: "Fantasy", ISBN: "978-0261103573"},
		{Title: "The Two Towers", Author: "J.R.R. Tolkien", Year: 1954, Genre: "Fantasy", ISBN: "978-0261103580"},
		{Title: "The Return of the King", Author: "J.R.R. Tolkien", Year: 1955, Genre: "Fantasy", ISBN: "978-0261103597"},
	}

	for _, book := range books {
		DB.Create(&book)
	}

	log.Println("Books added successfully")
}
