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
	"io"
	"log"
	"net/http"
)

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

	collection := database.Collection("OutofScopes")
	results, err := collection.CountDocuments(ctx, bson.M{"companyname": Scope.CompanyName})
	if err != nil {
		log.Print(err.Error())
	}
	fmt.Println(results)

	if results >= 1 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"companyname": Scope.CompanyName,
			"result":      "company name are in outofscope",
			"status":      http.StatusNotAcceptable,
		})
		c.Abort()
		return
	}
	c.Next()
}
