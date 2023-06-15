package auth

import (
	"ReconDB/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAuth is a simple token which is defined in configuration file
// it compares the authorization value in config file with Authorization header
// if it was valid c.Next() , otherwise c.Abort
func CheckAuth(c *gin.Context) {
	var configuration, _ = config.LoadConfig()
	var token = configuration.Authorization
	var header = c.Request.Header.Get("Authorization")

	if token != header {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unauthorized Access",
			"status": http.StatusUnauthorized,
		})
		return
	}
	c.Next()
}
