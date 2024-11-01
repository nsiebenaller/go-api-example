package books_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	services.Db.Find(&books)
	c.JSON(http.StatusOK, books)
}
