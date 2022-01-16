package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/middlewares"
)

func Route() {
	router := gin.Default()
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", Login)
			public.POST("/signup", Signup)
		}

		protected := api.Group("/protected").Use(middlewares.Authware())
		{
			protected.POST("/caesar", GetCaesarString)
			protected.POST("/vigenere", GetCaesarString)
			protected.POST("/diffiehellman", GetGeneratedDiffieHellman)
			protected.POST("/rsa", GetGeneratedRSA)
			protected.POST("/rsa/encrypt", UseRSA)
		}
	}
	router.Run("localhost:8080")
}
