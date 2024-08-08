package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/elangreza14/be-assignment/payment/dto"
	"github.com/gin-gonic/gin"
)

type (
	AuthMiddleware struct {
	}
)

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

const UserMiddlewareKey = "UserMiddlewareKey"

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
		fmt.Println(token)
		// TODO setup grpc here
		// user, err := am.authService.ProcessToken(c, token)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewBaseResponse(nil, errors.New("token unauthorize for this user")))
		// 	return
		// }

		c.Set(UserMiddlewareKey, nil)

		c.Next()
	}
}
