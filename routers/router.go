package routers

import (
	"ReconDB/api/test"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	v2 := router.Group("/api")
	{

		v2.GET("/ping", test.Hello)

	}
}
