package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/db"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type caesarRequest struct {
	Input   string `json:"input"`
	Shift   int32  `json:"shift"`
	Encrypt bool   `json:"encrypt"`
}

type caesarResponse struct {
	Shifted  string `json:"shifted"`
	Shift    int32  `json:"shift"`
	Original string `json:"original"`
}

func GetCaesarString(c *gin.Context) {
	user := db.GetUser(c)

	var req caesarRequest
	var strType rotationutils.StringType
	if err := c.BindJSON(&req); err != nil {
		return
	}
	if req.Encrypt {
		strType = rotationutils.PlainText
	} else {
		strType = rotationutils.CipherText
	}
	shifted, shift := rotationutils.DoCaesar(req.Input, req.Shift, strType)

	caesarOperation := db.CaesarEntry{
		StrToEncrypt: req.Input,
		Key:          shift,
		StrEncrypted: shifted,
		UserID:       user.ID,
		User:         user,
	}

	dbResult := db.GlobalDB.Create(&caesarOperation)
	if dbResult.Error != nil {
		c.JSON(500, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}

	res := caesarResponse{
		Shifted:  shifted,
		Shift:    shift,
		Original: req.Input,
	}
	c.JSON(http.StatusOK, res)
}
