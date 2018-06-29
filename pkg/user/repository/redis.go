package repository

import (
	"github.com/go-redis/redis"
)

// Redis database struct
type Redis struct{ Client *redis.Client }

// Get val by key
func (r *Redis) Get(key string) string {
	val, _ := r.Client.Get(key).Result()
	return val
}
