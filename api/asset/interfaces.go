package asset

import "github.com/gin-gonic/gin"

type Asset interface {
	AddAsset(c *gin.Context)
	GetAllAssets(c *gin.Context)
	GetAsset(c *gin.Context)
	DeleteAsset(c *gin.Context)
}
