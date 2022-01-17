package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
	"github.com/m0rdhau/Web_App_Sec/src/db"
)

type diffieHellmanRequest struct {
	Prime      uint64 `json:"prime"`
	Primitive  uint64 `json:"primitive"`
	UserSecret uint64 `json:"userSecret"`
}

type diffieHellmanResponse struct {
	ID           uint   `json:"id"`
	Prime        uint64 `json:"prime"`
	Primitive    uint64 `json:"primitive"`
	UserSecret   uint64 `json:"userSecret"`
	ServerSecret uint64 `json:"serverSecret"`
	SharedSecret uint64 `json:"sharedSecret"`
}

func GetDHEntries(c *gin.Context) {
	entries := make([]db.DHEntry, 10)
	user := db.GetUser(c)
	entryResult := db.GlobalDB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&entries).Limit(10)
	responseEntries := make([]diffieHellmanResponse, 0)
	for _, entry := range entries {
		responseEntries = append(responseEntries, diffieHellmanResponse{
			ID:           entry.ID,
			Prime:        entry.Prime,
			Primitive:    entry.Primitive,
			UserSecret:   entry.UserSecret,
			ServerSecret: entry.ServerSecret,
			SharedSecret: entry.SharedSecret,
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

func DeleteDHEntry(c *gin.Context) {
	var entry db.DHEntry
	user := db.GetUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Wrong type of id supplied",
		})
		c.Abort()
	}
	result := db.GlobalDB.First(&entry, id)

	db.HandleDBErrors(c, result, "Diffie-Hellman entry not found")
	db.HandleUserNotAllowed(c, user, entry.UserID)
	result = db.GlobalDB.Delete(&entry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Failure",
		})
		c.Abort()
	} else {
		c.String(http.StatusNoContent, "Deleted")
	}
}

func GetGeneratedDiffieHellman(c *gin.Context) {
	user := db.GetUser(c)
	var req diffieHellmanRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	serverSecret, sharedSecret, err := cryptoutils.GenerateDH(&req.Prime, &req.Primitive, req.UserSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		dhOperation := db.DHEntry{
			Prime:        req.Prime,
			Primitive:    req.Primitive,
			UserSecret:   req.UserSecret,
			ServerSecret: serverSecret,
			SharedSecret: sharedSecret,
			UserID:       user.ID,
			User:         user,
		}
		dbResult := db.GlobalDB.Create(&dhOperation)
		if dbResult.Error != nil {
			c.JSON(500, gin.H{
				"msg": "Database Failure",
			})
			c.Abort()
			return
		}
		res := diffieHellmanResponse{
			ID:           dhOperation.ID,
			Prime:        req.Prime,
			Primitive:    req.Primitive,
			UserSecret:   req.UserSecret,
			ServerSecret: serverSecret,
			SharedSecret: sharedSecret,
		}
		c.JSON(http.StatusOK, res)
	}

}
