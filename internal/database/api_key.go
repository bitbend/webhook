package database

import (
	"github.com/uptrace/bun"
	"time"
)

type ApiKey struct {
	bun.BaseModel `bun:"table:api_keys"`
	Id            string     `bun:"id,pk"`
	TenantId      string     `bun:"tenant_id"`
	ProjectId     string     `bun:"project_id"`
	ApiKey        string     `bun:"api_key"`
	RotatedAt     *time.Time `bun:"rotated_at"`
	ExpiresAt     *time.Time `bun:"expires_at"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
}
