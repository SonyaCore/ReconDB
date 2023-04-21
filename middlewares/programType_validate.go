package middlewares

import (
	"ReconDB/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func ProgramType(c *gin.Context) {
	var Company models.Company

	// Read the content
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, rawBody)
	}

	// Unmarshal rawBody to Company
	err = json.Unmarshal(rawBody, &Company)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	for i, _ := range ProgramTypes {

		if Company.ProgramType == ProgramTypes[i] {
			c.Next()
			return
		}

		c.JSON(http.StatusFailedDependency, gin.H{
			"error":       "program type is not valid",
			"valid_types": ProgramTypes,
			"status":      http.StatusFailedDependency,
		})

		c.Abort()
		return
	}

}