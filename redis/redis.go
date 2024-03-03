package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
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

func (r *Redis) SetFoo() {
	ctx := context.Background()
	err := r.rdb.Set(ctx, "foo", "bar", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
}
