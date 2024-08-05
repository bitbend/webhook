package model

import (
	"database/sql"
	"encoding/json"
	"github.com/uptrace/bun"
	"time"
)

type Event struct {
	bun.BaseModel  `bun:"table:events"`
	Id             string          `bun:"id,pk"`
	TenantId       string          `bun:"tenant_id"`
	StreamId       string          `bun:"stream_id"`
	Data           json.RawMessage `bun:"data"`
	IdempotencyKey sql.NullString  `bun:"idempotency_key"`
	CustomHeaders  json.RawMessage `bun:"custom_headers"`
	ScheduledAt    sql.NullTime    `bun:"scheduled_at"`
	CreatedAt      time.Time       `bun:"created_at"`
}
