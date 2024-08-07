package routes

import (
	"github.com/elangreza14/be-assignment/account/controller"
	"github.com/elangreza14/be-assignment/account/middleware"
	"github.com/gin-gonic/gin"
)

func AccountRoute(route *gin.RouterGroup,
	AuthMiddleware *middleware.AuthMiddleware,
	AccountController *controller.AccountController) {
	AccountRoutes := route.Group("/accounts")
	AccountRoutes.POST("", AuthMiddleware.MustAuthMiddleware(), AccountController.CreateAccount())
	AccountRoutes.GET("/me", AuthMiddleware.MustAuthMiddleware(), AccountController.GetAccountList())
}
