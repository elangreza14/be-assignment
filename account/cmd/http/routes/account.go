package routes

import (
	"github.com/elangreza14/be-assignment/account/controller"
	"github.com/elangreza14/be-assignment/account/middleware"
	"github.com/gin-gonic/gin"
)

func AccountRoute(route *gin.RouterGroup,
	AuthMiddleware *middleware.AuthMiddleware,
	AccountController *controller.AccountController) {
	route.POST("/accounts", AuthMiddleware.MustAuthMiddleware(), AccountController.CreateAccount())
	route.GET("/accounts/:accountID", AuthMiddleware.MustAuthMiddleware(), AccountController.GetAccountHistoriesList())
	route.GET("/accounts/me", AuthMiddleware.MustAuthMiddleware(), AccountController.GetAccountList())
}
