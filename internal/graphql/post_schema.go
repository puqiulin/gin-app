package graphql

import (
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"create_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"update_time": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

type PostField struct {
	Post   *graphql.Field
	Posts  *graphql.Field
	Create *graphql.Field
	Update *graphql.Field
	Delete *graphql.Field
}

func NewPostField(resolver *PostResolver) *PostField {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type:    UserType,
		Resolve: resolver.GetUserByPostID,
	})

	return &PostField{
		Post: &graphql.Field{
			Type:        PostType,
			Description: "Get Post by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.GetPostByID,
		},
		Posts: &graphql.Field{
			Type:        graphql.NewList(PostType),
			Description: "Get Post list",
			Resolve:     resolver.GetAllPosts,
		},
		Create: &graphql.Field{
			Type:        PostType,
			Description: "Create new Post",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"body": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.CreatePost,
		},
		Update: &graphql.Field{
			Type:        PostType,
			Description: "Update Post by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"body": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.UpdatePost,
		},
		Delete: &graphql.Field{
			Type:        PostType,
			Description: "Delete Post by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.DeletePost,
		},
	}

}
