package middlewares

import (
	"ReconDB/database"
	"ReconDB/models"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
)

func CompanyValidate(c *gin.Context) {
	var Company models.Company
	var ctx = context.TODO()

	// Read the content
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, rawBody)
	}

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Company)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	companyQuery := bson.M{
		"companyname": Company.CompanyName,
		"programtype": Company.ProgramType,
	}

	var collection *mongo.Collection
	var results int64

	collection = database.Collection("Company")
	results, err = collection.CountDocuments(ctx, companyQuery)

	if results >= 1 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"company": Company.CompanyName,
			"result":  "duplicate entry",
			"status":  http.StatusNotAcceptable,
		})
		c.Abort()
		return
	}

	c.Next()
}
