package database

import (
	"time"
)

type TenantOwner struct {
	Id        string    `db:"id"`
	TenantId  string    `db:"tenant_id"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
