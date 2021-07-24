package configs

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Ctx variable
var Ctx = context.Background()

// RedisConnection function
func RedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Env.RedisHost + ":" + Env.RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
