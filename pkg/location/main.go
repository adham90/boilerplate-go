package location

import (
	"time"

	"github.com/adham90/boilerplate/pkg/location/repository"

	"github.com/go-pg/pg"

	"github.com/adham90/boilerplate/pkg/location/usecase"
	"github.com/go-redis/redis"
)

func New(ctx time.Duration, db *pg.DB, inmem *redis.Client) *usecase.Datastore {
	return &usecase.Datastore{
		DB:    repository.Postgres{DB: db},
		// Inmem: repository.Redis{Client: inmem},
		Ctx:   ctx,
	}
}
