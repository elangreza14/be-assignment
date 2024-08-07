package controller

import (
	"context"
	"net/http"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/gin-gonic/gin"
)

type (
	currencyService interface {
		CurrencyList(ctx context.Context) (dto.CurrencyListResponse, error)
	}

	CurrencyController struct {
		currencyService
	}
)

func NewCurrencyController(currencyService currencyService) *CurrencyController {
	return &CurrencyController{currencyService}
}

func (cc *CurrencyController) CurrencyList() gin.HandlerFunc {
	return func(c *gin.Context) {
		currencies, err := cc.currencyService.CurrencyList(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(currencies, nil))
	}
}
