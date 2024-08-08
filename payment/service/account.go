package service

import (
	"context"
	"fmt"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/elangreza14/be-assignment/payment/model"
)

type (
	accountRepo interface {
		Create(ctx context.Context, entity model.Account) error
		Get(ctx context.Context, by string, val any) (*model.Account, error)
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
func (as *AccountService) CreateAccount(ctx context.Context, req dto.CreateAccountPayload) error {
	fmt.Println(req.ID)
	account, err := model.NewAccount(req.ID, req.UserID, req.Name, req.CurrencyCode, req.ProductID)
	if err != nil {
		return err
	}

	err = as.AccountRepo.Create(ctx, *account)
	if err != nil {
		return err
	}

	return nil
}
