package middlewares

import (
	"ReconDB/database"
	"ReconDB/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"regexp"
)

var scopeUri = "/api/scope"

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

	for i, _ := range Scopes {
		if Scope.ScopeType == Scopes[i] {
			c.Next()
			return
		}
		continue
	}
	c.JSON(http.StatusFailedDependency, gin.H{
		"error":       "scope type is not valid",
		"valid_types": Scopes,
		"status":      http.StatusFailedDependency,
	})
	c.Abort()
	return
}

func OutScopeCheck(c *gin.Context) {
	var Scope models.Scopes
	var ctx = context.TODO()

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

	ScopeQuery := bson.M{
		"companyname": Scope.CompanyName,
		"scopetype":   Scope.ScopeType,
		"scope": primitive.Regex{
			Pattern: "^" + regexp.QuoteMeta(Scope.Scope) + "$",
			Options: "i",
		},
	}

	var collection *mongo.Collection
	var results int64

	// only use this section if request uri was /api/scope
	if c.Request.RequestURI == scopeUri {
		collection = database.Collection("Scopes")
		results, err = collection.CountDocuments(ctx, ScopeQuery)

		if results >= 1 {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"result": "duplicate entry",
				"status": http.StatusNotAcceptable,
			})
			c.Abort()
			return
		}

	}

	collection = database.Collection("OutofScopes")
	results, err = collection.CountDocuments(ctx, ScopeQuery)
	if err != nil {
		log.Print(err.Error())
	}

	fmt.Println("document count", results)
	if results >= 1 {
		if c.Request.RequestURI == "/api/outscope" {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"companyname": Scope.CompanyName,
				"result":      "duplicate entry",
				"status":      http.StatusNotAcceptable,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusNotAcceptable, gin.H{
			"scope":  Scope.Scope,
			"result": "out of scope",
			"status": http.StatusNotAcceptable,
		})
		c.Abort()
		return
	}

	c.Next()
}
