package database

import (
	"database/sql"
	"encoding/json"
	"github.com/uptrace/bun"
	"time"
)

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_logs"`
	Id            string          `bun:"id,pk"`
	TenantId      sql.NullString  `bun:"tenant_id"`
	Name          string          `bun:"name"`
	Data          json.RawMessage `bun:"data"`
	CreatedAt     time.Time       `bun:"created_at"`
}
