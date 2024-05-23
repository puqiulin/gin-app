package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID         int64     `bun:"id,pk,autoincrement" ,json:"id"`
	UserID     int64     `bun:"user_id,notnull" ,json:"user_id"`
	Title      string    `bun:"title,notnull" ,json:"title"`
	Body       string    `bun:"body" ,json:"body"`
	CreateTime time.Time `bun:"create_time" ,json:"create_time"`
	UpdateTime time.Time `bun:"update_time" ,json:"update_time"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
}
