package servergrpc

import (
	"context"

	gen "github.com/elangreza14/be-assignment/gen/go"
	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	accountService interface {
		CreateAccount(ctx context.Context, req dto.CreateAccountPayload) error
	}

	paymentService interface {
		PaymentList(ctx context.Context, accountID int) (int, []dto.TransferHistoryResponse, error)
	}

	PaymentServerGrpc struct {
		accountService accountService
		paymentService paymentService
		gen.UnimplementedPaymentServer
	}
)

func NewPaymentServerGrpc(accountService accountService, paymentService paymentService) *PaymentServerGrpc {
	return &PaymentServerGrpc{
		accountService:             accountService,
		paymentService:             paymentService,
		UnimplementedPaymentServer: gen.UnimplementedPaymentServer{},
	}
}

func (psg PaymentServerGrpc) CreateAccount(ctx context.Context, req *gen.CreateAccountRequest) (*gen.CreateAccountReply, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	err = psg.accountService.CreateAccount(ctx, dto.CreateAccountPayload{
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

func (psg PaymentServerGrpc) GetAccountHistory(ctx context.Context, req *gen.GetAccountHistoryRequest) (*gen.GetAccountHistoriesReply, error) {
	balance, transfers, err := psg.paymentService.PaymentList(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	res := []*gen.GetAccountHistoryReply{}
	for _, payment := range transfers {
		res = append(res, &gen.GetAccountHistoryReply{
			Id:            uint32(payment.ID),
			ToAccountId:   uint32(payment.ToAccountID),
			FromAccountId: uint32(payment.FromAccountID),
			Amount:        int32(payment.Amount),
			Action:        payment.Action,
			CreatedAt:     timestamppb.New(payment.CreatedAt),
		})
	}
	return &gen.GetAccountHistoriesReply{
		Balance:   int32(balance),
		Histories: res,
	}, nil
}
