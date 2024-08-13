package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Stream struct {
	bun.BaseModel  `bun:"table:streams"`
	Id             string    `bun:"id,pk"`
	TenantId       string    `bun:"tenant_id"`
	ProjectId      string    `bun:"project_id"`
	EventType      string    `bun:"event_type"`
	Description    *string   `bun:"description"`
	Schema         []byte    `bun:"schema"`
	ForwardHeaders []byte    `bun:"forward_headers"`
	SourceType     string    `bun:"source_type"`
	SourceConfig   []byte    `bun:"source_config"`
	Status         string    `bun:"status"`
	CreatedAt      time.Time `bun:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at"`
}
