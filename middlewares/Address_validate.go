package middlewares

import (
	"ReconDB/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func CheckIPAddress(ip string) error {
	if net.ParseIP(ip) == nil {

		return fmt.Errorf("invalid IP Address: %s", ip)
	}
	return nil
}

func ValidateIPAddress(c *gin.Context) {
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

	if strings.ToLower(Scope.ScopeType) == "cidr" {
		ip, n, err := ParseCidr(Scope.Scope)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
				"status": http.StatusNotAcceptable,
			})
			c.Abort()
			return
		}

		if !n.IP.Equal(ip) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error":  fmt.Errorf("got %s; want %v\n", Scope.Scope, n),
				"status": http.StatusNotAcceptable,
			})
			c.Abort()
			return
		}
		c.Next()
		return
	}

	if strings.ToLower(Scope.ScopeType) == "ip" {
		err := CheckIPAddress(Scope.Scope)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"error":  err.Error(),
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

func ParseCidr(cidr string) (net.IP, *net.IPNet, error) {
	ip, n, err := net.ParseCIDR(cidr)
	return ip, n, err
}
