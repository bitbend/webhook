package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Project struct {
	bun.BaseModel      `bun:"table:projects"`
	Id                 string    `bun:"id,pk" db:"id"`
	TenantId           string    `bun:"tenant_id" db:"tenant_id"`
	OrganizationId     string    `bun:"organization_id" db:"organization_id"`
	Name               string    `bun:"name" db:"name"`
	Description        *string   `bun:"description" db:"description"`
	DisableEndpoints   bool      `bun:"disable_endpoints" db:"disable_endpoints"`
	RetentionPolicy    string    `bun:"retention_policy" db:"retention_policy"`
	AllowedPayloadSize int       `bun:"allowed_payload_size" db:"allowed_payload_size"`
	RateLimitCount     int       `bun:"rate_limit_count" db:"rate_limit_count"`
	RateLimitDuration  int       `bun:"rate_limit_duration" db:"rate_limit_duration"`
	RetryStrategy      string    `bun:"retry_strategy" db:"retry_strategy"`
	RetryCount         int       `bun:"retry_count" db:"retry_count"`
	RetryDuration      int       `bun:"retry_duration" db:"retry_duration"`
	Status             string    `bun:"status" db:"status"`
	CreatedAt          time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt          time.Time `bun:"updated_at" db:"updated_at"`
}
