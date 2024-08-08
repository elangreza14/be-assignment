package controller

import (
	"net/http"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/gin-gonic/gin"
)

type (
	PaymentService interface {
		// CreatePayment(ctx context.Context, userID uuid.UUID, name string, req dto.CreatePaymentPayload) error
		// GetPayments(ctx context.Context, userID uuid.UUID) (dto.PaymentListResponse, error)
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

		c.JSON(http.StatusOK, dto.NewBaseResponse("success", nil))
	}
}
