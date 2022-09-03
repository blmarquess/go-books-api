package main

import (
	"github.com/blmarquess/go-books-api/database"
	"github.com/blmarquess/go-books-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer("8000")
	server.Run()
}
