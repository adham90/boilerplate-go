package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

//User data
type User struct {
	ID        uint       `sql:"id"`
	UUID      uuid.UUID  `sql:"uuid"`
	GithubID  string     `sql:"github_id"`
	Name      string     `sql:"name"`
	Email     string     `sql:"email"`
	CreatedAt time.Time  `sql:"created_at"`
	UpdatedAt time.Time  `sql:"updated_at"`
	DeletedAt *time.Time `sql:"deleted_at"`
}

// NewUser Create new user object
func NewUser(githubID string) User {
	return User{UUID: uuid.NewV5(uuid.UUID{}, githubID)}
}
