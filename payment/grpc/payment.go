package servergrpc

import (
	"context"

	gen "github.com/elangreza14/be-assignment/gen/go"
	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/google/uuid"
)

type (
	paymentService interface {
		CreateAccount(ctx context.Context, req dto.CreateAccountPayload) error
	}

	PaymentServerGrpc struct {
		PaymentService paymentService
		gen.UnimplementedPaymentServer
	}
)

func NewPaymentServerGrpc(accountService paymentService) *PaymentServerGrpc {
	return &PaymentServerGrpc{
		PaymentService:             accountService,
		UnimplementedPaymentServer: gen.UnimplementedPaymentServer{},
	}
}

func (asg PaymentServerGrpc) CreateAccount(ctx context.Context, req *gen.CreateAccountRequest) (*gen.CreateAccountReply, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	err = asg.PaymentService.CreateAccount(ctx, dto.CreateAccountPayload{
		ID:           int(req.Id),
		UserID:       userID,
		Name:         req.Name,
		CurrencyCode: req.CurrencyCode,
		ProductID:    int(req.ProductId),
	})

	if err != nil {
		return nil, err
	}

	return &gen.CreateAccountReply{
		Status: "ok",
	}, nil
}
