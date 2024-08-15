package database

import (
	"github.com/uptrace/bun"
	"time"
)

type DeletedRecord struct {
	bun.BaseModel `bun:"table:deleted_records"`
	Id            string    `bun:"id,pk" db:"id"`
	TenantId      *string   `bun:"tenant_id" db:"tenant_id"`
	TableName     string    `bun:"table_name" db:"table_name"`
	RecordId      string    `bun:"record_id" db:"record_id"`
	Data          []byte    `bun:"data" db:"data"`
	DeletedAt     time.Time `bun:"deleted_at" db:"deleted_at"`
}
