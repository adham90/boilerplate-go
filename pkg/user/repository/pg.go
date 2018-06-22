package repository

import (
	"database/sql"
	"fmt"

	e "github.com/adham90/boilerplate/pkg/entity"
)

// PG postgres database struct
type PG struct {
	DB *sql.DB
}

// const fiels = [
// 	id,
// 	uuid,
// 	github_id,
// ]

func (pg *PG) Find(id string) (*e.User, error) {
	db := pg.DB

	r := e.User{}

	q := fmt.Sprintf("SELECT id, uuid, github_id FROM users WHERE id = %s limit 1", id)

	err := db.QueryRow(q).Scan(
		&r.ID,
		&r.UUID,
		&r.GithubID,
	)

	if err != nil {

	}

	fmt.Println(r)
	return &r, nil
}
