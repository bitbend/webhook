package model

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            string    `bun:"id,pk"`
	Name          string    `bun:"name"`
	Email         string    `bun:"email"`
	PasswordHash  string    `bun:"password_hash"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
}
