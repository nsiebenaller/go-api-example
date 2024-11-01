package members_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func CheckoutBook(c *gin.Context) {
	var checkoutBook models.MemberBooks
	err := c.ShouldBindJSON(&checkoutBook)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := services.Db.Table("member_books").Create(&checkoutBook)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
