package storage

import (
	"registiration/storage/postgres"
	repoi "registiration/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	UserRepo() repoi.UserRepoI
}

type storage struct {
	userRepo repoi.UserRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storage) UserRepo() repoi.UserRepoI {
	return s.userRepo
}
