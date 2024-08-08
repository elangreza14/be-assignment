package service

import (
	"context"

	"github.com/elangreza14/be-assignment/payment/dto"
)

type (
	entryRepo interface {
	}
	transferRepo interface {
	}

	PaymentService struct {
		AccountRepo  accountRepo
		EntryRepo    entryRepo
		TransferRepo transferRepo
	}
)

func NewPaymentService(accountRepo accountRepo, EntryRepo entryRepo, TransferRepo transferRepo) *PaymentService {
	return &PaymentService{
		AccountRepo:  accountRepo,
		EntryRepo:    EntryRepo,
		TransferRepo: TransferRepo,
	}
}

func (as *PaymentService) SendPayment(ctx context.Context, req dto.SendPayload) error {

	// todo
	// wrap tx
	// create transfer
	// create entry receiver
	// create entry sender
	// reduce account balance sender
	// add account balance receiver
	return nil
}

func (as *PaymentService) WithdrawPayment(ctx context.Context, req dto.WithdrawPayload) error {

	// todo
	// wrap tx
	// create entry current account
	// reduce account balance sender
	return nil
}

func (as *PaymentService) PaymentList(ctx context.Context, req dto.WithdrawPayload) error {

	// get entries

	return nil
}
