package entity

import (
	"encoding/json"
	"time"

	"github.com/satori/go.uuid"
)

//User data
type User struct {
	ID        int
	UUID      uuid.UUID
	GithubID  string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

// NewUser Create new user object
func NewUser(githubID string) User {
	return User{UUID: uuid.NewV5(uuid.UUID{}, githubID)}
}
