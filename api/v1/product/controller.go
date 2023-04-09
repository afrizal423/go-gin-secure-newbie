package product

import (
	"net/http"

	"github.com/afrizal423/go-gin-secure-newbie/app/business/product"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service product.IProductService
}

func NewProductController(service product.IProductService) *Controller {
	return &Controller{
		service,
	}
}

// Role Admin (1)
// GET, GET ALL, UPDATE, DELETE, POST
// Role User (2)
// GET, GET ALL, POST
func (handler *Controller) GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
