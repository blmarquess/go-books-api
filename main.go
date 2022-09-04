package main

import (
	"os"

	"github.com/blmarquess/go-books-api/database"
	"github.com/blmarquess/go-books-api/server"
)

var PORT = os.Getenv("PORT")

func main() {
	database.StartDB()
	server := server.NewServer("8000")
	server.Run()
}
