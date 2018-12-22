package utils

import "github.com/go-redis/redis"

type Inmem struct {
	Username, Password string
	Addr               string
	DB                 int
	LogMode            bool
}

// "localhost:6379"
func (i *Inmem) Open() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     i.Addr,
		Password: i.Password, // no password set
		DB:       i.DB,       // use default DB
	})
}
