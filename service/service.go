package service

import (
	"pkhub/sqlite"
)

type Service struct {
	Storage StorageInterface
}

func NewService(db *sqlite.Sqlite) *Service {
	return &Service{
		Storage: db,
	}
}
