package graphql

import (
	"context"

	"gin-app/internal/model"
	"gin-app/internal/repository"
	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

type PostResolver struct {
	db *repository.Repository
	l  *logrus.Logger
}

func NewPostResolver(db *repository.Repository, l *logrus.Logger) *PostResolver {
	return &PostResolver{
		db: db,
		l:  l,
	}
}

func (r *PostResolver) GetPostByID(p graphql.ResolveParams) (Post interface{}, err error) {
	id := p.Args["id"].(int)
	ctx := context.Background()
	Post, err = r.db.GetPostByID(ctx, int64(id))
	if err != nil {
		r.l.Error(err)
		return nil, err
	}
	return
}

func (r *PostResolver) GetAllPosts(p graphql.ResolveParams) (Posts interface{}, err error) {
	ctx := context.Background()
	Posts, err = r.db.GetAllPosts(ctx)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}
	return
}

func (r *PostResolver) CreatePost(p graphql.ResolveParams) (res interface{}, err error) {
	Post := model.Post{
		Title: p.Args["title"].(string),
		Body:  p.Args["body"].(string),
	}
	ctx := context.Background()

	res, err = r.db.CreatePost(ctx, &Post)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *PostResolver) UpdatePost(p graphql.ResolveParams) (res interface{}, err error) {
	id, idOk := p.Args["id"].(int)
	title, titleOk := p.Args["title"].(string)
	body, bodyOk := p.Args["body"].(string)

	Post := &model.Post{}

	if idOk {
		Post.ID = int64(id)
	} else {
		return "", nil
	}
	if titleOk {
		Post.Title = title
	}
	if bodyOk {
		Post.Body = body
	}

	ctx := context.Background()
	res, err = r.db.UpdatePost(ctx, Post)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *PostResolver) DeletePost(p graphql.ResolveParams) (res interface{}, err error) {
	id, idOk := p.Args["id"].(int)
	if !idOk {
		return "", nil
	}

	Post := &model.Post{ID: int64(id)}
	ctx := context.Background()
	res, err = r.db.DeletePost(ctx, Post)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}

func (r *PostResolver) GetUserByPostID(p graphql.ResolveParams) (user interface{}, err error) {
	post := p.Source.(*model.Post)
	user, err = r.db.GetUserByID(p.Context, post.UserID)
	if err != nil {
		r.l.Error(err)
		return nil, err
	}

	return
}
