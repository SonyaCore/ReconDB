package outscope

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

func GetAllOutofScopes(c *gin.Context) {
	var ctx = context.TODO()
	var OutofScopes []bson.M

	collection := database.Collection("OutofScopes")
	results, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &OutofScopes); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": OutofScopes,
		"status":  http.StatusOK,
	})
}

func GetOutofScopes(c *gin.Context) {
	var Param = c.Param("scope")
	var ctx = context.TODO()
	var OutofScopes []bson.M

	collection := database.Collection("OutofScopes")
	results, err := collection.Find(ctx, bson.M{"scope": Param})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &OutofScopes); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"results": OutofScopes,
		"status":  http.StatusOK,
	})
}
