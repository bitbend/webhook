package database

import (
	"github.com/uptrace/bun"
	"time"
)

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_logs"`
	Id            string    `bun:"id,pk" db:"id"`
	TenantId      *string   `bun:"tenant_id" db:"tenant_id"`
	Name          string    `bun:"name" db:"name"`
	Data          []byte    `bun:"data" db:"data"`
	CreatedAt     time.Time `bun:"created_at" db:"created_at"`
}
