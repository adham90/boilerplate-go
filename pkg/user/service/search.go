package service

import (
	e "github.com/adham90/boilerplate/pkg/entity"
)

// FindByID return user by it's id
func (s *Service) FindByID(id string) (*e.User, error) {
	return s.Postgres.Find(id)
}
