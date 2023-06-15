package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

// ReadBuffer reads the buffer of c.Request.Body then restore the data of buffer to request body
func ReadBuffer(c *gin.Context) ([]byte, error) {
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, rawBody)
		return nil, err
	}
	// restore the buffer to C.Request.Body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	return rawBody, nil
}

func ReturnError(c *gin.Context, err error, status int, asset, scope string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error":  err.Error(),
		"asset":  asset,
		"scope":  scope,
		"status": status,
	})
}

func WildCardToRegex(input string) string {
	input = strings.ReplaceAll(input, "*", ".*")
	return input
}
