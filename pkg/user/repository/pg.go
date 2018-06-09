package repository

import (
	e "github.com/adham90/boilerplate/pkg/entity"
	"github.com/adham90/boilerplate/pkg/user"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPGRepository(database *gorm.DB) user.Repository {
	return &repository{
		db: database,
	}
}

func (r *repository) Find(id string) (*e.User, error) {
	result := e.User{}

	err := r.db.Find(&result, id).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
