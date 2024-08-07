package repository

import (
	"github.com/elangreza14/be-assignment/account/model"
)

type currencyRepository struct {
	db QueryPgx
	*PostgresRepo[model.Currency]
}

func NewCurrencyRepository(
	dbPool QueryPgx,
) *currencyRepository {
	return &currencyRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Currency](dbPool),
	}
}
