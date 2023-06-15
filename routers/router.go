package routers

import (
	"ReconDB/api/asset"
	"ReconDB/api/company"
	"ReconDB/api/outscope"
	"ReconDB/api/scope"
	"ReconDB/config"
	"ReconDB/validation/address"
	assetMiddleware "ReconDB/validation/asset"
	auth "ReconDB/validation/auth"
	companyMiddleware "ReconDB/validation/company"
	scopeMiddleware "ReconDB/validation/scope"
	"github.com/gin-gonic/gin"
	"log"
)

// RegisterRouter registers the api routes with their validation
func RegisterRouter(router *gin.Engine) {
	routerURI, err := config.RouterConfig()
	if err != nil {
		log.Fatal(err)
	}
	api := router.Group(routerURI.API)
	{
		// scope router
		api.POST(routerURI.Scope, auth.CheckAuth,
			address.ValidateHost,
			scopeMiddleware.OutScopeCheck,
			scopeMiddleware.ValidateScopeType,
			scope.AddScope)

		api.GET(routerURI.Scope+"/:companyname", auth.CheckAuth, scope.GetScopes)
		api.GET(routerURI.Scope, auth.CheckAuth, scope.GetAllScopes)
		api.DELETE(routerURI.Scope+"/:companyname", auth.CheckAuth, scope.DeleteScopes)

		// out of scopes router
		api.POST(routerURI.OutofScope, auth.CheckAuth,
			address.ValidateHost,
			scopeMiddleware.ValidateScopeType,
			scopeMiddleware.OutScopeCheck,
			outscope.AddOutScope)

		api.GET(routerURI.OutofScope+"/:companyname", auth.CheckAuth, outscope.GetOutofScopes)
		api.GET(routerURI.OutofScope, auth.CheckAuth, outscope.GetAllOutofScopes)
		api.DELETE(routerURI.OutofScope+"/:companyname", auth.CheckAuth, outscope.DeleteOutofScopes)

		// company router
		api.POST(routerURI.Company, auth.CheckAuth,
			companyMiddleware.ProgramType,
			companyMiddleware.ValidateCompanyName,
			company.AddCompany)

		api.GET(routerURI.Company+"/:companyname", auth.CheckAuth, company.GetCompany)
		api.GET(routerURI.Company, auth.CheckAuth, company.GetAllCompanies)
		api.DELETE(routerURI.Company+"/:companyname", auth.CheckAuth, company.DeleteCompany)

		// asset router
		api.POST(routerURI.Asset, auth.CheckAuth,
			assetMiddleware.DuplicateValidate,
			assetMiddleware.OutScopeAssetValidate,
			asset.AddAsset)

		api.GET(routerURI.Asset+"/:asset", auth.CheckAuth, asset.GetAsset)
		api.GET(routerURI.Asset, auth.CheckAuth, asset.GetAllAssets)
		api.DELETE(routerURI.Asset+"/:asset", auth.CheckAuth, asset.DeleteAsset)

	}
}
