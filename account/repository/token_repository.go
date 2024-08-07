package repository

import (
	"github.com/elangreza14/be-assignment/account/model"
)

type tokenRepository struct {
	db QueryPgx
	*PostgresRepo[model.Token]
}

func NewTokenRepository(
	dbPool QueryPgx,
) *tokenRepository {
	return &tokenRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Token](dbPool),
	}
}
