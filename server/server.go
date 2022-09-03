package server

import (
	"log"

	routes "github.com/blmarquess/go-books-api/server/routers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Println("Server running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
