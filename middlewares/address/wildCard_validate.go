package address

import (
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func ValidateWildCard(c *gin.Context) {
	var Scope models.Scopes

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.ToLower(Scope.ScopeType) == "wildcard" {
		if !strings.Contains(Scope.Scope, "*") {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  "not a valid wildcard",
				"status": http.StatusNotAcceptable,
			})
		}

		if WildCardRegex(Scope.Scope) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  "host wildcard is not acceptable",
			"status": http.StatusNotAcceptable,
		})
		return
	}

	c.Next()
}

func WildCardRegex(query string) bool {
	regex := regexp.MustCompile(middlewares.WildCardPattern)
	if regex.MatchString(query) {
		return true
	}
	return false
}
