package model

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
	Phone         string `bun:"phone,unique"`
	Email         string `bun:"email,unique"`
	Gender        string `bun:"gender"`
	LastLoginTime string `bun:"last_login_time"`
	LastLoginIp   string `bun:"last_login_ip"`
	CreateTime    string `bun:"create_time"`
	UpdateTime    string `bun:"create_time"`
}
