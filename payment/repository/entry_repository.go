package repository

import (
	"github.com/elangreza14/be-assignment/payment/model"
)

type EntryRepository struct {
	db QueryPgx
	*PostgresRepo[model.Entry]
}

func NewEntryRepository(
	dbPool QueryPgx,
) *EntryRepository {
	return &EntryRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Entry](dbPool),
	}
}
