package database

import (
	"github.com/uptrace/bun"
	"time"
)

type EventDelivery struct {
	bun.BaseModel  `bun:"table:event_deliveries"`
	Id             string    `bun:"id,pk" db:"id"`
	TenantId       string    `bun:"tenant_id" db:"tenant_id"`
	EventId        string    `bun:"event_id" db:"event_id"`
	EndpointId     string    `bun:"endpoint_id" db:"endpoint_id"`
	SubscriptionId string    `bun:"subscription_id" db:"subscription_id"`
	AttemptedAt    time.Time `bun:"attempted_at" db:"attempted_at"`
	Url            string    `bun:"url" db:"url"`
	StatusCode     int       `bun:"status_code" db:"status_code"`
	Status         string    `bun:"status" db:"status"`
	CreatedAt      time.Time `bun:"created_at" db:"created_at"`
}
