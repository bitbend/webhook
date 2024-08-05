package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TenantOwner struct {
	bun.BaseModel `bun:"table:tenant_owners"`
	Id            string    `bun:"id,pk"`
	TenantId      string    `bun:"tenant_id"`
	UserId        string    `bun:"user_id"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
}
