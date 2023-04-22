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
		if WildCardRegex(Scope.Scope) {
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

func WildCardRegex(query string) bool {
	regex := regexp.MustCompile(middlewares.WildCardPattern)
	if regex.MatchString(query) {
		return true
	}
	return false
}
