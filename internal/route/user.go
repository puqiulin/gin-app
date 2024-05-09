package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	user := router.Group("/user")
	user.GET("", userHandler.GetUser)
}
