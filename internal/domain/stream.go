package domain

type StreamId string

func (si StreamId) String() string {
	return string(si)
}

type Stream struct {
	Id        StreamId
	TenantId  TenantId
	ProjectId ProjectId
	EventType EventType
}
