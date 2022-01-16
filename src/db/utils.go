package db

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) (user User) {

	username, _ := c.Get("username")

	userResult := GlobalDB.Where("username = ?", username.(string)).First(&user)

	if userResult.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if userResult.Error != nil {
		c.JSON(500, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}

	return user
}
