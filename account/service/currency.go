package service

import (
	"context"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
)

type (
	currencyRepo interface {
		GetAll(ctx context.Context) ([]model.Currency, error)
	}

	currencyService struct {
		currencyRepo currencyRepo
	}
)

func NewCurrencyService(currencyRepo currencyRepo) *currencyService {
	return &currencyService{
		currencyRepo: currencyRepo,
	}
}

func (cs *currencyService) CurrencyList(ctx context.Context) (dto.CurrencyListResponse, error) {
	currencies, err := cs.currencyRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.CurrencyListResponseElement, 0)
	for _, currency := range currencies {
		res = append(res, dto.CurrencyListResponseElement{
			Code:        currency.Code,
			Description: currency.Description,
		})
	}

	return res, nil
}
