package database

import (
	"time"
)

type Organization struct {
	Id        string    `db:"id"`
	TenantId  string    `db:"tenant_id"`
	OwnerId   string    `db:"owner_id"`
	Name      string    `db:"name"`
	SubDomain string    `db:"sub_domain"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
