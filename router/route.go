package router

import (
	"github.com/afrizal423/go-gin-secure-newbie/api/v1/product"
	"github.com/afrizal423/go-gin-secure-newbie/api/v1/user"
	"github.com/gin-gonic/gin"
)

func Route(productHandler *product.Controller,
	userHandler *user.Controller) *gin.Engine {
	r := gin.Default()

	return r
}
