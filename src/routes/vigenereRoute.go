package routes

import (
	"net/http"
	"strconv"

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
	ID          uint   `json:"id"`
	Shifted     string `json:"shifted"`
	ShiftString string `json:"shiftString"`
	Original    string `json:"original"`
}

func GetVigenereEntries(c *gin.Context) {
	entries := make([]db.VigenereEntry, 10)
	user := db.GetUser(c)
	entryResult := db.GlobalDB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&entries).Limit(10)
	responseEntries := make([]vigenereResponse, 0)
	for _, entry := range entries {
		responseEntries = append(responseEntries, vigenereResponse{
			ID:          entry.ID,
			Shifted:     entry.StrEncrypted,
			ShiftString: entry.Key,
			Original:    entry.StrToEncrypt,
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

func DeleteVigenereEntry(c *gin.Context) {
	var entry db.VigenereEntry
	user := db.GetUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong type of id supplied",
		})
		c.Abort()
	}
	result := db.GlobalDB.First(&entry, id)

	db.HandleDBErrors(c, result, "Vigenere entry not found")
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
		ID:          vigenereOperation.ID,
		Shifted:     shifted,
		ShiftString: req.ShiftString,
		Original:    req.Input,
	}
	c.JSON(http.StatusOK, res)
}
