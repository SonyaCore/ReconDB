package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/pkg/type"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func AddAsset(c *gin.Context) {
	var asset models.Assets
	var err error

	// Bind the JSON data to the Asset struct
	if err = c.ShouldBindJSON(&asset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid Request",
			"status": http.StatusBadRequest,
		})
		return
	}

	// Find the asset type of the given asset
	asset.AssetType, err = _type.FindAssetType(asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		return
	}

	// Insert the asset to the Asset collection
	collection := database.Collection("Assets")
	result, err := collection.InsertOne(database.Ctx, asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collectionID": result.InsertedID,
		"message":      "Asset Added",
		"result":       asset,
		"status":       http.StatusOK,
	})
}

func GetAllAssets(c *gin.Context) {
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

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"assets": Assets,
		"status": http.StatusOK,
	})
}

func GetAsset(c *gin.Context) {
	var Param = c.Param("asset")
	var ctx = context.TODO()
	var Assets []bson.M

	collection := database.Collection("Assets")
	filter, err := collection.Find(ctx, bson.M{"asset": Param})
	if err != nil {
		log.Print(err.Error())
	}

	if err = filter.All(ctx, &Assets); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"assets": Assets,
		"status": http.StatusOK,
	})
}

func DeleteAsset(c *gin.Context) {
	var Param = c.Param("scope")
	var ctx = context.TODO()

	collection := database.Collection("Assets")
	filter, err := collection.DeleteMany(ctx, bson.M{"scope": Param})
	if err != nil {
		log.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"asset_name":    Param,
		"deleted_count": filter.DeletedCount,
		"status":        http.StatusOK,
	})
}
