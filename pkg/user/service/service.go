package service

import (
	"database/sql"

	repo "github.com/adham90/boilerplate/pkg/user/repository"
)

// Service struct
type Service struct {
	PG repo.PG
}

// New return new service object
func New(db *sql.DB) Service {
	r := repo.PG{
		DB: db,
	}

	return Service{
		PG: r,
	}
}
