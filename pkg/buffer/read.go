package buffer

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
