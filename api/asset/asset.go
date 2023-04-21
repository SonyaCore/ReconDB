package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

func GetAssets(c *gin.Context) {
	var ctx = context.TODO()
	var Assets []bson.M

	collection := database.Collection("Assets")
	results, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &Assets); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"assets": Assets,
		"status": http.StatusOK,
	})
}
