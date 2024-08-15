package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Subscription struct {
	bun.BaseModel     `bun:"table:subscriptions"`
	Id                string    `bun:"id,pk" db:"id"`
	TenantId          string    `bun:"tenant_id" db:"tenant_id"`
	StreamId          string    `bun:"stream_id" db:"stream_id"`
	EndpointId        string    `bun:"endpoint_id" db:"endpoint_id"`
	RateLimitCount    int       `bun:"rate_limit_count" db:"rate_limit_count"`
	RateLimitDuration int       `bun:"rate_limit_duration" db:"rate_limit_duration"`
	RetryStrategy     string    `bun:"retry_strategy" db:"retry_strategy"`
	RetryCount        int       `bun:"retry_count" db:"retry_count"`
	RetryDuration     int       `bun:"retry_duration" db:"retry_duration"`
	CreatedAt         time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt         time.Time `bun:"updated_at" db:"updated_at"`
}
