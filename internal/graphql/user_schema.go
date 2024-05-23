package graphql

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.Int,
			},
			"last_login_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"last_login_ip": &graphql.Field{
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

type UserField struct {
	User   *graphql.Field
	Users  *graphql.Field
	Create *graphql.Field
	Update *graphql.Field
	Delete *graphql.Field
}

func NewUserField(resolver *UserResolver) *UserField {
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type:    graphql.NewList(PostType),
		Resolve: resolver.GetPostsByUserID,
	})

	return &UserField{
		User: &graphql.Field{
			Type:        UserType,
			Description: "Get user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.GetUserByID,
		},

		/* Get (read) user list
		   http://localhost:9999/api/graphql/user?query={list{id,name,email,phone,gender,last_login_ip,create_time,update_time}}
		*/
		Users: &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Get user list",
			Resolve:     resolver.GetAllUser,
		},

		Create: &graphql.Field{
			Type:        UserType,
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"phone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"gender": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.CreateUser,
		},

		/* Update user by id
		   http://localhost:9999/api/graphql/user?query=mutation+_{update(id:1,name:"asd",phone:"123456"){id,name,phone}}
		*/
		Update: &graphql.Field{
			Type:        UserType,
			Description: "Update user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"phone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"gender": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.UpdateUser,
		},

		/* Delete user by id
		   http://localhost:9999/api/graphql/user?query=mutation+_{delete(id:1){id}}
		*/
		Delete: &graphql.Field{
			Type:        UserType,
			Description: "Delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.DeleteUser,
		},
	}
}
