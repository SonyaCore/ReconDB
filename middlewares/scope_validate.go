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

		c.JSON(http.StatusFailedDependency, gin.H{
			"error":       "scope type is not valid",
			"valid_types": Scopes,
			"status":      http.StatusFailedDependency,
		})

		c.Abort()
		return
	}

}
