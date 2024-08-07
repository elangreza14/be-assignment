package routes

import (
	"github.com/elangreza14/be-assignment/account/controller"
	"github.com/gin-gonic/gin"
)

func CurrencyRoute(route *gin.RouterGroup, CurrencyController *controller.CurrencyController) {
	CurrencyRoutes := route.Group("/currencies")
	CurrencyRoutes.GET("", CurrencyController.CurrencyList())
}
