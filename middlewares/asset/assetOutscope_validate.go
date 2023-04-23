package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"ReconDB/pkg/type"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
)

// OutScopeAssetValidate checks for Scopes & OutofScope duplication
func OutScopeAssetValidate(c *gin.Context) {
	var Asset models.Assets

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Asset)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	Asset.AssetType, err = _type.FindAssetType(Asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		return
	}

	// Find the matching scope for the given Asset
	collectionScope := database.Collection("Scopes")

	// find company name based on scope
	scopeQuery := bson.M{
		//"scopetype": Asset.AssetType,

		"scope": Asset.Scope,
	}

	opts := options.FindOne().SetProjection(bson.M{"companyname": 1})

	var scopeResult struct {
		CompanyName string `bson:"companyname"`
	}

	if err = collectionScope.FindOne(context.Background(), scopeQuery, opts).Decode(&scopeResult); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "Company Name not found in Scope",
			"status": http.StatusBadRequest,
		})
		return
	}

	// Check if the Asset is in the out-of-scope collection for single type
	outOfScopeQuery := bson.M{
		"companyname": scopeResult.CompanyName,
		"scopetype":   Asset.AssetType,
		"scope":       Asset.Asset,
	}

	var outOfScope int64
	if outOfScope, err = database.CountDocuments("OutofScopes", outOfScopeQuery); err != nil {
		outOfScope = 0
	}
	if outOfScope > 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"error":  "Asset is out of Scope",
			"asset":  Asset.Asset,
			"scope":  Asset.Scope,
			"status": http.StatusNotAcceptable,
		})
		return
	}

	if strings.Contains(Asset.Scope, "*") && outOfScope == 0 {
		var outOfScopeWildCard int64
		outOfScopeWildCardQuery := bson.M{
			"companyname": scopeResult.CompanyName,
			"scopetype":   "wildcard",
			"scope":       Asset.Scope,
		}

		if outOfScopeWildCard, err = database.CountDocuments("OutofScopes", outOfScopeWildCardQuery); err != nil {
			fmt.Println(err.Error())
			outOfScopeWildCard = 0
		}
		if outOfScopeWildCard > 0 {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"error":  "Asset is out of Scope",
				"asset":  Asset.Asset,
				"scope":  Asset.Scope,
				"status": http.StatusNotAcceptable,
			})
			return
		}
	}
	c.Next()
	return
}
