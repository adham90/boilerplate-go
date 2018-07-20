package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

//User data
type User struct {
	ID        int        `jsonapi:"primary,users"`
	UUID      uuid.UUID  `sql:",type:uuid" jsonapi:"attr,uuid"`
	GithubID  string     `jsonapi:"attr,github_id"`
	Name      string     `jsonapi:"attr,name"`
	Email     string     `jsonapi:"attr,email"`
	CreatedAt time.Time  `sql:"default:now()" jsonapi:"attr,created_at"`
	UpdatedAt time.Time  `sql:"default:now()" jsonapi:"attr,updated_at"`
	DeletedAt *time.Time `jsonapi:"attr,deleted_at"`
}

// NewUser Create new user object
func NewUser(githubID string) User {
	return User{UUID: uuid.NewV5(uuid.UUID{}, githubID)}
}
