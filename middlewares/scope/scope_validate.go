package scope

import (
	"ReconDB/database"
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"regexp"
)

const scopeUri = "/api/scope"
const outScopeUri = "/api/outscope"

// errors
var CompanyNotRegister = "Scope are not registered for this company"
var DuplicateEntry = "Duplicate Entry"

func ValidateScopeType(c *gin.Context) {
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
		results, err = CompanyCheck(companyQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.CompanyName,
				"result": err.Error(),
				"status": http.StatusNotAcceptable,
			})
		}
		results, err = DuplicateCheck(ScopeQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"result": err.Error(),
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
		if c.Request.RequestURI == scopeUri {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"companyname": Scope.CompanyName,
				"result":      "Out of Scope",
				"status":      http.StatusNotAcceptable,
			})
			return
		}

		// outscope uri
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"scope":  Scope.Scope,
			"result": DuplicateEntry,
			"status": http.StatusNotAcceptable,
		})
		return
	}

	c.Next()
}

func CompanyCheck(query bson.M) (int64, error) {
	count, err := database.CountDocuments("Company", query)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, fmt.Errorf(CompanyNotRegister)
	}
	return 1, nil
}

func DuplicateCheck(query bson.M) (int64, error) {
	count, err := database.CountDocuments("Scopes", query)
	if err != nil {
		return 1, err
	}
	if count >= 1 {
		return count, fmt.Errorf("duplicate Entry")
	}

	return 0, nil
}
