package company

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func ValidateCompanyName(c *gin.Context) {
	var Company models.Company
	var results int64

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Company)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	companyQuery := bson.M{
		"companyname": Company.CompanyName,
		//"programtype": Company.ProgramType,
	}

	results, err = database.CountDocuments("Company", companyQuery)
	if results >= 1 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"company": Company.CompanyName,
			"result":  "Duplicate Entry",
			"status":  http.StatusNotAcceptable,
		})
		return
	}
	c.Next()
}
