package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=m dbname=library sslmodule=require")
	if err != nil {
		log.Fatal("Unable to connect to the database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}
}
