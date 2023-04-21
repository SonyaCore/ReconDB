package middlewares

import (
	"ReconDB/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func validateDomainName(domain string) bool {

	var RegExp = regexp.MustCompile(DomainPattern)

	return RegExp.MatchString(domain)
}

func ValidateSingleDomain(c *gin.Context) {
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

	if strings.ToLower(Scope.ScopeType) == "single" {
		if !validateDomainName(Scope.Scope) {
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
