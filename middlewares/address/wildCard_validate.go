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
