package graphql

import (
	"context"
	"net"

	"gin-app/internal/model"
	"gin-app/internal/repository"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

type UserResolver struct {
	db *repository.Repository
	l  *logrus.Logger
}

func NewUserResolver(db *repository.Repository, l *logrus.Logger) *UserResolver {
	return &UserResolver{
		db: db,
		l:  l,
	}
}

func (r *UserResolver) GetUserByID(p graphql.ResolveParams) (user interface{}, err error) {
	id := p.Args["id"].(int)
	ctx := context.Background()
	user, err = r.db.GetUserByID(ctx, int64(id))
	if err != nil {
		r.l.Error(err)
		return nil, err
	}
	return
}

func (r *UserResolver) GetAllUser(p graphql.ResolveParams) (users interface{}, err error) {
	ctx := context.Background()
	users, err = r.db.GetAllUser(ctx)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}
	return
}

func (r *UserResolver) GetPostsByUserID(p graphql.ResolveParams) (posts interface{}, err error) {
	user := p.Source.(*model.User)
	ctx := context.Background()

	posts, err = r.db.GetPostsByUserID(ctx, user.ID)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *UserResolver) CreateUser(p graphql.ResolveParams) (res interface{}, err error) {
	user := model.User{
		Name:        p.Args["name"].(string),
		Email:       p.Args["email"].(string),
		Phone:       p.Args["phone"].(string),
		Gender:      p.Args["gender"].(int),
		LastLoginIp: net.ParseIP("0.0.0.0"),
	}
	ctx := context.Background()

	res, err = r.db.CreateUser(ctx, &user)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *UserResolver) UpdateUser(p graphql.ResolveParams) (res interface{}, err error) {
	id, idOk := p.Args["id"].(int)
	name, nameOk := p.Args["name"].(string)
	email, emailOk := p.Args["email"].(string)
	phone, phoneOk := p.Args["phone"].(string)
	gender, genderOk := p.Args["gender"].(int)

	user := &model.User{}

	if idOk {
		user.ID = int64(id)
	} else {
		return "", nil
	}
	if nameOk {
		user.Name = name
	}
	if emailOk {
		user.Email = email
	}
	if phoneOk {
		user.Phone = phone
	}
	if genderOk {
		user.Gender = gender
	}

	ctx := context.Background()
	res, err = r.db.UpdateUser(ctx, user)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *UserResolver) DeleteUser(p graphql.ResolveParams) (res interface{}, err error) {
	id, idOk := p.Args["id"].(int)
	if !idOk {
		return "", nil
	}

	user := &model.User{ID: int64(id)}
	ctx := context.Background()
	res, err = r.db.DeleteUser(ctx, user)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return res, nil
}
