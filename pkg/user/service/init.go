package service

import (
	"github.com/adham90/boilerplate/pkg/user"
)

// Service struct
type service struct {
	repo user.Repository
}

//NewService create new service
func NewService(r user.Repository) user.Service {
	return &service{
		repo: r,
	}
}
