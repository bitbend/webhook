package database

import (
	"time"
)

type DeletedRecord struct {
	Id        string    `db:"id"`
	TenantId  *string   `db:"tenant_id"`
	TableName string    `db:"table_name"`
	RecordId  string    `db:"record_id"`
	Data      []byte    `db:"data"`
	DeletedAt time.Time `db:"deleted_at"`
}
