package routes

import (
	"net/http"

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
	Prime        uint64 `json:"prime"`
	Primitive    uint64 `json:"primitive"`
	UserSecret   uint64 `json:"userSecret"`
	ServerSecret uint64 `json:"serverSecret"`
	SharedSecret uint64 `json:"sharedSecret"`
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
			Prime:        req.Prime,
			Primitive:    req.Primitive,
			UserSecret:   req.UserSecret,
			ServerSecret: serverSecret,
			SharedSecret: sharedSecret,
		}
		c.JSON(http.StatusOK, res)
	}

}
