package controllers

import (
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
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
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

	if err := db.Model(&book).UpdateColumns(book).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on save data" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, book)
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDBInstance()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be an integer",
		})
		return
	}

	if err := db.Delete(&models.Book{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
