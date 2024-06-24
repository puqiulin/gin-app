package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func CacheRoutes(router *gin.RouterGroup, userHandler *handler.CacheHandler) {
	cache := router.Group("")
	cache.GET("", userHandler.GetRank)
}
