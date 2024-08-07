package service

import (
	"context"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
)

type (
	productsRepo interface {
		GetAll(ctx context.Context) ([]model.Product, error)
	}

	productsService struct {
		productsRepo productsRepo
	}
)

func NewProductService(productsRepo productsRepo) *productsService {
	return &productsService{
		productsRepo: productsRepo,
	}
}

func (cs *productsService) ProductsList(ctx context.Context) (dto.ProductListResponse, error) {
	currencies, err := cs.productsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ProductListResponseElement, 0)
	for _, products := range currencies {
		res = append(res, dto.ProductListResponseElement{
			ID:          products.ID,
			Name:        products.Name,
			Description: products.Description,
		})
	}

	return res, nil
}
