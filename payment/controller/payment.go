package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/elangreza14/be-assignment/payment/middleware"
	"github.com/gin-gonic/gin"
)

type (
	PaymentService interface {
		SendPayment(ctx context.Context, req dto.SendPayload) error
		WithdrawPayment(ctx context.Context, req dto.WithdrawPayload) error
	}

	PaymentController struct {
		PaymentService PaymentService
	}
)

func NewPaymentController(PaymentService PaymentService) *PaymentController {
	return &PaymentController{
		PaymentService: PaymentService,
	}
}

func (ac *PaymentController) SendPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.SendPayload{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse("success", nil))
	}
}

func (ac *PaymentController) WithdrawPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.WithdrawPayload{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		rawAccountIDs, ok := c.Get(middleware.UserAccountIDsMiddlewareKey)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse account id")))
			return
		}

		accountIDs, ok := rawAccountIDs.([]uint32)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse account id")))
			return
		}

		found := false
		for _, accountID := range accountIDs {
			if accountID == uint32(req.AccountID) {
				found = true
			}
		}

		if !found {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot find account id")))
			return
		}

		err = ac.PaymentService.WithdrawPayment(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		fmt.Println("cel")

		c.JSON(http.StatusOK, dto.NewBaseResponse("success", nil))
	}
}
