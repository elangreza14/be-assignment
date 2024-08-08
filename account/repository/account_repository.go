package repository

import (
	"context"

	"github.com/elangreza14/be-assignment/account/model"
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

func (pr *accountRepository) Create(ctx context.Context, req model.Account) (int, error) {
	q := `INSERT INTO accounts
			( user_id, currency_code, name, product_id)
			VALUES($1,$2,$3,$4 ) RETURNING id;`
	r := pr.db.QueryRow(ctx, q, req.UserID, req.CurrencyCode, req.Name, req.ProductID)
	var ID int
	err := r.Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (pr *PostgresRepo[T]) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]T, error) {
	q := `SELECT id, user_id, currency_code, name, product_id, created_at, updated_at
			FROM accounts
			WHERE user_id=$1;`
	v, err := pgxutil.Select(ctx, pr.db, q, []any{userID}, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return v, nil
}
