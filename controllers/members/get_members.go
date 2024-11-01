package members_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func GetMembers(c *gin.Context) {
	var members []models.Member
	services.Db.Find(&members)
	c.JSON(http.StatusOK, members)
}
