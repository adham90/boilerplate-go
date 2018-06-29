package repository

import (
	"database/sql"
	"fmt"

	e "github.com/adham90/boilerplate/pkg/entity"
)

// Postgres database struct
type Postgres struct{ DB *sql.DB }

// Find return user by id
func (pg *Postgres) Find(id string) (*e.User, error) {
	db := pg.DB

	r := e.User{}

	q := fmt.Sprintf("SELECT id, uuid, github_id FROM users WHERE id = %s limit 1", id)

	err := db.QueryRow(q).Scan(
		&r.ID,
		&r.UUID,
		&r.GithubID,
	)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
