package repository

import (
	"context"

	"github.com/elangreza14/be-assignment/payment/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

type accountRepository struct {
	db QueryPgx
	*PostgresRepo[model.Account]
}

func NewAccountRepository(
	dbPool QueryPgx,
) *accountRepository {
	return &accountRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Account](dbPool),
	}
}

func (pr *accountRepository) Create(ctx context.Context, req model.Account) error {
	q := `INSERT INTO accounts
			( id, user_id, currency_code, balance, name, product_id)
			VALUES($1,$2,$3,$4,$5,$6 );`
	_, err := pr.db.Exec(ctx, q, req.ID, req.UserID, req.CurrencyCode, req.Balance, req.Name, req.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostgresRepo[T]) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]T, error) {
	q := `SELECT id, user_id, currency_code, balance, "name", product_id, created_at, updated_at
			FROM accounts
			WHERE user_id=$1;`
	v, err := pgxutil.Select(ctx, pr.db, q, []any{userID}, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return v, nil
}
