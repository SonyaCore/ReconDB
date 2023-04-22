package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func CheckAuth(c *gin.Context) {
	var token = viper.GetString("authorization")
	var header = c.Request.Header.Get("Authorization")

	if token != header {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Unauthorized Access",
			"status": http.StatusUnauthorized,
		})
		c.Abort()
		return
	}
	c.Next()
}
