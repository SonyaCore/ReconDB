package company

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/utils"
	"ReconDB/validation"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

// ProgramType checks the input program with defined ProgramTypes []string
// if it not matches c.Abort with error else c.Next()
func ProgramType(c *gin.Context) {
	var Company models.Company

	rawBody, err := utils.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Company)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for i, _ := range validation.ProgramTypes {
		if strings.ToLower(Company.ProgramType) == validation.ProgramTypes[i] {
			c.Next()
			return
		}
		continue
	}

	c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
		"error":       "Program Type is not valid",
		"valid_types": validation.ProgramTypes,
		"status":      http.StatusFailedDependency,
	})
	return
}

// ValidateCompanyName checks if there was any duplication in Company collection
// if len results was more than 1 c.Abort with error message , otherwise c.Next()
func ValidateCompanyName(c *gin.Context) {
	var Company models.Company
	var results int64

	rawBody, err := utils.ReadBuffer(c)

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
