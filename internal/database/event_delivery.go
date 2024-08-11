package database

import (
	"github.com/uptrace/bun"
	"time"
)

type EventDelivery struct {
	bun.BaseModel  `bun:"table:event_deliveries"`
	Id             string    `bun:"id,pk"`
	TenantId       string    `bun:"tenant_id"`
	EventId        string    `bun:"event_id"`
	EndpointId     string    `bun:"endpoint_id"`
	SubscriptionId string    `bun:"subscription_id"`
	AttemptedAt    time.Time `bun:"attempted_at"`
	Url            string    `bun:"url"`
	StatusCode     int       `bun:"status_code"`
	Status         string    `bun:"status"`
	CreatedAt      time.Time `bun:"created_at"`
}
