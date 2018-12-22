package repository

import (
	"github.com/adham90/boilerplate/pkg/entity"
	"github.com/go-pg/pg"
)

// Postgres database struct
type Postgres struct{ DB *pg.DB }

// GetByID return user by id
func (p *Postgres) GetByID(id int) (*entity.Location, error) {
	db := p.DB
	location := entity.Location{}

	err := db.Model(&location).Where("id = ?", id).Select()

	if err != nil {
		return nil, err
	}

	return &location, nil
}

// Save return true or false with the error
func (p *Postgres) Save(l entity.Location) (*entity.Location, error) {
	db := p.DB

	err := db.Insert(&l)

	if err != nil {
		return &l, err
	}

	return &l, nil
}
