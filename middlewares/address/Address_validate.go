package address

import (
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"ReconDB/pkg/check"
	"ReconDB/pkg/parser"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ValidateIPAddress validates address in input if it was cidr or ip
func ValidateIPAddress(c *gin.Context) {
	var Scope models.Scopes

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.ToLower(Scope.ScopeType) == "cidr" {
		ip, n, err := parser.ParseCidr(Scope.Scope)
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
				"error":  fmt.Sprintf("got %s; want %v\n", Scope.Scope, n),
				"status": http.StatusNotAcceptable,
			})
			return
		}
		c.Next()
		return
	}

	if strings.ToLower(Scope.ScopeType) == "ip" {
		if err = check.IpAddress(Scope.Scope); err != nil {
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
