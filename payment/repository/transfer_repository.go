package repository

import (
	"context"

	"github.com/elangreza14/be-assignment/payment/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
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

func (pr *transferRepository) GetTransferByAccountID(ctx context.Context, accountID int) ([]model.Transfer, error) {
	q := `SELECT id, from_account_id, to_account_id, amount, created_at, updated_at
			FROM transfers
			WHERE from_account_id = $1 or to_account_id=$2 order by created_at DESC`
	v, err := pgxutil.Select(ctx, pr.db, q, []any{accountID, accountID}, pgx.RowToStructByNameLax[model.Transfer])
	if err != nil {
		return nil, err
	}
	return v, nil
}
