package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type ApiKey struct {
	bun.BaseModel `bun:"table:api_keys"`
	Id            string       `bun:"id,pk"`
	TenantId      string       `bun:"tenant_id"`
	ProjectId     string       `bun:"project_id"`
	ApiKey        string       `bun:"api_key"`
	RotatedAt     sql.NullTime `bun:"rotated_at"`
	ExpiresAt     sql.NullTime `bun:"expires_at"`
	CreatedAt     time.Time    `bun:"created_at"`
	UpdatedAt     time.Time    `bun:"updated_at"`
}
