package graphql

import (
	"github.com/graphql-go/graphql"
)

type RootSchema *graphql.Schema

func NewRootSchema(
	userField *UserField,
	postField *PostField,
) RootSchema {
	querySchema := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user":  userField.User,
				"users": userField.Users,

				"post":  postField.Post,
				"posts": postField.Posts,
			},
		})

	mutationSchema := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create_user": userField.Create,
			"update_user": userField.Update,
			"delete_user": userField.Delete,

			"create_post": userField.Create,
			"update_post": postField.Update,
			"delete_post": postField.Delete,
		},
	})

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    querySchema,
			Mutation: mutationSchema,
		},
	)
	if err != nil {
		panic(err)
	}

	return &schema
}
