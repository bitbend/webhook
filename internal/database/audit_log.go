package database

import (
	"database/sql"
	"time"
)

type AuditLog struct {
	Id        string         `db:"id"`
	TenantId  sql.NullString `db:"tenant_id"`
	Name      string         `db:"name"`
	Data      []byte         `db:"data"`
	CreatedAt time.Time      `db:"created_at"`
}
