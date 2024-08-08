package routes

import (
	"github.com/elangreza14/be-assignment/payment/controller"
	"github.com/elangreza14/be-assignment/payment/middleware"
	"github.com/gin-gonic/gin"
)

func PaymentRoute(route *gin.RouterGroup,
	AuthMiddleware *middleware.AuthMiddleware,
	PaymentController *controller.PaymentController) {
	PaymentRoutes := route.Group("/payments")
	PaymentRoutes.POST("/send", AuthMiddleware.MustAuthMiddleware(), PaymentController.SendPayment())
	PaymentRoutes.POST("/withdraw", AuthMiddleware.MustAuthMiddleware(), PaymentController.WithdrawPayment())
}
