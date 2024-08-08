package servergrpc

import (
	"context"

	genaccount "github.com/elangreza14/be-assignment/gen/go"
	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/google/uuid"
)

type (
	AccountService interface {
		CreateAccount(ctx context.Context, req dto.CreateAccountPayload) error
	}

	AccountServerGrpc struct {
		AccountService AccountService
		genaccount.UnimplementedPaymentServer
	}
)

func NewAccountServerGrpc(accountService AccountService) *AccountServerGrpc {
	return &AccountServerGrpc{
		AccountService:             accountService,
		UnimplementedPaymentServer: genaccount.UnimplementedPaymentServer{},
	}
}

func (asg AccountServerGrpc) CreateAccount(ctx context.Context, req *genaccount.AccountRequest) (*genaccount.AccountReply, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	err = asg.AccountService.CreateAccount(ctx, dto.CreateAccountPayload{
		UserID:       userID,
		Name:         req.Name,
		CurrencyCode: req.CurrencyCode,
		ProductID:    int(req.ProductId),
	})

	if err != nil {
		return nil, err
	}

	return &genaccount.AccountReply{
		Status: "ok",
	}, nil
}
