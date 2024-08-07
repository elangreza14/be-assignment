package routes

import (
	"github.com/elangreza14/be-assignment/account/controller"
	"github.com/gin-gonic/gin"
)

func ProductRoute(route *gin.RouterGroup, ProductController *controller.ProductController) {
	ProductRoutes := route.Group("/products")
	ProductRoutes.GET("", ProductController.ProductList())
}
