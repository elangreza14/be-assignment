package repository

import (
	"github.com/elangreza14/be-assignment/payment/model"
)

type transferRepository struct {
	db QueryPgx
	*PostgresRepo[model.Transfer]
}

func NewTransferRepository(
	dbPool QueryPgx,
) *transferRepository {
	return &transferRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Transfer](dbPool),
	}
}
