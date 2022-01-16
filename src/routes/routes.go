package routes

import (
	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/caesar", GetCaesarString)
		api.GET("/vigenere", GetCaesarString)
		api.POST("/login", Login)
		api.POST("/signup", Signup)
	}
	router.Run("localhost:8080")
}
