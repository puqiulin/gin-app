package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	g := router.Group("/user")
	g.GET("", userHandler.GetUser)
}
