package domain

import "time"

type SubscriptionId string

func (si SubscriptionId) String() string {
	return string(si)
}

type Subscription struct {
	Id         SubscriptionId
	TenantId   TenantId
	StreamId   StreamId
	EndpointId EndpointId
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
