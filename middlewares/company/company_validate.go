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

// ValidateCompanyName checks if there was any duplication in Company collection
// if len results was more than 1 c.Abort with error message , otherwise c.Next()
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
