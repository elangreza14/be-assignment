package service

import (
	"context"
	"errors"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/elangreza14/be-assignment/payment/model"
)

type (
	transferRepo interface {
		GetTransferByAccountID(ctx context.Context, accountID int) ([]model.Transfer, error)
	}

	accountTransferRepo interface {
		TopUpTX(ctx context.Context, req *model.Account, amount int) error
		WithdrawTX(ctx context.Context, req *model.Account, amount int) error
		SendTX(ctx context.Context, fromAccount *model.Account, toAccount *model.Account, amount int) error
	}

	PaymentService struct {
		AccountRepo         accountRepo
		TransferRepo        transferRepo
		accountTransferRepo accountTransferRepo
	}
)

func NewPaymentService(
	accountRepo accountRepo,
	TransferRepo transferRepo,
	accountTransferRepo accountTransferRepo) *PaymentService {
	return &PaymentService{
		AccountRepo:         accountRepo,
		TransferRepo:        TransferRepo,
		accountTransferRepo: accountTransferRepo,
	}
}

func (as *PaymentService) SendPayment(ctx context.Context, req dto.SendPayload) error {
	fromAccount, err := as.AccountRepo.Get(ctx, "id", req.AccountID)
	if err != nil {
		return err
	}

	toAccount, err := as.AccountRepo.Get(ctx, "id", req.ToAccountID)
	if err != nil {
		return err
	}

	if fromAccount.ID == toAccount.ID {
		err = as.accountTransferRepo.TopUpTX(ctx, fromAccount, req.Amount)
		if err != nil {
			return nil
		}

		return nil
	}

	if fromAccount.ProductID != toAccount.ProductID {
		return errors.New("cannot transfer with different product id")
	}

	if fromAccount.CurrencyCode != toAccount.CurrencyCode {
		return errors.New("cannot transfer with different currency")
	}

	err = as.accountTransferRepo.SendTX(ctx, fromAccount, toAccount, req.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (as *PaymentService) WithdrawPayment(ctx context.Context, req dto.WithdrawPayload) error {
	account, err := as.AccountRepo.Get(ctx, "id", req.AccountID)
	if err != nil {
		return err
	}

	err = as.accountTransferRepo.WithdrawTX(ctx, account, req.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (as *PaymentService) PaymentList(ctx context.Context, accountID int) (int, []dto.TransferHistoryResponse, error) {

	account, err := as.AccountRepo.Get(ctx, "id", accountID)
	if err != nil {
		return 0, nil, err
	}

	transfers, err := as.TransferRepo.GetTransferByAccountID(ctx, accountID)
	if err != nil {
		return 0, nil, err
	}
	res := make([]dto.TransferHistoryResponse, 0)
	for _, transfer := range transfers {
		action := ""
		if transfer.FromAccountID == transfer.ToAccountID {
			if transfer.Amount < 0 {
				action = "WITHDRAW"
			} else {
				action = "TOP_UP"
			}
		} else {
			if transfer.FromAccountID == accountID {
				action = "TRANSFER_OUT"
			} else {
				action = "TRANSFER_IN"
			}
		}

		res = append(res, dto.TransferHistoryResponse{
			ID:            transfer.ID,
			FromAccountID: transfer.FromAccountID,
			ToAccountID:   transfer.ToAccountID,
			Amount:        transfer.Amount,
			Action:        action,
			CreatedAt:     transfer.CreatedAt,
		})
	}

	return account.Balance, res, nil
}
