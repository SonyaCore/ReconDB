package routers

import (
	"ReconDB/api/asset"
	"ReconDB/api/company"
	"ReconDB/api/outscope"
	"ReconDB/api/scope"
	"ReconDB/middlewares/address"
	assetMiddleware "ReconDB/middlewares/asset"
	auth "ReconDB/middlewares/auth"
	companyMiddleware "ReconDB/middlewares/company"
	scopeMiddleware "ReconDB/middlewares/scope"
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

// AddressMiddleWare used to validate host input from client
var AddressMiddleWare = ChainMiddleware(address.ValidateSingleDomain, address.ValidateWildCard,
	address.ValidateIPAddress)

// CompanyMiddleWare validate company duplication and program type
//var CompanyMiddleWare = ChainMiddleware(companyMiddleware.ProgramType, companyMiddleware.CompanyValidate)

func RegisterRouter(router *gin.Engine) {

	api := router.Group("/api")
	{
		// scope router
		api.POST("/scope", auth.CheckAuth, AddressMiddleWare,
			scopeMiddleware.OutScopeCheck, scopeMiddleware.ValidateScopes, scope.AddScope)

		api.GET("/scope/:companyname", auth.CheckAuth, scope.GetScopes)
		api.GET("/scope", auth.CheckAuth, scope.GetAllScopes)
		api.DELETE("/scope/:companyname", auth.CheckAuth, scope.DeleteScopes)

		// out of scopes router
		api.POST("/outscope", auth.CheckAuth, AddressMiddleWare,
			scopeMiddleware.ValidateScopes, scopeMiddleware.OutScopeCheck, outscope.AddOutScope)

		api.GET("/outscope/:companyname", auth.CheckAuth, outscope.GetOutofScopes)
		api.GET("/outscope", auth.CheckAuth, outscope.GetAllOutofScopes)
		api.DELETE("/outscope/:companyname", auth.CheckAuth, outscope.DeleteOutofScopes)

		// company router
		api.POST("/company", auth.CheckAuth, companyMiddleware.ProgramType, companyMiddleware.CompanyValidate, company.AddCompany)
		api.GET("/company/:companyname", auth.CheckAuth, company.GetCompany)
		api.GET("/company", auth.CheckAuth, company.GetAllCompanies)
		api.DELETE("/company/:companyname", auth.CheckAuth, company.DeleteCompany)

		// asset router
		api.POST("/asset", auth.CheckAuth, assetMiddleware.DuplicateValidate, asset.AddAsset)
		api.GET("/asset/:asset", auth.CheckAuth, asset.GetAsset)
		api.GET("/asset", auth.CheckAuth, asset.GetAllAssets)
		api.DELETE("/asset/:asset", auth.CheckAuth, asset.DeleteAsset)

	}
}
