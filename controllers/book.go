package controllers

import (
	"log"
	"net/http"

	"github.com/blmarquess/go-books-api/database"
	"github.com/blmarquess/go-books-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDBInstance()
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}
