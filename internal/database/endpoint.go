package database

import (
	"database/sql"
	"time"
)

type Endpoint struct {
	Id                    string         `db:"id"`
	TenantId              string         `db:"tenant_id"`
	ProjectId             string         `db:"project_id"`
	Name                  string         `db:"name"`
	Description           sql.NullString `db:"description"`
	Url                   string         `db:"url"`
	RateLimitCount        int            `db:"rate_limit_count"`
	RateLimitDuration     int            `db:"rate_limit_duration"`
	RetryStrategy         string         `db:"retry_strategy"`
	RetryCount            int            `db:"retry_count"`
	RetryDuration         int            `db:"retry_duration"`
	AuthenticationType    string         `db:"authentication_type"`
	AuthenticationSecrets []byte         `db:"authentication_secrets"`
	Status                string         `db:"status"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             time.Time      `db:"updated_at"`
}
