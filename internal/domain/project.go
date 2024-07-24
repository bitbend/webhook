package domain

import "time"

type ProjectId string

func (pi ProjectId) String() string {
	return string(pi)
}

type Project struct {
	Id             ProjectId
	TenantId       TenantId
	OrganizationId OrganizationId
	Name           string
	Description    *string
	ApiKey         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
