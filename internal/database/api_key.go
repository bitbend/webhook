package database

import (
	"database/sql"
	"time"
)

type ApiKey struct {
	Id        string       `db:"id"`
	TenantId  string       `db:"tenant_id"`
	ProjectId string       `db:"project_id"`
	ApiKey    string       `db:"api_key"`
	RotatedAt sql.NullTime `db:"rotated_at"`
	ExpiresAt sql.NullTime `db:"expires_at"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
