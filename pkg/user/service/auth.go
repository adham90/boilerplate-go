package service

import (
	e "github.com/adham90/boilerplate/pkg/entity"
)

// Register new user
func (s *service) Register(user *e.User) (*e.User, error) {
	return s.repo.Find("test")
}
