package domain

import (
	"encoding/json"
	"time"
)

type EventId string

func (ei EventId) String() string {
	return string(ei)
}

type EventType string

func (et EventType) String() string {
	return string(et)
}

type Event struct {
	Id        EventId
	TenantId  TenantId
	StreamId  StreamId
	GroupId   *EndpointGroupId
	Data      json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}
