package outscope

import (
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddOutScope(c *gin.Context) {
	var Scope models.Scopes

	c.ShouldBindJSON(&Scope)

	c.JSON(http.StatusOK, gin.H{
		"message": "out of scope added",
		"result":  Scope,
		"status":  http.StatusOK,
	})
}
