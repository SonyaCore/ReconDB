package scope

import (
	"ReconDB/database"
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"regexp"
)

var scopeUri = "/api/scope"
var outScopeUri = "/api/outscope"

func ValidateScopes(c *gin.Context) {
	var Scope models.Scopes

	// Read the content
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, rawBody)
	}

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	for i, _ := range middlewares.Scopes {
		if Scope.ScopeType == middlewares.Scopes[i] {
			c.Next()
			return
		}
		continue
	}
	c.JSON(http.StatusFailedDependency, gin.H{
		"error":       "Scope Type is not valid",
		"valid_types": middlewares.Scopes,
		"status":      http.StatusFailedDependency,
	})
	c.Abort()
	return
}

func OutScopeCheck(c *gin.Context) {
	var Scope models.Scopes
	var results int64

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	ScopeQuery := bson.M{
		"companyname": Scope.CompanyName,
		"scopetype":   Scope.ScopeType,
		"scope": primitive.Regex{
			Pattern: "^" + regexp.QuoteMeta(Scope.Scope) + "$",
			Options: "i",
		},
	}

	// only use this section if request uri was /api/scope
	if c.Request.RequestURI == scopeUri {
		companyQuery := bson.M{
			"companyname": Scope.CompanyName,
		}

		results, err = database.CountDocuments("Company", companyQuery)
		if results == 0 {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"result": "Scope is not registered for this company.",
				"status": http.StatusNotAcceptable,
			})
			return
		}

		results, err = database.CountDocuments("Scopes", ScopeQuery)
		if results >= 1 {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"result": "Duplicate Entry",
				"status": http.StatusNotAcceptable,
			})
			return
		}
	}

	results, err = database.CountDocuments("OutofScopes", ScopeQuery)
	if err != nil {
		log.Print(err.Error())
	}

	if results >= 1 {
		if c.Request.RequestURI == outScopeUri {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"companyname": Scope.CompanyName,
				"result":      "Duplicate Entry",
				"status":      http.StatusNotAcceptable,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"scope":  Scope.Scope,
			"result": "Out of Scope",
			"status": http.StatusNotAcceptable,
		})
		return
	}

	c.Next()
}
