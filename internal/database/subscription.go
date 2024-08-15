package database

import (
	"time"
)

type Subscription struct {
	Id                string    `db:"id"`
	TenantId          string    `db:"tenant_id"`
	StreamId          string    `db:"stream_id"`
	EndpointId        string    `db:"endpoint_id"`
	RateLimitCount    int       `db:"rate_limit_count"`
	RateLimitDuration int       `db:"rate_limit_duration"`
	RetryStrategy     string    `db:"retry_strategy"`
	RetryCount        int       `db:"retry_count"`
	RetryDuration     int       `db:"retry_duration"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
