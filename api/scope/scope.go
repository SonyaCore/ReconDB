package scope

import (
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddScope(c *gin.Context) {
	var Scope models.Scopes

	c.ShouldBindJSON(&Scope)

	c.JSON(http.StatusOK, gin.H{
		"message": "scope added",
		"result":  Scope,
		"status":  http.StatusOK,
	})
}
