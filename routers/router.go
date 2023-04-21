package routers

import (
	"ReconDB/api/asset"
	"ReconDB/api/company"
	"ReconDB/api/outscope"
	"ReconDB/api/scope"
	"ReconDB/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		// scope router
		api.POST("/scope", middlewares.ValidateScopes, scope.AddScope)

		// out of scopes router
		api.POST("/outscope", middlewares.ValidateScopes, outscope.AddOutScope)

		// company router
		api.POST("/company", middlewares.ProgramType, company.AddCompany)

		// asset router
		api.POST("/asset", asset.AddAsset)

	}
}
