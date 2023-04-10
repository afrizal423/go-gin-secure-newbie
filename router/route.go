package router

import (
	"github.com/afrizal423/go-gin-secure-newbie/api/v1/product"
	"github.com/afrizal423/go-gin-secure-newbie/api/v1/user"
	"github.com/afrizal423/go-gin-secure-newbie/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func Route(productHandler *product.Controller,
	userHandler *user.Controller) *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
	}

	produkRouter := r.Group("/products")
	{
		produkRouter.Use(middlewares.Authentication())
		produkRouter.GET("/", productHandler.GetAllProducts)
		produkRouter.POST("/", middlewares.CreateProductAuthorizations(), productHandler.CreateProduct)
		produkRouter.GET("/:productId", middlewares.DetailDataProductAuthorizations(), productHandler.GetProduct)
		produkRouter.PUT("/:productId", middlewares.UpdateProductAuthorizations(), productHandler.UpdateProduct)
		produkRouter.DELETE("/:productId", middlewares.DeleteDataProductAuthorizations(), productHandler.DeleteProduct)

	}
	return r
}
