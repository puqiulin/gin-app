package route

import (
	"gin-app/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	r *gin.Engine,
	userHandler *handler.UserHandler,
	graphQLHandler *handler.GraphQLHandler,
	googleHandler *handler.GoogleHandler,
) {
	api := r.Group("/api")
	UserRoutes(api, userHandler)

	graphql := api.Group("/graphql")
	GraphqlRoutes(graphql, graphQLHandler)

	google := api.Group("/google")
	GoogleRoutes(google, googleHandler)
}
