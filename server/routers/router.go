package routes

import (
	"github.com/blmarquess/go-books-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetAllBooks)
			books.GET("/:id")
			books.POST("/")
			books.PUT("/")
			books.DELETE("/")
		}
	}
	return router
}
