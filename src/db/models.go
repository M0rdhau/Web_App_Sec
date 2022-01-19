package db

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	UserID uint
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"phash"`
}

type CaesarEntry struct {
	Model
	StrToEncrypt string
	Key          int32
	StrEncrypted string
	UserID       uint
	User         User
}

type VigenereEntry struct {
	Model
	StrToEncrypt string
	Key          string
	StrEncrypted string
	UserID       uint
	User         User
}

type DHEntry struct {
	Model
	Prime        uint64
	Primitive    uint64
	UserSecret   uint64
	ServerSecret uint64
	SharedSecret uint64
	UserID       uint
	User         User
}

type RSAEntry struct {
	Model
	PrimeP  uint64
	PrimeQ  uint64
	Private uint64
	Public  uint64
	Modulus uint64
	UserID  uint
	User    User
}

type RSAEncryption struct {
	Model
	Text     string
	Result   string
	Exponent uint64
	Modulus  uint64
	UserID   uint
	User     User
}

func (m *Model) DeleteModel(c *gin.Context, notFoundString string) {
	user := GetUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong type of id supplied",
		})
		c.Abort()
	}
	result := GlobalDB.First(&m, id)

	HandleDBErrors(c, result, notFoundString)
	HandleUserNotAllowed(c, user, m.UserID)
	result = GlobalDB.Delete(&m)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.String(http.StatusNoContent, "deleted")
	}
}
