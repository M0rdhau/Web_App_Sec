package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/middlewares"
)

func Route() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Origin", "content-type"}
	router.Use(cors.New(config))
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	api := router.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", Login)
			public.POST("/signup", Signup)
		}

		protected := api.Group("/protected").Use(middlewares.Authware())
		{
			protected.GET("/caesar", GetCaesarEntries)
			protected.GET("/vigenere", GetVigenereEntries)
			protected.GET("/diffiehellman", GetDHEntries)
			protected.GET("/rsa/generate", GetRSAEntries)
			protected.GET("/rsa/encrypt", GetRSAEncryptions)
			protected.POST("/caesar", GetCaesarString)
			protected.POST("/vigenere", GetVigenereString)
			protected.POST("/diffiehellman", GetGeneratedDiffieHellman)
			protected.POST("/rsa/generate", GetGeneratedRSA)
			protected.POST("/rsa/encrypt", UseRSA)
			protected.DELETE("/caesar/:id", DeleteCaesarEntry)
			protected.DELETE("/vigenere/:id", DeleteVigenereEntry)
			protected.DELETE("/diffiehellman/:id", DeleteDHEntry)
			protected.DELETE("/rsa/generate/:id", DeleteRSAEntry)
			protected.DELETE("/rsa/encrypt/:id", DeleteRSAEncryption)
		}
	}
	router.Run("0.0.0.0:8000")
}
