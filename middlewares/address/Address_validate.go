package address

import (
	"ReconDB/models"
	"ReconDB/pkg/buffer"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strconv"
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

	rawBody, err := buffer.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if strings.ToLower(Scope.ScopeType) == "cidr" {
		ip, n, err := ParseCidr(Scope.Scope)
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
		if strings.Contains(Scope.Scope, ":") {
			validateIpPort(c, Scope)
			return
		}

		if err := CheckIPAddress(Scope.Scope); err != nil {
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

func validateIpPort(c *gin.Context, Scope models.Scopes) {
	parts := strings.Split(Scope.Scope, ":")
	ip := parts[0]
	port := parts[1]
	if err := CheckIPAddress(ip); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"input":  Scope.Scope,
			"error":  err.Error(),
			"status": http.StatusNotAcceptable,
		})
		return
	}
	if err := CheckPort(port); err != nil {
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

func ParseCidr(cidr string) (net.IP, *net.IPNet, error) {
	ip, n, err := net.ParseCIDR(cidr)
	return ip, n, err
}

func CheckPort(port string) error {
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("invalid port number: %s", port)
	}
	if portNum < 1 || portNum > 65535 {
		return fmt.Errorf("port number out of range: %d", portNum)
	}
	return nil
}
