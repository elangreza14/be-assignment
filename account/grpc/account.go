package servergrpc

import (
	"context"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
	gen "github.com/elangreza14/be-assignment/gen/go"
	"github.com/google/uuid"
)

type (
	authService interface {
		ProcessToken(ctx context.Context, reqToken string) (*model.User, error)
	}

	accountService interface {
		GetAccounts(ctx context.Context, userID uuid.UUID) (dto.AccountListResponse, error)
	}

	accountServerGrpc struct {
		authService    authService
		accountService accountService
		gen.UnimplementedAccountServer
	}
)

func NewAccountServerGrpc(authService authService, accountService accountService) *accountServerGrpc {
	return &accountServerGrpc{
		authService:                authService,
		accountService:             accountService,
		UnimplementedAccountServer: gen.UnimplementedAccountServer{},
	}
}

func (asg accountServerGrpc) ValidateToken(ctx context.Context, req *gen.ValidateTokenRequest) (*gen.ValidateTokenResponse, error) {
	user, err := asg.authService.ProcessToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	accounts, err := asg.accountService.GetAccounts(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	accountIDs := make([]uint32, 0)
	for _, account := range accounts {
		accountIDs = append(accountIDs, uint32(account.ID))
	}

	res := &gen.ValidateTokenResponse{
		AccountIds: accountIDs,
		UserId:     user.ID.String(),
		Name:       user.Name,
	}

	return res, nil
}
