package database

import (
	"github.com/uptrace/bun"
	"time"
)

type Organization struct {
	bun.BaseModel `bun:"table:organizations"`
	Id            string    `bun:"id,pk"`
	TenantId      string    `bun:"tenant_id"`
	OwnerId       string    `bun:"owner_id"`
	Name          string    `bun:"name"`
	SubDomain     string    `bun:"sub_domain"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
}
