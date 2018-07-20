package entity

import "time"

type Location struct {
	ID   int
	Name string
	Lat  string
	Long string

	CreatedAt time.Time  `sql:"default:now()" jsonapi:"attr,created_at"`
	UpdatedAt time.Time  `sql:"default:now()" jsonapi:"attr,updated_at"`
	DeletedAt *time.Time `jsonapi:"attr,deleted_at"`
}
