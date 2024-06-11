package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func GoogleRoutes(router *gin.RouterGroup, GoogleHandler *handler.GoogleHandler) {
	google := router.Group("/")
	google.GET("", GoogleHandler.AuthorizedCallBack)
}
