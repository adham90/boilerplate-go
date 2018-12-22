package repository

import (
	e "github.com/adham90/boilerplate/pkg/entity"
	"github.com/go-pg/pg"
)

// Postgres database struct
type Postgres struct{ DB *pg.DB }

// GetByID return user by id
func (p *Postgres) GetByID(id int) (*e.User, error) {
	db := p.DB
	user := e.User{}

	err := db.Model(&user).Where("id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save return true or false with the error
func (p *Postgres) Store(u *e.User) (*e.User, error) {
	db := p.DB

	err := db.Insert(&u)

	if err != nil {
		return u, err
	}

	return u, nil
}
