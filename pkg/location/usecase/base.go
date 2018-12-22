package usecase

import (
	"time"

	"github.com/adham90/boilerplate/pkg/location/repository"
)

type Datastore struct {
	DB repository.Postgres
	// Inmem repository.Redis
	Ctx time.Duration
}
