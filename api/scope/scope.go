package scope

import (
	"ReconDB/database"
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddScope(c *gin.Context) {
	var Scope models.Scopes

	c.ShouldBindJSON(&Scope)

	// insert scope to db
	collection := database.Collection("Scopes")
	result, err := collection.InsertOne(database.Ctx, Scope)
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
		"message":      "scope added",
		"result":       Scope,
		"status":       http.StatusOK,
	})
}
