package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
)

type RSARequest struct {
	PrimeP uint64 `json:"primep"`
	PrimeQ uint64 `json:"primeq"`
}

type RSAResponse struct {
	Private uint64 `json:"private"`
	Public  uint64 `json:"public"`
	Modulus uint64 `json:"modulus"`
	Message string `json:"message"`
}

func GetGeneratedRSA(c *gin.Context) {
	var req RSARequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	modulus, e, d, message := cryptoutils.GenerateRSA(req.PrimeP, req.PrimeQ)
	res := RSAResponse{
		Private: e,
		Public:  d,
		Modulus: modulus,
		Message: message,
	}
	c.JSON(http.StatusOK, res)
}
