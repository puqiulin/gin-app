package handler

import (
	"gin-app/internal/graphql"
	"github.com/graphql-go/handler"
	"github.com/sirupsen/logrus"
)

type GraphQLHandler struct {
	Root *handler.Handler
}

func NewGraphQLHandler(rootSchema graphql.RootSchema, l *logrus.Logger) *GraphQLHandler {
	l.WithField("handler", "NewGraphQLHandler")

	h := handler.New(&handler.Config{
		Schema:   rootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	return &GraphQLHandler{
		Root: h,
	}
}
