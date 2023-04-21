package scope

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

func GetAllScopes(c *gin.Context) {
	var ctx = context.TODO()
	var Scopes []bson.M

	collection := database.Collection("Scopes")
	results, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &Scopes); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": Scopes,
		"status":  http.StatusOK,
	})
}

func GetScopes(c *gin.Context) {
	var Param = c.Param("scope")
	var ctx = context.TODO()
	var Scopes []bson.M

	collection := database.Collection("Scopes")
	results, err := collection.Find(ctx, bson.M{"scope": Param})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &Scopes); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": Scopes,
		"status":  http.StatusOK,
	})
}
