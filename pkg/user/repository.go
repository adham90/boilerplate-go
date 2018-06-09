// Package user provides connection between the database and entitys
package user

import e "github.com/adham90/boilerplate/pkg/entity"

//Repository repository interface
type Repository interface {
	Find(id string) (*e.User, error)
	// FindByEmail(email string) (*e.User, error)
	// FindAll() ([]*e.User, error)
	// Update(user *e.User) error
	// Store(user *e.User) (*e.User, error)
}
