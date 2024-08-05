package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Subscription struct {
	bun.BaseModel     `bun:"table:subscriptions"`
	Id                string    `bun:"id,pk"`
	TenantId          string    `bun:"tenant_id"`
	StreamId          string    `bun:"stream_id"`
	EndpointId        string    `bun:"endpoint_id"`
	RateLimitCount    int       `bun:"rate_limit_count"`
	RateLimitDuration int       `bun:"rate_limit_duration"`
	RetryStrategy     string    `bun:"retry_strategy"`
	RetryCount        int       `bun:"retry_count"`
	RetryDuration     int       `bun:"retry_duration"`
	CreatedAt         time.Time `bun:"created_at"`
	UpdatedAt         time.Time `bun:"updated_at"`
}
