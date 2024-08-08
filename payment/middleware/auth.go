package middleware

import (
	"errors"
	"net/http"
	"strings"

	gen "github.com/elangreza14/be-assignment/gen/go"
	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/gin-gonic/gin"
)

type (
	AuthMiddleware struct {
		accountClient gen.AccountClient
	}
)

func NewAuthMiddleware(accountClient gen.AccountClient) *AuthMiddleware {
	return &AuthMiddleware{
		accountClient: accountClient,
	}
}

const (
	UserNameMiddlewareKey       = "UserNameMiddlewareKey"
	UserIDMiddlewareKey         = "UserIDMiddlewareKey"
	UserAccountIDsMiddlewareKey = "UserAccountIDsMiddlewareKey"
)

func (am *AuthMiddleware) MustAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawAuthorization := c.Request.Header["Authorization"]
		if len(rawAuthorization) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("token not valid")))
			return
		}

		authorization := c.Request.Header["Authorization"][0]

		rawToken := strings.Split(authorization, " ")
		if len(rawToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("token not valid")))
			return
		}

		token := rawToken[1]

		userInfo, err := am.accountClient.ValidateToken(c, &gen.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("error validating token")))
			return
		}

		c.Set(UserNameMiddlewareKey, userInfo.Name)
		c.Set(UserIDMiddlewareKey, userInfo.UserId)
		c.Set(UserAccountIDsMiddlewareKey, userInfo.AccountIds)

		c.Next()
	}
}
