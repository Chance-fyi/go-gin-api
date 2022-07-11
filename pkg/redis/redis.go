package redis

import (
	"github.com/go-redis/redis/v9"
)

var Redis *redis.Client

func NewClient(address string, password string, db int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
