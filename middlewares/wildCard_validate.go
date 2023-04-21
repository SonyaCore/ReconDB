package middlewares

import (
	"ReconDB/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var WildCardPattern = `^(\*|(\*|\*\.)?\w+(\.\w+)*(\.\*|\*)?)$`

func ValidateWildCard(c *gin.Context) {
	var Scope models.Scopes

	// Read the content
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, rawBody)
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.ToLower(Scope.ScopeType) == "wildcard" {
		regex := regexp.MustCompile(WildCardPattern)
		if regex.MatchString(Scope.Scope) {
			c.Next()
			return
		}
		c.JSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  "host wildcard is not acceptable",
			"status": http.StatusNotAcceptable,
		})
		c.Abort()
		return
	}

	c.Next()
}
