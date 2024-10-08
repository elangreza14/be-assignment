package service

import (
	"context"
	"fmt"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/model"
	"github.com/jackc/pgx/v5"
)

type (
	userRepo interface {
		Create(ctx context.Context, entities ...model.User) error
		Get(ctx context.Context, by string, val any) (*model.User, error)
	}

	tokenRepo interface {
		Create(ctx context.Context, entities ...model.Token) error
		Get(ctx context.Context, by string, val any) (*model.Token, error)
	}

	AuthService struct {
		UserRepo  userRepo
		TokenRepo tokenRepo
	}
)

func NewAuthService(userRepo userRepo, tokenRepo tokenRepo) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
	}
}

func (as *AuthService) RegisterUser(ctx context.Context, req dto.RegisterPayload) error {
	user, err := model.NewUser(req.Email, req.Password, req.Name)
	if err != nil {
		return err
	}

	err = as.UserRepo.Create(ctx, *user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) LoginUser(ctx context.Context, req dto.LoginPayload) (string, error) {
	user, err := as.UserRepo.Get(ctx, "email", req.Email)
	if err != nil {
		return "", err
	}

	ok := user.IsPasswordValid(req.Password)
	if !ok {
		return "", err
	}

	token, err := as.TokenRepo.Get(ctx, "user_id", user.ID)
	if err != nil && err != pgx.ErrNoRows {
		fmt.Println("1")
		return "", err
	}

	if token != nil {
		_, err = token.IsTokenValid([]byte("test"))
		if err == nil {
			return token.Token, nil
		}
	}

	token, err = model.NewToken([]byte("test"), user.ID, "LOGIN")
	if err != nil {
		return "", err
	}

	err = as.TokenRepo.Create(ctx, *token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}

func (as *AuthService) ProcessToken(ctx context.Context, reqToken string) (*model.User, error) {
	token := &model.Token{Token: reqToken}

	id, err := token.IsTokenValid([]byte("test"))
	if err != nil {
		return nil, err
	}

	token, err = as.TokenRepo.Get(ctx, "id", id)
	if err != nil {
		return nil, err
	}

	return as.UserRepo.Get(ctx, "id", token.UserID)
}
