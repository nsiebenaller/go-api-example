package books_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func CreateBook(c *gin.Context) {
	var createBook models.CreateBook
	err := c.ShouldBindJSON(&createBook)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := services.Db.Table("books").Create(&createBook)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
