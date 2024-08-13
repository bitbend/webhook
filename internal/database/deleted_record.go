package database

import (
	"github.com/uptrace/bun"
	"time"
)

type DeletedRecord struct {
	bun.BaseModel `bun:"table:deleted_records"`
	Id            string    `bun:"id,pk"`
	TenantId      *string   `bun:"tenant_id"`
	TableName     string    `bun:"table_name"`
	RecordId      string    `bun:"record_id"`
	Data          []byte    `bun:"data"`
	DeletedAt     time.Time `bun:"deleted_at"`
}
