package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type caesarRequest struct {
	Input   string `json:"input"`
	Shift   int32  `json:"shift"`
	Encrypt bool   `json:"encrypt"`
}

type caesarResponse struct {
	Shifted string `json:"shifted"`
	Shift   int32  `json:"shift"`
}

func GetCaesarString(c *gin.Context) {
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
	res := caesarResponse{
		Shifted: shifted,
		Shift:   shift,
	}
	c.JSON(http.StatusOK, res)
}
