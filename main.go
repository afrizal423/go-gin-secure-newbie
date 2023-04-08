package main

import (
	productController "github.com/afrizal423/go-gin-secure-newbie/api/v1/product"
	userController "github.com/afrizal423/go-gin-secure-newbie/api/v1/user"
	productService "github.com/afrizal423/go-gin-secure-newbie/app/business/product"
	userService "github.com/afrizal423/go-gin-secure-newbie/app/business/user"
	productRepository "github.com/afrizal423/go-gin-secure-newbie/app/repository/product"
	userRepository "github.com/afrizal423/go-gin-secure-newbie/app/repository/user"
	"github.com/afrizal423/go-gin-secure-newbie/configs"
	"github.com/afrizal423/go-gin-secure-newbie/database"
	"github.com/afrizal423/go-gin-secure-newbie/router"
)

func main() {
	conn := configs.GormPostgresConn()
	// migrate db
	database.DbMigrate(conn)

	productHandler := productController.NewProductController(productService.NewProductService(
		productRepository.NewProductRepository(conn)))

	userHandler := userController.NewUserController(userService.NewUserService(
		userRepository.NewuserRepository(conn)))

	r := router.Route(productHandler,
		userHandler)
	r.Run(":8080")
}
