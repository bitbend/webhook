package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Endpoint struct {
	bun.BaseModel         `bun:"table:endpoints"`
	Id                    string    `bun:"id,pk" db:"id"`
	TenantId              string    `bun:"tenant_id" db:"tenant_id"`
	ProjectId             string    `bun:"project_id" db:"project_id"`
	Name                  string    `bun:"name" db:"name"`
	Description           *string   `bun:"description" db:"description"`
	Url                   string    `bun:"url" db:"url"`
	RateLimitCount        int       `bun:"rate_limit_count" db:"rate_limit_count"`
	RateLimitDuration     int       `bun:"rate_limit_duration" db:"rate_limit_duration"`
	RetryStrategy         string    `bun:"retry_strategy" db:"retry_strategy"`
	RetryCount            int       `bun:"retry_count" db:"retry_count"`
	RetryDuration         int       `bun:"retry_duration" db:"retry_duration"`
	AuthenticationType    string    `bun:"authentication_type" db:"authentication_type"`
	AuthenticationSecrets []byte    `bun:"authentication_secrets" db:"authentication_secrets"`
	Status                string    `bun:"status" db:"status"`
	CreatedAt             time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt             time.Time `bun:"updated_at" db:"updated_at"`
}
