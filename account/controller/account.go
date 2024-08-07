package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/elangreza14/be-assignment/account/dto"
	"github.com/elangreza14/be-assignment/account/middleware"
	"github.com/elangreza14/be-assignment/account/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	AccountService interface {
		CreateAccount(ctx context.Context, userID uuid.UUID, name string, req dto.CreateAccountPayload) error
		GetAccounts(ctx context.Context, userID uuid.UUID) (dto.AccountListResponse, error)
	}

	AccountController struct {
		AccountService AccountService
	}
)

func NewAccountController(AccountService AccountService) *AccountController {
	return &AccountController{
		AccountService: AccountService,
	}
}

func (ac *AccountController) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.CreateAccountPayload{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		rawUser, ok := c.Get(middleware.UserMiddlewareKey)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse middleware")))
			return
		}

		user, ok := rawUser.(*model.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse middleware payload")))
			return
		}

		err = ac.AccountService.CreateAccount(c, user.ID, user.Name, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusCreated, dto.NewBaseResponse("created", nil))
	}
}

func (ac *AccountController) GetAccountList() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawUser, ok := c.Get(middleware.UserMiddlewareKey)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse middleware")))
			return
		}

		user, ok := rawUser.(*model.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("cannot parse middleware payload")))
			return
		}

		res, err := ac.AccountService.GetAccounts(c, user.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(res, nil))
	}
}
