package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAsset(c *gin.Context) {
	var Asset models.Assets

	c.ShouldBindJSON(&Asset)

	// insert asset to db
	collection := database.Collection("Assets")
	result, err := collection.InsertOne(database.Ctx, Asset)
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
		"message":      "asset added",
		"result":       Asset,
		"status":       http.StatusOK,
	})

}
