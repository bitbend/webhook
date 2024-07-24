package domain

type StreamId string

func (si StreamId) String() string {
	return string(si)
}

type EventType string

func (et EventType) String() string {
	return string(et)
}

type Stream struct {
	Id        StreamId
	TenantId  TenantId
	ProjectId ProjectId
	EventType EventType
}
