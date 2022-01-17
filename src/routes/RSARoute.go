package routes

import (
	"net/http"
	"strconv"

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
	ID      uint   `json:"id"`
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
	ID       uint   `json:"id"`
	Exponent uint64 `json:"exp"`
	Modulus  uint64 `json:"mod"`
	Text     string `json:"text"`
	Result   string `json:"result"`
	Encrypt  bool   `json:"encrypt"`
}

func GetRSAEntries(c *gin.Context) {
	entries := make([]db.RSAEntry, 10)
	user := db.GetUser(c)
	entryResult := db.GlobalDB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&entries).Limit(10)
	responseEntries := make([]RSAResponse, 0)
	for _, entry := range entries {
		responseEntries = append(responseEntries, RSAResponse{
			ID:      entry.ID,
			PrimeP:  entry.PrimeP,
			PrimeQ:  entry.PrimeQ,
			Private: entry.Private,
			Public:  entry.Public,
			Modulus: entry.Modulus,
		})
	}
	if entryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, responseEntries)
	}

}

func GetRSAEncryptions(c *gin.Context) {
	entries := make([]db.RSAEncryption, 10)
	user := db.GetUser(c)
	entryResult := db.GlobalDB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&entries).Limit(10)
	responseEntries := make([]RSAEncResponse, 0)
	for _, entry := range entries {
		responseEntries = append(responseEntries, RSAEncResponse{
			ID:       entry.ID,
			Exponent: entry.Exponent,
			Modulus:  entry.Modulus,
			Text:     entry.Text,
			Result:   entry.Result,
		})
	}
	if entryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, responseEntries)
	}

}

func DeleteRSAEntry(c *gin.Context) {
	var entry db.RSAEntry
	user := db.GetUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong type of id supplied",
		})
		c.Abort()
	}
	result := db.GlobalDB.First(&entry, id)

	db.HandleDBErrors(c, result, "RSA Entry not found")
	db.HandleUserNotAllowed(c, user, entry.UserID)
	result = db.GlobalDB.Delete(&entry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.String(http.StatusNoContent, "deleted")
	}
}

func DeleteRSAEncryption(c *gin.Context) {
	var entry db.RSAEncryption
	user := db.GetUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong type of id supplied",
		})
		c.Abort()
	}
	result := db.GlobalDB.First(&entry, id)

	db.HandleDBErrors(c, result, "RSA Entry not found")
	db.HandleUserNotAllowed(c, user, entry.UserID)
	result = db.GlobalDB.Delete(&entry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.String(http.StatusNoContent, "deleted")
	}
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
		ID:       rsaEncryption.ID,
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
		ID:      rsaOperation.ID,
		PrimeP:  req.PrimeP,
		PrimeQ:  req.PrimeQ,
		Private: e,
		Public:  d,
		Modulus: modulus,
		Message: message,
	}
	c.JSON(http.StatusOK, res)
}
