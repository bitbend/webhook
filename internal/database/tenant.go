package database

import (
	"time"
)

type Tenant struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Domain    string    `db:"domain"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
