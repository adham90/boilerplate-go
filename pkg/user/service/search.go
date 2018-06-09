package service

import (
	e "github.com/adham90/boilerplate/pkg/entity"
)

//Find a bookmark
func (s *service) Find(id string) (*e.User, error) {
	return s.repo.Find(id)
}
