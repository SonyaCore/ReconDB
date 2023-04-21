package outscope

import (
	"ReconDB/database"
	"ReconDB/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddOutScope(c *gin.Context) {
	var Scope models.Scopes

	c.ShouldBindJSON(&Scope)

	// insert outofscope to db
	collection := database.Collection("OutofScopes")
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
		"message":      "out of scope added",
		"result":       Scope,
		"status":       http.StatusOK,
	})
}
