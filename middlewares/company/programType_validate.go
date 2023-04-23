package company

import (
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ProgramType checks the input program type with defined ProgramTypes []string
// if it not matches c.Abort with error else c.Next()
func ProgramType(c *gin.Context) {
	var Company models.Company

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Company)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for i, _ := range middlewares.ProgramTypes {
		if strings.ToLower(Company.ProgramType) == middlewares.ProgramTypes[i] {
			c.Next()
			return
		}
		continue
	}

	c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
		"error":       "Program Type is not valid",
		"valid_types": middlewares.ProgramTypes,
		"status":      http.StatusFailedDependency,
	})
	return
}
