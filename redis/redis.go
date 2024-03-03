package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func New() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Redis{
		rdb: rdb,
	}
}

func (r *Redis) Set(key string, value interface{}) {
	ctx := context.Background()
	err := r.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *Redis) Get(key string) (value string) {
	ctx := context.Background()
	value, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	} else {
		return
	}
}
