package database

import (
	"fmt"
	"log"
	"os"

	"github.com/blmarquess/go-books-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the Postgres Database\n Error: ", err)
	}

	db = database
	db.AutoMigrate(models.Book{})
	log.Println("Database connection successful")
}

func GetDBInstance() *gorm.DB {
	return db
}
