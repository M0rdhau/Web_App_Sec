package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleUserNotAllowed(c *gin.Context, user User, resourceUserID uint) {
	if user.ID != resourceUserID {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "Resource does not belong to user",
		})
		c.Abort()
	}
}

func HandleDBErrors(c *gin.Context, result *gorm.DB, notFoundString string) {
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": notFoundString,
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}
}

func GetUser(c *gin.Context) (user User) {

	username, _ := c.Get("username")

	userResult := GlobalDB.Where("username = ?", username.(string)).First(&user)
	HandleDBErrors(c, userResult, "Not found")
	return user
}
