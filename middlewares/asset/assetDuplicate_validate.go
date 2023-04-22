package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"regexp"
)

func DuplicateValidate(c *gin.Context) {
	var Asset models.Assets

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Asset)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Define the query to find scopes for the given company name and asset type
	assetQuery := bson.M{
		"asset": primitive.Regex{
			Pattern: "^" + regexp.QuoteMeta(Asset.Asset) + "$",
			Options: "i",
		},
		"scope": Asset.Scope,
	}

	// Check if the asset is already in the Asset collection
	count, collectionError := database.CountDocuments("Assets", assetQuery)
	if collectionError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to count documents in Assets collection",
			"status": http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "duplicate entry",
			"input":  Asset.Asset,
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	c.Next()
}
