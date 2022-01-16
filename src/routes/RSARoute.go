package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
	"github.com/m0rdhau/Web_App_Sec/src/db"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type RSARequest struct {
	PrimeP uint64 `json:"primep"`
	PrimeQ uint64 `json:"primeq"`
}

type RSAResponse struct {
	PrimeP  uint64 `json:"primep"`
	PrimeQ  uint64 `json:"primeq"`
	Private uint64 `json:"private"`
	Public  uint64 `json:"public"`
	Modulus uint64 `json:"modulus"`
	Message string `json:"message"`
}

type RSAEncRequest struct {
	Exponent uint64 `json:"exp"`
	Modulus  uint64 `json:"mod"`
	Text     string `json:"text"`
	Encrypt  bool   `json:"encrypt"`
}

type RSAEncResponse struct {
	Exponent uint64 `json:"exp"`
	Modulus  uint64 `json:"mod"`
	Text     string `json:"text"`
	Result   string `json:"result"`
	Encrypt  bool   `json:"encrypt"`
}

func UseRSA(c *gin.Context) {
	user := db.GetUser(c)
	var req RSAEncRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	stringType := rotationutils.PlainText
	if !req.Encrypt {
		stringType = rotationutils.CipherText
	}
	encrypted, err := rotationutils.DoRSA(req.Text, req.Modulus, req.Exponent, stringType)
	if err != nil {
		return
	}
	rsaEncryption := db.RSAEncryption{
		Text:     req.Text,
		Result:   encrypted,
		Exponent: req.Exponent,
		Modulus:  req.Modulus,
		UserID:   user.ID,
		User:     user,
	}
	dbResult := db.GlobalDB.Create(&rsaEncryption)
	if dbResult.Error != nil {
		c.JSON(500, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}
	res := RSAEncResponse{
		Text:     req.Text,
		Result:   encrypted,
		Exponent: req.Exponent,
		Modulus:  req.Modulus,
		Encrypt:  req.Encrypt,
	}
	c.JSON(http.StatusOK, res)
}

func GetGeneratedRSA(c *gin.Context) {
	user := db.GetUser(c)
	var req RSARequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	modulus, e, d, message := cryptoutils.GenerateRSA(&req.PrimeP, &req.PrimeQ)
	rsaOperation := db.RSAEntry{
		PrimeP:  req.PrimeP,
		PrimeQ:  req.PrimeQ,
		Private: e,
		Public:  d,
		Modulus: modulus,
		UserID:  user.ID,
		User:    user,
	}
	dbResult := db.GlobalDB.Create(&rsaOperation)
	if dbResult.Error != nil {
		c.JSON(500, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}
	res := RSAResponse{
		PrimeP:  req.PrimeP,
		PrimeQ:  req.PrimeQ,
		Private: e,
		Public:  d,
		Modulus: modulus,
		Message: message,
	}
	c.JSON(http.StatusOK, res)
}
