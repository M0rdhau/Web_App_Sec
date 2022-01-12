package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type vigenereRequest struct {
	Input       string `json:"input"`
	ShiftString string `json:"shiftString"`
	Encrypt     bool   `json:"encrypt"`
}

type vigenereResponse struct {
	Shifted string `json:"shifted"`
}

func GetVigenereString(c *gin.Context) {
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
	res := vigenereResponse{
		Shifted: shifted,
	}
	c.JSON(http.StatusOK, res)
}
