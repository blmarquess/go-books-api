package database

import (
	"log"

	"github.com/blmarquess/go-books-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := "host=localhost user=golang password=password dbname=books_api port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	db.AutoMigrate(models.Book{})
	log.Println("Database connection successful")
}

func GetDBInstance() *gorm.DB {
	return db
}
