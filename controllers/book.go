package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/blmarquess/go-books-api/database"
	"github.com/blmarquess/go-books-api/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDBInstance()
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	db := database.GetDBInstance()
	var book models.Book

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be an integer",
		})
		return
	}

	if err := db.First(&book, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDBInstance()
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON" + err.Error(),
		})
		return
	}

	if err := book.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := db.Create(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on save data" + err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDBInstance()
	var book models.Book
	log.Println("UpdateBook")
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be an integer",
		})
		return
	}

	if err := db.First(&book, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		log.Println(err)
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON" + err.Error(),
		})
		return
	}

	if err := book.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := db.Save(&book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on save data" + err.Error(),
		})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusAccepted, book)
}
