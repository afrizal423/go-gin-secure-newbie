package product

import "github.com/afrizal423/go-gin-secure-newbie/app/business/product"

type Controller struct {
	service product.IProductService
}

func NewProductController(service product.IProductService) *Controller {
	return &Controller{
		service,
	}
}
