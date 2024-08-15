package database

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            string    `bun:"id,pk" db:"id"`
	Name          string    `bun:"name" db:"name"`
	Email         string    `bun:"email" db:"email"`
	PasswordHash  string    `bun:"password_hash" db:"password_hash"`
	CreatedAt     time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at" db:"updated_at"`
}
