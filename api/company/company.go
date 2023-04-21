package company

import (
	"ReconDB/database"
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCompany(c *gin.Context) {
	var Company models.Company

	c.ShouldBindJSON(&Company)

	// insert company to db
	collection := database.Collection("Company")
	result, err := collection.InsertOne(database.Ctx, Company)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collectionID": result.InsertedID,
		"message":      "company added",
		"result":       Company,
		"status":       http.StatusOK,
	})
}
