package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
}

func New() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Redis{
		rdb: rdb,
	}
}

func (r *Redis) Set(key string, value interface{}) (err error) {
	ctx := context.Background()
	err = r.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	log.Print("Set to cache ", key)
	return
}

func (r *Redis) Get(key string) (value string, err error) {
	ctx := context.Background()
	value, err = r.rdb.Get(ctx, key).Result()
	if err != nil {
		return
	}
	log.Print("Get from cache ", key)
	return
}

func (r *Redis) DeleteCache() {
	ctx := context.Background()
	r.rdb.FlushAll(ctx)
}
