package domain

import "time"

type SubscriberId string

func (si SubscriberId) String() string {
	return string(si)
}

type Subscriber struct {
	Id           SubscriberId
	TenantId     TenantId
	EventId      EventId
	EndpointId   EndpointId
	SubscriberId SubscriberId
	AttemptedAt  time.Time
	Url          string
	StatusCode   int
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
