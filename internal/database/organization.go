package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Organization struct {
	bun.BaseModel `bun:"table:organizations"`
	Id            string    `bun:"id,pk" db:"id"`
	TenantId      string    `bun:"tenant_id" db:"tenant_id"`
	OwnerId       string    `bun:"owner_id" db:"owner_id"`
	Name          string    `bun:"name" db:"name"`
	SubDomain     string    `bun:"sub_domain" db:"sub_domain"`
	CreatedAt     time.Time `bun:"created_at" db:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at" db:"updated_at"`
}
