package domain

import "time"

type OrganizationId string

func (oi OrganizationId) String() string {
	return string(oi)
}

type Organization struct {
	Id        OrganizationId
	TenantId  TenantId
	OwnerId   UserId
	Name      string
	Version   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
