package model

import (
	"context"
	"net"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Phone         string    `bun:"phone,unique" json:"phone"`
	Email         string    `bun:"email,unique" json:"email"`
	Gender        int       `bun:"gender" json:"gender"`
	LastLoginTime time.Time `bun:"last_login_time" json:"last_login_time"`
	LastLoginIp   net.IP    `bun:"last_login_ip" json:"last_login_ip"`
	CreateTime    time.Time `bun:"create_time,notnull,default:current_timestamp" json:"create_time"`
	UpdateTime    time.Time `bun:"update_time" json:"update_time"`

	Posts []*Post `bun:"rel:has-many,join:id=user_id"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreateTime = time.Now()
	case *bun.UpdateQuery:
		u.UpdateTime = time.Now()
	}
	return nil
}
