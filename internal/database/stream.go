package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Stream struct {
	bun.BaseModel  `bun:"table:streams"`
	Id             string    `bun:"id,pk" db:"id"`
	TenantId       string    `bun:"tenant_id" db:"tenant_id"`
	ProjectId      string    `bun:"project_id" db:"project_id"`
	EventType      string    `bun:"event_type" db:"event_type"`
	Description    *string   `bun:"description" db:"description"`
	Schema         []byte    `bun:"schema" db:"schema"`
	ForwardHeaders []byte    `bun:"forward_headers" db:"forward_headers"`
	SourceType     string    `bun:"source_type" db:"source_type"`
	SourceConfig   []byte    `bun:"source_config" db:"source_config"`
	Status         string    `bun:"status" db:"status"`
	CreatedAt      time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at" db:"updated_at"`
}
