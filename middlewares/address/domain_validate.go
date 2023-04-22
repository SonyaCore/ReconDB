package address

import (
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func ValidateDomainName(domain string) bool {

	var RegExp = regexp.MustCompile(middlewares.DomainPattern)

	return RegExp.MatchString(domain)
}

func ValidateSingleDomain(c *gin.Context) {
	var Scope models.Scopes

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.ToLower(Scope.ScopeType) == "single" {
		if !ValidateDomainName(Scope.Scope) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  fmt.Sprintf("domain Name %s is invalid", Scope.Scope),
				"status": http.StatusNotAcceptable,
			})
			c.Abort()
			return
		}
		c.Next()
		return
	}
	c.Next()
}
