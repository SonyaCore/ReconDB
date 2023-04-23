package address

import (
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"ReconDB/pkg/check"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ValidateWildCard checks if the incoming request contains a valid wildcard scope in the request body.
// If the scope type is "wildcard", it checks whether the scope contains the '*' character using strings.Contains,
// and then validates the scope using a regular expression pattern check. If the scope is valid, it passes the request
// to the next middleware in the chain using c.Next(). Otherwise, it aborts the request and returns a JSON response
// indicating the error with a corresponding HTTP status code.
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
				"error":  "Not a valid Wildcard",
				"status": http.StatusNotAcceptable,
			})
		}

		if check.WildCardRegex(Scope.Scope) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  "Host wildcard is not acceptable",
			"status": http.StatusNotAcceptable,
		})
		return
	}

	c.Next()
}
