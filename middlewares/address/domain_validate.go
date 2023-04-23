package address

import (
	"ReconDB/middlewares"
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"ReconDB/pkg/domain"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// ValidateDomainName validates domain name with DomainPattern regex
func ValidateDomainName(domain string) bool {

	var RegExp = regexp.MustCompile(middlewares.DomainPattern)

	return RegExp.MatchString(domain)
}

// ValidateSingleDomain validates a domain:port string using the ValidateDomainName function to
// ensure that the domain is valid. It checks if the string is of the form "domain:port",
// splits the domain and port using the ":" delimiter, and then checks the port for validity.
// Finally, it calls ValidateDomainName to check the domain for validity as well. c.Next()
// if the domain:port or domain string is valid, and c.Abort() otherwise.
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
		if err = domain.CheckDomain(Scope.Scope); err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
				"status": http.StatusNotAcceptable,
			})
			return
		}
		c.Next()
		return
	}
	c.Next()
}
