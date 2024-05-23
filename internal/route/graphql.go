package route

import (
	"gin-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func GraphqlRoutes(
	router *gin.RouterGroup,
	h *handler.GraphQLHandler,
) {
	root := router.Group("")
	root.POST("", func(c *gin.Context) {
		h.Root.ServeHTTP(c.Writer, c.Request)
	})
}
