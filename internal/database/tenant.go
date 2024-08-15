package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Tenant struct {
	bun.BaseModel `bun:"table:tenants"`
	Id            string    `bun:"id,pk" db:"id"`
	Name          string    `bun:"name" db:"name"`
	Domain        string    `bun:"domain" db:"domain"`
	CreatedAt     time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at" db:"updated_at"`
}
