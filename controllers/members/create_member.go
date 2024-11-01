package members_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func CreateMember(c *gin.Context) {
	var createMember models.CreateMember
	err := c.ShouldBindJSON(&createMember)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := services.Db.Table("members").Create(&createMember)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
