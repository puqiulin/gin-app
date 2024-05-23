package repository

import (
	"context"

	"gin-app/internal/model"
)

func (r *Repository) GetUserByID(ctx context.Context, id int64) (user *model.User, err error) {
	//https://github.com/uptrace/bun/issues/336
	user = new(model.User)
	err = r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetAllUser(ctx context.Context) (users []*model.User, err error) {
	err = r.db.NewSelect().Model(&users).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (r *Repository) CreateUser(ctx context.Context, u *model.User) (user *model.User, err error) {
	_, err = r.db.NewInsert().Model(u).Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (r *Repository) UpdateUser(ctx context.Context, u *model.User) (user *model.User, err error) {
	query := r.db.NewUpdate().Model(u).WherePK()

	if user.Name != "" {
		query.Set("name = ?", user.Name)
	}
	if user.Email != "" {
		query.Set("email = ?", user.Email)
	}
	if user.Phone != "" {
		query.Set("phone = ?", user.Phone)
	}
	if user.Gender != 0 {
		query.Set("gender = ?", user.Gender)
	}

	_, err = query.Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (r *Repository) DeleteUser(ctx context.Context, u *model.User) (user *model.User, err error) {
	_, err = r.db.NewDelete().Model(u).WherePK().Returning("*").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return
}
