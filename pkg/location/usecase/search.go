package usecase

import "github.com/adham90/boilerplate/pkg/entity"

// FindByID return user by it's id
func (x *Datastore) FindByID(id int) (*entity.Location, error) {
	return x.DB.GetByID(id)
}
