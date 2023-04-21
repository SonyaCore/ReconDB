package company

import (
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCompany(c *gin.Context) {
	var Company models.Company

	c.ShouldBindJSON(&Company)

	c.JSON(http.StatusOK, gin.H{
		"message": "company added",
		"result":  Company,
		"status":  http.StatusOK,
	})
}
