package usecase

import "github.com/adham90/boilerplate/pkg/entity"

// FindByID return user by it's id
func (x *Datastore) FindByID(id int) (*entity.User, error) {
	return x.DB.GetByID(id)
}

// func (x *UserStore) CreateUser(e *entity.User) (*entity.User, error) {
// 	// return x.db.Save(e)
// }
