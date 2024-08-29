package database

import (
	"database/sql"
	"time"
)

type Stream struct {
	Id             string         `db:"id"`
	TenantId       string         `db:"tenant_id"`
	ProjectId      string         `db:"project_id"`
	EventType      string         `db:"event_type"`
	Description    sql.NullString `db:"description"`
	Schema         []byte         `db:"schema"`
	ForwardHeaders []byte         `db:"forward_headers"`
	SourceType     string         `db:"source_type"`
	SourceConfig   []byte         `db:"source_config"`
	Status         string         `db:"status"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
