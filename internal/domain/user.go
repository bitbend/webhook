package domain

import "time"

type UserId string

func (ui UserId) String() string {
	return string(ui)
}

type User struct {
	Id        UserId
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
