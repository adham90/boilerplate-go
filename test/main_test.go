package main

import (
	"fmt"
	"testing"

	"github.com/adham90/boilerplate/pkg/entity"
)

type usecase struct{}

func (uc *usecase) FindByID(int) (entity.User, error) {
	user := entity.User{
		ID:   1,
		Name: "adham",
	}
	return user, nil
}

func TestBooksIndex(t *testing.T) {
	uu := usecase{}
	u, _ := uu.FindByID(1)
	fmt.Println(u)
}
