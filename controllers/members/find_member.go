package members_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func FindMember(c *gin.Context) {
	id := c.Param("id")

	var member models.MemberWithBooks
	result := services.Db.Table("members").Model(&member).Where("id = ?", id).Preload("Books").Find(&member)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, member)
}
