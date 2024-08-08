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

func (ur *AccountTransferRepository) TopUpTX(ctx context.Context, account *model.Account, amount int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		accountRepo := NewAccountRepository(tx)
		var err error
		account.Balance = account.Balance + amount
		if err = accountRepo.Edit(ctx, *account); err != nil {
			return err
		}

		transferRepo := NewTransferRepository(tx)

		if err = transferRepo.Create(ctx, model.Transfer{
			ToAccountID:   account.ID,
			FromAccountID: account.ID,
			Amount:        +amount,
		}); err != nil {
			return err
		}

		return nil
	})
}

func (ur *AccountTransferRepository) SendTX(ctx context.Context, fromAccount *model.Account, toAccount *model.Account, amount int) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		transferRepo := NewTransferRepository(tx)

		if err := transferRepo.Create(ctx, model.Transfer{
			ToAccountID:   toAccount.ID,
			FromAccountID: fromAccount.ID,
			Amount:        -amount,
		}); err != nil {
			return err
		}

		accountRepo := NewAccountRepository(tx)

		fromAccount.Balance = fromAccount.Balance - amount
		if err := accountRepo.Edit(ctx, *fromAccount); err != nil {
			return err
		}

		toAccount.Balance = toAccount.Balance + amount
		if err := accountRepo.Edit(ctx, *toAccount); err != nil {
			return err
		}

		return nil
	})
}
