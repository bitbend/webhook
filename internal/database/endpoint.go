package database

import (
	"database/sql"
	"encoding/json"
	"github.com/uptrace/bun"
	"time"
)

type Endpoint struct {
	bun.BaseModel         `bun:"table:endpoints"`
	Id                    string          `bun:"id,pk"`
	TenantId              string          `bun:"tenant_id"`
	ProjectId             string          `bun:"project_id"`
	Name                  string          `bun:"name"`
	Description           sql.NullString  `bun:"description"`
	Url                   string          `bun:"url"`
	RateLimitCount        int             `bun:"rate_limit_count"`
	RateLimitDuration     int             `bun:"rate_limit_duration"`
	RetryStrategy         string          `bun:"retry_strategy"`
	RetryCount            int             `bun:"retry_count"`
	RetryDuration         int             `bun:"retry_duration"`
	AuthenticationType    string          `bun:"authentication_type"`
	AuthenticationSecrets json.RawMessage `bun:"authentication_secrets"`
	Status                string          `bun:"status"`
	CreatedAt             time.Time       `bun:"created_at"`
	UpdatedAt             time.Time       `bun:"updated_at"`
}
