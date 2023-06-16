package address

import (
	"ReconDB/models"
	"ReconDB/pkg/host"
	"ReconDB/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ValidateHost validates host according to their types to ensure the input is valid.
func ValidateHost(c *gin.Context) {
	var Scope models.Scopes

	rawBody, err := utils.ReadBuffer(c)
	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	switch strings.ToLower(Scope.ScopeType) {
	case "single":
		if err = host.CheckDomain(Scope.Scope); err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
				"status": http.StatusNotAcceptable,
			})
			return
		}
		c.Next()
		return
	case "wildcard":
		if !strings.Contains(Scope.Scope, "*") {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  "Not a valid Wildcard",
				"status": http.StatusNotAcceptable,
			})
		}
		if host.WildCardRegex(Scope.Scope) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  "Host wildcard is not acceptable",
			"status": http.StatusNotAcceptable,
		})
		return
	case "cidr":
		ip, n, err := host.ParseCidr(Scope.Scope)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
				"status": http.StatusNotAcceptable,
			})
			return
		}

		if !n.IP.Equal(ip) {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"error":  fmt.Sprintf("got %s; want %v", Scope.Scope, n),
				"status": http.StatusNotAcceptable,
			})
			return
		}
		c.Next()
		return
	case "ip":
		if err = host.IpAddress(Scope.Scope); err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
				"status": http.StatusNotAcceptable,
			})
			return
		}
		c.Next()
		return
	default:
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  "cannot find related Scope type",
			"status": http.StatusNotAcceptable,
		})
		return
	}
}
