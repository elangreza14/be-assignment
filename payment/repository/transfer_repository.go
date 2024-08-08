package repository

import (
	"context"

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

func (pr *transferRepository) Create(ctx context.Context, req model.Transfer) error {
	q := `INSERT INTO transfers
		( from_account_id, to_account_id, amount)
		VALUES($1,$2,$3);`
	_, err := pr.db.Exec(ctx, q, req.FromAccountID, req.ToAccountID, req.Amount)
	if err != nil {
		return err
	}

	return nil
}
