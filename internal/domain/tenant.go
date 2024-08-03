package domain

import "time"

type TenantId string

func (ti TenantId) String() string {
	return string(ti)
}

type Tenant struct {
	Id        TenantId
	UserId    UserId
	Domain    string
	CreatedAt time.Time
	UpdatedAt time.Time
}