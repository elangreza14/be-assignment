package service

import (
	"context"
	"fmt"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/elangreza14/be-assignment/payment/model"
)

type (
	entryRepo interface {
	}

	transferRepo interface {
	}

	accountTransferRepo interface {
		WithdrawTX(ctx context.Context, req *model.Account, amount int) error
	}

	PaymentService struct {
		AccountRepo         accountRepo
		EntryRepo           entryRepo
		TransferRepo        transferRepo
		accountTransferRepo accountTransferRepo
	}
)

func NewPaymentService(
	accountRepo accountRepo,
	EntryRepo entryRepo,
	TransferRepo transferRepo,
	accountTransferRepo accountTransferRepo) *PaymentService {
	return &PaymentService{
		AccountRepo:         accountRepo,
		EntryRepo:           EntryRepo,
		TransferRepo:        TransferRepo,
		accountTransferRepo: accountTransferRepo,
	}
}

func (as *PaymentService) SendPayment(ctx context.Context, req dto.SendPayload) error {
	// todo
	// wrap tx
	// create transfer
	// reduce account balance sender
	// add account balance receiver
	return nil
}

func (as *PaymentService) WithdrawPayment(ctx context.Context, req dto.WithdrawPayload) error {

	// todo
	// wrap tx
	// create entry current account
	// reduce account balance sender

	account, err := as.AccountRepo.Get(ctx, "id", req.AccountID)
	if err != nil {
		return err
	}

	fmt.Println(req.Amount, req.AccountID)

	err = as.accountTransferRepo.WithdrawTX(ctx, account, req.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (as *PaymentService) PaymentList(ctx context.Context, req dto.WithdrawPayload) error {

	// get entries

	return nil
}
