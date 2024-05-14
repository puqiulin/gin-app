package repository

import (
	"context"

	"gin-app/internal/model"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
	// Add more methods as needed
}

type userRepository struct {
	db  *bun.DB
	rdb *redis.Client
	l   *logrus.Logger
}

func NewUserRepository(db *bun.DB, rdb *redis.Client, l *logrus.Logger) UserRepository {
	return &userRepository{db: db, rdb: rdb, l: l}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = ?`
	err := r.db.NewRaw(query, id).Scan(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Phone, user.Gender).Scan(&user.ID)
	return err
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
