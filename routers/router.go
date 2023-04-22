package routers

import (
	"ReconDB/api/asset"
	"ReconDB/api/company"
	"ReconDB/api/outscope"
	"ReconDB/api/scope"
	"ReconDB/middlewares"
	"github.com/gin-gonic/gin"
)

func ChainMiddleware(handlers ...gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, handler := range handlers {
			handler(c)
			if c.IsAborted() {
				return
			}
		}
	}
}

func RegisterRouter(router *gin.Engine) {
	AddressMiddleWare := ChainMiddleware(middlewares.ValidateSingleDomain, middlewares.ValidateWildCard,
		middlewares.ValidateIPAddress)

	CompanyMiddleWare := ChainMiddleware(middlewares.ProgramType, middlewares.CompanyValidate)

	api := router.Group("/api")
	{
		// scope router
		api.POST("/scope", AddressMiddleWare,
			middlewares.OutScopeCheck, middlewares.ValidateScopes, scope.AddScope)

		api.GET("/scope/:companyname", scope.GetScopes)
		api.GET("/scope", scope.GetAllScopes)
		api.DELETE("/scope/:companyname", scope.DeleteScopes)

		// out of scopes router
		api.POST("/outscope", AddressMiddleWare,
			middlewares.ValidateScopes, middlewares.OutScopeCheck, outscope.AddOutScope)

		api.GET("/outscope/:companyname", outscope.GetOutofScopes)
		api.GET("/outscope", outscope.GetAllOutofScopes)
		api.DELETE("/outscope/:companyname", outscope.DeleteOutofScopes)

		// company router
		api.POST("/company", CompanyMiddleWare, company.AddCompany)
		api.GET("/company/:companyname", company.GetCompany)
		api.GET("/company", company.GetAllCompanies)
		api.DELETE("/company/:companyname", company.DeleteCompany)

		// asset router
		api.POST("/asset", asset.AddAsset)
		api.GET("/asset/:asset", asset.GetAsset)
		api.GET("/asset", asset.GetAllAssets)
		api.DELETE("/asset/:asset", asset.DeleteAsset)

	}
}
