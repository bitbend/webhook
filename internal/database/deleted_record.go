package database

import (
	"database/sql"
	"time"
)

type DeletedRecord struct {
	Id        string         `db:"id"`
	TenantId  sql.NullString `db:"tenant_id"`
	TableName string         `db:"table_name"`
	RecordId  string         `db:"record_id"`
	Data      []byte         `db:"data"`
	DeletedAt time.Time      `db:"deleted_at"`
}
