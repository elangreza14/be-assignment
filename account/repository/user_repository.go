package repository

import (
	"github.com/elangreza14/be-assignment/account/model"
)

type userRepository struct {
	db QueryPgx
	*PostgresRepo[model.User]
}

func NewUserRepository(
	dbPool QueryPgx,
) *userRepository {
	return &userRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.User](dbPool),
	}
}
