package repository

import "github.com/elangreza14/be-assignment/account/model"

type productRepository struct {
	db QueryPgx
	*PostgresRepo[model.Product]
}

func NewProductRepository(
	dbPool QueryPgx,
) *productRepository {
	return &productRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Product](dbPool),
	}
}
