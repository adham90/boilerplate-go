package service

import (
	repo "github.com/adham90/boilerplate/pkg/user/repository"
)

// Service struct
type Service struct {
	Postgres repo.Postgres
	Redis    repo.Redis
}
