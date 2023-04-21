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
		api.GET("/scope/:scope", scope.GetScopes)
		api.GET("/scope", scope.GetAllScopes)

		// out of scopes router
		api.POST("/outscope", middlewares.ValidateScopes, outscope.AddOutScope)
		api.GET("/outscope/:scope", outscope.GetOutofScopes)
		api.GET("/outscope", outscope.GetAllOutofScopes)

		// company router
		api.POST("/company", middlewares.ProgramType, company.AddCompany)
		api.GET("/company/:companyname", company.GetCompany)
		api.GET("/company", company.GetAllCompanies)

		// asset router
		api.POST("/asset", asset.AddAsset)
		api.GET("/asset/:asset", asset.GetAsset)
		api.GET("/asset", asset.GetAllAssets)

	}
}
