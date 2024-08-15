package database

import (
	"github.com/uptrace/bun"
	"time"
)

type TenantOwner struct {
	bun.BaseModel `bun:"table:tenant_owners"`
	Id            string    `bun:"id,pk" db:"id"`
	TenantId      string    `bun:"tenant_id" db:"tenant_id"`
	UserId        string    `bun:"user_id" db:"user_id"`
	CreatedAt     time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at" db:"updated_at"`
}
