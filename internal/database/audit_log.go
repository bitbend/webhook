package database

import (
	"github.com/uptrace/bun"
	"time"
)

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_logs"`
	Id            string    `bun:"id,pk"`
	TenantId      *string   `bun:"tenant_id"`
	Name          string    `bun:"name"`
	Data          []byte    `bun:"data"`
	CreatedAt     time.Time `bun:"created_at"`
}
