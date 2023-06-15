package asset

import (
	"ReconDB/api/asset"
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/pkg/host"
	"ReconDB/pkg/typeassert"
	"ReconDB/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net"
	"net/http"
	"strings"
)

// DuplicateValidate checks if an asset with the same asset name and scope already exists in the Assets collection.
// it defines a query to find the asset in the collection.
// If a document with the same asset name and scope exists, it aborts the request and returns a JSON
// response with a Bad Request status code and an error message indicating a duplicate entry. Otherwise, it passes the
// request to the next middleware in the chain using c.Next().
func DuplicateValidate(c *gin.Context) {
	var Asset models.Assets

	rawBody, err := utils.ReadBuffer(c)
	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Asset)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.Contains(Asset.Scope, "*") {
		Asset.Scope = utils.WildCardToRegex(Asset.Scope)
	}

	// Define the query to find scopes for the given company name and asset
	assetQuery := bson.M{
		"asset": Asset.Asset,
		"scope": primitive.Regex{
			Pattern: "^" + Asset.Scope + "$",
			Options: "i",
		}}

	// Check if the asset is already in the Asset collection
	count, err := database.CountDocuments("Assets", assetQuery)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to count documents in Assets collection",
			"status": http.StatusInternalServerError,
		})
		return
	}
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "Duplicate entry",
			"input":  Asset.Asset,
			"status": http.StatusBadRequest,
		})
		return
	}
	c.Next()
}

// OutScopeAssetValidate checks for Scopes & OutofScope duplication
func OutScopeAssetValidate(c *gin.Context) {
	var Asset models.Assets
	var ScopeCount int64
	var outOfScope int64
	var outOfScopeWildCard int64
	var outOfScopeCIDR int64

	rawBody, err := utils.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Asset)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Find Asset Type based on input
	Asset.AssetType, err = typeassert.FindAssetType(Asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		return
	}

	// Search for Company with Scope name if none found return error
	scopeResult, err := asset.FindCompanyName(Asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "Company Name not found in Scope",
			"status": http.StatusBadRequest,
		})
	}
	// Assert Scope result to Asset.CompanyName
	Asset.CompanyName = scopeResult

	if strings.Contains(Asset.Scope, "*") {
		if !host.MatchWildcard(Asset.Scope, Asset.Asset) {
			utils.ReturnError(c, errors.New("asset is not valid in that scope"),
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}
	}

	// Check if the Asset are in the out-of-scope of Scopes collection.
	ScopeQuery := bson.M{
		"companyname": Asset.CompanyName,
		"scope":       Asset.Scope,
	}

	if outOfScope, err = database.CountDocuments("OutofScopes", ScopeQuery); err != nil {
		outOfScope = 0
	}

	if outOfScope > 0 {
		utils.ReturnError(c, errors.New("asset is out of Scope"),
			http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
		return
	}

	if ScopeCount, err = database.CountDocuments("Scopes", ScopeQuery); err != nil {
		ScopeCount = 0
	}

	if ScopeCount < 1 {
		utils.ReturnError(c, errors.New("asset not found in Scope collection"),
			http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
		return
	}

	// Check Asset if Asset.Scope contain wildcard
	if strings.Contains(Asset.Scope, "*") && outOfScope == 0 {
		var Scope = utils.WildCardToRegex(Asset.Scope)

		outOfScopeWildCardQuery := bson.M{
			"companyname": Asset.CompanyName,
			"scopetype":   "wildcard",
			"scope": primitive.Regex{
				Pattern: "^" + Scope + "$",
				Options: "i",
			}}

		if outOfScopeWildCard, err = database.CountDocuments("OutofScopes", outOfScopeWildCardQuery); err != nil {
			utils.ReturnError(c, err,
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}
		if outOfScopeWildCard > 0 {
			utils.ReturnError(c, errors.New("asset is out of Scope"),
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}
	}

	// Check Asset if Asset.Scope contain wildcard
	if Asset.AssetType == "ip" && outOfScope == 0 {
		// Check if cidr exist in OutofScopes Collection
		outOfScopeCIDRQuery := bson.M{
			"companyname": Asset.CompanyName,
			"scope":       Asset.Scope}

		if outOfScopeCIDR, err = database.CountDocuments("OutofScopes", outOfScopeCIDRQuery); err != nil {
			utils.ReturnError(c, err,
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}
		if outOfScopeCIDR > 0 {
			utils.ReturnError(c, errors.New("asset is out of Scope"),
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}

		ip := net.ParseIP(Asset.Asset)
		_, subnet, _ := net.ParseCIDR(Asset.Scope)
		switch subnet.Contains(ip) {
		case true:
			c.Next()
			break
		case false:
			utils.ReturnError(c, errors.New("IP is not within CIDR range"),
				http.StatusNotAcceptable, Asset.Asset, Asset.Scope)
			return
		}
	}

	c.Next()
	return
}
