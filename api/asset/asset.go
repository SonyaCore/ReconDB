package asset

import (
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAsset(c *gin.Context) {
	var Asset models.Assets

	c.ShouldBindJSON(&Asset)

	c.JSON(http.StatusOK, gin.H{
		"message": "asset added",
		"result":  Asset,
		"status":  http.StatusOK,
	})
}
