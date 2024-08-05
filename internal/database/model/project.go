package model

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type Project struct {
	bun.BaseModel      `bun:"table:projects"`
	Id                 string         `bun:"id,pk"`
	TenantId           string         `bun:"tenant_id"`
	OrganizationId     string         `bun:"organization_id"`
	Name               string         `bun:"name"`
	Description        sql.NullString `bun:"description"`
	DisableEndpoints   bool           `bun:"disable_endpoints"`
	RetentionPolicy    string         `bun:"retention_policy"`
	AllowedPayloadSize int            `bun:"allowed_payload_size"`
	RateLimitCount     int            `bun:"rate_limit_count"`
	RateLimitDuration  int            `bun:"rate_limit_duration"`
	RetryStrategy      string         `bun:"retry_strategy"`
	RetryCount         int            `bun:"retry_count"`
	RetryDuration      int            `bun:"retry_duration"`
	Status             string         `bun:"status"`
	CreatedAt          time.Time      `bun:"created_at"`
	UpdatedAt          time.Time      `bun:"updated_at"`
}
