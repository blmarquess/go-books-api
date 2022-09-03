package server

import "github.com/gin-gonic/gin"

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port:   "8000",
		server: gin.Default(),
	}
}
