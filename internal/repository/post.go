package repository

import (
	"context"

	"gin-app/internal/model"
)

func (r *Repository) GetPostByID(ctx context.Context, id int64) (*model.Post, error) {
	var Post model.Post
	query := `SELECT * FROM Posts WHERE id = ?`
	err := r.db.NewRaw(query, id).Scan(ctx, &Post)
	if err != nil {
		return nil, err
	}
	return &Post, nil
}

func (r *Repository) GetAllPosts(ctx context.Context) (posts []*model.Post, err error) {
	err = r.db.NewSelect().Model(&posts).Order("create_time DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (r *Repository) CreatePost(ctx context.Context, Post *model.Post) (*model.Post, error) {
	_, err := r.db.NewInsert().Model(Post).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return Post, nil
}

func (r *Repository) UpdatePost(ctx context.Context, Post *model.Post) (*model.Post, error) {
	query := r.db.NewUpdate().Model(Post).WherePK()

	if Post.Title != "" {
		query.Set("title = ?", Post.Title)
	}
	if Post.Body != "" {
		query.Set("email = ?", Post.Body)
	}

	_, err := query.Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return Post, nil
}

func (r *Repository) DeletePost(ctx context.Context, Post *model.Post) (*model.Post, error) {
	_, err := r.db.NewDelete().Model(Post).WherePK().Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return Post, nil
}

func (r *Repository) GetPostsByUserID(ctx context.Context, id int64) (posts []*model.Post, err error) {
	err = r.db.NewSelect().Model(&posts).Where("user_id = ?", id).Order("create_time DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return
}
