package service

import (
	"context"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
	"github.com/google/uuid"
)

type (
	accountRepo interface {
		Create(ctx context.Context, entity model.Account) error
		GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Account, error)
	}

	AccountService struct {
		AccountRepo accountRepo
	}
)

func NewAccountService(accountRepo accountRepo) *AccountService {
	return &AccountService{
		AccountRepo: accountRepo,
	}
}

func (as *AccountService) CreateAccount(ctx context.Context, userID uuid.UUID, name string, req dto.CreateAccountPayload) error {
	account, err := model.NewAccount(userID, name, req.CurrencyCode, req.ProductID)
	if err != nil {
		return err
	}

	err = as.AccountRepo.Create(ctx, *account)
	if err != nil {
		return err
	}

	return nil
}

func (as *AccountService) GetAccounts(ctx context.Context, userID uuid.UUID) (dto.AccountListResponse, error) {

	accounts, err := as.AccountRepo.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]dto.AccountListResponseElement, 0)
	for _, account := range accounts {
		res = append(res, dto.AccountListResponseElement{
			CurrencyCode: account.CurrencyCode,
			Balance:      account.Balance,
			Name:         account.Name,
			Status:       account.Status,
			ProductID:    account.ProductID,
		})
	}

	return res, nil
}
