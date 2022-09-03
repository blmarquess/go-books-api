package main

import "github.com/blmarquess/go-books-api/server"

func main() {
	server := server.NewServer("8000")
	server.Run()
}
