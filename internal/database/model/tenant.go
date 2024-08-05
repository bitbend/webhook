package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Tenant struct {
	bun.BaseModel `bun:"table:tenants"`
	Id            string    `bun:"id,pk"`
	Name          string    `bun:"name"`
	Domain        string    `bun:"domain"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
}
