package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Event struct {
	bun.BaseModel  `bun:"table:events"`
	Id             string     `bun:"id,pk" db:"id"`
	TenantId       string     `bun:"tenant_id" db:"tenant_id"`
	StreamId       string     `bun:"stream_id" db:"stream_id"`
	Data           []byte     `bun:"data" db:"data"`
	IdempotencyKey *string    `bun:"idempotency_key" db:"idempotency_key"`
	CustomHeaders  []byte     `bun:"custom_headers" db:"custom_headers"`
	ScheduledAt    *time.Time `bun:"scheduled_at" db:"scheduled_at"`
	CreatedAt      time.Time  `bun:"created_at" db:"created_at"`
}
