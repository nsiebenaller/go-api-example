package members_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/models"
	"github.com/nsiebenaller/go-api-example/services"
)

func DeleteMember(c *gin.Context) {
	id := c.Param("id")

	tx := services.Db.Begin()

	// Delete book associations owned by this member
	var memberBooks models.MemberBooks
	r1 := tx.Where("member_id = ?", id).Delete(&memberBooks)
	if r1.Error != nil {
		tx.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}

	// Delete member
	var member models.Member
	r2 := tx.Where("id = ?", id).Delete(&member)
	if r2.Error != nil {
		tx.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}
	if r2.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	r3 := tx.Commit()
	if r3.Error != nil {
		tx.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
