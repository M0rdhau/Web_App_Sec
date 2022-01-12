package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Route() {
	router := gin.Default()
	router.GET("/", GetHello)
	router.GET("/caesar", GetCaesarString)
	router.Run("localhost:8080")
}
