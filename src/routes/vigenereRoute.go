package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/db"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type vigenereRequest struct {
	Input       string `json:"input"`
	ShiftString string `json:"shiftString"`
	Encrypt     bool   `json:"encrypt"`
}

type vigenereResponse struct {
	Shifted     string `json:"shifted"`
	ShiftString string `json:"shiftString"`
	Original    string `json:"original"`
}

func GetVigenereString(c *gin.Context) {
	user := db.GetUser(c)
	var req vigenereRequest
	var strType rotationutils.StringType
	if err := c.BindJSON(&req); err != nil {
		return
	}
	if req.Encrypt {
		strType = rotationutils.PlainText
	} else {
		strType = rotationutils.CipherText
	}
	shifted := rotationutils.DoVigenere(req.Input, req.ShiftString, strType)
	vigenereOperation := db.VigenereEntry{
		StrToEncrypt: req.Input,
		Key:          req.ShiftString,
		StrEncrypted: shifted,
		UserID:       user.ID,
		User:         user,
	}

	dbResult := db.GlobalDB.Create(&vigenereOperation)
	if dbResult.Error != nil {
		c.JSON(500, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}

	res := vigenereResponse{
		Shifted:     shifted,
		ShiftString: req.ShiftString,
		Original:    req.Input,
	}
	c.JSON(http.StatusOK, res)
}
