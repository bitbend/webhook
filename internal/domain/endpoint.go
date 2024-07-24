package domain

import "time"

type EndpointId string

func (ei EndpointId) String() string {
	return string(ei)
}

type EndpointGroupId string

func (egi EndpointGroupId) String() string {
	return string(egi)
}

type EndpointStatus string

func (es EndpointStatus) String() string {
	return string(es)
}

const (
	EndpointStatusActive   EndpointStatus = "active"
	EndpointStatusInactive EndpointStatus = "inactive"
	EndpointStatusPending  EndpointStatus = "pending"
	EndpointStatusPaused   EndpointStatus = "paused"
)

type Endpoint struct {
	Id          EndpointId
	TenantId    TenantId
	ProjectId   ProjectId
	GroupId     EndpointGroupId
	Name        string
	Url         string
	Description *string
	Status      EndpointStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
