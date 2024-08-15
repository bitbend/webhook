package database

import (
	"github.com/uptrace/bun"
	"time"
)

type ApiKey struct {
	bun.BaseModel `bun:"table:api_keys"`
	Id            string     `bun:"id,pk" db:"id"`
	TenantId      string     `bun:"tenant_id" db:"tenant_id"`
	ProjectId     string     `bun:"project_id" db:"project_id"`
	ApiKey        string     `bun:"api_key" db:"api_key"`
	RotatedAt     *time.Time `bun:"rotated_at" db:"rotated_at"`
	ExpiresAt     *time.Time `bun:"expires_at" db:"expires_at"`
	CreatedAt     time.Time  `bun:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at" db:"updated_at"`
}
