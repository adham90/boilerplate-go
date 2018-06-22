package service

import (
	e "github.com/adham90/boilerplate/pkg/entity"
)

//Find user
func (s *Service) Find(id string) (*e.User, error) {
	return s.PG.Find(id)
}
