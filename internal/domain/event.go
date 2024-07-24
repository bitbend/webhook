package domain

import (
	"encoding/json"
	"time"
)

type EventId string

func (ei EventId) String() string {
	return string(ei)
}

type Event struct {
	Id        EventId
	TenantId  TenantId
	StreamId  StreamId
	Data      json.RawMessage
	CreatedAt time.Time
	UpdatedAt time.Time
}
