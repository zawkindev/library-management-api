package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     string `gorm:"type:uuid;primaryKey"`
	Title  string `gorm:"type:varchar(255);not null"`
	Author string `gorm:"type:varchar(255);not null"`
	Year   int    `gorm:"type:int"`
	Genre  string `gorm:"type:varchar(100)"`
	ISBN   string `gorm:"type:varchar(20);unique"`
}

// BeforeCreate is a GORM hook that is called before a record is inserted into the database.
func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	book.ID = uuid.New().String() // Generate a new UUID for the ID
	return
}
