package service

import (
	"context"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
	gen "github.com/elangreza14/be-assignment/gen/go"
	genaccount "github.com/elangreza14/be-assignment/gen/go"
	"github.com/google/uuid"
)

type (
	accountRepo interface {
		Create(ctx context.Context, entity model.Account) (int, error)
		GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Account, error)
	}

	AccountService struct {
		AccountRepo   accountRepo
		paymentClient genaccount.PaymentClient
	}
)

func NewAccountService(accountRepo accountRepo, paymentClient genaccount.PaymentClient) *AccountService {
	return &AccountService{
		AccountRepo:   accountRepo,
		paymentClient: paymentClient,
	}
}

func (as *AccountService) CreateAccount(ctx context.Context, userID uuid.UUID, name string, req dto.CreateAccountPayload) error {
	account, err := model.NewAccount(userID, name, req.CurrencyCode, req.ProductID)
	if err != nil {
		return err
	}

	id, err := as.AccountRepo.Create(ctx, *account)
	if err != nil {
		return err
	}

	_, err = as.paymentClient.CreateAccount(ctx, &genaccount.CreateAccountRequest{
		Id:           uint32(id),
		UserId:       userID.String(),
		CurrencyCode: req.CurrencyCode,
		ProductId:    uint32(req.ProductID),
		Balance:      0,
		Name:         name,
	})

	if err != nil {
		return err
	}

	return nil
}

func (as *AccountService) GetAccounts(ctx context.Context, userID uuid.UUID) (dto.AccountListResponse, error) {

	accounts, err := as.AccountRepo.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]dto.AccountListResponseElement, 0)
	for _, account := range accounts {
		res = append(res, dto.AccountListResponseElement{
			CurrencyCode: account.CurrencyCode,
			Name:         account.Name,
			ProductID:    account.ProductID,
			ID:           account.ID,
		})
	}

	return res, nil
}

func (as *AccountService) GetAccountHistories(ctx context.Context, userID uuid.UUID, accountID int) (*dto.TransferHistoryResponse, error) {
	// todo check userid
	transfers, err := as.paymentClient.GetAccountHistory(ctx, &gen.GetAccountHistoryRequest{
		Id: uint32(accountID),
	})
	if err != nil {
		return nil, err
	}

	res := make([]dto.TransferHistoryResponseElement, 0)
	for _, transfer := range transfers.Histories {
		res = append(res,
			dto.TransferHistoryResponseElement{
				ID:            int(transfer.Id),
				FromAccountID: int(transfer.FromAccountId),
				ToAccountID:   int(transfer.ToAccountId),
				Amount:        int(transfer.Amount),
				Action:        transfer.Action,
				CreatedAt:     transfer.CreatedAt.AsTime(),
			})
	}

	return &dto.TransferHistoryResponse{
		Histories: res,
		Balance:   int(transfers.Balance),
	}, nil
}
