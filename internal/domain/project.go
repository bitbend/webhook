package domain

import "time"

type ProjectId string

func (pi ProjectId) String() string {
	return string(pi)
}

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusInactive ProjectStatus = "inactive"
)

type ProjectConfiguration struct {
	DisableEndpoints   bool
	RetentionPolicy    string
	AllowedPayloadSize int
	RateLimitCount     int
	RateLimitDuration  time.Duration
	RetryStrategy      string
	RetryCount         int
	RetryDuration      time.Duration
}

type Project struct {
	Id             ProjectId
	TenantId       TenantId
	OrganizationId OrganizationId
	Name           string
	Description    *string
	Configuration  *ProjectConfiguration
	Status         ProjectStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
