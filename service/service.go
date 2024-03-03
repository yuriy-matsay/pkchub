package service

import (
	"pkhub/redis"
	"pkhub/sqlite"
)

type Service struct {
	Storage StorageInterface
	Cache   CacheInteface
}

func NewService(db *sqlite.Sqlite, ch *redis.Redis) *Service {
	return &Service{
		Storage: db,
		Cache:   ch,
	}
}
