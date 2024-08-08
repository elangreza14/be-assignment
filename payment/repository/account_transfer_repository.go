package repository

import (
	"context"

	"github.com/elangreza14/be-assignment/payment/model"
)

type AccountTransferRepository struct {
	txRepo *PostgresTransactionRepo
}

func NewAccountTransferRepository(
	tx PgxTXer,
) *AccountTransferRepository {
	return &AccountTransferRepository{
		txRepo: NewPostgresTransactionRepo(tx),
	}
}

func (ur *AccountTransferRepository) WithdrawTX(ctx context.Context, account *model.Account, amount int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		accountRepo := NewAccountRepository(tx)
		var err error
		account.Balance = account.Balance - amount
		if err = accountRepo.Edit(ctx, *account); err != nil {
			return err
		}

		transferRepo := NewTransferRepository(tx)

		if err = transferRepo.Create(ctx, model.Transfer{
			ToAccountID:   account.ID,
			FromAccountID: account.ID,
			Amount:        -amount,
		}); err != nil {
			return err
		}

		return nil
	})
}
