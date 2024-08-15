package database

import (
	"time"
)

type AuditLog struct {
	Id        string    `db:"id"`
	TenantId  *string   `db:"tenant_id"`
	Name      string    `db:"name"`
	Data      []byte    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
}
