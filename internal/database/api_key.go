package database

import (
	"time"
)

type ApiKey struct {
	Id        string     `db:"id"`
	TenantId  string     `db:"tenant_id"`
	ProjectId string     `db:"project_id"`
	ApiKey    string     `db:"api_key"`
	RotatedAt *time.Time `db:"rotated_at"`
	ExpiresAt *time.Time `db:"expires_at"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
