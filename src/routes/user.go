package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/auth"
	"github.com/m0rdhau/Web_App_Sec/src/db"
	"gorm.io/gorm"
)

// LoginPayload login body
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

// Signup creates a user in db
func Signup(c *gin.Context) {
	var userReq LoginPayload

	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	user := db.User{
		Username:     userReq.Username,
		PasswordHash: "",
	}

	err = user.HashPassword(userReq.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		c.JSON(500, gin.H{
			"msg": "error creating user",
		})
		c.Abort()

		return
	}

	c.JSON(200, user)
}

// Login logs users in
func Login(c *gin.Context) {
	var payload LoginPayload
	var user db.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := db.GlobalDB.Where("username = ?", payload.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Username)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

	return
}
