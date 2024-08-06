package domain

import (
	"time"
)

type EndpointId string

func (ei EndpointId) String() string {
	return string(ei)
}

type EndpointStatus string

func (es EndpointStatus) String() string {
	return string(es)
}

const (
	EndpointStatusActive   EndpointStatus = "active"
	EndpointStatusInactive EndpointStatus = "inactive"
	EndpointStatusPaused   EndpointStatus = "paused"
)

type EndpointConfiguration struct {
	RateLimitCount    int
	RateLimitDuration time.Duration
	RetryStrategy     string
	RetryCount        int
	RetryDuration     time.Duration
}

type Endpoint struct {
	Id            EndpointId
	TenantId      TenantId
	ProjectId     ProjectId
	Name          string
	Description   *string
	Url           string
	Configuration *EndpointConfiguration
	Status        EndpointStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
