package database

import (
	"time"
)

type Project struct {
	Id                 string    `db:"id"`
	TenantId           string    `db:"tenant_id"`
	OrganizationId     string    `db:"organization_id"`
	Name               string    `db:"name"`
	Description        *string   `db:"description"`
	DisableEndpoints   bool      `db:"disable_endpoints"`
	RetentionPolicy    string    `db:"retention_policy"`
	AllowedPayloadSize int       `db:"allowed_payload_size"`
	RateLimitCount     int       `db:"rate_limit_count"`
	RateLimitDuration  int       `db:"rate_limit_duration"`
	RetryStrategy      string    `db:"retry_strategy"`
	RetryCount         int       `db:"retry_count"`
	RetryDuration      int       `db:"retry_duration"`
	Status             string    `db:"status"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
