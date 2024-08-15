package database

import (
	"time"
)

type Event struct {
	Id             string     `db:"id"`
	TenantId       string     `db:"tenant_id"`
	StreamId       string     `db:"stream_id"`
	Data           []byte     `db:"data"`
	IdempotencyKey *string    `db:"idempotency_key"`
	CustomHeaders  []byte     `db:"custom_headers"`
	ScheduledAt    *time.Time `db:"scheduled_at"`
	CreatedAt      time.Time  `db:"created_at"`
}
