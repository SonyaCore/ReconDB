package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func AddAsset(c *gin.Context) {
	var asset models.Assets
	var err error

	// Bind the JSON data to the Asset struct
	if err = c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request",
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}

	// Find the asset type of the given asset
	asset.AssetType, err = FindAssetType(asset)
	if err != nil {
		c.JSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		c.Abort()
		return
	}

	// Find the matching scope for the given asset
	collectionScope := database.Collection("Scopes")

	// find company name based on scope
	scopeQuery := bson.M{
		//"scopetype": asset.AssetType,
		"scope": asset.Scope,
	}

	opts := options.FindOne().SetProjection(bson.M{"companyname": 1})

	var scopeResult struct {
		CompanyName string `bson:"companyname"`
	}

	if err = collectionScope.FindOne(context.Background(), scopeQuery, opts).Decode(&scopeResult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "company name not found in scope",
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}

	// Check if the asset is in the out-of-scope collection
	outOfScopeQuery := bson.M{
		"companyname": scopeResult.CompanyName,
		"scopetype":   asset.AssetType,
		"scope":       asset.Asset,
	}

	var outOfScope int64
	if outOfScope, err = database.CountDocuments("OutofScopes", outOfScopeQuery); err != nil {
		outOfScope = 0
	}
	if outOfScope > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":  "asset is out of scope",
			"asset":  asset.Asset,
			"status": http.StatusNotAcceptable,
		})
		c.Abort()
		return
	}

	// Insert the asset to the Asset collection
	collection := database.Collection("Assets")
	result, err := collection.InsertOne(database.Ctx, asset)
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

	c.JSON(http.StatusOK, gin.H{
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
	var Param = c.Param("asset")
	var ctx = context.TODO()

	collection := database.Collection("Assets")
	filter, err := collection.DeleteMany(ctx, bson.M{"asset": Param})
	if err != nil {
		log.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"AssetName":    Param,
		"DeletedCount": filter.DeletedCount,
		"status":       http.StatusOK,
	})
}
