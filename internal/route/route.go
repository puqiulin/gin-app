package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler *handler.UserHandler) {
	api := r.Group("/api")
	UserRoutes(api, userHandler)
}
