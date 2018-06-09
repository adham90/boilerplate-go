package user

import e "github.com/adham90/boilerplate/pkg/entity"

//Service service interface
type Service interface {
	Find(id string) (*e.User, error)
	// Register(user *e.User) (e.User, error)
	// Validate(user *e.User) error
	// Auth(user *e.User) error
	// GetRepo() Repository
}
