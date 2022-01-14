package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
)

type diffieHellmanRequest struct {
	Prime      uint64 `json:"prime"`
	Primitive  uint64 `json:"primitive"`
	UserSecret uint64 `json:"userSecret"`
}

type diffieHellmanResponse struct {
	ServerSecret uint64 `json:"serverSecret"`
	SharedSecret uint64 `json:"sharedSecret"`
}

func GetGeneratedDiffieHellman(c *gin.Context) {
	var req diffieHellmanRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}
	serverSecret, sharedSecret, err := cryptoutils.GenerateDH(req.Prime, req.Primitive, req.UserSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		res := diffieHellmanResponse{
			ServerSecret: serverSecret,
			SharedSecret: sharedSecret,
		}
		c.JSON(http.StatusOK, res)
	}

}
