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
	ID       uint   `json:"id"`
	Shifted  string `json:"shifted"`
	Shift    int32  `json:"shift"`
	Original string `json:"original"`
}

func GetCaesarEntries(c *gin.Context) {
	entries := make([]db.CaesarEntry, 10)
	user := db.GetUser(c)
	entryResult := db.GlobalDB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&entries).Limit(10)
	if entryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		responseEntries := make([]caesarResponse, 0)
		for _, entry := range entries {
			responseEntries = append(responseEntries, caesarResponse{
				ID:       entry.ID,
				Shifted:  entry.StrEncrypted,
				Shift:    entry.Key,
				Original: entry.StrToEncrypt,
			})
		}
		c.JSON(http.StatusOK, responseEntries)
	}

}

func DeleteCaesarEntry(c *gin.Context) {
	var entry db.CaesarEntry
	entry.DeleteModel(c, "Caesar string not found")
}

func GetCaesarString(c *gin.Context) {
	user := db.GetUser(c)

	var req caesarRequest
	var strType rotationutils.StringType
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "JSON fields unsupported",
		})
		c.Abort()
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
		return
	}

	res := caesarResponse{
		ID:       caesarOperation.ID,
		Shifted:  shifted,
		Shift:    shift,
		Original: req.Input,
	}
	c.JSON(http.StatusOK, res)
}
