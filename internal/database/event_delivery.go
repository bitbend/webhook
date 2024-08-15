package database

import (
	"time"
)

type EventDelivery struct {
	Id             string    `db:"id"`
	TenantId       string    `db:"tenant_id"`
	EventId        string    `db:"event_id"`
	EndpointId     string    `db:"endpoint_id"`
	SubscriptionId string    `db:"subscription_id"`
	AttemptedAt    time.Time `db:"attempted_at"`
	Url            string    `db:"url"`
	StatusCode     int       `db:"status_code"`
	Status         string    `db:"status"`
	CreatedAt      time.Time `db:"created_at"`
}
