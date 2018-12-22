package usecase

import (
	"github.com/adham90/boilerplate/pkg/entity"
)

func (x *Datastore) Store(e *entity.User) (*entity.User, error) {
	db := x.DB

	user, err := db.Store(e)
	if err != nil {
		return nil, err
	}

	return user, nil
}
